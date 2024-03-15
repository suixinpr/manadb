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
*  Uint64 init function
*
********************************************************************************/

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncUint64Input, Uint64Input)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64Output, Uint64Output)

	fmngr.InitBuiltinFuncs(metadata.FuncUint64LT, Uint64LT)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64EQ, Uint64EQ)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64GT, Uint64GT)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64LE, Uint64LE)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64NE, Uint64NE)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64GE, Uint64GE)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64Add, Uint64Add)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64Sub, Uint64Sub)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64Mul, Uint64Mul)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64Div, Uint64Div)

	fmngr.InitBuiltinFuncs(metadata.FuncUint64ToInt8, Uint64ToInt8)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64ToInt16, Uint64ToInt16)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64ToInt32, Uint64ToInt32)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64ToInt64, Uint64ToInt64)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64ToOID, Uint64ToOID)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64ToUint8, Uint64ToUint8)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64ToUint16, Uint64ToUint16)
	fmngr.InitBuiltinFuncs(metadata.FuncUint64ToUint32, Uint64ToUint32)
}

/********************************************************************************
*
*  Uint64 Input and Ouput
*
********************************************************************************/

/*
 * Uint64Input - converts "num" to uint64
 *
 * args:    Bytes
 * results: Uint64
 */
func Uint64Input(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetString(finfo.GetArg(0))
	res, err := strconv.ParseUint(arg, 10, 64)
	return datum.ValueGetDatum(res), err
}

/*
 * Uint64Output - converts uint64 to "num"
 *
 * args:    Uint64
 * results: Bytes
 */
func Uint64Output(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValue[uint64](finfo.GetArg(0))
	res := fmt.Appendf(nil, "%d", arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Uint64 operators
*
********************************************************************************/

/* uint64 < uint64 */
func Uint64LT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* uint64 == uint64 */
func Uint64EQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* uint64 > uint64 */
func Uint64GT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* uint64 <= uint64 */
func Uint64LE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* uint64 <> uint64 */
func Uint64NE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* uint64 >= uint64 */
func Uint64GE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* uint64 + uint64 */
func Uint64Add(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 + arg1), nil
}

/* uint64 - uint64 */
func Uint64Sub(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 - arg1), nil
}

/* uint64 * uint64 */
func Uint64Mul(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 * arg1), nil
}

/* uint64 / uint64 */
func Uint64Div(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[uint64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 / arg1), nil
}

/********************************************************************************
*
*  Convert Uint64 to ...
*
********************************************************************************/

/* uint64 to int8 */
func Uint64ToInt8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	return datum.ValueGetDatum(int8(arg0)), nil
}

/* uint64 to int16 */
func Uint64ToInt16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	return datum.ValueGetDatum(int16(arg0)), nil
}

/* uint64 to int32 */
func Uint64ToInt32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	return datum.ValueGetDatum(int32(arg0)), nil
}

/* uint64 to int64 */
func Uint64ToInt64(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	return datum.ValueGetDatum(int64(arg0)), nil
}

/* uint64 to oid */
func Uint64ToOID(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	return datum.ValueGetDatum(metadata.OID(arg0)), nil
}

/* uint64 to uint8 */
func Uint64ToUint8(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	return datum.ValueGetDatum(uint8(arg0)), nil
}

/* uint64 to uint16 */
func Uint64ToUint16(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	return datum.ValueGetDatum(uint16(arg0)), nil
}

/* uint64 to uint32 */
func Uint64ToUint32(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[uint64](finfo.GetArg(0))
	return datum.ValueGetDatum(uint32(arg0)), nil
}
