/*
 * 对 entry 的操作, 相比较于 table 包中的操作, 这里的
 * 处理额外考虑了相应操作对系统表需要进行的处理.
 */
package command

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/relation/catalog"
	"github/suixinpr/manadb/internal/relation/table"
	"github/suixinpr/manadb/internal/storage/page"
)

/*
 * 当向系统表插入 entry 时, 需要知道当前 entry 的 oid 应该
 * 是多少, 这个函数可以返回应该赋值的 oid 的值.
 */
func registerCatalogEntryOid(oid metadata.OID) (metadata.OID, error) {
	meta, err := catalog.SearchCatalogOne(metadata.ManaTablesID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaTablesTblID, metadata.OIDID, datum.ValueGetDatum(oid)),
	})
	if err != nil {
		return metadata.InvalidID, err
	}
	res := datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaTablesTblRowsNum])
	return res + 1, nil
}

func InsertTableEntry(tbl *table.Table, values []datum.Datum) error {
	/* 打开 mana_tables 系统表 */
	manaTables, err := table.Open(metadata.ManaTablesID)
	if err != nil {
		return err
	}
	defer manaTables.Close()

	/* 更新 mana_tables 系统表 */
	meta, err := catalog.SearchCatalogOne(metadata.ManaTablesID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaTablesTblID, metadata.OIDID, datum.ValueGetDatum(tbl.TableID)),
	})
	if err != nil {
		return err
	}

	/* tblrowsnum 加 1 */
	err = manaTables.UpdataEntry(meta.Pos, []datum.Datum{
		meta.Values[metadata.ManaTablesTblID],
		meta.Values[metadata.ManaTablesTblName],
		datum.ValueGetDatum(datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaTablesTblRowsNum]) + 1),
	})
	if err != nil {
		return err
	}

	/* 插入 entry */
	return tbl.InsertEntry(values)
}

func UpdateTableEntry(tbl *table.Table, pos *page.EntryPos, values []datum.Datum) error {
	/* do nothing in advance */

	/* 更新 entry */
	return tbl.UpdataEntry(pos, values)
}

func DeleteTableEntry(tbl *table.Table, pos *page.EntryPos) error {
	/* 打开 mana_tables 系统表 */
	manaTables, err := table.Open(metadata.ManaTablesID)
	if err != nil {
		return err
	}
	defer manaTables.Close()

	/* 更新 mana_tables 系统表 */
	meta, err := catalog.SearchCatalogOne(metadata.ManaTablesID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaTablesTblID, metadata.OIDID, datum.ValueGetDatum(tbl.TableID)),
	})
	if err != nil {
		return err
	}

	/* tblrowsnum 减 1 */
	err = manaTables.UpdataEntry(meta.Pos, []datum.Datum{
		meta.Values[metadata.ManaTablesTblID],
		meta.Values[metadata.ManaTablesTblName],
		datum.ValueGetDatum(datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaTablesTblRowsNum]) - 1),
	})
	if err != nil {
		return err
	}

	/* 删除 entry */
	return tbl.DeleteEntry(pos)
}
