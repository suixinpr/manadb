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
*  Int8 init function
*
********************************************************************************/

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncInt8Input, Int8Input)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8Output, Int8Output)

	fmngr.InitBuiltinFuncs(metadata.FuncInt8LT, Int8LT)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8EQ, Int8EQ)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8GT, Int8GT)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8LE, Int8LE)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8NE, Int8NE)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8GE, Int8GE)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8Add, Int8Add)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8Sub, Int8Sub)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8Mul, Int8Mul)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8Div, Int8Div)

	fmngr.InitBuiltinFuncs(metadata.FuncInt8ToInt16, Int8ToInt16)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8ToInt32, Int8ToInt32)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8ToInt64, Int8ToInt64)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8ToOID, Int8ToOID)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8ToUint8, Int8ToUint8)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8ToUint16, Int8ToUint16)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8ToUint32, Int8ToUint32)
	fmngr.InitBuiltinFuncs(metadata.FuncInt8ToUint64, Int8ToUint64)
}

/********************************************************************************
*
*  Int8 Input and Ouput
*
********************************************************************************/

/*
 * Int8Input - converts "num" to int8
 *
 * args:    Bytes
 * results: Int8
 */
func Int8Input(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetString(finfo.GetArg(0))
	res, err := strconv.ParseInt(arg, 10, 8)
	return datum.ValueGetDatum(int8(res)), err
}

/*
 * Int8Output - converts int8 to "num"
 *
 * args:    Int8
 * results: Bytes
 */
func Int8Output(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValue[int8](finfo.GetArg(0))
	res := fmt.Appendf(nil, "%d", arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Int8 operators
*
********************************************************************************/

/* int8 < int8 */
func Int8LT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* int8 == int8 */
func Int8EQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* int8 > int8 */
func Int8GT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* int8 <= int8 */
func Int8LE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* int8 <> int8 */
func Int8NE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* int8 >= int8 */
func Int8GE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* int8 + int8 */
func Int8Add(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 + arg1), nil
}

/* int8 - int8 */
func Int8Sub(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 - arg1), nil
}

/* int8 * int8 */
func Int8Mul(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 * arg1), nil
}

/* int8 / int8 */
func Int8Div(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int8](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 / arg1), nil
}

/********************************************************************************
*
*  Convert Int8 to ...
*
********************************************************************************/

/* int8 to int16 */
func Int8ToInt16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	return datum.ValueGetDatum(int16(arg0)), nil
}

/* int8 to int32 */
func Int8ToInt32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	return datum.ValueGetDatum(int32(arg0)), nil
}

/* int8 to int64 */
func Int8ToInt64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	return datum.ValueGetDatum(int64(arg0)), nil
}

/* int8 to oid */
func Int8ToOID(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	return datum.ValueGetDatum(metadata.OID(arg0)), nil
}

/* int8 to uint8 */
func Int8ToUint8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	return datum.ValueGetDatum(uint8(arg0)), nil
}

/* int8 to uint16 */
func Int8ToUint16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	return datum.ValueGetDatum(uint16(arg0)), nil
}

/* int8 to uint32 */
func Int8ToUint32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	return datum.ValueGetDatum(uint32(arg0)), nil
}

/* int8 to uint64 */
func Int8ToUint64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int8](finfo.GetArg(0))
	return datum.ValueGetDatum(uint64(arg0)), nil
}
