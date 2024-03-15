package executor

import (
	"context"
	"fmt"
	"github/suixinpr/manadb/internal/analyzer/logic"
	"github/suixinpr/manadb/internal/executor/command"
	"github/suixinpr/manadb/internal/executor/row"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/errlog"
)

type CommandExec struct {
	executor
	plan logic.LogicPlan
}

func (exec *CommandExec) run(ctx context.Context, output chan<- *Result, outer, inner <-chan *Result) error {
	switch plan := exec.plan.(type) {
	case *logic.CreateTablePlan:
		if err := command.ExecCreateTablePlan(plan); err != nil {
			return err
		}
	case *logic.DropTablePlan:
		if err := command.ExecDropTablePlan(plan); err != nil {
			return err
		}
	default:
		return errlog.New(fmt.Sprintf("not known %T plan", plan))
	}
	output <- exec.message("success")
	return nil
}

func (exec *CommandExec) message(s string) *Result {
	return &Result{Slot: &row.EntrySlot{Values: []datum.Datum{datum.StringGetDatum(s)}}}
}
