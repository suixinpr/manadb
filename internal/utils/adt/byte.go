package adt

import (
	"fmt"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/utils/fmngr"
	"strconv"
)

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncByteInput, ByteInput)
	fmngr.InitBuiltinFuncs(metadata.FuncByteOutput, ByteOutput)
	fmngr.InitBuiltinFuncs(metadata.FuncByteLT, ByteLT)
	fmngr.InitBuiltinFuncs(metadata.FuncByteEQ, ByteEQ)
	fmngr.InitBuiltinFuncs(metadata.FuncByteGT, ByteGT)
	fmngr.InitBuiltinFuncs(metadata.FuncByteLE, ByteLE)
	fmngr.InitBuiltinFuncs(metadata.FuncByteNE, ByteNE)
	fmngr.InitBuiltinFuncs(metadata.FuncByteGE, ByteGE)
	fmngr.InitBuiltinFuncs(metadata.FuncByteAdd, ByteAdd)
	fmngr.InitBuiltinFuncs(metadata.FuncByteSub, ByteSub)
	fmngr.InitBuiltinFuncs(metadata.FuncByteMul, ByteMul)
	fmngr.InitBuiltinFuncs(metadata.FuncByteDiv, ByteDiv)
}

/*
 * ByteInput - converts "num" to byte
 *
 * args:    Bytes
 * results: Byte
 */
func ByteInput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetString(finfo.GetArg(0))
	res, err := strconv.ParseUint(arg, 10, 8)
	return datum.ValueGetDatum(byte(res)), err
}

/*
 * ByteOutput - converts byte to "num"
 *
 * args:    Byte
 * results: Bytes
 */
func ByteOutput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValue[byte](finfo.GetArg(0))
	res := fmt.Appendf(nil, "%d", arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Byte operators
*
********************************************************************************/

/* byte < byte */
func ByteLT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[byte](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[byte](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* byte == byte */
func ByteEQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[byte](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[byte](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* byte > byte */
func ByteGT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[byte](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[byte](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* byte <= byte */
func ByteLE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[byte](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[byte](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* byte <> byte */
func ByteNE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[byte](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[byte](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* byte >= byte */
func ByteGE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[byte](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[byte](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* byte + byte */
func ByteAdd(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[byte](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[byte](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 + arg1), nil
}

/* byte - byte */
func ByteSub(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[byte](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[byte](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 - arg1), nil
}

/* byte * byte */
func ByteMul(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[byte](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[byte](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 * arg1), nil
}

/* byte / byte */
func ByteDiv(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[byte](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[byte](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 / arg1), nil
}
