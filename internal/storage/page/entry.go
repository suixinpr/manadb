package page

import (
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/utils/misc"
	"unsafe"
)

type (
	/*
	 * 表中数据的 Entry, 结构如下:
	 * +---------------+------------+------------------------------------+
	 * |  EntryHeader  | NullBitmap |            filedsDatum             |
	 * +---------------+------------+------------------------------------+
	 * 每个 Entry 的头部都是 8 字节对齐的.
	 *
	 * 在 filedsDatum 中存储着所有的字段信息, 它们按照内存进行对齐.
	 * 如果是变长字段，则前两个字节表示长度.
	 */
	Entry []byte

	EntryHeader struct {
		FiledNum uint16 /* 字段数据的数量 */
		Flags    uint16
	}
)

const (
	/* Entry Hearder 的长度*/
	EntryHeaderSize = int(unsafe.Sizeof(EntryHeader{}))
)

/* EntryHeader 的 Flags 中的标志位 */
const (
	EntryHasNull = 0x0001
)

/* 获取 entry 的头部 */
func (e Entry) GetHeader() *EntryHeader {
	return (*EntryHeader)(unsafe.Pointer(&e[0]))
}

/* 形成并返回 entry */
func NewEntry(cols []*metadata.ManaColumns, values []datum.Datum) Entry {
	flags := uint16(0)
	filedNum := len(values)

	/* 计算 entry 中各部分的长度 */
	dataSize := 0
	for i := 0; i < filedNum; i++ {
		value := values[i]
		if value == nil {
			flags |= EntryHasNull
			continue
		}
		col := cols[i]
		sz := int(col.ColTypeLen)
		if sz == -1 {
			dataSize = misc.AlignUp(dataSize, 2) + 2 /* 需要对齐变长类型长度的 2 个字节*/
			sz = len(value)
		}
		dataSize = misc.AlignUp(dataSize, int(col.ColTypeAlign))
		dataSize += sz
	}

	/* bitmap 长度*/
	bitmapSize := 0
	if flags&EntryHasNull != 0 {
		bitmapSize = misc.AlignUp(filedNum, 8) / 8
	}

	totalSize := misc.AlignUp(EntryHeaderSize+bitmapSize, 8) + misc.AlignUp(dataSize, 8)

	/* 填充 entry 各部分信息 */
	entry := make(Entry, totalSize)
	h := entry.GetHeader()
	h.FiledNum = uint16(filedNum)
	h.Flags = flags

	if flags&EntryHasNull != 0 {
		bitmap := entry[EntryHeaderSize : EntryHeaderSize+bitmapSize]
		for i, value := range values {
			if value == nil {
				bitmap[i/8] |= (1 << (i % 8))
			}
		}
	}

	pos := EntryHeaderSize + bitmapSize
	for i := 0; i < filedNum; i++ {
		value := values[i]
		if value == nil {
			continue
		}
		col := cols[i]
		sz := int(col.ColTypeLen)
		if sz == -1 {
			pos = misc.AlignUp(pos, 2)
			sz = len(value)
			*(*uint16)(unsafe.Pointer(&entry[pos])) = uint16(sz)
			pos += 2
		}
		pos = misc.AlignUp(pos, int(col.ColTypeAlign))
		copy(entry[pos:], value)
		pos += sz
	}

	return entry
}

/* 如果目标列为空, 则返回值为 nil */
func (e Entry) GetColumnValue(cols []*metadata.ManaColumns, no int16) datum.Datum {
	filedNum := int(e.GetHeader().FiledNum)
	if len(cols) != filedNum {
		panic("columns num is not equal to field num")
	}

	var pos int
	var bitmap Entry
	if e.GetHeader().Flags&EntryHasNull != 0 {
		bitmapSize := misc.AlignUp(filedNum, 8) / 8
		pos = EntryHeaderSize + bitmapSize
		bitmap = e[EntryHeaderSize:pos]
	} else {
		pos = EntryHeaderSize
		bitmap = nil
	}

	/* 目标列为空 */
	if (bitmap != nil) && (bitmap[no/8]&(1<<(no%8)) != 0) {
		return nil
	}

	sz := 0
	for i, col := range cols {
		sz = int(col.ColTypeLen)
		if (bitmap != nil) && (bitmap[i/8]&(1<<(i%8)) != 0) {
			continue
		} else if sz == -1 {
			pos = misc.AlignUp(pos, 2)
			sz = int(*(*uint16)(unsafe.Pointer(&e[pos])))
			pos += 2
		}
		pos = misc.AlignUp(pos, int(col.ColTypeAlign))
		if i == int(no) {
			break
		}
		pos += sz
	}

	return datum.Datum(e[pos : pos+sz])
}

func (e Entry) GetAllColumnsValue(cols []*metadata.ManaColumns) []datum.Datum {
	filedNum := int(e.GetHeader().FiledNum)
	if len(cols) != filedNum {
		panic("columns num is not equal to field num")
	}

	var pos int
	var bitmap Entry
	if e.GetHeader().Flags&EntryHasNull != 0 {
		bitmapSize := misc.AlignUp(filedNum, 8) / 8
		pos = EntryHeaderSize + bitmapSize
		bitmap = e[EntryHeaderSize:pos]
	} else {
		pos = EntryHeaderSize
		bitmap = nil
	}

	values := make([]datum.Datum, filedNum)
	for i, col := range cols {
		sz := int(col.ColTypeLen)
		isnull := (bitmap != nil) && (bitmap[i/8]&(1<<(i%8)) != 0)
		if isnull {
			sz = 0
		} else if sz == -1 {
			pos = misc.AlignUp(pos, 2)
			sz = int(*(*uint16)(unsafe.Pointer(&e[pos])))
			pos += 2
		}
		pos = misc.AlignUp(pos, int(col.ColTypeAlign))
		if isnull {
			values[i] = nil
		} else {
			values[i] = datum.Datum(e[pos : pos+sz])
		}
		pos += sz
	}

	return values
}
