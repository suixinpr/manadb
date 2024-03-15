package table

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/access/heapam"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/relation/catalog"
	"github/suixinpr/manadb/internal/storage/page"
	"github/suixinpr/manadb/internal/storage/smngr"
)

type Table struct {
	TableID metadata.OID
	Desc    *common.EntryDesc

	smgr *smngr.StorageManager
}

func Create(tblID metadata.OID) error {
	smgr := smngr.OpenStorageManager(tblID)
	err := smgr.Create()
	if err != nil {
		return err
	}
	return smgr.Close()
}

func Drop(tblID metadata.OID) error {
	smgr := smngr.OpenStorageManager(tblID)
	return smgr.Remove()
}

func Open(tblID metadata.OID) (*Table, error) {
	smgr := smngr.OpenStorageManager(tblID)
	if err := smgr.Open(); err != nil {
		return nil, err
	}

	desc, err := OpenEntryDesc(tblID)
	if err != nil {
		return nil, err
	}

	tbl := &Table{TableID: tblID, Desc: desc, smgr: smgr}
	return tbl, nil
}

func (tbl *Table) Close() error {
	tbl.smgr.Close()
	tbl.smgr = nil
	tbl.Desc = nil
	return nil
}

func (tbl *Table) InsertEntry(values []datum.Datum) error {
	e := page.NewEntry(tbl.Desc.Cols, values)
	return heapam.InsertEntry(tbl.smgr, e)
}

func (tbl *Table) DeleteEntry(pos *page.EntryPos) error {
	return heapam.DeleteEntry(tbl.smgr, pos)
}

func (tbl *Table) UpdataEntry(pos *page.EntryPos, values []datum.Datum) error {
	e := page.NewEntry(tbl.Desc.Cols, values)
	return heapam.UpdataEntry(tbl.smgr, pos, e)
}

func NewScanner(tbl *Table) *heapam.Scanner {
	return heapam.NewScanner(tbl.TableID, tbl.Desc, tbl.smgr, nil)
}

func OpenEntryDesc(tblID metadata.OID) (*common.EntryDesc, error) {
	metas, err := catalog.SearchCatalogList(metadata.ManaColumnsID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaColumnsColTblID, metadata.OIDID, datum.ValueGetDatum(tblID)),
	})
	if err != nil {
		return nil, err
	}
	desc := &common.EntryDesc{}
	for _, meta := range metas {
		desc.Cols = append(desc.Cols, &metadata.ManaColumns{
			ColID:        datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaColumnsColID]),
			ColTblID:     datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaColumnsColTblID]),
			ColName:      datum.DatumGetString(meta.Values[metadata.ManaColumnsColName]),
			ColNo:        datum.DatumGetValue[int16](meta.Values[metadata.ManaColumnsColNo]),
			ColMod:       datum.DatumGetValue[int16](meta.Values[metadata.ManaColumnsColMod]),
			ColTypeID:    datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaColumnsColTypeID]),
			ColTypeLen:   datum.DatumGetValue[int16](meta.Values[metadata.ManaColumnsColTypeLen]),
			ColTypeAlign: datum.DatumGetValue[uint8](meta.Values[metadata.ManaColumnsColTypeAlign]),
			ColNotNULL:   datum.DatumGetValue[bool](meta.Values[metadata.ManaColumnsColNotNULL]),
		})
	}
	return desc, nil
}
