package expression

import (
	"fmt"
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/relation/catalog"
	"github/suixinpr/manadb/internal/utils/fmngr"
)

type Cast struct {
	CastID   metadata.OID
	CastIn   Expression
	CastOut  metadata.OID
	CastFunc metadata.OID
	CastType byte
}

func (c *Cast) TypeID() metadata.OID {
	return c.CastOut
}

func (c *Cast) Walker(walk Walker) {
	walk(c.CastIn)
}

func (c *Cast) Mutator(mutate Mutator) Expression {
	result := *c
	result.CastIn = mutate(c.CastIn)
	return &result
}

func (c *Cast) Evaluate(outer, inner []datum.Datum) (datum.Datum, error) {
	in, err := c.CastIn.Evaluate(outer, inner)
	if err != nil {
		return nil, err
	}
	return fmngr.Call(c.CastFunc, in)
}

func (c *Cast) ToString() (string, error) {
	in, err := c.CastIn.ToString()
	if err != nil {
		return "", err
	}

	meta, err := catalog.SearchCatalogOne(metadata.ManaTypesID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaTypesTypeID, metadata.OIDID, datum.ValueGetDatum(c.CastOut)),
	})
	if err != nil {
		return "", err
	}
	out := datum.DatumGetString(meta.Values[metadata.ManaTypesTypeName])

	return fmt.Sprintf("CAST(%s AS %s)", in, out), nil
}
