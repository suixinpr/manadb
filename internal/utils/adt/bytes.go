package adt

import (
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/utils/fmngr"
)

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncBytesInput, BytesInput)
	fmngr.InitBuiltinFuncs(metadata.FuncBytesOutput, BytesOutput)
}

/*
 * BytesInput - converts Bytes to Bytes
 *
 * args:    Bytes
 * results: Bytes
 */
func BytesInput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetBytes(finfo.GetArg(0))
	res := make([]byte, len(arg))
	copy(res, arg)
	return datum.BytesGetDatum(res), nil
}

/*
 * BytesOutput - converts Bytes to Bytes
 *
 * args:    Bytes
 * results: Bytes
 */
func BytesOutput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetBytes(finfo.GetArg(0))
	res := make([]byte, len(arg))
	copy(res, arg)
	return datum.BytesGetDatum(res), nil
}
