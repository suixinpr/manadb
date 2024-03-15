package adt

import (
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/utils/fmngr"
)

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncCharInput, CharInput)
	fmngr.InitBuiltinFuncs(metadata.FuncCharOutput, CharOutput)
	fmngr.InitBuiltinFuncs(metadata.FuncCharLT, CharLT)
	fmngr.InitBuiltinFuncs(metadata.FuncCharEQ, CharEQ)
	fmngr.InitBuiltinFuncs(metadata.FuncCharGT, CharGT)
	fmngr.InitBuiltinFuncs(metadata.FuncCharLE, CharLE)
	fmngr.InitBuiltinFuncs(metadata.FuncCharNE, CharNE)
	fmngr.InitBuiltinFuncs(metadata.FuncCharGE, CharGE)
	fmngr.InitBuiltinFuncs(metadata.FuncCharAdd, CharAdd)
}

/*
 * CharInput - converts Bytes to char(n)
 *
 * args:    Bytes, Int32
 * results: Char
 */
func CharInput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetBytes(finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int32](finfo.GetArg(1))
	res := make([]byte, arg1)
	copy(res, arg0)
	return datum.BytesGetDatum(res), nil
}

/*
 * CharOutput - converts char(n) to Bytes
 *
 * args:    Char
 * results: Bytes
 */
func CharOutput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetBytes(finfo.GetArg(0))
	res := make([]byte, len(arg))
	copy(res, arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Char operators
*
********************************************************************************/

/* varchar < varchar */
func CharLT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* varchar == varchar */
func CharEQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* varchar > varchar */
func CharGT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* varchar <= varchar */
func CharLE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* varchar <> varchar */
func CharNE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* varchar >= varchar */
func CharGE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* varchar + varchar */
func CharAdd(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.StringGetDatum(arg0 + arg1), nil
}
