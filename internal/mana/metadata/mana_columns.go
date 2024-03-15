package metadata

type ManaColumns struct {
	ColID        OID    /* mana_columns: colid */
	ColTblID     OID    /* mana_columns: coltblid */
	ColName      string /* mana_columns: colname */
	ColNo        int16  /* mana_columns: colno */
	ColMod       int16  /* mana_columns: colmod */
	ColTypeID    OID    /* mana_columns: coltypeid */
	ColTypeLen   int16  /* mana_columns: coltypelen */
	ColTypeAlign uint8  /* mana_columns: coltypealign */
	ColNotNULL   bool   /* mana_columns: colnotnull */
}

const (
	ManaColumnsColID int16 = iota
	ManaColumnsColTblID
	ManaColumnsColName
	ManaColumnsColNo
	ManaColumnsColMod
	ManaColumnsColTypeID
	ManaColumnsColTypeLen
	ManaColumnsColTypeAlign
	ManaColumnsColNotNULL
	NumManaColumns /* number of mana_columns columns */
)

func InitManaColumns() (metas []*ManaColumns) {
	defer func() {
		for oid, meta := range metas {
			meta.ColID = OID(oid) + 1
		}
	}()

	return []*ManaColumns{
		/* mana_cast */
		{InvalidID, ManaCastID, "castid", ManaCastCastID, -1, OIDID, 8, 8, true},
		{InvalidID, ManaCastID, "castin", ManaCastCastIn, -1, OIDID, 8, 8, true},
		{InvalidID, ManaCastID, "castout", ManaCastCastOut, -1, OIDID, 8, 8, true},
		{InvalidID, ManaCastID, "castfunc", ManaCastCastFunc, -1, OIDID, 8, 8, true},
		{InvalidID, ManaCastID, "casttype", ManaCastCastType, -1, OIDID, 8, 8, true},

		/* mana_columns */
		{InvalidID, ManaColumnsID, "colid", ManaColumnsColID, -1, OIDID, 8, 8, true},
		{InvalidID, ManaColumnsID, "coltblid", ManaColumnsColTblID, -1, OIDID, 8, 8, true},
		{InvalidID, ManaColumnsID, "colname", ManaColumnsColName, -1, TextID, -1, 1, true},
		{InvalidID, ManaColumnsID, "colno", ManaColumnsColNo, -1, Int16ID, 2, 2, true},
		{InvalidID, ManaColumnsID, "colmod", ManaColumnsColMod, -1, Int16ID, 2, 2, true},
		{InvalidID, ManaColumnsID, "coltypeid", ManaColumnsColTypeID, -1, OIDID, 8, 8, true},
		{InvalidID, ManaColumnsID, "coltypelen", ManaColumnsColTypeLen, -1, Int16ID, 2, 2, true},
		{InvalidID, ManaColumnsID, "coltypealign", ManaColumnsColTypeAlign, -1, Uint8ID, 1, 1, true},
		{InvalidID, ManaColumnsID, "colnotnull", ManaColumnsColNotNULL, -1, BooleanID, 1, 1, true},

		/* mana_funcs */
		{InvalidID, ManaFuncsID, "funcid", ManaFuncsFuncID, -1, OIDID, 8, 8, true},
		{InvalidID, ManaFuncsID, "funcname", ManaFuncsFuncName, -1, TextID, -1, 1, true},
		{InvalidID, ManaFuncsID, "funcargs", ManaFuncsFuncArgs, -1, Int8ID, 1, 1, true},
		{InvalidID, ManaFuncsID, "funcargtypes", ManaFuncsFuncArgTypes, -1, OIDArrayID, -1, 8, true},
		{InvalidID, ManaFuncsID, "funcrettype", ManaFuncsFuncRetType, -1, OIDID, 8, 8, true},

		/* mana_operators */
		{InvalidID, ManaOperatorsID, "oprid", ManaOperatorsOprID, -1, OIDID, 8, 8, true},
		{InvalidID, ManaOperatorsID, "oprname", ManaOperatorsOprName, -1, TextID, -1, 1, true},
		{InvalidID, ManaOperatorsID, "oprleft", ManaOperatorsOprLeft, -1, OIDID, 8, 8, true},
		{InvalidID, ManaOperatorsID, "oprright", ManaOperatorsOprRight, -1, OIDID, 8, 8, true},
		{InvalidID, ManaOperatorsID, "oprresult", ManaOperatorsOprResult, -1, OIDID, 8, 8, true},
		{InvalidID, ManaOperatorsID, "oprfuncid", ManaOperatorsOprFuncID, -1, OIDID, 8, 8, true},

		/* mana_tables */
		{InvalidID, ManaTablesID, "tblid", ManaTablesTblID, -1, OIDID, 8, 8, true},
		{InvalidID, ManaTablesID, "tblname", ManaTablesTblName, -1, TextID, -1, 1, true},
		{InvalidID, ManaTablesID, "tblrowsnum", ManaTablesTblRowsNum, -1, Uint64ID, -1, 1, true},

		/* mana_types */
		{InvalidID, ManaTypesID, "typeid", ManaTypesTypeID, -1, OIDID, 8, 8, true},
		{InvalidID, ManaTypesID, "typename", ManaTypesTypeName, -1, TextID, -1, 1, true},
		{InvalidID, ManaTypesID, "typelen", ManaTypesTypeLen, -1, Int16ID, 4, 4, true},
		{InvalidID, ManaTypesID, "typealign", ManaTypesTypeAlign, -1, Uint8ID, 1, 1, true},
		{InvalidID, ManaTypesID, "typeinput", ManaTypesTypeInput, -1, OIDID, 8, 8, true},
		{InvalidID, ManaTypesID, "typeoutput", ManaTypesTypeOutput, -1, Uint8ID, 1, 1, true},
	}
}

func init() {
	metas := InitManaColumns()
	h := make(map[OID]struct{})
	for _, meta := range metas {
		if _, ok := h[meta.ColID]; ok {
			panic("duplicate metadata exists in system tables: mana_columns")
		}
		if meta.ColID == InvalidID {
			panic("Invalid OID metadata exists in system tables: mana_columns")
		}
	}
}
