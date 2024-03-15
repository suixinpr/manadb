package heapam_test

import (
	"fmt"
	"github/suixinpr/manadb/internal/access/heapam"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/storage/page"
	"github/suixinpr/manadb/internal/storage/smngr"
	"github/suixinpr/manadb/test"
	"testing"
)

func TestInsertEntry(t *testing.T) {
	err := test.InitDB()
	if err != nil {
		fmt.Println(err)
	}
	smngr.Init(test.DbDir())

	smgr := smngr.OpenStorageManager(metadata.ManaColumnsID)
	values := make([]datum.Datum, metadata.NumManaTables)
	values[0] = datum.ValueGetDatum(metadata.OID(10))
	values[1] = datum.StringGetDatum("test")
	e := page.NewEntry([]*metadata.ManaColumns{
		{
			ColID:        26,
			ColTblID:     metadata.ManaTablesID,
			ColName:      "tblid",
			ColNo:        metadata.ManaTablesTblID,
			ColMod:       -1,
			ColTypeID:    metadata.OIDID,
			ColTypeLen:   8,
			ColTypeAlign: 8,
			ColNotNULL:   true,
		},
		{
			ColID:        27,
			ColTblID:     metadata.ManaTablesID,
			ColName:      "tblname",
			ColNo:        metadata.ManaTablesTblName,
			ColMod:       -1,
			ColTypeID:    metadata.TextID,
			ColTypeLen:   -1,
			ColTypeAlign: 1,
			ColNotNULL:   true,
		},
	}, values)
	err = heapam.InsertEntry(smgr, e)
	t.Error(err)
}
