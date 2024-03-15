package physical

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/relation/catalog"
)

/* 生成一个系统表中的标准列信息 */
func NewTemplateColumn(no int16, name string, typeID metadata.OID, mod int16) (*metadata.ManaColumns, error) {
	meta, err := catalog.SearchCatalogOne(metadata.ManaTypesID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaTypesTypeID, metadata.OIDID, datum.ValueGetDatum(typeID)),
	})
	if err != nil {
		return nil, err
	}
	return &metadata.ManaColumns{
		ColID:        metadata.InvalidID,
		ColTblID:     metadata.InvalidID,
		ColName:      name,
		ColNo:        no,
		ColMod:       mod,
		ColTypeID:    typeID,
		ColTypeLen:   datum.DatumGetValue[int16](meta.Values[metadata.ManaTypesTypeLen]),
		ColTypeAlign: datum.DatumGetValue[uint8](meta.Values[metadata.ManaTypesTypeAlign]),
		ColNotNULL:   false,
	}, nil
}

func inDomain(no int, domain []int) bool {
	for _, x := range domain {
		if no == x {
			return true
		}
	}
	return false
}
