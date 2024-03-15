package adt

import (
	"fmt"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/utils/fmngr"
	"strconv"
	"strings"
)

func init() {
	fmngr.InitBuiltinFuncs(metadata.FuncOIDArrayInput, OIDArrayInput)
	fmngr.InitBuiltinFuncs(metadata.FuncOIDArrayOutput, OIDArrayOutput)
}

/*
 * OIDArrayInput - converts "num num ..." to []metadata.OID
 *
 * args:    Bytes
 * results: OIDArray
 */
func OIDArrayInput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetString(finfo.GetArg(0))
	res := make([]metadata.OID, 0, 64)
	for _, s := range strings.Split(arg, " ") {
		oid, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return nil, err
		}
		res = append(res, metadata.OID(oid))
	}
	return datum.ValueArrayGetDatum(res), nil
}

/*
 * OIDArrayOutput - converts OIDArray to "num num ..."
 *
 * args:    OIDArray
 * results: Bytes
 */
func OIDArrayOutput(finfo *fmngr.FuncInfo) (datum.Datum, error) {
	arg := datum.DatumGetValueArray[metadata.OID](finfo.GetArg(0))
	res := make([]byte, 0, len(arg)*2)
	for _, oid := range arg {
		res = fmt.Appendf(res, "%d ", oid)
	}
	res = res[:len(res)-1]
	return datum.BytesGetDatum(res), nil
}
