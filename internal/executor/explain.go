package executor

import (
	"context"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/executor/row"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/planner/physical"
	"strings"
)

type ExplainExec struct {
	executor

	plan physical.PhysicalPlan

	/* runtime information */
	id   int
	node int
}

func (exec *ExplainExec) run(ctx context.Context, output chan<- *Result, outer, inner <-chan *Result) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		exec.id = -1
		exec.node = 0
		exec.explain(output, exec.plan)
		return nil
	}
}

/*
 * ------------------------------------------------------
 * | id | node | operation | qual | start | run | total |
 * ------------------------------------------------------
 */
func (exec *ExplainExec) explain(output chan<- *Result, plan physical.PhysicalPlan) error {
	exec.id++
	switch plan := plan.(type) {
	case *physical.DeletePlan:
		return exec.explainDeletePlan(output, plan)
	case *physical.JoinPlan:
		return exec.explainJoinPlan(output, plan)
	case *physical.InsertPlan:
		return exec.explainInsertPlan(output, plan)
	case *physical.LimitPlan:
		return exec.explainLimitPlan(output, plan)
	case *physical.ResultPlan:
		return exec.explainResultPlan(output, plan)
	case *physical.SeqScanPlan:
		return exec.explainSeqScanPlan(output, plan)
	case *physical.UpdatePlan:
		return exec.explainUpdatePlan(output, plan)
	case *physical.ValuesPlan:
		return exec.explainValues(output, plan)
	default:
		return errlog.New("explain error")
	}
}

func (exec *ExplainExec) explainDeletePlan(output chan<- *Result, plan *physical.DeletePlan) error {
	output <- explain(exec.id, exec.node, "Delete", "", plan.Cost.Startup, plan.Cost.Run, plan.Cost.Total)

	exec.node++
	defer func() { exec.node-- }()
	return exec.explain(output, plan.From)
}

func (exec *ExplainExec) explainJoinPlan(output chan<- *Result, plan *physical.JoinPlan) error {
	operation := "Nestloop Join"
	if plan.On != nil {
		on, err := plan.On.ToString()
		if err != nil {
			return err
		}
		operation += " On " + on
	}

	qual, err := explainQual(plan.Qual)
	if err != nil {
		return err
	}

	output <- explain(exec.id, exec.node, operation, qual, plan.Cost.Startup, plan.Cost.Run, plan.Cost.Total)

	/* children nodes */
	exec.node++
	defer func() { exec.node-- }()
	err = exec.explain(output, plan.Outer)
	if err != nil {
		return err
	}
	err = exec.explain(output, plan.Inner)
	if err != nil {
		return err
	}
	return nil
}

func (exec *ExplainExec) explainInsertPlan(output chan<- *Result, plan *physical.InsertPlan) error {
	output <- explain(exec.id, exec.node, "Insert", "", plan.Cost.Startup, plan.Cost.Run, plan.Cost.Total)

	exec.node++
	defer func() { exec.node-- }()
	return exec.explain(output, plan.From)
}

func (exec *ExplainExec) explainLimitPlan(output chan<- *Result, plan *physical.LimitPlan) error {
	operation := ""
	if plan.Count != nil {
		res, err := plan.Count.ToString()
		if err != nil {
			return err
		}
		operation += "Limit " + res
	}
	if plan.Offset != nil {
		res, err := plan.Offset.ToString()
		if err != nil {
			return err
		}
		operation += "Offset " + res
	}

	output <- explain(exec.id, exec.node, operation, "", plan.Cost.Startup, plan.Cost.Run, plan.Cost.Total)

	return exec.explain(output, plan.From)
}

func (exec *ExplainExec) explainResultPlan(output chan<- *Result, plan *physical.ResultPlan) error {
	qual, err := explainQual(plan.Qual)
	if err != nil {
		return err
	}

	output <- explain(exec.id, exec.node, "Result", qual, plan.Cost.Startup, plan.Cost.Run, plan.Cost.Total)
	return nil
}

func (exec *ExplainExec) explainSeqScanPlan(output chan<- *Result, plan *physical.SeqScanPlan) error {
	qual, err := explainQual(plan.Qual)
	if err != nil {
		return err
	}

	output <- explain(exec.id, exec.node, "SeqScan", qual, plan.Cost.Startup, plan.Cost.Run, plan.Cost.Total)
	return nil
}

func (exec *ExplainExec) explainUpdatePlan(output chan<- *Result, plan *physical.UpdatePlan) error {
	output <- explain(exec.id, exec.node, "Update", "", plan.Cost.Startup, plan.Cost.Run, plan.Cost.Total)

	exec.node++
	defer func() { exec.node-- }()
	return exec.explain(output, plan.From)
}

func (exec *ExplainExec) explainValues(output chan<- *Result, plan *physical.ValuesPlan) error {
	output <- explain(exec.id, exec.node, "Values", "", plan.Cost.Startup, plan.Cost.Run, plan.Cost.Total)
	return nil
}

func explain(id, node int, operation string, qual string, startup, run, total int) *Result {
	values := make([]datum.Datum, 7)
	values[0] = datum.ValueGetDatum[int32](int32(id))
	values[1] = datum.ValueGetDatum[int32](int32(node))
	values[2] = datum.StringGetDatum(strings.Repeat("  ", node) + operation)
	values[3] = datum.StringGetDatum(qual)
	values[4] = datum.ValueGetDatum[int32](int32(startup))
	values[5] = datum.ValueGetDatum[int32](int32(run))
	values[6] = datum.ValueGetDatum[int32](int32(total))
	return &Result{Slot: &row.EntrySlot{Values: values}}
}

func explainQual(qual expression.Expression) (string, error) {
	if qual == nil {
		return "", nil
	}
	return qual.ToString()
}
