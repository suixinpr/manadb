package physical

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/analyzer/logic"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/parser/types"
)

/*
 * 物理表, 用来抽象表示物理计划中用到的每一个表
 */
type PhysicalTable interface {
	/* 返回该表及所有子节点中包含的所有物理表 ID */
	Domain() []int

	/* 向物理表注册一个逻辑计划中的变量, 并返回它在物理表中的位置索引. */
	RegisterVar(expr *expression.Variable) (int, error)

	/*
	 * 向物理表注册一个逻辑计划中的表达式, 并返回一个新的表达式, 这个
	 * 表达式从原有的直接引用基表, 改为了引用该物理表的列.
	 */
	RegisterExpr(expr expression.Expression) (expression.Expression, error)

	/* 设置物理表的谓词条件 */
	RegisterQual(qual expression.Expression) error

	/* 为物理表创建所有的基础路径 */
	CreatePathList()

	/* 为物理表添加父路径 */
	AppendPath(create func(phyTbl *BasePhysicalTable, subpath Path) Path)

	/* 为该物理表生成物理计划 */
	MakePhysicalPlan() PathPhysicalPlan
}

/*
 * 基础的物理表, 记录了生成计划需要的所有信息
 */
type BasePhysicalTable struct {
	/* 用来标识一个物理表, 每个物理表在同一个物理计划中 ID 唯一 */
	ID int

	/* 表的基础信息, 从逻辑计划中获取 */
	Info *logic.TableInfo

	/* 逻辑计划中的变量映射到物理表中变量的位置 */
	LogicVarMap map[expression.Variable]int

	/*
	 * 物理表内存在的所有变量, 这里的变量是自己朝自己的引用. 如果物理表中有
	 * 三个列分别是 int, int, text. 那么 Project 的值如下所示:
	 *
	 * Project = []expression.Expression{
	 *     &expression.Variable{VarNo: expression.OuterVar, VarColNo: 0, VarType: metadata.Int32ID},
	 *     &expression.Variable{VarNo: expression.OuterVar, VarColNo: 1, VarType: metadata.Int32ID},
	 *     &expression.Variable{VarNo: expression.OuterVar, VarColNo: 2, VarType: metadata.TextID},
	 * }
	 */
	Project []expression.Expression

	/*
	 * 物理表中的所有变量 Project 的生成方式, 它里面的位置和 Project 的位置一一
	 * 对应.
	 *
	 * 具体来说有 3 种情况: 1. 这个物理表 是一个假的物理表, 此时 FromExprs 的值
	 * 为 nil, 比如 select 1; 2. 这个物理 表是真的物理表, 但是没有引用下层物理
	 * 表, 此时 FromExprs 的值和 Project 的值完全相同, 比如对应顺序扫描生成的物
	 * 理表; 3. 这个物理表引用了下层物理表, 这 个时候 FromExprs 的值表示如何从下
	 * 层物理表生成当前物理表中的变量, 比如 Join 连接后生成的物理表.
	 */
	FromExprs []expression.Expression

	/* 物理表列的描述符, 记录了物理表所有投影前的列信息 */
	Desc *common.EntryDesc

	/* Where 条件 */
	Qual expression.Expression

	/* 生成该表的所有路径 */
	PathList []Path
}

/*
 * 新建一个基础表
 */
func NewBasePhysicalTable(id int, info *logic.TableInfo) *BasePhysicalTable {
	return &BasePhysicalTable{
		ID:          id,
		Info:        info,
		LogicVarMap: make(map[expression.Variable]int),
		Desc:        new(common.EntryDesc),
	}
}

/*
 * 在物理表中追加一个输出并返回其索引号, 输入的
 * 参数 from 表示这个变量从哪获取.
 */
func (phyTbl *BasePhysicalTable) appendProject(from *expression.Variable) (int, error) {
	no := len(phyTbl.Project)
	expr := &expression.Variable{VarNo: expression.OuterVar, VarColNo: no, VarType: from.VarType}
	phyTbl.Project = append(phyTbl.Project, expr)
	phyTbl.FromExprs = append(phyTbl.FromExprs, from)

	col, err := NewTemplateColumn(int16(no), "", from.VarType, -1)
	if err != nil {
		return -1, err
	}
	phyTbl.Desc.Cols = append(phyTbl.Desc.Cols, col)
	return no, nil
}

/*
* 对于基础物理表, 它包含的表只有自己.
 */
func (phyTbl *BasePhysicalTable) Domain() []int {
	return []int{phyTbl.ID}
}

/*
 * 注册逻辑计划中的变量
 */
func (phyTbl *BasePhysicalTable) RegisterVar(expr *expression.Variable) (int, error) {
	if expr.VarNo != phyTbl.ID {
		return 0, errlog.New("not found var in TableItem")
	}
	if no, ok := phyTbl.LogicVarMap[*expr]; ok {
		return no, nil
	}
	from := &expression.Variable{VarNo: expression.OuterVar, VarColNo: expr.VarColNo, VarType: expr.VarType}
	no, err := phyTbl.appendProject(from)
	if err != nil {
		return -1, err
	}
	phyTbl.LogicVarMap[*expr] = no
	return no, nil
}

func (phyTbl *BasePhysicalTable) RegisterExpr(expr expression.Expression) (expression.Expression, error) {
	var err error
	var mutate expression.Mutator
	mutate = func(expr expression.Expression) expression.Expression {
		if err != nil || expr == nil {
			return nil
		}
		switch expr := expr.(type) {
		case *expression.Variable:
			var no int
			no, err = phyTbl.RegisterVar(expr)
			return phyTbl.Project[no]
		default:
			return expr.Mutator(mutate)
		}
	}
	return mutate(expr), nil
}

