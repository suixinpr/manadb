package planner

import (
	"github/suixinpr/manadb/internal/analyzer"
	"github/suixinpr/manadb/internal/parser/ast"
	"github/suixinpr/manadb/internal/planner/builder"
	"github/suixinpr/manadb/internal/planner/physical"
	"github/suixinpr/manadb/pkg/debug"
)

func BuildPlan(stmt ast.StmtNode) (physical.PhysicalPlan, error) {
	/* 语义分析 */
	plan, err := analyzer.Analyze(stmt)
	if err != nil {
		return nil, err
	}

	/* 逻辑优化 */
	plan, err = Optimize(plan)
	if err != nil {
		return nil, err
	}

	debug.Print("logic plan", plan)

	/* 创建物理计划 */
	return builder.BuildPhysicalPlan(plan)
}
