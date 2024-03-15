package physical

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/analyzer/logic"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/parser/types"
)

/*
 * 物理计划
 *
 * 包括两种类型, 分别是 NonPathPhysicalPlan 和 PathPhysicalPlan.
 * 前者表示不是通过 Path 生成的物理计划, 这些物理计划各有特性, 自己
 * 创建单独管理.
 *
 * 另一类 PathPhysicalPlan 表示通过 Path 生成的计划, 这部分计划都会
 * 生成一个 entry 向后传递, 也都会执行 project.
 */
type PhysicalPlan interface {
	/* 设置物理计划的描述符 */
	ModifyDescription(desc *common.EntryDesc)
}

type physicalPlan struct {
	Cost *Cost             /* 执行计划该节点花费代价 */
	Desc *common.EntryDesc /* 生成行的描述符 */
}

func (plan *physicalPlan) ModifyDescription(desc *common.EntryDesc) {
	plan.Desc = desc
}

/********************************************************************************
*
*  NonPathPhysicalPlan
*
********************************************************************************/

/*
 * Command Plan
 */
type CommandPlan struct {
	physicalPlan

	Plan logic.LogicPlan
}

func NewCommandPlan(desc *common.EntryDesc, plan logic.LogicPlan) *CommandPlan {
	return &CommandPlan{
		physicalPlan: physicalPlan{Cost: CalculateCost(0, 1), Desc: desc},
		Plan:         plan,
	}
}

/*
 * Delete Plan
 */
type DeletePlan struct {
	physicalPlan

	TableID metadata.OID
	From    PhysicalPlan
}

func NewDeletePlan(desc *common.EntryDesc, tblID metadata.OID, from PhysicalPlan) *DeletePlan {
	return &DeletePlan{
		physicalPlan: physicalPlan{Cost: CalculateCost(0, 1), Desc: desc},
		TableID:      tblID,
		From:         from,
	}
}

/*
 * Explain Plan
 */
type ExplainPlan struct {
	physicalPlan

	Plan PhysicalPlan
}

func NewExplainPlan(desc *common.EntryDesc, plan PhysicalPlan) *ExplainPlan {
	return &ExplainPlan{
		physicalPlan: physicalPlan{Cost: CalculateCost(0, 1), Desc: desc},
		Plan:         plan,
	}
}

/*
 * Insert Plan
 */
type InsertPlan struct {
	physicalPlan

	TableID metadata.OID
	From    PhysicalPlan
}

func NewInsertPlan(desc *common.EntryDesc, tblID metadata.OID, from PhysicalPlan) *InsertPlan {
	return &InsertPlan{
		physicalPlan: physicalPlan{Cost: CalculateCost(0, 1), Desc: desc},
		TableID:      tblID,
		From:         from,
	}
}

/*
 * Update Plan
 */
type UpdatePlan struct {
	physicalPlan

	TableID metadata.OID
	Project []expression.Expression
	From    PhysicalPlan
}

func NewUpdatePlan(desc *common.EntryDesc, tblID metadata.OID, project []expression.Expression, from PhysicalPlan) *UpdatePlan {
	return &UpdatePlan{
		physicalPlan: physicalPlan{Cost: CalculateCost(0, 1), Desc: desc},
		TableID:      tblID,
		Project:      project,
		From:         from,
	}
}

/*
 * Values Plan
 */
type ValuesPlan struct {
	physicalPlan

	ValuesList [][]expression.Expression
}

func NewValuesPlan(desc *common.EntryDesc, exprsList [][]expression.Expression) *ValuesPlan {
	return &ValuesPlan{
		physicalPlan: physicalPlan{Cost: CalculateCost(0, 1), Desc: desc},
		ValuesList:   exprsList,
	}
}

/********************************************************************************
*
*  PathPhysicalPlan
*
********************************************************************************/

type PathPhysicalPlan interface {
	PhysicalPlan

	/* 将一个物理表中的表达式变为计划中的表达式 */
	AdjustExpression(expr expression.Expression) expression.Expression

	/* 设置物理计划的投影列 */
	ModifyProject(exprs []expression.Expression)
}

type pathPhysicalPlan struct {
	physicalPlan
	Project []expression.Expression /* 投影列生成表达式 */
}

func (plan *pathPhysicalPlan) AdjustExpression(expr expression.Expression) expression.Expression {
	var mutate expression.Mutator
	mutate = func(expr expression.Expression) expression.Expression {
		if expr == nil {
			return nil
		}
		switch expr := expr.(type) {
		case *expression.Variable:
			return plan.Project[expr.VarColNo]
		default:
			return expr.Mutator(mutate)
		}
	}
	return mutate(expr)
}

func (plan *pathPhysicalPlan) ModifyProject(exprs []expression.Expression) {
	project := make([]expression.Expression, len(exprs))
	for i, expr := range exprs {
		project[i] = plan.AdjustExpression(expr)
	}
	plan.Project = project
}

/*
 * Join Plan
 */
type JoinPlan struct {
	pathPhysicalPlan

	Type  types.JoinType
	Outer PhysicalPlan
	Inner PhysicalPlan
	On    expression.Expression
	Qual  expression.Expression
}

/*
 * Limit Plan
 */
type LimitPlan struct {
	pathPhysicalPlan

	From   PhysicalPlan
	Count  expression.Expression
	Offset expression.Expression
}

/*
 * Result Plan
 */
type ResultPlan struct {
	pathPhysicalPlan

	Qual expression.Expression /* 谓词条件 */
}

/*
 * Seq Scan Plan
 */
type SeqScanPlan struct {
	pathPhysicalPlan

	TableID metadata.OID
	Qual    expression.Expression /* 谓词条件 */
}
