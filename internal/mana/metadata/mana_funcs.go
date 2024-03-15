package metadata

type ManaFuncs struct {
	FuncID       OID    /* mana_funcs: funcid */
	FuncName     string /* mana_funcs: funcname */
	FuncArgs     int8   /* mana_funcs: funcargs */
	FuncArgTypes []OID  /* mana_funcs: funcargtypes */
	FuncRetType  OID    /* mana_funcs: funcrettype */
}

const (
	ManaFuncsFuncID int16 = iota
	ManaFuncsFuncName
	ManaFuncsFuncArgs
	ManaFuncsFuncArgTypes
	ManaFuncsFuncRetType
	NumManaFuncs /* number of mana_funcs columns */
)

func InitManaFuncs() []*ManaFuncs {
	return []*ManaFuncs{
		/* boolean */
		{FuncBooleanInput, "boolean_input", 1, []OID{BytesID}, BooleanID},
		{FuncBooleanOutput, "boolean_output", 1, []OID{BooleanID}, BytesID},
		{FuncBooleanEQ, "boolean_eq", 2, []OID{BooleanID, BooleanID}, BooleanID},
		{FuncBooleanNE, "boolean_ne", 2, []OID{BooleanID, BooleanID}, BooleanID},
		{FuncBooleanAnd, "boolean_and", 2, []OID{BooleanID, BooleanID}, BooleanID},
		{FuncBooleanOr, "boolean_or", 2, []OID{BooleanID, BooleanID}, BooleanID},
		{FuncBooleanNot, "boolean_not", 1, []OID{BooleanID}, BooleanID},

		/* byte */
		{FuncByteInput, "byte_input", 1, []OID{BytesID}, ByteID},
		{FuncByteOutput, "byte_output", 1, []OID{ByteID}, BytesID},
		{FuncByteLT, "byte_lt", 2, []OID{ByteID, ByteID}, BooleanID},
		{FuncByteEQ, "byte_eq", 2, []OID{ByteID, ByteID}, BooleanID},
		{FuncByteGT, "byte_gt", 2, []OID{ByteID, ByteID}, BooleanID},
		{FuncByteLE, "byte_le", 2, []OID{ByteID, ByteID}, BooleanID},
		{FuncByteNE, "byte_ne", 2, []OID{ByteID, ByteID}, BooleanID},
		{FuncByteGE, "byte_ge", 2, []OID{ByteID, ByteID}, BooleanID},
		{FuncByteAdd, "byte_add", 2, []OID{ByteID, ByteID}, ByteID},
		{FuncByteSub, "byte_sub", 2, []OID{ByteID, ByteID}, ByteID},
		{FuncByteMul, "byte_mul", 2, []OID{ByteID, ByteID}, ByteID},
		{FuncByteDiv, "byte_div", 2, []OID{ByteID, ByteID}, ByteID},

		/* bytes */
		{FuncBytesInput, "bytes_input", 1, []OID{BytesID}, BytesID},
		{FuncBytesOutput, "bytes_output", 1, []OID{BytesID}, BytesID},

		/* char */
		{FuncCharInput, "char_input", 2, []OID{BytesID, Int32ID}, CharID},
		{FuncCharOutput, "char_output", 1, []OID{CharID}, BytesID},
		{FuncCharLT, "char_lt", 2, []OID{CharID, CharID}, BooleanID},
		{FuncCharEQ, "char_eq", 2, []OID{CharID, CharID}, BooleanID},
		{FuncCharGT, "char_gt", 2, []OID{CharID, CharID}, BooleanID},
		{FuncCharLE, "char_le", 2, []OID{CharID, CharID}, BooleanID},
		{FuncCharNE, "char_ne", 2, []OID{CharID, CharID}, BooleanID},
		{FuncCharGE, "char_ge", 2, []OID{CharID, CharID}, BooleanID},
		{FuncCharAdd, "char_add", 2, []OID{CharID, CharID}, CharID},

		/* float32 */
		{FuncFloat32Input, "float32_input", 1, []OID{BytesID}, Float32ID},
		{FuncFloat32Output, "float32_output", 1, []OID{Float32ID}, BytesID},
		{FuncFloat32LT, "float32_lt", 2, []OID{Float32ID, Float32ID}, BooleanID},
		{FuncFloat32EQ, "float32_eq", 2, []OID{Float32ID, Float32ID}, BooleanID},
		{FuncFloat32GT, "float32_gt", 2, []OID{Float32ID, Float32ID}, BooleanID},
		{FuncFloat32LE, "float32_le", 2, []OID{Float32ID, Float32ID}, BooleanID},
		{FuncFloat32NE, "float32_ne", 2, []OID{Float32ID, Float32ID}, BooleanID},
		{FuncFloat32GE, "float32_ge", 2, []OID{Float32ID, Float32ID}, BooleanID},
		{FuncFloat32Add, "float32_add", 2, []OID{Float32ID, Float32ID}, Float32ID},
		{FuncFloat32Sub, "float32_sub", 2, []OID{Float32ID, Float32ID}, Float32ID},
		{FuncFloat32Mul, "float32_mul", 2, []OID{Float32ID, Float32ID}, Float32ID},
		{FuncFloat32Div, "float32_div", 2, []OID{Float32ID, Float32ID}, Float32ID},

		/* float64 */
		{FuncFloat64Input, "float64_input", 1, []OID{BytesID}, Float64ID},
		{FuncFloat64Output, "float64_output", 1, []OID{Float64ID}, BytesID},
		{FuncFloat64LT, "float64_lt", 2, []OID{Float64ID, Float64ID}, BooleanID},
		{FuncFloat64EQ, "float64_eq", 2, []OID{Float64ID, Float64ID}, BooleanID},
		{FuncFloat64GT, "float64_gt", 2, []OID{Float64ID, Float64ID}, BooleanID},
		{FuncFloat64LE, "float64_le", 2, []OID{Float64ID, Float64ID}, BooleanID},
		{FuncFloat64NE, "float64_ne", 2, []OID{Float64ID, Float64ID}, BooleanID},
		{FuncFloat64GE, "float64_ge", 2, []OID{Float64ID, Float64ID}, BooleanID},
		{FuncFloat64Add, "float64_add", 2, []OID{Float64ID, Float64ID}, Float64ID},
		{FuncFloat64Sub, "float64_sub", 2, []OID{Float64ID, Float64ID}, Float64ID},
		{FuncFloat64Mul, "float64_mul", 2, []OID{Float64ID, Float64ID}, Float64ID},
		{FuncFloat64Div, "float64_div", 2, []OID{Float64ID, Float64ID}, Float64ID},

		/* int8 */
		{FuncInt8Input, "int8_input", 1, []OID{BytesID}, Int8ID},
		{FuncInt8Output, "int8_output", 1, []OID{Int8ID}, BytesID},
		{FuncInt8LT, "int8_lt", 2, []OID{Int8ID, Int8ID}, BooleanID},
		{FuncInt8EQ, "int8_eq", 2, []OID{Int8ID, Int8ID}, BooleanID},
		{FuncInt8GT, "int8_gt", 2, []OID{Int8ID, Int8ID}, BooleanID},
		{FuncInt8LE, "int8_le", 2, []OID{Int8ID, Int8ID}, BooleanID},
		{FuncInt8NE, "int8_ne", 2, []OID{Int8ID, Int8ID}, BooleanID},
		{FuncInt8GE, "int8_ge", 2, []OID{Int8ID, Int8ID}, BooleanID},
		{FuncInt8Add, "int8_add", 2, []OID{Int8ID, Int8ID}, Int8ID},
		{FuncInt8Sub, "int8_sub", 2, []OID{Int8ID, Int8ID}, Int8ID},
		{FuncInt8Mul, "int8_mul", 2, []OID{Int8ID, Int8ID}, Int8ID},
		{FuncInt8Div, "int8_div", 2, []OID{Int8ID, Int8ID}, Int8ID},
		{FuncInt8ToInt16, "int8_to_int16", 1, []OID{Int8ID}, Int16ID},
		{FuncInt8ToInt32, "int8_to_int32", 1, []OID{Int8ID}, Int32ID},
		{FuncInt8ToInt64, "int8_to_int64", 1, []OID{Int8ID}, Int64ID},
		{FuncInt8ToOID, "int8_to_oid", 1, []OID{Int8ID}, OIDID},
		{FuncInt8ToUint8, "int8_to_uint8", 1, []OID{Int8ID}, Uint8ID},
		{FuncInt8ToUint16, "int8_to_uint16", 1, []OID{Int8ID}, Uint16ID},
		{FuncInt8ToUint32, "int8_to_uint32", 1, []OID{Int8ID}, Uint32ID},
		{FuncInt8ToUint64, "int8_to_uint64", 1, []OID{Int8ID}, Uint64ID},

		/* int16 */
		{FuncInt16Input, "int16_input", 1, []OID{BytesID}, Int16ID},
		{FuncInt16Output, "int16_output", 1, []OID{Int16ID}, BytesID},
		{FuncInt16LT, "int16_lt", 2, []OID{Int16ID, Int16ID}, BooleanID},
		{FuncInt16EQ, "int16_eq", 2, []OID{Int16ID, Int16ID}, BooleanID},
		{FuncInt16GT, "int16_gt", 2, []OID{Int16ID, Int16ID}, BooleanID},
		{FuncInt16LE, "int16_le", 2, []OID{Int16ID, Int16ID}, BooleanID},
		{FuncInt16NE, "int16_ne", 2, []OID{Int16ID, Int16ID}, BooleanID},
		{FuncInt16GE, "int16_ge", 2, []OID{Int16ID, Int16ID}, BooleanID},
		{FuncInt16Add, "int16_add", 2, []OID{Int16ID, Int16ID}, Int16ID},
		{FuncInt16Sub, "int16_sub", 2, []OID{Int16ID, Int16ID}, Int16ID},
		{FuncInt16Mul, "int16_mul", 2, []OID{Int16ID, Int16ID}, Int16ID},
		{FuncInt16Div, "int16_div", 2, []OID{Int16ID, Int16ID}, Int16ID},
		{FuncInt16ToInt8, "int16_to_int8", 1, []OID{Int16ID}, Int8ID},
		{FuncInt16ToInt32, "int16_to_int32", 1, []OID{Int16ID}, Int32ID},
		{FuncInt16ToInt64, "int16_to_int64", 1, []OID{Int16ID}, Int64ID},
		{FuncInt16ToOID, "int16_to_oid", 1, []OID{Int16ID}, OIDID},
		{FuncInt16ToUint8, "int16_to_uint8", 1, []OID{Int16ID}, Uint8ID},
		{FuncInt16ToUint16, "int16_to_uint16", 1, []OID{Int16ID}, Uint16ID},
		{FuncInt16ToUint32, "int16_to_uint32", 1, []OID{Int16ID}, Uint32ID},
		{FuncInt16ToUint64, "int16_to_uint64", 1, []OID{Int16ID}, Uint64ID},

		/* int32 */
		{FuncInt32Input, "int32_input", 1, []OID{BytesID}, Int32ID},
		{FuncInt32Output, "int32_output", 1, []OID{Int32ID}, BytesID},
		{FuncInt32LT, "int32_lt", 2, []OID{Int32ID, Int32ID}, BooleanID},
		{FuncInt32EQ, "int32_eq", 2, []OID{Int32ID, Int32ID}, BooleanID},
		{FuncInt32GT, "int32_gt", 2, []OID{Int32ID, Int32ID}, BooleanID},
		{FuncInt32LE, "int32_le", 2, []OID{Int32ID, Int32ID}, BooleanID},
		{FuncInt32NE, "int32_ne", 2, []OID{Int32ID, Int32ID}, BooleanID},
		{FuncInt32GE, "int32_ge", 2, []OID{Int32ID, Int32ID}, BooleanID},
		{FuncInt32Add, "int32_add", 2, []OID{Int32ID, Int32ID}, Int32ID},
		{FuncInt32Sub, "int32_sub", 2, []OID{Int32ID, Int32ID}, Int32ID},
		{FuncInt32Mul, "int32_mul", 2, []OID{Int32ID, Int32ID}, Int32ID},
		{FuncInt32Div, "int32_div", 2, []OID{Int32ID, Int32ID}, Int32ID},
		{FuncInt32ToInt8, "int32_to_int8", 1, []OID{Int32ID}, Int8ID},
		{FuncInt32ToInt16, "int32_to_int16", 1, []OID{Int32ID}, Int16ID},
		{FuncInt32ToInt64, "int32_to_int64", 1, []OID{Int32ID}, Int64ID},
		{FuncInt32ToOID, "int32_to_oid", 1, []OID{Int32ID}, OIDID},
		{FuncInt32ToUint8, "int32_to_uint8", 1, []OID{Int32ID}, Uint8ID},
		{FuncInt32ToUint16, "int32_to_uint16", 1, []OID{Int32ID}, Uint16ID},
		{FuncInt32ToUint32, "int32_to_uint32", 1, []OID{Int32ID}, Uint32ID},
		{FuncInt32ToUint64, "int32_to_uint64", 1, []OID{Int32ID}, Uint64ID},

		/* int64 */
		{FuncInt64Input, "int64_input", 1, []OID{BytesID}, Int64ID},
		{FuncInt64Output, "int64_output", 1, []OID{Int64ID}, BytesID},
		{FuncInt64LT, "int64_lt", 2, []OID{Int64ID, Int64ID}, BooleanID},
		{FuncInt64EQ, "int64_eq", 2, []OID{Int64ID, Int64ID}, BooleanID},
		{FuncInt64GT, "int64_gt", 2, []OID{Int64ID, Int64ID}, BooleanID},
		{FuncInt64LE, "int64_le", 2, []OID{Int64ID, Int64ID}, BooleanID},
		{FuncInt64NE, "int64_ne", 2, []OID{Int64ID, Int64ID}, BooleanID},
		{FuncInt64GE, "int64_ge", 2, []OID{Int64ID, Int64ID}, BooleanID},
		{FuncInt64Add, "int64_add", 2, []OID{Int64ID, Int64ID}, Int64ID},
		{FuncInt64Sub, "int64_sub", 2, []OID{Int64ID, Int64ID}, Int64ID},
		{FuncInt64Mul, "int64_mul", 2, []OID{Int64ID, Int64ID}, Int64ID},
		{FuncInt64Div, "int64_div", 2, []OID{Int64ID, Int64ID}, Int64ID},
		{FuncInt64ToInt8, "int64_to_int8", 1, []OID{Int64ID}, Int8ID},
		{FuncInt64ToInt16, "int64_to_int16", 1, []OID{Int64ID}, Int16ID},
		{FuncInt64ToInt32, "int64_to_int32", 1, []OID{Int64ID}, Int32ID},
		{FuncInt64ToOID, "int64_to_oid", 1, []OID{Int64ID}, OIDID},
		{FuncInt64ToUint8, "int64_to_uint8", 1, []OID{Int64ID}, Uint8ID},
		{FuncInt64ToUint16, "int64_to_uint16", 1, []OID{Int64ID}, Uint16ID},
		{FuncInt64ToUint32, "int64_to_uint32", 1, []OID{Int64ID}, Uint32ID},
		{FuncInt64ToUint64, "int64_to_uint64", 1, []OID{Int64ID}, Uint64ID},

		/* oid */
		{FuncOIDInput, "oid_input", 1, []OID{BytesID}, OIDID},
		{FuncOIDOutput, "oid_output", 1, []OID{OIDID}, BytesID},
		{FuncOIDLT, "oid_lt", 2, []OID{OIDID, OIDID}, BooleanID},
		{FuncOIDEQ, "oid_eq", 2, []OID{OIDID, OIDID}, BooleanID},
		{FuncOIDGT, "oid_gt", 2, []OID{OIDID, OIDID}, BooleanID},
		{FuncOIDLE, "oid_le", 2, []OID{OIDID, OIDID}, BooleanID},
		{FuncOIDNE, "oid_ne", 2, []OID{OIDID, OIDID}, BooleanID},
		{FuncOIDGE, "oid_ge", 2, []OID{OIDID, OIDID}, BooleanID},
		{FuncOIDAdd, "oid_add", 2, []OID{OIDID, OIDID}, OIDID},
		{FuncOIDSub, "oid_sub", 2, []OID{OIDID, OIDID}, OIDID},
		{FuncOIDMul, "oid_mul", 2, []OID{OIDID, OIDID}, OIDID},
		{FuncOIDDiv, "oid_div", 2, []OID{OIDID, OIDID}, OIDID},

		/* oidarray */
		{FuncOIDArrayInput, "oidarray_input", 1, []OID{BytesID}, OIDArrayID},
		{FuncOIDArrayOutput, "oidarray_output", 1, []OID{OIDArrayID}, BytesID},

		/* text */
		{FuncTextInput, "text_input", 2, []OID{BytesID, Int32ID}, TextID},
		{FuncTextOutput, "text_output", 1, []OID{TextID}, BytesID},
		{FuncTextLT, "text_lt", 2, []OID{TextID, TextID}, BooleanID},
		{FuncTextEQ, "text_eq", 2, []OID{TextID, TextID}, BooleanID},
		{FuncTextGT, "text_gt", 2, []OID{TextID, TextID}, BooleanID},
		{FuncTextLE, "text_le", 2, []OID{TextID, TextID}, BooleanID},
		{FuncTextNE, "text_ne", 2, []OID{TextID, TextID}, BooleanID},
		{FuncTextGE, "text_ge", 2, []OID{TextID, TextID}, BooleanID},
		{FuncTextAdd, "text_add", 2, []OID{TextID, TextID}, TextID},

		/* uint8 */
		{FuncUint8Input, "uint8_input", 1, []OID{BytesID}, Uint8ID},
		{FuncUint8Output, "uint8_output", 1, []OID{Uint8ID}, BytesID},
		{FuncUint8LT, "uint8_lt", 2, []OID{Uint8ID, Uint8ID}, BooleanID},
		{FuncUint8EQ, "uint8_eq", 2, []OID{Uint8ID, Uint8ID}, BooleanID},
		{FuncUint8GT, "uint8_gt", 2, []OID{Uint8ID, Uint8ID}, BooleanID},
		{FuncUint8LE, "uint8_le", 2, []OID{Uint8ID, Uint8ID}, BooleanID},
		{FuncUint8NE, "uint8_ne", 2, []OID{Uint8ID, Uint8ID}, BooleanID},
		{FuncUint8GE, "uint8_ge", 2, []OID{Uint8ID, Uint8ID}, BooleanID},
		{FuncUint8Add, "uint8_add", 2, []OID{Uint8ID, Uint8ID}, Uint8ID},
		{FuncUint8Sub, "uint8_sub", 2, []OID{Uint8ID, Uint8ID}, Uint8ID},
		{FuncUint8Mul, "uint8_mul", 2, []OID{Uint8ID, Uint8ID}, Uint8ID},
		{FuncUint8Div, "uint8_div", 2, []OID{Uint8ID, Uint8ID}, Uint8ID},
		{FuncUint8ToInt8, "uint8_to_int8", 1, []OID{Uint8ID}, Int8ID},
		{FuncUint8ToInt16, "uint8_to_int16", 1, []OID{Uint8ID}, Int16ID},
		{FuncUint8ToInt32, "uint8_to_int32", 1, []OID{Uint8ID}, Int32ID},
		{FuncUint8ToInt64, "uint8_to_int64", 1, []OID{Uint8ID}, Int64ID},
		{FuncUint8ToOID, "uint8_to_oid", 1, []OID{Uint8ID}, OIDID},
		{FuncUint8ToUint16, "uint8_to_uint16", 1, []OID{Uint8ID}, Uint16ID},
		{FuncUint8ToUint32, "uint8_to_uint32", 1, []OID{Uint8ID}, Uint32ID},
		{FuncUint8ToUint64, "uint8_to_uint64", 1, []OID{Uint8ID}, Uint64ID},

		/* uint16 */
		{FuncUint16Input, "uint16_input", 1, []OID{BytesID}, Uint16ID},
		{FuncUint16Output, "uint16_output", 1, []OID{Uint16ID}, BytesID},
		{FuncUint16LT, "uint16_lt", 2, []OID{Uint16ID, Uint16ID}, BooleanID},
		{FuncUint16EQ, "uint16_eq", 2, []OID{Uint16ID, Uint16ID}, BooleanID},
		{FuncUint16GT, "uint16_gt", 2, []OID{Uint16ID, Uint16ID}, BooleanID},
		{FuncUint16LE, "uint16_le", 2, []OID{Uint16ID, Uint16ID}, BooleanID},
		{FuncUint16NE, "uint16_ne", 2, []OID{Uint16ID, Uint16ID}, BooleanID},
		{FuncUint16GE, "uint16_ge", 2, []OID{Uint16ID, Uint16ID}, BooleanID},
		{FuncUint16Add, "uint16_add", 2, []OID{Uint16ID, Uint16ID}, Uint16ID},
		{FuncUint16Sub, "uint16_sub", 2, []OID{Uint16ID, Uint16ID}, Uint16ID},
		{FuncUint16Mul, "uint16_mul", 2, []OID{Uint16ID, Uint16ID}, Uint16ID},
		{FuncUint16Div, "uint16_div", 2, []OID{Uint16ID, Uint16ID}, Uint16ID},
		{FuncUint16ToInt8, "uint16_to_int8", 1, []OID{Uint16ID}, Int8ID},
		{FuncUint16ToInt16, "uint16_to_int16", 1, []OID{Uint16ID}, Int16ID},
		{FuncUint16ToInt32, "uint16_to_int32", 1, []OID{Uint16ID}, Int32ID},
		{FuncUint16ToInt64, "uint16_to_int64", 1, []OID{Uint16ID}, Int64ID},
		{FuncUint16ToOID, "uint16_to_oid", 1, []OID{Uint16ID}, OIDID},
		{FuncUint16ToUint8, "uint16_to_uint16", 1, []OID{Uint16ID}, Uint8ID},
		{FuncUint16ToUint32, "uint16_to_uint32", 1, []OID{Uint16ID}, Uint32ID},
		{FuncUint16ToUint64, "uint16_to_uint64", 1, []OID{Uint16ID}, Uint64ID},

		/* uint32 */
		{FuncUint32Input, "uint32_input", 1, []OID{BytesID}, Uint32ID},
		{FuncUint32Output, "uint32_output", 1, []OID{Uint32ID}, BytesID},
		{FuncUint32LT, "uint32_lt", 2, []OID{Uint32ID, Uint32ID}, BooleanID},
		{FuncUint32EQ, "uint32_eq", 2, []OID{Uint32ID, Uint32ID}, BooleanID},
		{FuncUint32GT, "uint32_gt", 2, []OID{Uint32ID, Uint32ID}, BooleanID},
		{FuncUint32LE, "uint32_le", 2, []OID{Uint32ID, Uint32ID}, BooleanID},
		{FuncUint32NE, "uint32_ne", 2, []OID{Uint32ID, Uint32ID}, BooleanID},
		{FuncUint32GE, "uint32_ge", 2, []OID{Uint32ID, Uint32ID}, BooleanID},
		{FuncUint32Add, "uint32_add", 2, []OID{Uint32ID, Uint32ID}, Uint32ID},
		{FuncUint32Sub, "uint32_sub", 2, []OID{Uint32ID, Uint32ID}, Uint32ID},
		{FuncUint32Mul, "uint32_mul", 2, []OID{Uint32ID, Uint32ID}, Uint32ID},
		{FuncUint32Div, "uint32_div", 2, []OID{Uint32ID, Uint32ID}, Uint32ID},
		{FuncUint32ToInt8, "uint32_to_int8", 1, []OID{Uint32ID}, Int8ID},
		{FuncUint32ToInt16, "uint32_to_int16", 1, []OID{Uint32ID}, Int16ID},
		{FuncUint32ToInt32, "uint32_to_int32", 1, []OID{Uint32ID}, Int32ID},
		{FuncUint32ToInt64, "uint32_to_int64", 1, []OID{Uint32ID}, Int64ID},
		{FuncUint32ToOID, "uint32_to_oid", 1, []OID{Uint32ID}, OIDID},
		{FuncUint32ToUint8, "uint32_to_uint8", 1, []OID{Uint32ID}, Uint8ID},
		{FuncUint32ToUint16, "uint32_to_uint16", 1, []OID{Uint32ID}, Uint16ID},
		{FuncUint32ToUint64, "uint32_to_uint64", 1, []OID{Uint32ID}, Uint64ID},

		/* uint64 */
		{FuncUint64Input, "uint64_input", 1, []OID{BytesID}, Uint64ID},
		{FuncUint64Output, "uint64_output", 1, []OID{Uint64ID}, BytesID},
		{FuncUint64LT, "uint64_lt", 2, []OID{Uint64ID, Uint64ID}, BooleanID},
		{FuncUint64EQ, "uint64_eq", 2, []OID{Uint64ID, Uint64ID}, BooleanID},
		{FuncUint64GT, "uint64_gt", 2, []OID{Uint64ID, Uint64ID}, BooleanID},
		{FuncUint64LE, "uint64_le", 2, []OID{Uint64ID, Uint64ID}, BooleanID},
		{FuncUint64NE, "uint64_ne", 2, []OID{Uint64ID, Uint64ID}, BooleanID},
		{FuncUint64GE, "uint64_ge", 2, []OID{Uint64ID, Uint64ID}, BooleanID},
		{FuncUint64Add, "uint64_add", 2, []OID{Uint64ID, Uint64ID}, Uint64ID},
		{FuncUint64Sub, "uint64_sub", 2, []OID{Uint64ID, Uint64ID}, Uint64ID},
		{FuncUint64Mul, "uint64_mul", 2, []OID{Uint64ID, Uint64ID}, Uint64ID},
		{FuncUint64Div, "uint64_div", 2, []OID{Uint64ID, Uint64ID}, Uint64ID},
		{FuncUint64ToInt8, "uint64_to_int8", 1, []OID{Uint64ID}, Int8ID},
		{FuncUint64ToInt16, "uint64_to_int16", 1, []OID{Uint64ID}, Int16ID},
		{FuncUint64ToInt32, "uint64_to_int32", 1, []OID{Uint64ID}, Int32ID},
		{FuncUint64ToInt64, "uint64_to_int64", 1, []OID{Uint64ID}, Int64ID},
		{FuncUint64ToOID, "uint64_to_oid", 1, []OID{Uint64ID}, OIDID},
		{FuncUint64ToUint8, "uint64_to_uint8", 1, []OID{Uint64ID}, Uint8ID},
		{FuncUint64ToUint16, "uint64_to_uint16", 1, []OID{Uint64ID}, Uint16ID},
		{FuncUint64ToUint32, "uint64_to_uint32", 1, []OID{Uint64ID}, Uint32ID},

		/* varchar */
		{FuncVarcharInput, "varchar_input", 2, []OID{BytesID, Int32ID}, VarcharID},
		{FuncVarcharOutput, "varchar_output", 1, []OID{VarcharID}, BytesID},
		{FuncVarcharLT, "varchar_lt", 2, []OID{VarcharID, VarcharID}, BooleanID},
		{FuncVarcharEQ, "varchar_eq", 2, []OID{VarcharID, VarcharID}, BooleanID},
		{FuncVarcharGT, "varchar_gt", 2, []OID{VarcharID, VarcharID}, BooleanID},
		{FuncVarcharLE, "varchar_le", 2, []OID{VarcharID, VarcharID}, BooleanID},
		{FuncVarcharNE, "varchar_ne", 2, []OID{VarcharID, VarcharID}, BooleanID},
		{FuncVarcharGE, "varchar_ge", 2, []OID{VarcharID, VarcharID}, BooleanID},
		{FuncVarcharAdd, "varchar_add", 2, []OID{VarcharID, VarcharID}, VarcharID},
	}
}

