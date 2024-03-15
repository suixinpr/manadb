package adt

import (
	"fmt"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/utils/fmngr"
	"strconv"
)

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncFloat64Input, Float64Input)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat64Output, Float64Output)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat64LT, Float64LT)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat64EQ, Float64EQ)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat64GT, Float64GT)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat64LE, Float64LE)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat64NE, Float64NE)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat64GE, Float64GE)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat64Add, Float64Add)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat64Sub, Float64Sub)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat64Mul, Float64Mul)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat64Div, Float64Div)
}

/*
 * Float64Input - converts "num" to float64
 *
 * args:    Bytes
 * results: Float64
 */
func Float64Input(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetString(finfo.GetArg(0))
	res, err := (strconv.ParseFloat(arg, 64))
	return datum.ValueGetDatum(res), err
}

/*
 * Float64Output - converts float64 to "num"
 *
 * args:    Float64
 * results: Bytes
 */
func Float64Output(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValue[float64](finfo.GetArg(0))
	res := fmt.Appendf(nil, "%f", arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Float64 operators
*
********************************************************************************/

/* float64 < float64 */
func Float64LT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* float64 == float64 */
func Float64EQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* float64 > float64 */
func Float64GT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* float64 <= float64 */
func Float64LE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* float64 <> float64 */
func Float64NE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* float64 >= float64 */
func Float64GE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* float64 + float64 */
func Float64Add(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 + arg1), nil
}

/* float64 - float64 */
func Float64Sub(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 - arg1), nil
}

/* float64 * float64 */
func Float64Mul(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 * arg1), nil
}

/* float64 / float64 */
func Float64Div(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float64](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float64](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 / arg1), nil
}
