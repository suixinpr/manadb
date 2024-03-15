package executor

import (
	"context"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/executor/row"
	"github/suixinpr/manadb/internal/relation/table"
)

type SeqScanExec struct {
	executor

	project []expression.Expression
	tbl     *table.Table
	qual    expression.Expression
}

func (exec *SeqScanExec) run(ctx context.Context, output chan<- *Result, outer, inner <-chan *Result) error {
	defer exec.tbl.Close()

	sc := table.NewScanner(exec.tbl)
	if err := sc.BeginScan(false); err != nil {
		return err
	}
	defer sc.EndScan()

	for {
		he, err := sc.GetNextEntry()
		if err != nil {
			return err
		}
		if he == nil {
			return nil
		}

		values := he.Entry.GetAllColumnsValue(exec.tbl.Desc.Cols)
		if exec.qual != nil {
			b, err := row.ExecQual(exec.qual, values, nil)
			if err != nil {
				return err
			}
			if !b {
				continue
			}
		}

		values, err = row.ExecProject(exec.project, values, nil)
		if err != nil {
			return err
		}

		output <- &Result{Slot: &row.EntrySlot{Values: values, Pos: he.Pos}}
	}
}
