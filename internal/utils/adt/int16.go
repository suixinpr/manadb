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
*  Int16 init function
*
********************************************************************************/

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncInt16Input, Int16Input)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16Output, Int16Output)

	fmngr.InitBuiltinFuncs(metadata.FuncInt16LT, Int16LT)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16EQ, Int16EQ)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16GT, Int16GT)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16LE, Int16LE)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16NE, Int16NE)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16GE, Int16GE)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16Add, Int16Add)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16Sub, Int16Sub)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16Mul, Int16Mul)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16Div, Int16Div)

	fmngr.InitBuiltinFuncs(metadata.FuncInt16ToInt8, Int16ToInt8)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16ToInt32, Int16ToInt32)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16ToInt64, Int16ToInt64)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16ToOID, Int16ToOID)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16ToUint8, Int16ToUint8)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16ToUint16, Int16ToUint16)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16ToUint32, Int16ToUint32)
	fmngr.InitBuiltinFuncs(metadata.FuncInt16ToUint64, Int16ToUint64)
}

/********************************************************************************
*
*  Int16 Input and Ouput
*
********************************************************************************/

/*
 * Int16Input - converts "num" to int16
 *
 * args:    Bytes
 * results: Int16
 */
func Int16Input(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetString(finfo.GetArg(0))
	res, err := strconv.ParseInt(arg, 10, 16)
	return datum.ValueGetDatum(int16(res)), err
}

/*
 * Int16Output - converts int16 to "num"
 *
 * args:    Int16
 * results: Bytes
 */
func Int16Output(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValue[int16](finfo.GetArg(0))
	res := fmt.Appendf(nil, "%d", arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Int16 operators
*
********************************************************************************/

/* int16 < int16 */
func Int16LT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* int16 == int16 */
func Int16EQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* int16 > int16 */
func Int16GT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* int16 <= int16 */
func Int16LE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* int16 <> int16 */
func Int16NE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* int16 >= int16 */
func Int16GE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* int16 + int16 */
func Int16Add(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 + arg1), nil
}

/* int16 - int16 */
func Int16Sub(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 - arg1), nil
}

/* int16 * int16 */
func Int16Mul(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 * arg1), nil
}

/* int16 / int16 */
func Int16Div(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int16](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 / arg1), nil
}

/********************************************************************************
*
*  Convert Int16 to ...
*
********************************************************************************/

/* int16 to int8 */
func Int16ToInt8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	return datum.ValueGetDatum(int16(arg0)), nil
}

/* int16 to int32 */
func Int16ToInt32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	return datum.ValueGetDatum(int32(arg0)), nil
}

/* int16 to int64 */
func Int16ToInt64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	return datum.ValueGetDatum(int64(arg0)), nil
}

/* int16 to oid */
func Int16ToOID(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	return datum.ValueGetDatum(metadata.OID(arg0)), nil
}

/* int16 to uint8 */
func Int16ToUint8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	return datum.ValueGetDatum(uint8(arg0)), nil
}

/* int16 to uint16 */
func Int16ToUint16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	return datum.ValueGetDatum(uint16(arg0)), nil
}

/* int16 to uint32 */
func Int16ToUint32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	return datum.ValueGetDatum(uint32(arg0)), nil
}

/* int16 to uint64 */
func Int16ToUint64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[int16](finfo.GetArg(0))
	return datum.ValueGetDatum(uint64(arg0)), nil
}
