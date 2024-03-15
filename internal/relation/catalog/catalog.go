package catalog

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/storage/smngr"
	_ "github/suixinpr/manadb/internal/utils/adt"
	"sort"
)

type Catalog struct {
	OID  metadata.OID
	Desc *common.EntryDesc
	Smgr *smngr.StorageManager
}

var catalogs = make(map[metadata.OID]*Catalog)

func init() {
	initCatalog(metadata.ManaCastID)
	initCatalog(metadata.ManaColumnsID)
	initCatalog(metadata.ManaFuncsID)
	initCatalog(metadata.ManaOperatorsID)
	initCatalog(metadata.ManaTablesID)
	initCatalog(metadata.ManaTypesID)
}

func initCatalog(oid metadata.OID) {
	catalogs[oid] = &Catalog{
		OID:  oid,
		Desc: initDesc(oid),
		Smgr: initSmgr(oid),
	}
}

func initDesc(oid metadata.OID) *common.EntryDesc {
	var cols []*metadata.ManaColumns
	for _, col := range metadata.InitManaColumns() {
		if col.ColTblID == oid {
			cols = append(cols, col)
		}
	}
	sort.Slice(cols, func(i, j int) bool {
		return cols[i].ColNo < cols[j].ColNo
	})
	return &common.EntryDesc{Cols: cols}
}

func initSmgr(oid metadata.OID) *smngr.StorageManager {
	return smngr.OpenStorageManager(oid)
}

func open(oid metadata.OID) (*Catalog, error) {
	cat := catalogs[oid]
	return cat, cat.Smgr.Open()
}
