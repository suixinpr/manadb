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
*  Int64 init function
*
********************************************************************************/

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncInt64Input, Int64Input)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64Output, Int64Output)

	fmngr.InitBuiltinFuncs(metadata.FuncInt64LT, Int64LT)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64EQ, Int64EQ)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64GT, Int64GT)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64LE, Int64LE)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64NE, Int64NE)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64GE, Int64GE)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64Add, Int64Add)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64Sub, Int64Sub)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64Mul, Int64Mul)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64Div, Int64Div)

	fmngr.InitBuiltinFuncs(metadata.FuncInt64ToInt8, Int64ToInt8)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64ToInt16, Int64ToInt16)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64ToInt32, Int64ToInt32)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64ToOID, Int64ToOID)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64ToUint8, Int64ToUint8)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64ToUint16, Int64ToUint16)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64ToUint32, Int64ToUint32)
	fmngr.InitBuiltinFuncs(metadata.FuncInt64ToUint64, Int64ToUint64)
}

/********************************************************************************
*
*  Int64 Input and Ouput
*
********************************************************************************/

/*
 * Int64Input - converts "num" to int64
 *
 * args:    Bytes
 * results: Int64
 */
func Int64Input(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetString(finfo.GetArg(0))
	res, err := strconv.ParseInt(arg, 10, 64)
	return datum.ValueGetDatum(res), err
}

/*
 * Int64Output - converts int64 to "num"
 *
 * args:    Int64
 * results: Bytes
 */
func Int64Output(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValue[int64](finfo.GetArg(0))
	res := fmt.Appendf(nil, "%d", arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Int64 operators
*
********************************************************************************/

/* int64 < int64 */
func Int64LT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* int64 == int64 */
func Int64EQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* int64 > int64 */
func Int64GT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* int64 <= int64 */
func Int64LE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* int64 <> int64 */
func Int64NE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* int64 >= int64 */
func Int64GE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* int64 + int64 */
func Int64Add(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 + arg1), nil
}

/* int64 - int64 */
func Int64Sub(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 - arg1), nil
}

/* int64 * int64 */
func Int64Mul(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 * arg1), nil
}

/* int64 / int64 */
func Int64Div(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 / arg1), nil
}

/********************************************************************************
*
*  Convert Int64 to ...
*
********************************************************************************/

/* int64 to int8 */
func Int64ToInt8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	return datum.ValueGetDatum(int8(arg0)), nil
}

/* int64 to int16 */
func Int64ToInt16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	return datum.ValueGetDatum(int16(arg0)), nil
}

/* int64 to int32 */
func Int64ToInt32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	return datum.ValueGetDatum(int32(arg0)), nil
}

/* int64 to oid */
func Int64ToOID(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	return datum.ValueGetDatum(metadata.OID(arg0)), nil
}

/* int64 to uint8 */
func Int64ToUint8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	return datum.ValueGetDatum(uint8(arg0)), nil
}

/* int64 to uint16 */
func Int64ToUint16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	return datum.ValueGetDatum(uint16(arg0)), nil
}

/* int64 to uint32 */
func Int64ToUint32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	return datum.ValueGetDatum(uint32(arg0)), nil
}

/* int64 to uint64 */
func Int64ToUint64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int64](finfo.GetArg(0))
	return datum.ValueGetDatum(uint64(arg0)), nil
}
