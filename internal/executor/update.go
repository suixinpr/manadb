package executor

import (
	"context"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/executor/command"
	"github/suixinpr/manadb/internal/executor/row"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/relation/table"
)

type UpdateExec struct {
	executor

	tbl     *table.Table
	project []expression.Expression
	num     int
}

func (exec *UpdateExec) run(ctx context.Context, output chan<- *Result, outer, inner <-chan *Result) error {
	defer exec.tbl.Close()

	for {
		select {
		case <-ctx.Done():
			return nil
		case res, ok := <-outer:
			if !ok {
				output <- exec.message(exec.num)
				return nil
			}
			if res.Err != nil {
				return res.Err
			}

			values, err := row.ExecProject(exec.project, res.Slot.Values, nil)
			if err != nil {
				return err
			}

			err = command.UpdateTableEntry(exec.tbl, res.Slot.Pos, values)
			if err != nil {
				return err
			}
			exec.num++
		}
	}
}

func (exec *UpdateExec) message(num int) *Result {
	return &Result{Slot: &row.EntrySlot{Values: []datum.Datum{datum.ValueGetDatum[int32](int32(num))}}}
}
