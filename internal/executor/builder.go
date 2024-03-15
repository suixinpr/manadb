package executor

import (
	"fmt"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/planner/physical"
	"github/suixinpr/manadb/internal/relation/table"
)

type ExecutorBuilder struct {
}

func NewExecutorBuilder() *ExecutorBuilder {
	return &ExecutorBuilder{}
}

func (b *ExecutorBuilder) Build(plan physical.PhysicalPlan) (Executor, error) {
	switch plan := plan.(type) {
	case nil:
		return nil, errlog.New("nil physical plan")
	case *physical.CommandPlan:
		return b.buildCommandExec(plan)
	case *physical.DeletePlan:
		return b.buildDeleteExec(plan)
	case *physical.ExplainPlan:
		return b.buildExplainExec(plan)
	case *physical.InsertPlan:
		return b.buildInsertExec(plan)
	case *physical.JoinPlan:
		return b.buildJoinExec(plan)
	case *physical.LimitPlan:
		return b.buildLimitExec(plan)
	case *physical.ResultPlan:
		return b.buildResultExec(plan)
	case *physical.SeqScanPlan:
		return b.buildSeqScanExec(plan)
	case *physical.UpdatePlan:
		return b.buildUpdateExec(plan)
	case *physical.ValuesPlan:
		return b.buildValuesExec(plan)
	default:
		return nil, errlog.New(fmt.Sprintf("Build not support PhysicalPlan type %T", plan))
	}
}

func (b *ExecutorBuilder) buildCommandExec(plan *physical.CommandPlan) (*CommandExec, error) {
	exec := new(CommandExec)
	exec.desc = plan.Desc

	exec.plan = plan.Plan
	return exec, nil
}

func (b *ExecutorBuilder) buildDeleteExec(plan *physical.DeletePlan) (*DeleteExec, error) {
	tbl, err := table.Open(plan.TableID)
	if err != nil {
		return nil, err
	}

	from, err := b.Build(plan.From)
	if err != nil {
		return nil, err
	}

	exec := new(DeleteExec)
	exec.desc = plan.Desc
	exec.outer = from

	exec.tbl = tbl
	return exec, nil
}

func (b *ExecutorBuilder) buildExplainExec(plan *physical.ExplainPlan) (*ExplainExec, error) {
	exec := new(ExplainExec)
	exec.desc = plan.Desc

	exec.plan = plan.Plan
	return exec, nil
}

func (b *ExecutorBuilder) buildInsertExec(plan *physical.InsertPlan) (*InsertExec, error) {
	tbl, err := table.Open(plan.TableID)
	if err != nil {
		return nil, err
	}

	from, err := b.Build(plan.From)
	if err != nil {
		return nil, err
	}

	exec := new(InsertExec)
	exec.desc = plan.Desc
	exec.outer = from

	exec.tbl = tbl
	return exec, nil
}

func (b *ExecutorBuilder) buildJoinExec(plan *physical.JoinPlan) (*JoinExec, error) {
	outer, err := b.Build(plan.Outer)
	if err != nil {
		return nil, err
	}
	inner, err := b.Build(plan.Inner)
	if err != nil {
		return nil, err
	}

	exec := new(JoinExec)
	exec.outer = outer
	exec.inner = inner
	exec.project = plan.Project
	exec.desc = plan.Desc

	exec.tp = plan.Type
	exec.on = plan.On
	exec.qual = plan.Qual
	return exec, nil
}

func (b *ExecutorBuilder) buildLimitExec(plan *physical.LimitPlan) (*LimitExec, error) {
	from, err := b.Build(plan.From)
	if err != nil {
		return nil, err
	}

	exec := new(LimitExec)
	exec.outer = from
	exec.project = plan.Project
	exec.desc = plan.Desc

	exec.count = plan.Count
	exec.offset = plan.Offset
	return exec, nil
}

func (b *ExecutorBuilder) buildResultExec(plan *physical.ResultPlan) (*ResultExec, error) {
	exec := new(ResultExec)
	exec.project = plan.Project
	exec.desc = plan.Desc

	exec.qual = plan.Qual
	return exec, nil
}

func (b *ExecutorBuilder) buildSeqScanExec(plan *physical.SeqScanPlan) (*SeqScanExec, error) {
	tbl, err := table.Open(plan.TableID)
	if err != nil {
		return nil, err
	}

	exec := new(SeqScanExec)
	exec.project = plan.Project
	exec.desc = plan.Desc

	exec.tbl = tbl
	exec.qual = plan.Qual
	return exec, nil
}

func (b *ExecutorBuilder) buildUpdateExec(plan *physical.UpdatePlan) (*UpdateExec, error) {
	tbl, err := table.Open(plan.TableID)
	if err != nil {
		return nil, err
	}

	from, err := b.Build(plan.From)
	if err != nil {
		return nil, err
	}

	exec := new(UpdateExec)
	exec.desc = plan.Desc
	exec.outer = from

	exec.tbl = tbl
	exec.project = plan.Project
	return exec, nil
}

func (b *ExecutorBuilder) buildValuesExec(plan *physical.ValuesPlan) (*ValuesExec, error) {
	exec := new(ValuesExec)
	exec.valuesList = plan.ValuesList
	return exec, nil
}
