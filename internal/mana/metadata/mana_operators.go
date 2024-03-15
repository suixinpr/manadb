package metadata

type ManaOperators struct {
	OprID     OID    /* mana_operators: oprid */
	OprName   string /* mana_operators: oprname */
	OprLeft   OID    /* mana_operators: oprleft */
	OprRight  OID    /* mana_operators: oprright */
	OprResult OID    /* mana_operators: oprresult */
	OprFuncID OID    /* mana_operators: oprfuncid */
}

const (
	ManaOperatorsOprID int16 = iota
	ManaOperatorsOprName
	ManaOperatorsOprLeft
	ManaOperatorsOprRight
	ManaOperatorsOprResult
	ManaOperatorsOprFuncID
	NumManaOperators /* number of mana_operators columns */
)

func InitManaOperators() (metas []*ManaOperators) {
	defer func() {
		for oid, meta := range metas {
			meta.OprID = OID(oid) + 1
		}
	}()

	return []*ManaOperators{
		/* Boolean */
		{InvalidID, "==", BooleanID, BooleanID, BooleanID, FuncBooleanEQ},
		{InvalidID, "<>", BooleanID, BooleanID, BooleanID, FuncBooleanNE},
		{InvalidID, "and", BooleanID, BooleanID, BooleanID, FuncBooleanAnd},
		{InvalidID, "or", BooleanID, BooleanID, BooleanID, FuncBooleanOr},
		{InvalidID, "not", BooleanID, InvalidID, BooleanID, FuncBooleanNot},

		/* Byte */
		{InvalidID, "<", ByteID, ByteID, BooleanID, FuncByteLT},
		{InvalidID, "=", ByteID, ByteID, BooleanID, FuncByteEQ},
		{InvalidID, ">", ByteID, ByteID, BooleanID, FuncByteGT},
		{InvalidID, "<=", ByteID, ByteID, BooleanID, FuncByteLE},
		{InvalidID, "<>", ByteID, ByteID, BooleanID, FuncByteNE},
		{InvalidID, ">=", ByteID, ByteID, BooleanID, FuncByteGE},

		{InvalidID, "+", ByteID, ByteID, ByteID, FuncByteAdd},
		{InvalidID, "-", ByteID, ByteID, ByteID, FuncByteSub},
		{InvalidID, "*", ByteID, ByteID, ByteID, FuncByteMul},
		{InvalidID, "/", ByteID, ByteID, ByteID, FuncByteDiv},

		/* Vhar */
		{InvalidID, "<", CharID, CharID, BooleanID, FuncCharLT},
		{InvalidID, "=", CharID, CharID, BooleanID, FuncCharEQ},
		{InvalidID, ">", CharID, CharID, BooleanID, FuncCharGT},
		{InvalidID, "<=", CharID, CharID, BooleanID, FuncCharLE},
		{InvalidID, "<>", CharID, CharID, BooleanID, FuncCharNE},
		{InvalidID, ">=", CharID, CharID, BooleanID, FuncCharGE},

		{InvalidID, "+", CharID, CharID, CharID, FuncCharAdd},

		/* Float32 */
		{InvalidID, "<", Float32ID, Float32ID, BooleanID, FuncFloat32LT},
		{InvalidID, "=", Float32ID, Float32ID, BooleanID, FuncFloat32EQ},
		{InvalidID, ">", Float32ID, Float32ID, BooleanID, FuncFloat32GT},
		{InvalidID, "<=", Float32ID, Float32ID, BooleanID, FuncFloat32LE},
		{InvalidID, "<>", Float32ID, Float32ID, BooleanID, FuncFloat32NE},
		{InvalidID, ">=", Float32ID, Float32ID, BooleanID, FuncFloat32GE},

		{InvalidID, "+", Float32ID, Float32ID, Float32ID, FuncFloat32Add},
		{InvalidID, "-", Float32ID, Float32ID, Float32ID, FuncFloat32Sub},
		{InvalidID, "*", Float32ID, Float32ID, Float32ID, FuncFloat32Mul},
		{InvalidID, "/", Float32ID, Float32ID, Float32ID, FuncFloat32Div},

		/* Float64 */
		{InvalidID, "<", Float64ID, Float64ID, BooleanID, FuncFloat64LT},
		{InvalidID, "=", Float64ID, Float64ID, BooleanID, FuncFloat64EQ},
		{InvalidID, ">", Float64ID, Float64ID, BooleanID, FuncFloat64GT},
		{InvalidID, "<=", Float64ID, Float64ID, BooleanID, FuncFloat64LE},
		{InvalidID, "<>", Float64ID, Float64ID, BooleanID, FuncFloat64NE},
		{InvalidID, ">=", Float64ID, Float64ID, BooleanID, FuncFloat64GE},

		{InvalidID, "+", Float64ID, Float64ID, Float64ID, FuncFloat64Add},
		{InvalidID, "-", Float64ID, Float64ID, Float64ID, FuncFloat64Sub},
		{InvalidID, "*", Float64ID, Float64ID, Float64ID, FuncFloat64Mul},
		{InvalidID, "/", Float64ID, Float64ID, Float64ID, FuncFloat64Div},

		/* Int8 */
		{InvalidID, "<", Int8ID, Int8ID, BooleanID, FuncInt8LT},
		{InvalidID, "=", Int8ID, Int8ID, BooleanID, FuncInt8EQ},
		{InvalidID, ">", Int8ID, Int8ID, BooleanID, FuncInt8GT},
		{InvalidID, "<=", Int8ID, Int8ID, BooleanID, FuncInt8LE},
		{InvalidID, "<>", Int8ID, Int8ID, BooleanID, FuncInt8NE},
		{InvalidID, ">=", Int8ID, Int8ID, BooleanID, FuncInt8GE},

		{InvalidID, "+", Int8ID, Int8ID, Int8ID, FuncInt8Add},
		{InvalidID, "-", Int8ID, Int8ID, Int8ID, FuncInt8Sub},
		{InvalidID, "*", Int8ID, Int8ID, Int8ID, FuncInt8Mul},
		{InvalidID, "/", Int8ID, Int8ID, Int8ID, FuncInt8Div},

		/* Int16 */
		{InvalidID, "<", Int16ID, Int16ID, BooleanID, FuncInt16LT},
		{InvalidID, "=", Int16ID, Int16ID, BooleanID, FuncInt16EQ},
		{InvalidID, ">", Int16ID, Int16ID, BooleanID, FuncInt16GT},
		{InvalidID, "<=", Int16ID, Int16ID, BooleanID, FuncInt16LE},
		{InvalidID, "<>", Int16ID, Int16ID, BooleanID, FuncInt16NE},
		{InvalidID, ">=", Int16ID, Int16ID, BooleanID, FuncInt16GE},

		{InvalidID, "+", Int16ID, Int16ID, Int16ID, FuncInt16Add},
		{InvalidID, "-", Int16ID, Int16ID, Int16ID, FuncInt16Sub},
		{InvalidID, "*", Int16ID, Int16ID, Int16ID, FuncInt16Mul},
		{InvalidID, "/", Int16ID, Int16ID, Int16ID, FuncInt16Div},

		/* Int32 */
		{InvalidID, "<", Int32ID, Int32ID, BooleanID, FuncInt32LT},
		{InvalidID, "=", Int32ID, Int32ID, BooleanID, FuncInt32EQ},
		{InvalidID, ">", Int32ID, Int32ID, BooleanID, FuncInt32GT},
		{InvalidID, "<=", Int32ID, Int32ID, BooleanID, FuncInt32LE},
		{InvalidID, "<>", Int32ID, Int32ID, BooleanID, FuncInt32NE},
		{InvalidID, ">=", Int32ID, Int32ID, BooleanID, FuncInt32GE},

		{InvalidID, "+", Int32ID, Int32ID, Int32ID, FuncInt32Add},
		{InvalidID, "-", Int32ID, Int32ID, Int32ID, FuncInt32Sub},
		{InvalidID, "*", Int32ID, Int32ID, Int32ID, FuncInt32Mul},
		{InvalidID, "/", Int32ID, Int32ID, Int32ID, FuncInt32Div},

		/* Int64 */
		{InvalidID, "<", Int64ID, Int64ID, BooleanID, FuncInt64LT},
		{InvalidID, "=", Int64ID, Int64ID, BooleanID, FuncInt64EQ},
		{InvalidID, ">", Int64ID, Int64ID, BooleanID, FuncInt64GT},
		{InvalidID, "<=", Int64ID, Int64ID, BooleanID, FuncInt64LE},
		{InvalidID, "<>", Int64ID, Int64ID, BooleanID, FuncInt64NE},
		{InvalidID, ">=", Int64ID, Int64ID, BooleanID, FuncInt64GE},

		{InvalidID, "+", Int64ID, Int64ID, Int64ID, FuncInt64Add},
		{InvalidID, "-", Int64ID, Int64ID, Int64ID, FuncInt64Sub},
		{InvalidID, "*", Int64ID, Int64ID, Int64ID, FuncInt64Mul},
		{InvalidID, "/", Int64ID, Int64ID, Int64ID, FuncInt64Div},

		/* OID */
		{InvalidID, "<", OIDID, OIDID, BooleanID, FuncOIDLT},
		{InvalidID, "=", OIDID, OIDID, BooleanID, FuncOIDEQ},
		{InvalidID, ">", OIDID, OIDID, BooleanID, FuncOIDGT},
		{InvalidID, "<=", OIDID, OIDID, BooleanID, FuncOIDLE},
		{InvalidID, "<>", OIDID, OIDID, BooleanID, FuncOIDNE},
		{InvalidID, ">=", OIDID, OIDID, BooleanID, FuncOIDGE},

		{InvalidID, "+", OIDID, OIDID, OIDID, FuncOIDAdd},
		{InvalidID, "-", OIDID, OIDID, OIDID, FuncOIDSub},
		{InvalidID, "*", OIDID, OIDID, OIDID, FuncOIDMul},
		{InvalidID, "/", OIDID, OIDID, OIDID, FuncOIDDiv},

		/* Text */
		{InvalidID, "<", TextID, TextID, BooleanID, FuncTextLT},
		{InvalidID, "=", TextID, TextID, BooleanID, FuncTextEQ},
		{InvalidID, ">", TextID, TextID, BooleanID, FuncTextGT},
		{InvalidID, "<=", TextID, TextID, BooleanID, FuncTextLE},
		{InvalidID, "<>", TextID, TextID, BooleanID, FuncTextNE},
		{InvalidID, ">=", TextID, TextID, BooleanID, FuncTextGE},

		{InvalidID, "+", TextID, TextID, TextID, FuncTextAdd},

		/* Uint8 */
		{InvalidID, "<", Uint8ID, Uint8ID, BooleanID, FuncUint8LT},
		{InvalidID, "=", Uint8ID, Uint8ID, BooleanID, FuncUint8EQ},
		{InvalidID, ">", Uint8ID, Uint8ID, BooleanID, FuncUint8GT},
		{InvalidID, "<=", Uint8ID, Uint8ID, BooleanID, FuncUint8LE},
		{InvalidID, "<>", Uint8ID, Uint8ID, BooleanID, FuncUint8NE},
		{InvalidID, ">=", Uint8ID, Uint8ID, BooleanID, FuncUint8GE},

		{InvalidID, "+", Uint8ID, Uint8ID, Uint8ID, FuncUint8Add},
		{InvalidID, "-", Uint8ID, Uint8ID, Uint8ID, FuncUint8Sub},
		{InvalidID, "*", Uint8ID, Uint8ID, Uint8ID, FuncUint8Mul},
		{InvalidID, "/", Uint8ID, Uint8ID, Uint8ID, FuncUint8Div},

		/* Uint16 */
		{InvalidID, "<", Uint16ID, Uint16ID, BooleanID, FuncUint16LT},
		{InvalidID, "=", Uint16ID, Uint16ID, BooleanID, FuncUint16EQ},
		{InvalidID, ">", Uint16ID, Uint16ID, BooleanID, FuncUint16GT},
		{InvalidID, "<=", Uint16ID, Uint16ID, BooleanID, FuncUint16LE},
		{InvalidID, "<>", Uint16ID, Uint16ID, BooleanID, FuncUint16NE},
		{InvalidID, ">=", Uint16ID, Uint16ID, BooleanID, FuncUint16GE},

		{InvalidID, "+", Uint16ID, Uint16ID, Uint16ID, FuncUint16Add},
		{InvalidID, "-", Uint16ID, Uint16ID, Uint16ID, FuncUint16Sub},
		{InvalidID, "*", Uint16ID, Uint16ID, Uint16ID, FuncUint16Mul},
		{InvalidID, "/", Uint16ID, Uint16ID, Uint16ID, FuncUint16Div},

		/* Uint32 */
		{InvalidID, "<", Uint32ID, Uint32ID, BooleanID, FuncUint32LT},
		{InvalidID, "=", Uint32ID, Uint32ID, BooleanID, FuncUint32EQ},
		{InvalidID, ">", Uint32ID, Uint32ID, BooleanID, FuncUint32GT},
		{InvalidID, "<=", Uint32ID, Uint32ID, BooleanID, FuncUint32LE},
		{InvalidID, "<>", Uint32ID, Uint32ID, BooleanID, FuncUint32NE},
		{InvalidID, ">=", Uint32ID, Uint32ID, BooleanID, FuncUint32GE},

		{InvalidID, "+", Uint32ID, Uint32ID, Uint32ID, FuncUint32Add},
		{InvalidID, "-", Uint32ID, Uint32ID, Uint32ID, FuncUint32Sub},
		{InvalidID, "*", Uint32ID, Uint32ID, Uint32ID, FuncUint32Mul},
		{InvalidID, "/", Uint32ID, Uint32ID, Uint32ID, FuncUint32Div},

		/* Uint64 */
		{InvalidID, "<", Uint64ID, Uint64ID, BooleanID, FuncUint64LT},
		{InvalidID, "=", Uint64ID, Uint64ID, BooleanID, FuncUint64EQ},
		{InvalidID, ">", Uint64ID, Uint64ID, BooleanID, FuncUint64GT},
		{InvalidID, "<=", Uint64ID, Uint64ID, BooleanID, FuncUint64LE},
		{InvalidID, "<>", Uint64ID, Uint64ID, BooleanID, FuncUint64NE},
		{InvalidID, ">=", Uint64ID, Uint64ID, BooleanID, FuncUint64GE},

		{InvalidID, "+", Uint64ID, Uint64ID, Uint64ID, FuncUint64Add},
		{InvalidID, "-", Uint64ID, Uint64ID, Uint64ID, FuncUint64Sub},
		{InvalidID, "*", Uint64ID, Uint64ID, Uint64ID, FuncUint64Mul},
		{InvalidID, "/", Uint64ID, Uint64ID, Uint64ID, FuncUint64Div},

		/* Varchar */
		{InvalidID, "<", VarcharID, VarcharID, BooleanID, FuncVarcharLT},
		{InvalidID, "=", VarcharID, VarcharID, BooleanID, FuncVarcharEQ},
		{InvalidID, ">", VarcharID, VarcharID, BooleanID, FuncVarcharGT},
		{InvalidID, "<=", VarcharID, VarcharID, BooleanID, FuncVarcharLE},
		{InvalidID, "<>", VarcharID, VarcharID, BooleanID, FuncVarcharNE},
		{InvalidID, ">=", VarcharID, VarcharID, BooleanID, FuncVarcharGE},

		{InvalidID, "+", VarcharID, VarcharID, VarcharID, FuncVarcharAdd},
	}
}

func init() {
	metas := InitManaOperators()
	h := make(map[OID]struct{})
	for _, meta := range metas {
		if _, ok := h[meta.OprID]; ok {
			panic("duplicate metadata exists in system tables: mana_operators")
		}
		if meta.OprID == InvalidID {
			panic("Invalid OID metadata exists in system tables: mana_operators")
		}
	}
}