/* builtin function id */
const (
	_ OID = iota
	FuncBooleanInput
	FuncBooleanOutput
	FuncBooleanEQ
	FuncBooleanNE
	FuncBooleanAnd
	FuncBooleanOr
	FuncBooleanNot
	FuncByteInput
	FuncByteOutput
	FuncByteLT
	FuncByteEQ
	FuncByteGT
	FuncByteLE
	FuncByteNE
	FuncByteGE
	FuncByteAdd
	FuncByteSub
	FuncByteMul
	FuncByteDiv
	FuncBytesInput
	FuncBytesOutput
	FuncCharInput
	FuncCharOutput
	FuncCharLT
	FuncCharEQ
	FuncCharGT
	FuncCharLE
	FuncCharNE
	FuncCharGE
	FuncCharAdd
	FuncFloat32Input
	FuncFloat32Output
	FuncFloat32LT
	FuncFloat32EQ
	FuncFloat32GT
	FuncFloat32LE
	FuncFloat32NE
	FuncFloat32GE
	FuncFloat32Add
	FuncFloat32Sub
	FuncFloat32Mul
	FuncFloat32Div
	FuncFloat64Input
	FuncFloat64Output
	FuncFloat64LT
	FuncFloat64EQ
	FuncFloat64GT
	FuncFloat64LE
	FuncFloat64NE
	FuncFloat64GE
	FuncFloat64Add
	FuncFloat64Sub
	FuncFloat64Mul
	FuncFloat64Div
	FuncInt8Input
	FuncInt8Output
	FuncInt8LT
	FuncInt8EQ
	FuncInt8GT
	FuncInt8LE
	FuncInt8NE
	FuncInt8GE
	FuncInt8Add
	FuncInt8Sub
	FuncInt8Mul
	FuncInt8Div
	FuncInt8ToInt16
	FuncInt8ToInt32
	FuncInt8ToInt64
	FuncInt8ToOID
	FuncInt8ToUint8
	FuncInt8ToUint16
	FuncInt8ToUint32
	FuncInt8ToUint64
	FuncInt16Input
	FuncInt16Output
	FuncInt16LT
	FuncInt16EQ
	FuncInt16GT
	FuncInt16LE
	FuncInt16NE
	FuncInt16GE
	FuncInt16Add
	FuncInt16Sub
	FuncInt16Mul
	FuncInt16Div
	FuncInt16ToInt8
	FuncInt16ToInt32
	FuncInt16ToInt64
	FuncInt16ToOID
	FuncInt16ToUint8
	FuncInt16ToUint16
	FuncInt16ToUint32
	FuncInt16ToUint64
	FuncInt32Input
	FuncInt32Output
	FuncInt32LT
	FuncInt32EQ
	FuncInt32GT
	FuncInt32LE
	FuncInt32NE
	FuncInt32GE
	FuncInt32Add
	FuncInt32Sub
	FuncInt32Mul
	FuncInt32Div
	FuncInt32ToInt8
	FuncInt32ToInt16
	FuncInt32ToInt64
	FuncInt32ToOID
	FuncInt32ToUint8
	FuncInt32ToUint16
	FuncInt32ToUint32
	FuncInt32ToUint64
	FuncInt64Input
	FuncInt64Output
	FuncInt64LT
	FuncInt64EQ
	FuncInt64GT
	FuncInt64LE
	FuncInt64NE
	FuncInt64GE
	FuncInt64Add
	FuncInt64Sub
	FuncInt64Mul
	FuncInt64Div
	FuncInt64ToInt8
	FuncInt64ToInt16
	FuncInt64ToInt32
	FuncInt64ToOID
	FuncInt64ToUint8
	FuncInt64ToUint16
	FuncInt64ToUint32
	FuncInt64ToUint64
	FuncOIDInput
	FuncOIDOutput
	FuncOIDLT
	FuncOIDEQ
	FuncOIDGT
	FuncOIDLE
	FuncOIDNE
	FuncOIDGE
	FuncOIDAdd
	FuncOIDSub
	FuncOIDMul
	FuncOIDDiv
	FuncOIDArrayInput
	FuncOIDArrayOutput
	FuncTextInput
	FuncTextOutput
	FuncTextLT
	FuncTextEQ
	FuncTextGT
	FuncTextLE
	FuncTextNE
	FuncTextGE
	FuncTextAdd
	FuncUint8Input
	FuncUint8Output
	FuncUint8LT
	FuncUint8EQ
	FuncUint8GT
	FuncUint8LE
	FuncUint8NE
	FuncUint8GE
	FuncUint8Add
	FuncUint8Sub
	FuncUint8Mul
	FuncUint8Div
	FuncUint8ToInt8
	FuncUint8ToInt16
	FuncUint8ToInt32
	FuncUint8ToInt64
	FuncUint8ToOID
	FuncUint8ToUint16
	FuncUint8ToUint32
	FuncUint8ToUint64
	FuncUint16Input
	FuncUint16Output
	FuncUint16LT
	FuncUint16EQ
	FuncUint16GT
	FuncUint16LE
	FuncUint16NE
	FuncUint16GE
	FuncUint16Add
	FuncUint16Sub
	FuncUint16Mul
	FuncUint16Div
	FuncUint16ToInt8
	FuncUint16ToInt16
	FuncUint16ToInt32
	FuncUint16ToInt64
	FuncUint16ToOID
	FuncUint16ToUint8
	FuncUint16ToUint32
	FuncUint16ToUint64
	FuncUint32Input
	FuncUint32Output
	FuncUint32LT
	FuncUint32EQ
	FuncUint32GT
	FuncUint32LE
	FuncUint32NE
	FuncUint32GE
	FuncUint32Add
	FuncUint32Sub
	FuncUint32Mul
	FuncUint32Div
	FuncUint32ToInt8
	FuncUint32ToInt16
	FuncUint32ToInt32
	FuncUint32ToInt64
	FuncUint32ToOID
	FuncUint32ToUint8
	FuncUint32ToUint16
	FuncUint32ToUint64
	FuncUint64Input
	FuncUint64Output
	FuncUint64LT
	FuncUint64EQ
	FuncUint64GT
	FuncUint64LE
	FuncUint64NE
	FuncUint64GE
	FuncUint64Add
	FuncUint64Sub
	FuncUint64Mul
	FuncUint64Div
	FuncUint64ToInt8
	FuncUint64ToInt16
	FuncUint64ToInt32
	FuncUint64ToInt64
	FuncUint64ToOID
	FuncUint64ToUint8
	FuncUint64ToUint16
	FuncUint64ToUint32
	FuncVarcharInput
	FuncVarcharOutput
	FuncVarcharLT
	FuncVarcharEQ
	FuncVarcharGT
	FuncVarcharLE
	FuncVarcharNE
	FuncVarcharGE
	FuncVarcharAdd
	numFuncID
)

func init() {
	metas := InitManaFuncs()
	h := make(map[OID]struct{})
	for _, meta := range metas {
		if _, ok := h[meta.FuncID]; ok {
			panic("duplicate metadata exists in system tables: mana_funcs")
		}
	}
	if len(metas) != int(numFuncID)-1 {
		panic("there is uninitialized metadata in system tables: mana_funcs")
	}
}
