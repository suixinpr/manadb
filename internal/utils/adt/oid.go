package adt

import (
	"fmt"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/utils/fmngr"
	"strconv"
)

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncOIDInput, OIDInput)
	fmngr.InitBuiltinFuncs(metadata.FuncOIDOutput, OIDOutput)
	fmngr.InitBuiltinFuncs(metadata.FuncOIDLT, OIDLT)
	fmngr.InitBuiltinFuncs(metadata.FuncOIDEQ, OIDEQ)
	fmngr.InitBuiltinFuncs(metadata.FuncOIDGT, OIDGT)
	fmngr.InitBuiltinFuncs(metadata.FuncOIDLE, OIDLE)
	fmngr.InitBuiltinFuncs(metadata.FuncOIDNE, OIDNE)
	fmngr.InitBuiltinFuncs(metadata.FuncOIDGE, OIDGE)
	fmngr.InitBuiltinFuncs(metadata.FuncOIDAdd, OIDAdd)
	fmngr.InitBuiltinFuncs(metadata.FuncOIDSub, OIDSub)
	fmngr.InitBuiltinFuncs(metadata.FuncOIDMul, OIDMul)
	fmngr.InitBuiltinFuncs(metadata.FuncOIDDiv, OIDDiv)
}

/*
 * OIDInput - converts "num" to OID
 *
 * args:    Bytes
 * results: OID
 */
func OIDInput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetString(finfo.GetArg(0))
	res, err := strconv.ParseUint(arg, 10, 64)
	return datum.ValueGetDatum(metadata.OID(res)), err
}

/*
 * OIDOutput - converts OID to "num"
 *
 * args:    OID
 * results: Bytes
 */
func OIDOutput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValue[metadata.OID](finfo.GetArg(0))
	res := fmt.Appendf(nil, "%d", arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  OID operators
*
********************************************************************************/

/* oid < oid */
func OIDLT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[metadata.OID](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[metadata.OID](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* oid == oid */
func OIDEQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[metadata.OID](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[metadata.OID](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* oid > oid */
func OIDGT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[metadata.OID](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[metadata.OID](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* oid <= oid */
func OIDLE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[metadata.OID](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[metadata.OID](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* oid <> oid */
func OIDNE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[metadata.OID](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[metadata.OID](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* oid >= oid */
func OIDGE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[metadata.OID](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[metadata.OID](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* oid + oid */
func OIDAdd(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[metadata.OID](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[metadata.OID](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 + arg1), nil
}

/* oid - oid */
func OIDSub(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[metadata.OID](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[metadata.OID](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 - arg1), nil
}

/* oid * oid */
func OIDMul(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[metadata.OID](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[metadata.OID](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 * arg1), nil
}

/* oid / oid */
func OIDDiv(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[metadata.OID](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[metadata.OID](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 / arg1), nil
}