/*
 * 注册物理表的 Where 条件
 */
func (phyTbl *BasePhysicalTable) RegisterQual(qual expression.Expression) error {
	qual, err := phyTbl.RegisterExpr(qual)
	phyTbl.Qual = qual
	return err
}

/*
 * 创建该物理表所能拥有的的所有路径
 */
func (phyTbl *BasePhysicalTable) CreatePathList() {
	if len(phyTbl.Desc.Cols) == 0 {
		path := CreateResultPath(phyTbl)
		phyTbl.PathList = append(phyTbl.PathList, path)
		return
	}

	path := CreateSeqScanPath(phyTbl)
	phyTbl.PathList = append(phyTbl.PathList, path)
}

/*
 * 向物理表中新增路径作为所有原有路径的父亲节点
 */
func (phyTbl *BasePhysicalTable) AppendPath(create func(phyTbl *BasePhysicalTable, subpath Path) Path) {
	for i, subpath := range phyTbl.PathList {
		path := create(phyTbl, subpath)
		phyTbl.PathList[i] = path
	}
}

/*
 * 为物理表生成物理计划
 */
func (phyTbl *BasePhysicalTable) MakePhysicalPlan() PathPhysicalPlan {
	path := phyTbl.PathList[0]
	return path.GeneratePlan()
}

/*
 * 一个 Join 类型的物理表, 附带额外的 Join 相关的信息和操作
 */
type JoinPhysicalTable struct {
	BasePhysicalTable

	/* 该 Item 及所有子节点所包含的 Var */
	domain []int

	/* join 相关类型 */
	Type  types.JoinType
	Outer PhysicalTable
	Inner PhysicalTable
	On    expression.Expression /* join ... on 条件 */
}

/*
 * 新建一个 Join 表
 */
func NewJoinPhysicalTable(id int, info *logic.TableInfo, obj *logic.JoinObject, outer, inner PhysicalTable) (*JoinPhysicalTable, error) {
	phyTbl := &JoinPhysicalTable{
		BasePhysicalTable: BasePhysicalTable{
			ID:          id,
			Info:        info,
			LogicVarMap: make(map[expression.Variable]int),
			Desc:        new(common.EntryDesc),
		},
	}

	phyTbl.domain = []int{id}
	phyTbl.domain = append(phyTbl.domain, outer.Domain()...)
	phyTbl.domain = append(phyTbl.domain, inner.Domain()...)

	phyTbl.Type = obj.Type
	phyTbl.Outer = outer
	phyTbl.Inner = inner

	/* 注册 on 表达式 */
	on, err := phyTbl.RegisterExpr(obj.On)
	if err != nil {
		return nil, err
	}
	phyTbl.On = on

	return phyTbl, nil
}

/*
 * 对于 Join 表, 它包含的自己以及所有子节点的 ID.
 */
func (phyTbl *JoinPhysicalTable) Domain() []int {
	return phyTbl.domain
}

/*
 * 新增输出变量, 需要同时在 outer 和 inner 节点新增输出变量
 */
func (phyTbl *JoinPhysicalTable) RegisterVar(expr *expression.Variable) (int, error) {
	if !inDomain(expr.VarNo, phyTbl.domain) {
		return 0, errlog.New("not found var in JoinItem")
	}
	if no, ok := phyTbl.LogicVarMap[*expr]; ok {
		return no, nil
	}

	/* 生成 from 变量 */
	from := &expression.Variable{VarType: expr.VarType}
	if inDomain(expr.VarNo, phyTbl.Outer.Domain()) {
		no, err := phyTbl.Outer.RegisterVar(expr)
		if err != nil {
			return 0, err
		}
		from.VarNo = expression.OuterVar
		from.VarColNo = no
	} else if inDomain(expr.VarNo, phyTbl.Inner.Domain()) {
		no, err := phyTbl.Inner.RegisterVar(expr)
		if err != nil {
			return 0, err
		}
		from.VarNo = expression.InnerVar
		from.VarColNo = no
	} else {
		return 0, errlog.New("not found var in JoinItem")
	}

	no, err := phyTbl.appendProject(from)
	if err != nil {
		return -1, err
	}
	phyTbl.LogicVarMap[*expr] = no
	return no, nil
}

func (phyTbl *JoinPhysicalTable) RegisterExpr(expr expression.Expression) (expression.Expression, error) {
	var err error
	var mutate expression.Mutator
	mutate = func(expr expression.Expression) expression.Expression {
		if err != nil || expr == nil {
			return nil
		}
		switch expr := expr.(type) {
		case *expression.Variable:
			var no int
			no, err = phyTbl.RegisterVar(expr)
			return phyTbl.Project[no]
		default:
			return expr.Mutator(mutate)
		}
	}
	return mutate(expr), nil
}

/*
 * 注册物理表的 Where 条件
 */
func (phyTbl *JoinPhysicalTable) RegisterQual(qual expression.Expression) error {
	qual, err := phyTbl.RegisterExpr(qual)
	phyTbl.Qual = qual
	return err
}

/*
 * 创建该物理表所能拥有的的所有路径
 */
func (phyTbl *JoinPhysicalTable) CreatePathList() {
	phyTbl.Outer.CreatePathList()
	phyTbl.Inner.CreatePathList()

	path := CreateJoinPath(phyTbl)
	phyTbl.PathList = append(phyTbl.PathList, path)
}
