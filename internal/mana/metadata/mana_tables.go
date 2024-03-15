package metadata

type ManaTables struct {
	TblID      OID    /* mana_tables: tblid */
	TblName    string /* mana_tables: tblname */
	TblRowsNum uint64 /* mana_tables: tblrowsnum */
}

const (
	ManaTablesTblID int16 = iota
	ManaTablesTblName
	ManaTablesTblRowsNum
	NumManaTables /* number of mana_tables columns */
)

func InitManaTables() []*ManaTables {
	return []*ManaTables{
		/* mana_cast */
		{ManaCastID, "mana_cast", uint64(len(InitManaCast()))},

		/* mana_columns */
		{ManaColumnsID, "mana_columns", uint64(len(InitManaColumns()))},

		/* mana_funcs */
		{ManaFuncsID, "mana_funcs", uint64(numFuncID) - 1},

		/* mana_operators */
		{ManaOperatorsID, "mana_operators", uint64(len(InitManaOperators()))},

		/* mana_tables */
		{ManaTablesID, "mana_tables", uint64(numTblID) - 1},

		/* mana_types */
		{ManaTypesID, "mana_types", uint64(numTypeID) - 1},
	}
}

const (
	_               OID = iota
	ManaCastID          /* mana_cast */
	ManaColumnsID       /* mana_columns */
	ManaFuncsID         /* mana_funcs */
	ManaOperatorsID     /* mana_operators */
	ManaTablesID        /* mana_tables */
	ManaTypesID         /* mana_types */
	numTblID
)

func init() {
	metas := InitManaTables()
	h := make(map[OID]struct{})
	for _, meta := range metas {
		if _, ok := h[meta.TblID]; ok {
			panic("duplicate metadata exists in system tables: mana_tables")
		}
	}
	if len(metas) != int(numTblID)-1 {
		panic("there is uninitialized metadata in system tables: mana_tables")
	}
}
