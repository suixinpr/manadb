package adt

import (
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/utils/fmngr"
)

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncVarcharInput, VarcharInput)
	fmngr.InitBuiltinFuncs(metadata.FuncVarcharOutput, VarcharOutput)
	fmngr.InitBuiltinFuncs(metadata.FuncVarcharLT, VarcharLT)
	fmngr.InitBuiltinFuncs(metadata.FuncVarcharEQ, VarcharEQ)
	fmngr.InitBuiltinFuncs(metadata.FuncVarcharGT, VarcharGT)
	fmngr.InitBuiltinFuncs(metadata.FuncVarcharLE, VarcharLE)
	fmngr.InitBuiltinFuncs(metadata.FuncVarcharNE, VarcharNE)
	fmngr.InitBuiltinFuncs(metadata.FuncVarcharGE, VarcharGE)
	fmngr.InitBuiltinFuncs(metadata.FuncVarcharAdd, VarcharAdd)
}

/*
 * VarcharInput - converts Bytes to varchar(n)
 *
 * args:    Bytes, Int32
 * results: Varchar
 */
func VarcharInput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetBytes(finfo.GetArg(0))
	arg1 := datum.DatumGetValue[int32](finfo.GetArg(1))
	res := make([]byte, min(int(arg1), len(arg0)))
	copy(res, arg0)
	return datum.BytesGetDatum(res), nil
}

/*
 * VarcharOutput - converts varchar(n) to Bytes
 *
 * args:    Varchar
 * results: Bytes
 */
func VarcharOutput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetBytes(finfo.GetArg(0))
	res := make([]byte, len(arg))
	copy(res, arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Varchar operators
*
********************************************************************************/

/* varchar < varchar */
func VarcharLT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* varchar == varchar */
func VarcharEQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* varchar > varchar */
func VarcharGT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* varchar <= varchar */
func VarcharLE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* varchar <> varchar */
func VarcharNE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* varchar >= varchar */
func VarcharGE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* varchar + varchar */
func VarcharAdd(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.StringGetDatum(arg0 + arg1), nil
}
