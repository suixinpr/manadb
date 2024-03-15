package physical

import (
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/parser/types"
)

/*
 * 路径
 *
 * 表示生成该物理表的步骤, 分为两种类型: producer 和 medium. producer
 * 物理表中的最底层的路径, medium 表示物理表的路径列表里中间的路径.
 * medium 存在 subpath 而 producer 没有
 */
type Path interface {
	/* 将该路径转化成物理计划 */
	GeneratePlan() PathPhysicalPlan
}

type path struct {
	Ownner *BasePhysicalTable
	Cost   *Cost /* 路径花费代价 */
}

/********************************************************************************
*
*  生产者
*
********************************************************************************/

type producer struct {
	path
}

/*
 * JoinPath
 */
type JoinPath struct {
	producer

	Type  types.JoinType
	Outer PhysicalTable
	Inner PhysicalTable
	On    expression.Expression
}

func CreateJoinPath(phyTbl *JoinPhysicalTable) *JoinPath {
	path := new(JoinPath)
	path.Ownner = &phyTbl.BasePhysicalTable
	path.Cost = CalculateCost(0, 1)

	path.Type = phyTbl.Type
	path.Outer = phyTbl.Outer
	path.Inner = phyTbl.Inner
	path.On = phyTbl.On
	return path
}

func (path *JoinPath) GeneratePlan() PathPhysicalPlan {
	plan := new(JoinPlan)
	plan.Cost = path.Cost
	plan.Desc = path.Ownner.Desc
	plan.Project = path.Ownner.FromExprs

	/* extra information */
	plan.Type = path.Type
	plan.Outer = path.Outer.MakePhysicalPlan()
	plan.Inner = path.Inner.MakePhysicalPlan()
	plan.On = plan.AdjustExpression(path.On)
	plan.Qual = plan.AdjustExpression(path.Ownner.Qual)
	return plan
}

/*
 * ResultPath
 */
type ResultPath struct {
	path
}

func CreateResultPath(phyTbl *BasePhysicalTable) *ResultPath {
	path := new(ResultPath)
	path.Ownner = phyTbl
	path.Cost = CalculateCost(0, 1)

	return path
}

func (path *ResultPath) GeneratePlan() PathPhysicalPlan {
	plan := new(ResultPlan)
	plan.Cost = path.Cost
	plan.Desc = path.Ownner.Desc
	plan.Project = path.Ownner.FromExprs

	/* extra information */
	plan.Qual = plan.AdjustExpression(path.Ownner.Qual)
	return plan
}

/*
 * SeqScanPath
 */
type SeqScanPath struct {
	path
}

func CreateSeqScanPath(phyTbl *BasePhysicalTable) *SeqScanPath {
	path := new(SeqScanPath)
	path.Ownner = phyTbl
	path.Cost = CalculateCost(0, 1)

	return path
}

func (path *SeqScanPath) GeneratePlan() PathPhysicalPlan {
	plan := new(SeqScanPlan)
	plan.Cost = path.Cost
	plan.Desc = path.Ownner.Desc
	plan.Project = path.Ownner.FromExprs

	/* extra information */
	plan.Qual = plan.AdjustExpression(path.Ownner.Qual)
	plan.TableID = path.Ownner.Info.TableID
	return plan
}

/********************************************************************************
*
*  中间商
*
********************************************************************************/

type medium struct {
	path
	Subpath Path
}

/* Limit 路径 */
type LimitPath struct {
	medium

	Count  expression.Expression
	Offset expression.Expression
}

func CreateLimitPath(count, offset expression.Expression) func(phyTbl *BasePhysicalTable, subpath Path) Path {
	return func(phyTbl *BasePhysicalTable, subpath Path) Path {
		path := new(LimitPath)
		path.Ownner = phyTbl
		path.Cost = CalculateCost(0, 1)
		path.Subpath = subpath

		path.Count = count
		path.Offset = offset
		return path
	}
}

func (path *LimitPath) GeneratePlan() PathPhysicalPlan {
	plan := new(LimitPlan)
	plan.Cost = path.Cost
	plan.Desc = path.Ownner.Desc
	plan.Project = path.Ownner.Project

	/* extra information */
	plan.From = path.Subpath.GeneratePlan()
	plan.Count = path.Count
	plan.Offset = path.Offset
	return plan
}
