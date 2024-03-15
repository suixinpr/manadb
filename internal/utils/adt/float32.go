package adt

import (
	"fmt"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/utils/fmngr"
	"strconv"
)

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncFloat32Input, Float32Input)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat32Output, Float32Output)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat32LT, Float32LT)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat32EQ, Float32EQ)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat32GT, Float32GT)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat32LE, Float32LE)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat32NE, Float32NE)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat32GE, Float32GE)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat32Add, Float32Add)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat32Sub, Float32Sub)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat32Mul, Float32Mul)
	fmngr.InitBuiltinFuncs(metadata.FuncFloat32Div, Float32Div)
}

/*
 * Float32Input - converts "num" to float32
 *
 * args:    Bytes
 * results: Float32
 */
func Float32Input(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetString(finfo.GetArg(0))
	res, err := (strconv.ParseFloat(arg, 32))
	return datum.ValueGetDatum(float32(res)), err
}

/*
 * Float32Output - converts float32 to "num"
 *
 * args:    Float32
 * results: Bytes
 */
func Float32Output(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValue[float32](finfo.GetArg(0))
	res := fmt.Appendf(nil, "%f", arg)
	return datum.BytesGetDatum(res), nil
}

/********************************************************************************
*
*  Float32 operators
*
********************************************************************************/

/* float32 < float32 */
func Float32LT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 < arg1), nil
}

/* float32 == float32 */
func Float32EQ(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 == arg1), nil
}

/* float32 > float32 */
func Float32GT(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 > arg1), nil
}

/* float32 <= float32 */
func Float32LE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 <= arg1), nil
}

/* float32 <> float32 */
func Float32NE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 != arg1), nil
}

/* float32 >= float32 */
func Float32GE(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 >= arg1), nil
}

/* float32 + float32 */
func Float32Add(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 + arg1), nil
}

/* float32 - float32 */
func Float32Sub(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 - arg1), nil
}

/* float32 * float32 */
func Float32Mul(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 * arg1), nil
}

/* float32 / float32 */
func Float32Div(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg0 := datum.DatumGetValue[float32](finfo.GetArg(0))
	arg1 := datum.DatumGetValue[float32](finfo.GetArg(1))
	return datum.ValueGetDatum(arg0 / arg1), nil
}
