package executor

import (
	"context"
	"github/suixinpr/manadb/internal/executor/command"
	"github/suixinpr/manadb/internal/executor/row"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/relation/table"
)

type DeleteExec struct {
	executor

	tbl *table.Table
	num int
}

func (exec *DeleteExec) run(ctx context.Context, output chan<- *Result, outer, inner <-chan *Result) error {
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

			err := command.DeleteTableEntry(exec.tbl, res.Slot.Pos)
			if err != nil {
				return err
			}
			exec.num++
		}
	}
}

func (exec *DeleteExec) message(num int) *Result {
	return &Result{Slot: &row.EntrySlot{Values: []datum.Datum{datum.ValueGetDatum[int32](int32(num))}}}
}
