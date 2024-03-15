package builder

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/analyzer/logic"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/planner/physical"
)

func BuildPhysicalPlan(lp logic.LogicPlan) (physical.PhysicalPlan, error) {
	b := new(PlanBuilder)
	return b.build(lp)
}

/* 物理计划生成器 */
type PlanBuilder struct {
	/* 逻辑计划中的所有表信息 */
	tables []*logic.TableInfo

	/* 整个物理计划的顶层物理表 */
	phyTbl physical.PhysicalTable

	/* 物理计划的投影 */
	project []expression.Expression

	/* 物理计划执行输出的描述符 */
	desc *common.EntryDesc
}

func (b *PlanBuilder) build(lp logic.LogicPlan) (physical.PhysicalPlan, error) {
	switch lp := lp.(type) {
	case *logic.DeletePlan:
		return b.buildDeletePlan(lp)
	case *logic.ExplainPlan:
		return b.buildExplainPlan(lp)
	case *logic.InsertPlan:
		return b.buildInsertPlan(lp)
	case *logic.SelectPlan:
		return b.buildSelectPlan(lp)
	case *logic.UpdatePlan:
		return b.buildUpdatePlan(lp)
	default:
		return b.buildCommandPlan(lp)
	}
}

func (b *PlanBuilder) buildCommandPlan(lp logic.LogicPlan) (*physical.CommandPlan, error) {
	err := b.collectCommandDesc()
	if err != nil {
		return nil, err
	}

	return physical.NewCommandPlan(b.desc, lp), nil
}

func (b *PlanBuilder) buildDeletePlan(lp *logic.DeletePlan) (*physical.DeletePlan, error) {
	err := b.collectQueryPlan(lp.Query)
	if err != nil {
		return nil, err
	}

	/* 获取处理完后的物理表 */
	final := b.phyTbl

	/* 创建基础路径 */
	final.CreatePathList()

	/* 生成物理计划 */
	from := final.MakePhysicalPlan()

	err = b.collectDeleteDesc()
	if err != nil {
		return nil, err
	}

	return physical.NewDeletePlan(b.desc, lp.Table.TableID, from), nil
}

func (b *PlanBuilder) buildExplainPlan(lp *logic.ExplainPlan) (*physical.ExplainPlan, error) {
	plan, err := b.build(lp.Plan)
	if err != nil {
		return nil, err
	}

	err = b.collectExplainDesc()
	if err != nil {
		return nil, err
	}

	return physical.NewExplainPlan(b.desc, plan), nil
}

func (b *PlanBuilder) buildInsertPlan(lp *logic.InsertPlan) (*physical.InsertPlan, error) {
	err := b.collectInsertDesc()
	if err != nil {
		return nil, err
	}

	if lp.ExprsList != nil {
		from := physical.NewValuesPlan(nil, lp.ExprsList)
		return physical.NewInsertPlan(b.desc, lp.Table.TableID, from), nil
	}

	return nil, errlog.New("not support")
}

func (b *PlanBuilder) buildSelectPlan(lp *logic.SelectPlan) (physical.PathPhysicalPlan, error) {
	err := b.collectQueryPlan(lp.Query)
	if err != nil {
		return nil, err
	}

	/* 获取处理完后的物理表 */
	final := b.phyTbl

	/* 创建基础路径 */
	final.CreatePathList()

	/* 处理 Limit 关键字 */
	if lp.Limit != nil {
		final.AppendPath(physical.CreateLimitPath(lp.Limit.Count, lp.Limit.Offset))
	}

	/* 生成物理计划 */
	plan := final.MakePhysicalPlan()
	plan.ModifyDescription(b.desc)
	plan.ModifyProject(b.project)
	return plan, nil
}

func (b *PlanBuilder) buildUpdatePlan(lp *logic.UpdatePlan) (*physical.UpdatePlan, error) {
	err := b.collectQueryPlan(lp.Query)
	if err != nil {
		return nil, err
	}

	/* 获取处理完后的物理表 */
	final := b.phyTbl

	/* 创建基础路径 */
	final.CreatePathList()

	/* 生成物理计划 */
	from := final.MakePhysicalPlan()

	err = b.collectUpdateDesc()
	if err != nil {
		return nil, err
	}

	var mutate expression.Mutator
	mutate = func(expr expression.Expression) expression.Expression {
		if expr == nil {
			return nil
		}
		switch expr := expr.(type) {
		case *expression.Variable:
			return &expression.Variable{VarNo: expression.OuterVar, VarColNo: expr.VarColNo, VarType: expr.VarType}
		default:
			return expr.Mutator(mutate)
		}
	}
	project := make([]expression.Expression, len(lp.SetList))
	for i, expr := range lp.SetList {
		project[i] = mutate(expr)
	}

	return physical.NewUpdatePlan(b.desc, lp.Table.TableID, project, from), nil
}
