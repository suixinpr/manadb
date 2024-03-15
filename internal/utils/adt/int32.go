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
*  Int32 init function
*
********************************************************************************/

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncInt32Input, Int32Input)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32Output, Int32Output)

	fmngr.InitBuiltinFuncs(metadata.FuncInt32LT, Int32LT)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32EQ, Int32EQ)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32GT, Int32GT)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32LE, Int32LE)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32NE, Int32NE)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32GE, Int32GE)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32Add, Int32Add)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32Sub, Int32Sub)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32Mul, Int32Mul)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32Div, Int32Div)

	fmngr.InitBuiltinFuncs(metadata.FuncInt32ToInt8, Int32ToInt8)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32ToInt16, Int32ToInt16)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32ToInt64, Int32ToInt64)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32ToOID, Int32ToOID)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32ToUint8, Int32ToUint8)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32ToUint16, Int32ToUint16)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32ToUint32, Int32ToUint32)
	fmngr.InitBuiltinFuncs(metadata.FuncInt32ToUint64, Int32ToUint64)
}

/********************************************************************************
*
*  Int32 Input and Ouput
*
********************************************************************************/

/*
 * Int32Input - converts "num" to int32
 *
 * args:    Bytes
 * results: Int32
 */
func Int32Input(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetString(finfo.GetArg(0))
	res, err := strconv.ParseInt(arg, 10, 32)
	return datum.ValueGetDatum(int32(res)), err
}

/*
 * Int32Output - converts int32 to "num"
 *
 * args:    Int32
 * results: Bytes
 */
func Int32Output(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValue[int32](finfo.GetArg(0))
	res := fmt.Appendf(nil, "%d", arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Int32 operators
*
********************************************************************************/

/* int32 < int32 */
func Int32LT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* int32 == int32 */
func Int32EQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* int32 > int32 */
func Int32GT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* int32 <= int32 */
func Int32LE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* int32 <> int32 */
func Int32NE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* int32 >= int32 */
func Int32GE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* int32 + int32 */
func Int32Add(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 + arg1), nil
}

/* int32 - int32 */
func Int32Sub(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 - arg1), nil
}

/* int32 * int32 */
func Int32Mul(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 * arg1), nil
}

/* int32 / int32 */
func Int32Div(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 / arg1), nil
}

/********************************************************************************
*
*  Convert Int32 to ...
*
********************************************************************************/

/* int32 to int8 */
func Int32ToInt8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	return datum.ValueGetDatum(int8(arg0)), nil
}

/* int32 to int16 */
func Int32ToInt16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	return datum.ValueGetDatum(int16(arg0)), nil
}

/* int32 to int64 */
func Int32ToInt64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	return datum.ValueGetDatum(int64(arg0)), nil
}

/* int32 to oid */
func Int32ToOID(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	return datum.ValueGetDatum(metadata.OID(arg0)), nil
}

/* int32 to uint8 */
func Int32ToUint8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	return datum.ValueGetDatum(uint8(arg0)), nil
}

/* int32 to uint16 */
func Int32ToUint16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	return datum.ValueGetDatum(uint16(arg0)), nil
}

/* int32 to uint32 */
func Int32ToUint32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	return datum.ValueGetDatum(uint32(arg0)), nil
}

/* int32 to uint64 */
func Int32ToUint64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int32](finfo.GetArg(0))
	return datum.ValueGetDatum(uint64(arg0)), nil
}
