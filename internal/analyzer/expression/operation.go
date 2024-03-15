package expression

import (
	"fmt"
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/relation/catalog"
	"github/suixinpr/manadb/internal/utils/fmngr"
)

/*
 * 操作符, 表示一个简单运算.
 */
type Operation struct {
	OprID     metadata.OID
	OprLeft   Expression
	OprRight  Expression
	OprResult metadata.OID
	OprFuncID metadata.OID
}

func (op *Operation) TypeID() metadata.OID {
	return op.OprResult
}

func (op *Operation) Walker(walk Walker) {
	walk(op.OprLeft)
	walk(op.OprRight)
}

func (op *Operation) Mutator(mutate Mutator) Expression {
	result := *op
	result.OprLeft = mutate(op.OprLeft)
	result.OprRight = mutate(op.OprRight)
	return &result
}

func (op *Operation) Evaluate(outer, inner []datum.Datum) (datum.Datum, error) {
	left, err := op.OprLeft.Evaluate(outer, inner)
	if err != nil {
		return nil, err
	}
	right, err := op.OprRight.Evaluate(outer, inner)
	if err != nil {
		return nil, err
	}
	return fmngr.Call(op.OprFuncID, left, right)
}

func (op *Operation) ToString() (string, error) {
	left, err := op.OprLeft.ToString()
	if err != nil {
		return "", err
	}

	right, err := op.OprRight.ToString()
	if err != nil {
		return "", err
	}

	meta, err := catalog.SearchCatalogOne(metadata.ManaOperatorsID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaOperatorsOprID, metadata.OIDID, datum.ValueGetDatum(op.OprID)),
	})
	if err != nil {
		return "", err
	}
	name := datum.DatumGetString(meta.Values[metadata.ManaOperatorsOprName])

	return fmt.Sprintf("%s %s %s", left, name, right), nil
}
