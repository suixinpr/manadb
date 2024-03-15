package metadata

type ManaTypes struct {
	TypeID     OID    /* mana_types: typeid */
	TypeName   string /* mana_types: typename */
	TypeLen    int16  /* mana_types: typelen */
	TypeAlign  uint8  /* mana_types: typealign */
	TypeInput  OID    /* mana_types: typeinput */
	TypeOutput OID    /* mana_types: typeoutput */
}

const (
	ManaTypesTypeID int16 = iota
	ManaTypesTypeName
	ManaTypesTypeLen
	ManaTypesTypeAlign
	ManaTypesTypeInput
	ManaTypesTypeOutput
	NumManaTypes /* number of mana_types columns */
)

func InitManaTypes() []*ManaTypes {
	return []*ManaTypes{
		/* boolean: unknown, false, true */
		{BooleanID, "boolean", 1, 1, FuncBooleanInput, FuncBooleanOutput},

		/* enumeration */
		{ByteID, "byte", 1, 1, FuncByteInput, FuncByteOutput},

		/* builtin string */
		{BytesID, "bytes", -1, 1, FuncBytesInput, FuncBytesOutput},

		/* char(length), fixed storage length */
		{CharID, "char", -1, 1, FuncCharInput, FuncCharOutput},

		/* single-precision float point number */
		{Float32ID, "float32", 4, 4, FuncFloat32Input, FuncFloat32Output},

		/* double-precision float point number */
		{Float64ID, "float64", 8, 8, FuncFloat64Input, FuncFloat64Output},

		/* -128 ~ 127 */
		{Int8ID, "int8", 1, 1, FuncInt8Input, FuncInt8Output},

		/* -32768 ~ 32767 */
		{Int16ID, "int16", 2, 2, FuncInt16Input, FuncInt16Output},

		/* -2 billion ~ 2 billion */
		{Int32ID, "int32", 4, 4, FuncInt32Input, FuncInt32Output},

		/* 19 digit integer */
		{Int64ID, "int64", 8, 8, FuncInt64Input, FuncInt64Output},

		/* row identifier, 8-byte storage */
		{OIDID, "oid", 8, 8, FuncOIDInput, FuncOIDOutput},

		/* row identifier, 8-byte storage */
		{OIDArrayID, "oidarray", -1, 8, FuncOIDArrayInput, FuncOIDArrayOutput},

		/* variable-length string */
		{TextID, "text", -1, 1, FuncTextInput, FuncTextOutput},

		/* 0 ~ 255 */
		{Uint8ID, "uint8", 1, 1, FuncUint8Input, FuncUint8Output},

		/* 0 ~ 65535 */
		{Uint16ID, "uint16", 2, 2, FuncUint16Input, FuncUint16Output},

		/* maximum 4 billion */
		{Uint32ID, "uint32", 4, 4, FuncUint32Input, FuncUint32Output},

		/* 20 digit integer */
		{Uint64ID, "uint64", 8, 8, FuncUint64Input, FuncUint64Output},

		/* varchar(length), variable storage length */
		{VarcharID, "varchar", -1, 1, FuncVarcharInput, FuncVarcharOutput},
	}
}

/* abstract type id */
const (
	_ OID = iota
	BooleanID
	ByteID
	BytesID
	CharID
	Float32ID
	Float64ID
	Int8ID
	Int16ID
	Int32ID
	Int64ID
	OIDID
	OIDArrayID
	TextID
	Uint8ID
	Uint16ID
	Uint32ID
	Uint64ID
	VarcharID
	numTypeID
)

func init() {
	metas := InitManaTypes()
	h := make(map[OID]struct{})
	for _, meta := range metas {
		if _, ok := h[meta.TypeID]; ok {
			panic("duplicate metadata exists in system tables: mana_types")
		}
	}
	if len(metas) != int(numTypeID)-1 {
		panic("there is uninitialized metadata in system tables: mana_types")
	}
}
