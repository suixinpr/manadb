package main

import (
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/storage/page"
	"github/suixinpr/manadb/internal/storage/smngr"
	"sort"
)

var catalogs = make(map[metadata.OID][]*metadata.ManaColumns)

func init() {
	for _, col := range metadata.InitManaColumns() {
		catalogs[col.ColTblID] = append(catalogs[col.ColTblID], col)
	}
	for _, cols := range catalogs {
		sort.Slice(cols, func(i, j int) bool {
			return cols[i].ColNo < cols[j].ColNo
		})
	}
}

/* 初始化系统表 */
func initMetadata() {
	initManaCast()
	initManaColumns()
	initManaFuncs()
	initManaOperators()
	initManaTables()
	initManaTypes()
}

func initManaCast() {
	var es []page.Entry
	metas := metadata.InitManaCast()
	for _, meta := range metas {
		values := make([]datum.Datum, metadata.NumManaCast)
		values[metadata.ManaCastCastID] = datum.ValueGetDatum(meta.CastID)
		values[metadata.ManaCastCastIn] = datum.ValueGetDatum(meta.CastIn)
		values[metadata.ManaCastCastOut] = datum.ValueGetDatum(meta.CastOut)
		values[metadata.ManaCastCastFunc] = datum.ValueGetDatum(meta.CastFunc)
		values[metadata.ManaCastCastType] = datum.ValueGetDatum(meta.CastType)
		es = append(es, page.NewEntry(catalogs[metadata.ManaCastID], values))
	}
	ps := entrysGetPages(es)

	smgr := smngr.OpenStorageManager(metadata.ManaCastID)
	smgr.Create()
	for i, p := range ps {
		smgr.Write(p, page.PageNumber(i))
	}
	smgr.Close()
}

func initManaColumns() {
	var es []page.Entry
	metas := metadata.InitManaColumns()
	for _, meta := range metas {
		values := make([]datum.Datum, metadata.NumManaColumns)
		values[metadata.ManaColumnsColID] = datum.ValueGetDatum(meta.ColID)
		values[metadata.ManaColumnsColTblID] = datum.ValueGetDatum(meta.ColTblID)
		values[metadata.ManaColumnsColName] = datum.StringGetDatum(meta.ColName)
		values[metadata.ManaColumnsColNo] = datum.ValueGetDatum(meta.ColNo)
		values[metadata.ManaColumnsColMod] = datum.ValueGetDatum(meta.ColMod)
		values[metadata.ManaColumnsColTypeID] = datum.ValueGetDatum(meta.ColTypeID)
		values[metadata.ManaColumnsColTypeLen] = datum.ValueGetDatum(meta.ColTypeLen)
		values[metadata.ManaColumnsColTypeAlign] = datum.ValueGetDatum(meta.ColTypeAlign)
		values[metadata.ManaColumnsColNotNULL] = datum.ValueGetDatum(meta.ColNotNULL)
		es = append(es, page.NewEntry(catalogs[metadata.ManaColumnsID], values))
	}
	ps := entrysGetPages(es)

	smgr := smngr.OpenStorageManager(metadata.ManaColumnsID)
	smgr.Create()
	for i, p := range ps {
		smgr.Write(p, page.PageNumber(i))
	}
	smgr.Close()
}

func initManaFuncs() {
	var es []page.Entry
	metas := metadata.InitManaFuncs()
	for _, meta := range metas {
		values := make([]datum.Datum, metadata.NumManaFuncs)
		values[metadata.ManaFuncsFuncID] = datum.ValueGetDatum(meta.FuncID)
		values[metadata.ManaFuncsFuncName] = datum.StringGetDatum(meta.FuncName)
		values[metadata.ManaFuncsFuncArgs] = datum.ValueGetDatum(meta.FuncArgs)
		values[metadata.ManaFuncsFuncArgTypes] = datum.ValueArrayGetDatum(meta.FuncArgTypes)
		values[metadata.ManaFuncsFuncRetType] = datum.ValueGetDatum(meta.FuncRetType)
		es = append(es, page.NewEntry(catalogs[metadata.ManaFuncsID], values))
	}
	ps := entrysGetPages(es)

	smgr := smngr.OpenStorageManager(metadata.ManaFuncsID)
	smgr.Create()
	for i, p := range ps {
		smgr.Write(p, page.PageNumber(i))
	}
	smgr.Close()
}

