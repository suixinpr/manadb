package catalog

import (
	"bytes"
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/storage/smngr"
	"log"
	"os/user"
	"path/filepath"
	"testing"
)

func TestSearchCatalog(t *testing.T) {
	usr, err := user.Current()
	if err != nil {
		log.Panic(err)
	}
	smngr.Init(filepath.Join(usr.HomeDir, "manadata"))

	test := []struct {
		oid  metadata.OID
		keys []*common.ScanKey
		meta *Meta
	}{
		{
			oid: metadata.ManaOperatorsID,
			keys: []*common.ScanKey{
				NewScanKey(metadata.ManaOperatorsOprName, metadata.TextID, datum.StringGetDatum("+")),
				NewScanKey(metadata.ManaOperatorsOprLeft, metadata.OIDID, datum.ValueGetDatum(metadata.Int64ID)),
				NewScanKey(metadata.ManaOperatorsOprRight, metadata.OIDID, datum.ValueGetDatum(metadata.Int64ID)),
			},
			meta: &Meta{
				Values: []datum.Datum{
					datum.ValueGetDatum[metadata.OID](77),
					datum.StringGetDatum("+"),
					datum.ValueGetDatum[metadata.OID](metadata.Int64ID),
					datum.ValueGetDatum[metadata.OID](metadata.Int64ID),
					datum.ValueGetDatum[metadata.OID](metadata.Int64ID),
					datum.ValueGetDatum[metadata.OID](metadata.FuncInt64Add),
				},
				Pos: nil,
			},
		},
	}

	for no, tt := range test {
		meta, err := SearchCatalogOne(tt.oid, tt.keys)
		if err != nil {
			t.Errorf("(%v) SearchCatalogOne: err = %v", no, err)
		}

		if meta == nil {
			if tt.meta != nil {
				t.Errorf("(%v) SearchCatalogOne: meta is nil", no)
			}
			continue
		}

		num := len(meta.Values)
		for i := 0; i < num; i++ {
			if !bytes.Equal(tt.meta.Values[i], meta.Values[i]) {
				t.Errorf("(%v) SearchCatalogOne: meta[%v] expected %v, got = %v", no, i, tt.meta.Values[i], meta.Values[i])
			}
		}
	}
}
