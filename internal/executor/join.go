package executor

import (
	"context"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/executor/row"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/parser/types"
)

type JoinExec struct {
	executor

	project []expression.Expression
	tp      types.JoinType
	on      expression.Expression
	qual    expression.Expression
}

func (exec *JoinExec) run(ctx context.Context, output chan<- *Result, outer, inner <-chan *Result) error {
	type StoragedResult struct {
		slot *row.EntrySlot
		used bool
	}
	var outerResults, innerResults []*StoragedResult

	open1, open2 := true, true
	for open1 || open2 {
		select {
		case <-ctx.Done():
			return nil
		case res, ok := <-outer:
			if !ok {
				open1 = false
				continue
			}
			if res.Err != nil {
				return res.Err
			}
			outerResults = append(outerResults, &StoragedResult{slot: res.Slot})
		case res, ok := <-inner:
			if !ok {
				open2 = false
				continue
			}
			if res.Err != nil {
				return res.Err
			}
			innerResults = append(innerResults, &StoragedResult{slot: res.Slot})
		}
	}

	for _, outerResult := range outerResults {
		outerSlot := outerResult.slot
		for _, innerResult := range innerResults {
			innerSlot := innerResult.slot
			if exec.on != nil {
				b, err := row.ExecQual(exec.on, outerSlot.Values, innerSlot.Values)
				if err != nil {
					return err
				}
				if !b {
					continue
				}
			}
			if exec.qual != nil {
				b, err := row.ExecQual(exec.qual, outerSlot.Values, innerSlot.Values)
				if err != nil {
					return err
				}
				if !b {
					continue
				}
			}
			values, err := row.ExecProject(exec.project, outerSlot.Values, innerSlot.Values)
			if err != nil {
				return err
			}
			outerResult.used, innerResult.used = true, true
			output <- &Result{Slot: &row.EntrySlot{Values: values}}
		}
	}

	if exec.tp == types.LeftJoin || exec.tp == types.FullJoin {
		NullValues := make([]datum.Datum, len(exec.inner.base().desc.Cols))
		for _, outerResult := range outerResults {
			if outerResult.used {
				continue
			}
			outerSlot := outerResult.slot
			if exec.qual != nil {
				b, err := row.ExecQual(exec.qual, outerSlot.Values, NullValues)
				if err != nil {
					return err
				}
				if !b {
					continue
				}
			}
			values, err := row.ExecProject(exec.project, outerSlot.Values, NullValues)
			if err != nil {
				return err
			}
			output <- &Result{Slot: &row.EntrySlot{Values: values}}
		}
	}

	if exec.tp == types.RightJoin || exec.tp == types.FullJoin {
		NullValues := make([]datum.Datum, len(exec.inner.base().desc.Cols))
		for _, innerResult := range innerResults {
			if innerResult.used {
				continue
			}
			innerSlot := innerResult.slot
			if exec.qual != nil {
				b, err := row.ExecQual(exec.qual, NullValues, innerSlot.Values)
				if err != nil {
					return err
				}
				if !b {
					continue
				}
			}
			values, err := row.ExecProject(exec.project, NullValues, innerSlot.Values)
			if err != nil {
				return err
			}
			output <- &Result{Slot: &row.EntrySlot{Values: values}}
		}
	}

	return nil
}
