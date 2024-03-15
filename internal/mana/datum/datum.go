package datum

import "unsafe"

/*
 * Datum 包含所有数据库内置的数据对象.
 *
 * 这些数据对象不能直接被使用, 只能通过下面的接口来
 * 与 Golang 的类型进行转化, 或通过内置函数来进行操作.
 *
 * Datum 对应的 Input 操作是把数据从文本格式变成二进制格式.
 * Datum 对应的 Output 操作是把数据从二进制格式变成文本格式.
 * 数据库系统内的数据都是二进制格式表示的.
 * 不同数据之间相互转换时需要通过文本格式作为中间格式.
 *
 * Datum{} 表示空值, 如空串. 定义为 var d Datum = Datum{}
 * nil 表示 NULL. 定义为 var d Datum = nil
 * 前者在进行 d == nil 时结果为 false, 后者结果为 true
 * 空切片和 nil 切片的区别
 */
type Datum []byte

func (d Datum) Copy() Datum {
	if d == nil {
		return nil
	}
	res := make(Datum, len(d))
	copy(res, d)
	return res
}

func EmtryDatum() Datum {
	return Datum{}
}

func NullDatum() Datum {
	return nil
}

func IsNull(datum Datum) bool {
	return datum == nil
}

/*
 * Datum 中的值类型, 不包括切片和字符串等引用类型
 */
type Value interface {
	~bool | ~float32 | ~float64 | ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

/*
 * 将 Datum 转化为对应的 Golang 的值类型
 */
func DatumGetValue[T Value](data Datum) T {
	return *(*T)(unsafe.Pointer(&data[0]))
}

/*
 * 将 Datum 转化为对应的 Golang 的值切片类型
 */
func DatumGetValueArray[T Value](data Datum) []T {
	var x T
	return unsafe.Slice((*T)(unsafe.Pointer(&data[0])), len(data)/int(unsafe.Sizeof(x)))
}

/*
 * 将 Datum 转化为对应的 Golang 的 []byte 类型
 */
func DatumGetBytes(data Datum) []byte {
	return []byte(data)
}

/*
 * 将 Datum 转化为对应的 Golang 的 string 类型
 */
func DatumGetString(data Datum) string {
	return unsafe.String(unsafe.SliceData(data), len(data))
}

/*
 * 将 Golang 的值类型转化为 Datum 类型
 */
func ValueGetDatum[T Value](value T) Datum {
	return unsafe.Slice((*byte)(unsafe.Pointer(&value)), unsafe.Sizeof(value))
}

/*
 * 将 Golang 的值切片类型转化为 Datum 类型
 */
func ValueArrayGetDatum[T Value](value []T) Datum {
	return unsafe.Slice((*byte)(unsafe.Pointer(&value[0])), int(unsafe.Sizeof(value[0]))*len(value))
}

/*
 * 将 Golang 的 []byte 类型转化为 Datum 类型
 */
func BytesGetDatum(value []byte) Datum {
	return Datum(value)
}

/*
 * 将 Golang 的 string 类型转化为 Datum 类型
 */
func StringGetDatum(s string) Datum {
	if s == "" {
		/* 是空串, 不是 NULL */
		return EmtryDatum()
	}
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
