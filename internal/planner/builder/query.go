package builder

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/analyzer/logic"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/planner/physical"
)

func (b *PlanBuilder) collectQueryPlan(query *logic.QueryPlan) error {
	var err error

	err = b.collectTables(query.Tables)
	if err != nil {
		return err
	}

	err = b.collectFromObject(query.From)
	if err != nil {
		return err
	}

	err = b.collectTargetList(query.TargetList)
	if err != nil {
		return err
	}

	err = b.collectQual(query.Qual)
	if err != nil {
		return err
	}

	return nil
}

/* 收集逻辑计划中表的信息 */
func (b *PlanBuilder) collectTables(tables []*logic.TableInfo) error {
	b.tables = tables
	return nil
}

/*
 * 收集逻辑计划中的 from 对象的信息, 并生成物理表信息. 在调用该函数前, 需要
 * 先调用 collectTables 收集表信息.
 */
func (b *PlanBuilder) collectFromObject(from logic.FromObject) error {
	/* 返回一个虚拟的物理表 */
	if from == nil {
		b.phyTbl = physical.NewBasePhysicalTable(0, nil)
		return nil
	}

	var collect func(obj logic.FromObject) (physical.PhysicalTable, error)
	collect = func(obj logic.FromObject) (physical.PhysicalTable, error) {
		switch obj := obj.(type) {
		case *logic.JoinObject:
			outer, err := collect(obj.Left)
			if err != nil {
				return nil, err
			}
			inner, err := collect(obj.Right)
			if err != nil {
				return nil, err
			}
			return physical.NewJoinPhysicalTable(obj.Index, b.tables[obj.Index], obj, outer, inner)
		case *logic.TableObject:
			return physical.NewBasePhysicalTable(obj.Index, b.tables[obj.Index]), nil
		default:
			return nil, errlog.New("not support")
		}
	}
	phyTbl, err := collect(from)
	b.phyTbl = phyTbl
	return err
}

/*
 * 收集逻辑计划中 TargetList 的信息, 获取物理表的投影和输出的描述符.
 */
func (b *PlanBuilder) collectTargetList(targetList []*logic.TargetInfo) error {
	if targetList == nil {
		return nil
	}

	b.project = make([]expression.Expression, len(targetList))
	for i, target := range targetList {
		expr, err := b.phyTbl.RegisterExpr(target.Expr)
		if err != nil {
			return err
		}
		b.project[i] = expr
	}

	b.desc = new(common.EntryDesc)
	b.desc.Cols = make([]*metadata.ManaColumns, len(targetList))
	for i, target := range targetList {
		col, err := physical.NewTemplateColumn(int16(i), target.Name, b.project[i].TypeID(), -1)
		if err != nil {
			return err
		}
		b.desc.Cols[i] = col
	}

	return nil
}

/* 收集谓词信息 */
func (b *PlanBuilder) collectQual(qual expression.Expression) error {
	if qual == nil {
		return nil
	}
	return b.phyTbl.RegisterQual(qual)
}
