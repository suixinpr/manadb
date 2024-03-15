package analyzer

import (
	"fmt"
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/relation/catalog"
)

func (ana *Analyzer) castToTargetType(expr expression.Expression, targetID metadata.OID) (expression.Expression, error) {
	if expr.TypeID() == targetID {
		return expr, nil
	}

	meta, err := catalog.SearchCatalog(metadata.ManaCastID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaCastCastIn, metadata.OIDID, datum.ValueGetDatum(expr.TypeID())),
		catalog.NewScanKey(metadata.ManaCastCastOut, metadata.OIDID, datum.ValueGetDatum(targetID)),
	})
	if err != nil {
		return nil, err
	}
	if meta == nil {
		in, err := expr.ToString()
		if err != nil {
			return nil, err
		}
		meta, err := catalog.SearchCatalogOne(metadata.ManaTypesID, []*common.ScanKey{
			catalog.NewScanKey(metadata.ManaTypesTypeID, metadata.OIDID, datum.ValueGetDatum(targetID)),
		})
		if err != nil {
			return nil, err
		}
		out := datum.DatumGetString(meta.Values[metadata.ManaTypesTypeName])
		return nil, errlog.New(fmt.Sprintf("cannot cast %s to %s", in, out))
	}

	return &expression.Cast{
		CastID:   datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaCastCastID]),
		CastIn:   expr,
		CastOut:  datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaCastCastOut]),
		CastFunc: datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaCastCastFunc]),
		CastType: datum.DatumGetValue[byte](meta.Values[metadata.ManaCastCastType]),
	}, nil
}

func (ana *Analyzer) castImplicitType(expr1, expr2 expression.Expression) (expression.Expression, expression.Expression, error) {
	typeID1 := expr1.TypeID()
	typeID2 := expr2.TypeID()

	if typeID1 == typeID2 {
		return expr1, expr2, nil
	}

	/* 尝试转换 expr1 -> expr2 */
	meta, err := catalog.SearchCatalog(metadata.ManaCastID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaCastCastIn, metadata.OIDID, datum.ValueGetDatum(typeID1)),
		catalog.NewScanKey(metadata.ManaCastCastOut, metadata.OIDID, datum.ValueGetDatum(typeID2)),
	})
	if err != nil {
		return nil, nil, err
	}
	if meta != nil {
		cast := &expression.Cast{
			CastID:   datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaCastCastID]),
			CastIn:   expr1,
			CastOut:  datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaCastCastOut]),
			CastFunc: datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaCastCastFunc]),
			CastType: datum.DatumGetValue[byte](meta.Values[metadata.ManaCastCastType]),
		}
		return cast, expr2, nil
	}

	/* 尝试转换 expr2 -> expr1 */
	meta, err = catalog.SearchCatalog(metadata.ManaCastID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaCastCastIn, metadata.OIDID, datum.ValueGetDatum(typeID2)),
		catalog.NewScanKey(metadata.ManaCastCastOut, metadata.OIDID, datum.ValueGetDatum(typeID1)),
	})
	if err != nil {
		return nil, nil, err
	}
	if meta != nil {
		cast := &expression.Cast{
			CastID:   datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaCastCastID]),
			CastIn:   expr2,
			CastOut:  datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaCastCastOut]),
			CastFunc: datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaCastCastFunc]),
			CastType: datum.DatumGetValue[byte](meta.Values[metadata.ManaCastCastType]),
		}
		return expr1, cast, nil
	}

	return nil, nil, errlog.New("cannot do implicit cast")
}
