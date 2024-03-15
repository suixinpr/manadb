package executor

import (
	"context"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/executor/row"
	"github/suixinpr/manadb/internal/mana/datum"
)

type LimitExec struct {
	executor

	project []expression.Expression
	count   expression.Expression
	offset  expression.Expression
}

func (exec *LimitExec) run(ctx context.Context, output chan<- *Result, outer, inner <-chan *Result) error {
	var count, offset uint64

	if exec.count != nil {
		res, err := exec.count.Evaluate(nil, nil)
		if err != nil {
			return err
		}
		count = datum.DatumGetValue[uint64](res)
	}

	if exec.offset != nil {
		res, err := exec.offset.Evaluate(nil, nil)
		if err != nil {
			return err
		}
		offset = datum.DatumGetValue[uint64](res)
	}

	/* offset */
	for i := uint64(0); i < offset; i++ {
		select {
		case <-ctx.Done():
			return nil
		case res, ok := <-outer:
			if !ok {
				return nil
			}

			if res.Err != nil {
				return res.Err
			}
		}
	}

	/* count */
	for i := uint64(0); i < count; i++ {
		select {
		case <-ctx.Done():
			return nil
		case res, ok := <-outer:
			if !ok {
				return nil
			}
			if res.Err != nil {
				return res.Err
			}

			slot := res.Slot
			values, err := row.ExecProject(exec.project, slot.Values, nil)
			if err != nil {
				return err
			}
			output <- &Result{Slot: &row.EntrySlot{Values: values, Pos: slot.Pos}}
		}
	}
	return nil
}
