package page

import (
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/utils/misc"
	"hash/fnv"
	"unsafe"
)

/*
 * The structure of the page is as follows
 *
 * +------------------------------------------+-------+-------+
 * |                pageHeader                | item1 | item2 |
 * +-------+----------------------------------+-------+-------+
 * | item3 |                                                  |
 * +-------+--------------------------------------------------+
 * |                                                          |
 * +---------+------------+------------+------------+---------+
 * |         |   entry3   |   entry2   |   entry1   | special |
 * +---------+------------+------------+------------+---------+
 */
var PageSize int = 1 << 13

type (
	/* PageNumber 是页面号 */
	PageNumber uint32

	/* OffsetNumber 标记页面内数据的偏移量, 最大值为 64KB */
	OffsetNumber uint16

	/* Page */
	Page []byte

	/* Page 头部信息 */
	PageHeader struct {
		Checksum uint64

		Lower   OffsetNumber
		Upper   OffsetNumber
		Special OffsetNumber

		Flags uint16
	}

	/* 表示页面中的一个项, 指向一个页面内的 entry */
	Item struct {
		/*
		 * 指向的 entry 的位置, 由于 entry 8 字节
		 * 对齐，所以取最后 3 个 bit 表示 entry 的状态
		 */
		off OffsetNumber

		/* 指向的 entry 的长度 */
		len OffsetNumber
	}

	EntryPos struct {
		PageID PageNumber
		ItemID OffsetNumber
	}
)

const (
	/* Page Header 的长度 */
	PageHeaderSize = int(unsafe.Sizeof(PageHeader{}))

	/* Item 的大小 */
	itemSize = int(unsafe.Sizeof(Item{}))

	offMask  = 0xfff8
	flagMask = 0x0007

	/* entry 的状态, EntryPtrData 的 off 的后 3 位表示 */
	isUnused = 0
	isLive   = 1
	isDead   = 2
)

/* 初始化一个新的没有任何数据的 page */
func NewPage() Page {
	p := make(Page, PageSize)
	return p
}

/* 初始化一个新的没有任何数据的 page */
func (p *Page) Init() {
	h := p.GetHeader()
	h.Lower = OffsetNumber(PageHeaderSize)
	h.Upper = OffsetNumber(PageSize)
	h.Special = OffsetNumber(PageSize)
	h.Flags = 0
	p.CalculateChecksum()
}

/* 返回页面的头部信息 */
func (p Page) GetHeader() *PageHeader {
	return (*PageHeader)(unsafe.Pointer(&p[0]))
}

func (p Page) GetMaxItemID() OffsetNumber {
	lower := int(p.GetHeader().Lower)
	if lower <= PageHeaderSize {
		return 0
	}
	return OffsetNumber((lower - PageHeaderSize) / itemSize)
}

func (p Page) GetItem(id OffsetNumber) Item {
	return *(*Item)(unsafe.Add(unsafe.Pointer(&p[0]), PageHeaderSize+int(id)*itemSize))
}

func (p Page) updateItem(id OffsetNumber, item Item) {
	*(*Item)(unsafe.Add(unsafe.Pointer(&p[0]), PageHeaderSize+int(id)*itemSize)) = item
}

/* 获取 off 处的 entry */
func (p Page) GetEntry(item Item) Entry {
	off := item.off & offMask
	len := item.len
	return Entry(p[off : off+len])
}

/* 计算校验和 */
func (p Page) CalculateChecksum() {
	f := fnv.New64()
	f.Write(p[8:])
	p.GetHeader().Checksum = f.Sum64()
}

/* 检查校验和 */
func (p Page) VerifyChecksum() bool {
	f := fnv.New64()
	f.Write(p[8:])
	return p.GetHeader().Checksum == f.Sum64()
}

/* 返回页面剩余的空间大小 */
func (p Page) FreeSpace() int {
	h := p.GetHeader()
	return int(h.Upper) - int(h.Lower)
}

/* 页面是否可以插入 entry */
func (p Page) CanInsertEntry(e Entry) bool {
	return len(e)+itemSize <= p.FreeSpace()
}

/* 在页面中新插入一个 entry */
func (p Page) InsertEntry(e Entry) error {
	h := p.GetHeader()

	/* 判断 entry 能否放入页面 */
	if !p.CanInsertEntry(e) {
		return errlog.New("err")
	}

	/* 放入 entry */
	sz := len(e)
	off := misc.AlignDown(int(h.Upper)-sz, 8)
	copy(p[off:], e)

	/* 放入 entryPtr*/
	*(*Item)(unsafe.Pointer(&p[h.Lower])) = Item{OffsetNumber(off) | isLive, OffsetNumber(sz)}

	h.Lower += OffsetNumber(itemSize)
	h.Upper = OffsetNumber(off)

	return nil
}

func (p Page) DeleteEntry(itemID OffsetNumber) error {
	if p.GetMaxItemID() < itemID {
		return errlog.New("err")
	}

	item := p.GetItem(itemID)
	p.updateItem(itemID, Item{item.off&offMask | isDead, item.len})
	return nil
}

func (p Page) TryUpdateEntry(itemID OffsetNumber, new Entry) bool {
	item := p.GetItem(itemID)
	old := p.GetEntry(item)
	if len(new) > len(old) {
		return false
	}
	copy(old, new)
	return true
}

func (item Item) IsLive() bool {
	return item.off&flagMask == isLive
}
