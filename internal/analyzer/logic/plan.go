package logic

import (
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/mana/metadata"
)

type LogicPlan interface {
	logicPlanImpl()
}

type logicPlan struct{}

func (plan *logicPlan) logicPlanImpl() {}

/*
 * 查询计划, 用来表达从物理表中获取 entry 的计划. 对于 SelectPlan,
 * InsertPlan, UpdatePlan 和 DeletePlan 都会用到该计划, 因为他们
 * 的本质都是查询计划, 只是分别在查询到 entry 后的操作不用.
 *   - SelectPlan: 输出
 *   - InsertPlan: 插入
 *   - UpdatePlan: 更新
 *   - DeletePlan: 删除
 * 所以这里使用一个抽象的 QueryPlan 来表示获取 entry 这个过程.
 */
type QueryPlan struct {
	logicPlan

	Tables     []*TableInfo          /* 查询中会用到的表 */
	TargetList []*TargetInfo         /* 查询中的投影列 */
	From       FromObject            /* 查询对象 */
	Qual       expression.Expression /* where 条件 */
}

/*
 * 表示一个 SELECT 语句.
 */
type SelectPlan struct {
	logicPlan

	Query *QueryPlan
	Limit *LimitPlan
}

type InsertPlan struct {
	logicPlan

	Table     *TableInfo
	ExprsList [][]expression.Expression
	Query     *QueryPlan
}

type UpdatePlan struct {
	logicPlan

	Table   *TableInfo
	SetList []expression.Expression
	Query   *QueryPlan
}

type DeletePlan struct {
	logicPlan

	Table *TableInfo
	Query *QueryPlan
}

type CreateTablePlan struct {
	logicPlan

	TableName string
	Columns   []*ColumnInfo
}

type DropTablePlan struct {
	logicPlan

	Tables []metadata.OID
}

type ExplainPlan struct {
	logicPlan

	Plan LogicPlan
}

type LimitPlan struct {
	logicPlan

	Count  expression.Expression
	Offset expression.Expression
}
