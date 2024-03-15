package adt

import (
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/utils/fmngr"
)

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncTextInput, TextInput)
	fmngr.InitBuiltinFuncs(metadata.FuncTextOutput, TextOutput)
	fmngr.InitBuiltinFuncs(metadata.FuncTextLT, TextLT)
	fmngr.InitBuiltinFuncs(metadata.FuncTextEQ, TextEQ)
	fmngr.InitBuiltinFuncs(metadata.FuncTextGT, TextGT)
	fmngr.InitBuiltinFuncs(metadata.FuncTextLE, TextLE)
	fmngr.InitBuiltinFuncs(metadata.FuncTextNE, TextNE)
	fmngr.InitBuiltinFuncs(metadata.FuncTextGE, TextGE)
	fmngr.InitBuiltinFuncs(metadata.FuncTextAdd, TextAdd)
}

/*
 * TextInput - converts Bytes to text
 *
 * args:    Bytes, Int32
 * results: Text
 */
func TextInput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetBytes(finfo.GetArg(0))
	res := make([]byte, len(arg))
	copy(res, arg)
	return datum.BytesGetDatum(res), nil
}

/*
 * TextOutput - converts text to Bytes
 *
 * args:    Text
 * results: Bytes
 */
func TextOutput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetBytes(finfo.GetArg(0))
	res := make([]byte, len(arg))
	copy(res, arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Text operators
*
********************************************************************************/

/* text < text */
func TextLT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* text == text */
func TextEQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* text > text */
func TextGT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* text <= text */
func TextLE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* text <> text */
func TextNE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* text >= text */
func TextGE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* text + text */
func TextAdd(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetString(finfo.GetArg(0))
	arg1 := datum.DatumGetString(finfo.GetArg(1))
	return datum.StringGetDatum(arg0 + arg1), nil
}
