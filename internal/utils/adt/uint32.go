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
*  Uint32 init function
*
********************************************************************************/

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncUint32Input, Uint32Input)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32Output, Uint32Output)

	fmngr.InitBuiltinFuncs(metadata.FuncUint32LT, Uint32LT)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32EQ, Uint32EQ)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32GT, Uint32GT)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32LE, Uint32LE)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32NE, Uint32NE)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32GE, Uint32GE)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32Add, Uint32Add)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32Sub, Uint32Sub)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32Mul, Uint32Mul)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32Div, Uint32Div)

	fmngr.InitBuiltinFuncs(metadata.FuncUint32ToInt8, Uint32ToInt8)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32ToInt16, Uint32ToInt16)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32ToInt32, Uint32ToInt32)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32ToInt64, Uint32ToInt64)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32ToOID, Uint32ToOID)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32ToUint8, Uint32ToUint8)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32ToUint16, Uint32ToUint16)
	fmngr.InitBuiltinFuncs(metadata.FuncUint32ToUint64, Uint32ToUint64)
}

/********************************************************************************
*
*  Uint32 Input and Ouput
*
********************************************************************************/

/*
 * Uint32Input - converts "num" to uint32
 *
 * args:    Bytes
 * results: Uint32
 */
func Uint32Input(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetString(finfo.GetArg(0))
	res, err := strconv.ParseUint(arg, 10, 32)
	return datum.ValueGetDatum(uint32(res)), err
}

/*
 * Uint32Output - converts uint32 to "num"
 *
 * args:    Uint32
 * results: Bytes
 */
func Uint32Output(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValue[uint32](finfo.GetArg(0))
	res := fmt.Appendf(nil, "%d", arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Uint32 operators
*
********************************************************************************/

/* uint32 < uint32 */
func Uint32LT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* uint32 == uint32 */
func Uint32EQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* uint32 > uint32 */
func Uint32GT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* uint32 <= uint32 */
func Uint32LE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* uint32 <> uint32 */
func Uint32NE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* uint32 >= uint32 */
func Uint32GE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* uint32 + uint32 */
func Uint32Add(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 + arg1), nil
}

/* uint32 - uint32 */
func Uint32Sub(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 - arg1), nil
}

/* uint32 * uint32 */
func Uint32Mul(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 * arg1), nil
}

/* uint32 / uint32 */
func Uint32Div(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 / arg1), nil
}

/********************************************************************************
*
*  Convert Uint32 to ...
*
********************************************************************************/

/* uint32 to int8 */
func Uint32ToInt8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	return datum.ValueGetDatum(int8(arg0)), nil
}

/* uint32 to int16 */
func Uint32ToInt16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	return datum.ValueGetDatum(int16(arg0)), nil
}

/* uint32 to int32 */
func Uint32ToInt32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	return datum.ValueGetDatum(int32(arg0)), nil
}

/* uint32 to int64 */
func Uint32ToInt64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	return datum.ValueGetDatum(int64(arg0)), nil
}

/* uint32 to oid */
func Uint32ToOID(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	return datum.ValueGetDatum(metadata.OID(arg0)), nil
}

/* uint32 to uint8 */
func Uint32ToUint8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	return datum.ValueGetDatum(uint8(arg0)), nil
}

/* uint32 to uint16 */
func Uint32ToUint16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	return datum.ValueGetDatum(uint16(arg0)), nil
}

/* uint32 to uint64 */
func Uint32ToUint64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint32](finfo.GetArg(0))
	return datum.ValueGetDatum(uint64(arg0)), nil
}
