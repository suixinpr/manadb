package adt

import (
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/utils/fmngr"
)

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncBooleanInput, BooleanInput)
	fmngr.InitBuiltinFuncs(metadata.FuncBooleanOutput, BooleanOutput)
	fmngr.InitBuiltinFuncs(metadata.FuncBooleanEQ, BooleanEQ)
	fmngr.InitBuiltinFuncs(metadata.FuncBooleanNE, BooleanNE)
	fmngr.InitBuiltinFuncs(metadata.FuncBooleanAnd, BooleanAnd)
	fmngr.InitBuiltinFuncs(metadata.FuncBooleanOr, BooleanOr)
	fmngr.InitBuiltinFuncs(metadata.FuncBooleanNot, BooleanNot)
}

/*
 * BooleanInput - converts "t" or "f" to true or false
 *
 * args:    Bytes
 * results: Boolean
 */
func BooleanInput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetBytes(finfo.GetArg(0))
	if len(arg) != 1 {
		return nil, errlog.New("")
	}

	res := false
	if arg[0] == 't' {
		res = true
	}

	return datum.ValueGetDatum(res), nil
}

/*
 * BooleanOutput - converts true or false to "t" or "f"
 *
 * args:    Boolean
 * results: Bytes
 */
func BooleanOutput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValue[bool](finfo.GetArg(0))
	result := make([]byte, 1)
	if arg {
		result[0] = 't'
	} else {
		result[0] = 'f'
	}
	return datum.BytesGetDatum(result), nil
}

/********************************************************************************
*
*  Boolean operators
*
********************************************************************************/

/* bool == bool */
func BooleanEQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[bool](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[bool](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* bool <> bool */
func BooleanNE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[bool](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[bool](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* bool and bool */
func BooleanAnd(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[bool](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[bool](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 && arg1), nil
}

/* bool or bool */
func BooleanOr(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[bool](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[bool](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 || arg1), nil
}

/* not bool */
func BooleanNot(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[bool](finfo.GetArg(0))
	return datum.ValueGetDatum(!arg0), nil
}
