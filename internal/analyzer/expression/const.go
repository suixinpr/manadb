package expression

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/relation/catalog"
	"github/suixinpr/manadb/internal/utils/fmngr"
)

/*
 * 常量, 表示一个具体的值.
 */
type Const struct {
	ConstType  metadata.OID
	ConstValue datum.Datum
}

func (con *Const) TypeID() metadata.OID {
	return con.ConstType
}

func (con *Const) Walker(walk Walker) {

}

func (con *Const) Mutator(mutate Mutator) Expression {
	result := *con
	return &result
}

func (con *Const) Evaluate(outer, inner []datum.Datum) (datum.Datum, error) {
	return con.ConstValue, nil
}

func (con *Const) ToString() (string, error) {
	meta, err := catalog.SearchCatalogOne(metadata.ManaTypesID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaTypesTypeID, metadata.OIDID, datum.ValueGetDatum(con.ConstType)),
	})
	if err != nil {
		return "", err
	}
	res, err := fmngr.Call(datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaTypesTypeOutput]), con.ConstValue)
	if err != nil {
		return "", err
	}
	return "'" + datum.DatumGetString(res) + "'", nil
}
