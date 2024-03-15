package fmngr

import (
	"fmt"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/mana/metadata"
)

/* 所有的内置函数 */
var builtin = make(map[metadata.OID]FuncAddr)

/* 内置函数地址 */
type FuncAddr func(*FuncInfo) (datum.Datum, error)

func InitBuiltinFuncs(oid metadata.OID, addr FuncAddr) {
	builtin[oid] = addr
}

func Call(funcID metadata.OID, args ...datum.Datum) (datum.Datum, error) {
	addr, ok := builtin[funcID]
	if !ok {
		return nil, errlog.New(fmt.Sprintf("not found function %d", funcID))
	}

	/* 目前的所有函数都是严格函数 */
	for _, arg := range args {
		if datum.IsNull(arg) {
			return datum.NullDatum(), nil
		}
	}

	return InternalCall(addr, args...)
}

func InternalCall(addr FuncAddr, args ...datum.Datum) (datum.Datum, error) {
	finfo := &FuncInfo{args: args}
	return addr(finfo)
}

type FuncInfo struct {
	args []datum.Datum
}

func (finfo *FuncInfo) GetArg(n int) datum.Datum {
	return finfo.args[n]
}
