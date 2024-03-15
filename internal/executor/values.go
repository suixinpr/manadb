package executor

import (
	"context"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/executor/row"
)

type ValuesExec struct {
	executor

	valuesList [][]expression.Expression
}

func (exec *ValuesExec) run(ctx context.Context, output chan<- *Result, outer, inner <-chan *Result) error {
	for _, project := range exec.valuesList {
		select {
		case <-ctx.Done():
			return nil
		default:
			values, err := row.ExecProject(project, nil, nil)
			if err != nil {
				return err
			}
			output <- &Result{Slot: &row.EntrySlot{Values: values}}
		}
	}
	return nil
}
