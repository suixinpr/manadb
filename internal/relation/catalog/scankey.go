package catalog

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
)

/* type comparison equality function */
var eqFuncs = make(map[metadata.OID]metadata.OID)

func init() {
	eqFuncs[metadata.BooleanID] = metadata.FuncBooleanEQ
	eqFuncs[metadata.ByteID] = metadata.FuncByteEQ
	eqFuncs[metadata.CharID] = metadata.FuncCharEQ
	eqFuncs[metadata.Float32ID] = metadata.FuncFloat32EQ
	eqFuncs[metadata.Float64ID] = metadata.FuncFloat64EQ
	eqFuncs[metadata.Int8ID] = metadata.FuncInt8EQ
	eqFuncs[metadata.Int16ID] = metadata.FuncInt16EQ
	eqFuncs[metadata.Int32ID] = metadata.FuncInt32EQ
	eqFuncs[metadata.Int64ID] = metadata.FuncInt64EQ
	eqFuncs[metadata.OIDID] = metadata.FuncOIDEQ
	eqFuncs[metadata.TextID] = metadata.FuncTextEQ
	eqFuncs[metadata.Uint8ID] = metadata.FuncUint8EQ
	eqFuncs[metadata.Uint16ID] = metadata.FuncUint16EQ
	eqFuncs[metadata.Uint32ID] = metadata.FuncUint32EQ
	eqFuncs[metadata.Uint64ID] = metadata.FuncUint64EQ
}

func NewScanKey(colNo int16, typeID metadata.OID, argument datum.Datum) *common.ScanKey {
	return &common.ScanKey{
		ColNo:     colNo,
		Strategy:  common.SKEqualStrategy,
		Argument:  argument,
		CmpFuncID: eqFuncs[typeID],
	}
}
