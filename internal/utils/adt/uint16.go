package adt

import (
	"fmt"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/utils/fmngr"
	"strconv"
)

/********************************************************************************
*
*  Uint16 init function
*
********************************************************************************/

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncUint16Input, Uint16Input)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16Output, Uint16Output)

	fmngr.InitBuiltinFuncs(metadata.FuncUint16LT, Uint16LT)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16EQ, Uint16EQ)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16GT, Uint16GT)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16LE, Uint16LE)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16NE, Uint16NE)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16GE, Uint16GE)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16Add, Uint16Add)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16Sub, Uint16Sub)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16Mul, Uint16Mul)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16Div, Uint16Div)

	fmngr.InitBuiltinFuncs(metadata.FuncUint16ToInt8, Uint16ToInt8)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16ToInt16, Uint16ToInt16)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16ToInt32, Uint16ToInt32)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16ToInt64, Uint16ToInt64)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16ToOID, Uint16ToOID)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16ToUint8, Uint16ToUint8)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16ToUint32, Uint16ToUint32)
	fmngr.InitBuiltinFuncs(metadata.FuncUint16ToUint64, Uint16ToUint64)
}

/********************************************************************************
*
*  Uint16 Input and Ouput
*
********************************************************************************/

/*
 * Uint16Input - converts "num" to uint16
 *
 * args:    Bytes
 * results: Uint16
 */
func Uint16Input(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetString(finfo.GetArg(0))
	res, err := strconv.ParseUint(arg, 10, 16)
	return datum.ValueGetDatum(uint16(res)), err
}

/*
 * Uint16Output - converts uint16 to "num"
 *
 * args:    Uint16
 * results: Bytes
 */
func Uint16Output(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValue[uint16](finfo.GetArg(0))
	res := fmt.Appendf(nil, "%d", arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Uint16 operators
*
********************************************************************************/

/* uint16 < uint16 */
func Uint16LT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* uint16 == uint16 */
func Uint16EQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* uint16 > uint16 */
func Uint16GT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* uint16 <= uint16 */
func Uint16LE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* uint16 <> uint16 */
func Uint16NE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* uint16 >= uint16 */
func Uint16GE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* uint16 + uint16 */
func Uint16Add(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 + arg1), nil
}

/* uint16 - uint16 */
func Uint16Sub(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 - arg1), nil
}

/* uint16 * uint16 */
func Uint16Mul(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 * arg1), nil
}

/* uint16 / uint16 */
func Uint16Div(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 / arg1), nil
}

/********************************************************************************
*
*  Convert Uint16 to ...
*
********************************************************************************/

/* uint16 to int8 */
func Uint16ToInt8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	return datum.ValueGetDatum(int8(arg0)), nil
}

/* uint16 to int16 */
func Uint16ToInt16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	return datum.ValueGetDatum(int16(arg0)), nil
}

/* uint16 to int32 */
func Uint16ToInt32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	return datum.ValueGetDatum(int32(arg0)), nil
}

/* uint16 to int64 */
func Uint16ToInt64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	return datum.ValueGetDatum(int64(arg0)), nil
}

/* uint16 to oid */
func Uint16ToOID(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	return datum.ValueGetDatum(metadata.OID(arg0)), nil
}

/* uint16 to uint8 */
func Uint16ToUint8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	return datum.ValueGetDatum(uint8(arg0)), nil
}

/* uint16 to uint32 */
func Uint16ToUint32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	return datum.ValueGetDatum(uint32(arg0)), nil
}

/* uint16 to uint64 */
func Uint16ToUint64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint16](finfo.GetArg(0))
	return datum.ValueGetDatum(uint64(arg0)), nil
}
