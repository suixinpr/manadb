package command

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/analyzer/logic"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/relation/catalog"
	"github/suixinpr/manadb/internal/relation/table"
)

func ExecDropTablePlan(plan *logic.DropTablePlan) error {
	for _, oid := range plan.Tables {
		if err := dropTable(oid); err != nil {
			return err
		}
	}
	return nil
}

/* 删除表 */
func dropTable(tblID metadata.OID) error {
	/* 处理 mana_tables */
	err := dropTableInManaTables(tblID)
	if err != nil {
		return err
	}

	/* 处理 mana_columns */
	err = dropTableInManaColumns(tblID)
	if err != nil {
		return err
	}

	/* 删除物理文件 */
	return table.Drop(tblID)
}

/* 在 mana_tables 中删除表的信息 */
func dropTableInManaTables(tblID metadata.OID) error {
	manaTables, err := table.Open(metadata.ManaTablesID)
	if err != nil {
		return err
	}
	defer manaTables.Close()

	meta, err := catalog.SearchCatalogOne(metadata.ManaTablesID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaTablesTblID, metadata.OIDID, datum.ValueGetDatum(tblID)),
	})
	if err != nil {
		return err
	}

	return DeleteTableEntry(manaTables, meta.Pos)
}

/* 在 mana_columns 中删除表的信息 */
func dropTableInManaColumns(tblID metadata.OID) error {
	manaColumns, err := table.Open(metadata.ManaColumnsID)
	if err != nil {
		return err
	}
	defer manaColumns.Close()

	metas, err := catalog.SearchCatalogList(metadata.ManaColumnsID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaColumnsColTblID, metadata.OIDID, datum.ValueGetDatum(tblID)),
	})
	if err != nil {
		return err
	}

	for _, meta := range metas {
		err := DeleteTableEntry(manaColumns, meta.Pos)
		if err != nil {
			return err
		}
	}
	return nil
}
