package executor

import (
	"context"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/executor/row"
)

type ResultExec struct {
	executor

	project []expression.Expression
	qual    expression.Expression
}

func (exec *ResultExec) run(ctx context.Context, output chan<- *Result, outer, inner <-chan *Result) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		values, err := row.ExecProject(exec.project, nil, nil)
		if err != nil {
			return err
		}
		if exec.qual != nil {
			b, err := row.ExecQual(exec.qual, values, nil)
			if err != nil {
				return err
			}
			if !b {
				return nil
			}
		}
		output <- &Result{Slot: &row.EntrySlot{Values: values}}
		return nil
	}
}
