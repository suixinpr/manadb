package builder

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/planner/physical"
)

/* 收集 command 的描述符 */
func (b *PlanBuilder) collectCommandDesc() (err error) {
	b.desc = new(common.EntryDesc)
	b.desc.Cols = make([]*metadata.ManaColumns, 1)
	b.desc.Cols[0], err = physical.NewTemplateColumn(0, "result", metadata.TextID, -1)
	return
}

/* 收集 explain 的描述符 */
func (b *PlanBuilder) collectExplainDesc() error {
	var err error
	b.desc = new(common.EntryDesc)
	b.desc.Cols = make([]*metadata.ManaColumns, 7)

	b.desc.Cols[0], err = physical.NewTemplateColumn(0, "id", metadata.Int32ID, -1)
	if err != nil {
		return err
	}
	b.desc.Cols[1], err = physical.NewTemplateColumn(1, "node", metadata.Int32ID, -1)
	if err != nil {
		return err
	}
	b.desc.Cols[2], err = physical.NewTemplateColumn(2, "operation", metadata.TextID, -1)
	if err != nil {
		return err
	}
	b.desc.Cols[3], err = physical.NewTemplateColumn(3, "qual", metadata.TextID, -1)
	if err != nil {
		return err
	}
	b.desc.Cols[4], err = physical.NewTemplateColumn(4, "startup", metadata.Int32ID, -1)
	if err != nil {
		return err
	}
	b.desc.Cols[5], err = physical.NewTemplateColumn(5, "run", metadata.Int32ID, -1)
	if err != nil {
		return err
	}
	b.desc.Cols[6], err = physical.NewTemplateColumn(6, "total", metadata.Int32ID, -1)
	if err != nil {
		return err
	}
	return nil
}

func (b *PlanBuilder) collectInsertDesc() (err error) {
	b.desc = new(common.EntryDesc)
	b.desc.Cols = make([]*metadata.ManaColumns, 1)
	b.desc.Cols[0], err = physical.NewTemplateColumn(0, "num", metadata.Int32ID, -1)
	return
}

func (b *PlanBuilder) collectDeleteDesc() (err error) {
	b.desc = new(common.EntryDesc)
	b.desc.Cols = make([]*metadata.ManaColumns, 1)
	b.desc.Cols[0], err = physical.NewTemplateColumn(0, "num", metadata.Int32ID, -1)
	return
}

func (b *PlanBuilder) collectUpdateDesc() (err error) {
	b.desc = new(common.EntryDesc)
	b.desc.Cols = make([]*metadata.ManaColumns, 1)
	b.desc.Cols[0], err = physical.NewTemplateColumn(0, "num", metadata.Int32ID, -1)
	return
}
