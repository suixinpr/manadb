package metadata

type ManaCast struct {
	CastID   OID  /* mana_cast: castid */
	CastIn   OID  /* mana_cast: castin */
	CastOut  OID  /* mana_cast: castout */
	CastFunc OID  /* mana_cast: castfunc */
	CastType byte /* mana_cast: casttype */
}

const (
	ManaCastCastID int16 = iota
	ManaCastCastIn
	ManaCastCastOut
	ManaCastCastFunc
	ManaCastCastType
	NumManaCast /* number of mana_cast columns */
)

func InitManaCast() (metas []*ManaCast) {
	defer func() {
		for oid, meta := range metas {
			meta.CastID = OID(oid) + 1
		}
	}()

	return []*ManaCast{
		/* Int8 */
		{InvalidID, Int8ID, Int16ID, FuncInt8ToInt16, 'a'},
		{InvalidID, Int8ID, Int32ID, FuncInt8ToInt32, 'a'},
		{InvalidID, Int8ID, Int64ID, FuncInt8ToInt64, 'a'},
		{InvalidID, Int8ID, OIDID, FuncInt8ToOID, 'a'},
		{InvalidID, Int8ID, Uint8ID, FuncInt8ToUint8, 'a'},
		{InvalidID, Int8ID, Uint16ID, FuncInt8ToUint16, 'a'},
		{InvalidID, Int8ID, Uint32ID, FuncInt8ToUint32, 'a'},
		{InvalidID, Int8ID, Uint64ID, FuncInt8ToUint64, 'a'},

		/* Int16 */
		{InvalidID, Int16ID, Int8ID, FuncInt16ToInt8, 'a'},
		{InvalidID, Int16ID, Int32ID, FuncInt16ToInt32, 'a'},
		{InvalidID, Int16ID, Int64ID, FuncInt16ToInt64, 'a'},
		{InvalidID, Int16ID, OIDID, FuncInt16ToOID, 'a'},
		{InvalidID, Int16ID, Uint8ID, FuncInt16ToUint8, 'a'},
		{InvalidID, Int16ID, Uint16ID, FuncInt16ToUint16, 'a'},
		{InvalidID, Int16ID, Uint32ID, FuncInt16ToUint32, 'a'},
		{InvalidID, Int16ID, Uint64ID, FuncInt16ToUint64, 'a'},

		/* Int32 */
		{InvalidID, Int32ID, Int8ID, FuncInt32ToInt8, 'a'},
		{InvalidID, Int32ID, Int16ID, FuncInt32ToInt16, 'a'},
		{InvalidID, Int32ID, Int64ID, FuncInt32ToInt64, 'a'},
		{InvalidID, Int32ID, OIDID, FuncInt32ToOID, 'a'},
		{InvalidID, Int32ID, Uint8ID, FuncInt32ToUint8, 'a'},
		{InvalidID, Int32ID, Uint16ID, FuncInt32ToUint16, 'a'},
		{InvalidID, Int32ID, Uint32ID, FuncInt32ToUint32, 'a'},
		{InvalidID, Int32ID, Uint64ID, FuncInt32ToUint64, 'a'},

		/* Int64 */
		{InvalidID, Int64ID, Int8ID, FuncInt64ToInt8, 'a'},
		{InvalidID, Int64ID, Int16ID, FuncInt64ToInt16, 'a'},
		{InvalidID, Int64ID, Int32ID, FuncInt64ToInt32, 'a'},
		{InvalidID, Int64ID, OIDID, FuncInt64ToOID, 'a'},
		{InvalidID, Int64ID, Uint8ID, FuncInt64ToUint8, 'a'},
		{InvalidID, Int64ID, Uint16ID, FuncInt64ToUint16, 'a'},
		{InvalidID, Int64ID, Uint32ID, FuncInt64ToUint32, 'a'},
		{InvalidID, Int64ID, Uint64ID, FuncInt64ToUint64, 'a'},

		/* Uint8 */
		{InvalidID, Uint8ID, Int8ID, FuncUint8ToInt8, 'a'},
		{InvalidID, Uint8ID, Int16ID, FuncUint8ToInt16, 'a'},
		{InvalidID, Uint8ID, Int32ID, FuncUint8ToInt32, 'a'},
		{InvalidID, Uint8ID, Int64ID, FuncUint8ToInt64, 'a'},
		{InvalidID, Uint8ID, OIDID, FuncUint8ToOID, 'a'},
		{InvalidID, Uint8ID, Uint16ID, FuncUint8ToUint16, 'a'},
		{InvalidID, Uint8ID, Uint32ID, FuncUint8ToUint32, 'a'},
		{InvalidID, Uint8ID, Uint64ID, FuncUint8ToUint64, 'a'},

		/* Uint16 */
		{InvalidID, Uint16ID, Int8ID, FuncUint16ToInt8, 'a'},
		{InvalidID, Uint16ID, Int16ID, FuncUint16ToInt16, 'a'},
		{InvalidID, Uint16ID, Int32ID, FuncUint16ToInt32, 'a'},
		{InvalidID, Uint16ID, Int64ID, FuncUint16ToInt64, 'a'},
		{InvalidID, Uint16ID, OIDID, FuncUint16ToOID, 'a'},
		{InvalidID, Uint16ID, Uint8ID, FuncUint16ToUint8, 'a'},
		{InvalidID, Uint16ID, Uint32ID, FuncUint16ToUint32, 'a'},
		{InvalidID, Uint16ID, Uint64ID, FuncUint16ToUint64, 'a'},

		/* Uint32 */
		{InvalidID, Uint32ID, Int8ID, FuncUint32ToInt8, 'a'},
		{InvalidID, Uint32ID, Int16ID, FuncUint32ToInt16, 'a'},
		{InvalidID, Uint32ID, Int32ID, FuncUint32ToInt32, 'a'},
		{InvalidID, Uint32ID, Int64ID, FuncUint32ToInt64, 'a'},
		{InvalidID, Uint32ID, OIDID, FuncUint32ToOID, 'a'},
		{InvalidID, Uint32ID, Uint8ID, FuncUint32ToUint8, 'a'},
		{InvalidID, Uint32ID, Uint16ID, FuncUint32ToUint16, 'a'},
		{InvalidID, Uint32ID, Uint64ID, FuncUint32ToUint64, 'a'},

		/* Uint64 */
		{InvalidID, Uint64ID, Int8ID, FuncUint64ToInt8, 'a'},
		{InvalidID, Uint64ID, Int16ID, FuncUint64ToInt16, 'a'},
		{InvalidID, Uint64ID, Int32ID, FuncUint64ToInt32, 'a'},
		{InvalidID, Uint64ID, Int64ID, FuncUint64ToInt64, 'a'},
		{InvalidID, Uint64ID, OIDID, FuncUint64ToOID, 'a'},
		{InvalidID, Uint64ID, Uint8ID, FuncUint64ToUint8, 'a'},
		{InvalidID, Uint64ID, Uint16ID, FuncUint64ToUint16, 'a'},
		{InvalidID, Uint64ID, Uint32ID, FuncUint64ToUint32, 'a'},
	}
}

func init() {
	metas := InitManaCast()
	h := make(map[OID]struct{})
	for _, meta := range metas {
		if _, ok := h[meta.CastID]; ok {
			panic("duplicate metadata exists in system tables: mana_cast")
		}
		if meta.CastID == InvalidID {
			panic("Invalid OID metadata exists in system tables: mana_cast")
		}
	}
}
