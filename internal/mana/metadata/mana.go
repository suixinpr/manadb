package metadata

/* builtin type oid */
type OID uint64

const InvalidID OID = 0

func IsCatalog(oid OID) bool {
	return oid > InvalidID && oid < numTblID
}