func initManaOperators() {
	var es []page.Entry
	metas := metadata.InitManaOperators()
	for _, meta := range metas {
		values := make([]datum.Datum, metadata.NumManaOperators)
		values[metadata.ManaOperatorsOprID] = datum.ValueGetDatum(meta.OprID)
		values[metadata.ManaOperatorsOprName] = datum.StringGetDatum(meta.OprName)
		values[metadata.ManaOperatorsOprLeft] = datum.ValueGetDatum(meta.OprLeft)
		values[metadata.ManaOperatorsOprRight] = datum.ValueGetDatum(meta.OprRight)
		values[metadata.ManaOperatorsOprResult] = datum.ValueGetDatum(meta.OprResult)
		values[metadata.ManaOperatorsOprFuncID] = datum.ValueGetDatum(meta.OprFuncID)
		es = append(es, page.NewEntry(catalogs[metadata.ManaOperatorsID], values))
	}
	ps := entrysGetPages(es)

	smgr := smngr.OpenStorageManager(metadata.ManaOperatorsID)
	smgr.Create()
	for i, p := range ps {
		smgr.Write(p, page.PageNumber(i))
	}
	smgr.Close()
}

func initManaTables() {
	var es []page.Entry
	metas := metadata.InitManaTables()
	for _, meta := range metas {
		values := make([]datum.Datum, metadata.NumManaTables)
		values[metadata.ManaTablesTblID] = datum.ValueGetDatum(meta.TblID)
		values[metadata.ManaTablesTblName] = datum.StringGetDatum(meta.TblName)
		values[metadata.ManaTablesTblRowsNum] = datum.ValueGetDatum(meta.TblRowsNum)
		es = append(es, page.NewEntry(catalogs[metadata.ManaTablesID], values))
	}
	ps := entrysGetPages(es)

	smgr := smngr.OpenStorageManager(metadata.ManaTablesID)
	smgr.Create()
	for i, p := range ps {
		smgr.Write(p, page.PageNumber(i))
	}
	smgr.Close()
}

func initManaTypes() {
	var es []page.Entry
	metas := metadata.InitManaTypes()
	for _, meta := range metas {
		values := make([]datum.Datum, metadata.NumManaTypes)
		values[metadata.ManaTypesTypeID] = datum.ValueGetDatum(meta.TypeID)
		values[metadata.ManaTypesTypeName] = datum.StringGetDatum(meta.TypeName)
		values[metadata.ManaTypesTypeLen] = datum.ValueGetDatum(meta.TypeLen)
		values[metadata.ManaTypesTypeAlign] = datum.ValueGetDatum(meta.TypeAlign)
		values[metadata.ManaTypesTypeInput] = datum.ValueGetDatum(meta.TypeInput)
		values[metadata.ManaTypesTypeOutput] = datum.ValueGetDatum(meta.TypeOutput)
		es = append(es, page.NewEntry(catalogs[metadata.ManaTypesID], values))
	}
	ps := entrysGetPages(es)

	smgr := smngr.OpenStorageManager(metadata.ManaTypesID)
	smgr.Create()
	for i, p := range ps {
		smgr.Write(p, page.PageNumber(i))
	}
	smgr.Close()
}

func entrysGetPages(es []page.Entry) []page.Page {
	var ps []page.Page

	p := page.NewPage()
	p.Init()
	ps = append(ps, p)
	for _, e := range es {
		if p.CanInsertEntry(e) {
			_ = p.InsertEntry(e)
			continue
		}

		p = page.NewPage()
		p.Init()
		ps = append(ps, p)
		_ = p.InsertEntry(e)
	}

	return ps
}
