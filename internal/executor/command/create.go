package command

import (
	"fmt"
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/analyzer/logic"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/planner/physical"
	"github/suixinpr/manadb/internal/relation/catalog"
	"github/suixinpr/manadb/internal/relation/table"
)

func ExecCreateTablePlan(plan *logic.CreateTablePlan) error {
	return createTable(plan.TableName, plan.Columns)
}

/* 创建一个表 */
func createTable(name string, columns []*logic.ColumnInfo) error {
	/* 检查表是否已经存在 */
	meta, err := catalog.SearchCatalog(metadata.ManaTablesID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaTablesTblName, metadata.TextID, datum.StringGetDatum(name)),
	})
	if err != nil {
		return err
	}
	if meta != nil {
		return errlog.New(fmt.Sprintf("table %s already exists", name))
	}

	/* 处理 mana_tables */
	tblID, err := createTableInManaTables(name)
	if err != nil {
		return err
	}

	/* 处理 mana_columns */
	err = createTableInManaColumns(tblID, columns)
	if err != nil {
		return err
	}

	return table.Create(tblID)
}

func createTableInManaTables(name string) (metadata.OID, error) {
	/* 获取表 oid */
	oid, err := registerCatalogEntryOid(metadata.ManaTablesID)
	if err != nil {
		return metadata.InvalidID, err
	}

	/* mana_tables */
	manaTables, err := table.Open(metadata.ManaTablesID)
	if err != nil {
		return metadata.InvalidID, err
	}
	defer manaTables.Close()

	values := make([]datum.Datum, metadata.NumManaTables)
	values[0] = datum.ValueGetDatum(oid)
	values[1] = datum.StringGetDatum(name)
	values[2] = datum.ValueGetDatum[uint64](0)

	err = InsertTableEntry(manaTables, values)
	if err != nil {
		return metadata.InvalidID, err
	}
	return oid, nil
}

func createTableInManaColumns(tblID metadata.OID, columns []*logic.ColumnInfo) error {
	desc, err := buildEntryDesc(columns)
	if err != nil {
		return err
	}

	manaColumns, err := table.Open(metadata.ManaColumnsID)
	if err != nil {
		return err
	}
	defer manaColumns.Close()

	for i, col := range desc.Cols {
		oid, err := registerCatalogEntryOid(metadata.ManaColumnsID)
		if err != nil {
			return err
		}

		values := make([]datum.Datum, metadata.NumManaColumns)
		values[0] = datum.ValueGetDatum(oid)
		values[1] = datum.ValueGetDatum(tblID)
		values[2] = datum.StringGetDatum(col.ColName)
		values[3] = datum.ValueGetDatum(int16(i))
		values[4] = datum.ValueGetDatum(col.ColMod)
		values[5] = datum.ValueGetDatum(col.ColTypeID)
		values[6] = datum.ValueGetDatum(col.ColTypeLen)
		values[7] = datum.ValueGetDatum(col.ColTypeAlign)
		values[8] = datum.ValueGetDatum(col.ColNotNULL)
		err = InsertTableEntry(manaColumns, values)
		if err != nil {
			return err
		}
	}
	return nil
}

func buildEntryDesc(colDefs []*logic.ColumnInfo) (*common.EntryDesc, error) {
	desc := new(common.EntryDesc)
	desc.Cols = make([]*metadata.ManaColumns, len(colDefs))
	for i, colDef := range colDefs {
		col, err := physical.NewTemplateColumn(int16(i), colDef.Name, colDef.TypeID, colDef.Mod)
		if err != nil {
			return nil, err
		}
		desc.Cols[i] = col
	}
	return desc, nil
}
