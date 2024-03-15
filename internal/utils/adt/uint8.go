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
*  Uint8 init function
*
********************************************************************************/

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncUint8Input, Uint8Input)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8Output, Uint8Output)

	fmngr.InitBuiltinFuncs(metadata.FuncUint8LT, Uint8LT)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8EQ, Uint8EQ)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8GT, Uint8GT)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8LE, Uint8LE)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8NE, Uint8NE)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8GE, Uint8GE)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8Add, Uint8Add)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8Sub, Uint8Sub)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8Mul, Uint8Mul)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8Div, Uint8Div)

	fmngr.InitBuiltinFuncs(metadata.FuncUint8ToInt8, Uint8ToInt8)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8ToInt16, Uint8ToInt16)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8ToInt32, Uint8ToInt32)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8ToInt64, Uint8ToInt64)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8ToOID, Uint8ToOID)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8ToUint16, Uint8ToUint16)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8ToUint32, Uint8ToUint32)
	fmngr.InitBuiltinFuncs(metadata.FuncUint8ToUint64, Uint8ToUint64)
}

/********************************************************************************
*
*  Uint8 Input and Ouput
*
********************************************************************************/

/*
 * Uint8Input - converts "num" to uint8
 *
 * args:    Bytes
 * results: Uint8
 */
func Uint8Input(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetString(finfo.GetArg(0))
	res, err := strconv.ParseUint(arg, 10, 8)
	return datum.ValueGetDatum(uint8(res)), err
}

/*
 * Uint8Output - converts uint8 to "num"
 *
 * args:    Uint8
 * results: Bytes
 */
func Uint8Output(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValue[uint8](finfo.GetArg(0))
	res := fmt.Appendf(nil, "%d", arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Uint8 operators
*
********************************************************************************/

/* uint8 < uint8 */
func Uint8LT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* uint8 == uint8 */
func Uint8EQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* uint8 > uint8 */
func Uint8GT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* uint8 <= uint8 */
func Uint8LE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* uint8 <> uint8 */
func Uint8NE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* uint8 >= uint8 */
func Uint8GE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* uint8 + uint8 */
func Uint8Add(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 + arg1), nil
}

/* uint8 - uint8 */
func Uint8Sub(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 - arg1), nil
}

/* uint8 * uint8 */
func Uint8Mul(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 * arg1), nil
}

/* uint8 / uint8 */
func Uint8Div(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 / arg1), nil
}

/********************************************************************************
*
*  Convert Uint8 to ...
*
********************************************************************************/

/* uint8 to int8 */
func Uint8ToInt8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	return datum.ValueGetDatum(int8(arg0)), nil
}

/* uint8 to int16 */
func Uint8ToInt16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	return datum.ValueGetDatum(int16(arg0)), nil
}

/* uint8 to int32 */
func Uint8ToInt32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	return datum.ValueGetDatum(int32(arg0)), nil
}

/* uint8 to int64 */
func Uint8ToInt64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	return datum.ValueGetDatum(int64(arg0)), nil
}

/* uint8 to oid */
func Uint8ToOID(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	return datum.ValueGetDatum(metadata.OID(arg0)), nil
}

/* uint8 to uint16 */
func Uint8ToUint16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	return datum.ValueGetDatum(uint16(arg0)), nil
}

/* uint8 to uint32 */
func Uint8ToUint32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	return datum.ValueGetDatum(uint32(arg0)), nil
}

/* uint8 to uint64 */
func Uint8ToUint64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint8](finfo.GetArg(0))
	return datum.ValueGetDatum(uint64(arg0)), nil
}
