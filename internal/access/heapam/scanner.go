package heapam

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/storage/buffer"
	"github/suixinpr/manadb/internal/storage/page"
	"github/suixinpr/manadb/internal/storage/smngr"
	"github/suixinpr/manadb/internal/utils/fmngr"
)

/* 对下层的 Entry 进行封装后的结构, 包括了一些其他的相关信息 */
type HeapEntry struct {
	Entry page.Entry     /* entry 本身 */
	Pos   *page.EntryPos /* entry 的位置 */
}

func NewHeapEntry(e page.Entry, pageID page.PageNumber, itemID page.OffsetNumber) *HeapEntry {
	return &HeapEntry{
		Entry: e,
		Pos: &page.EntryPos{
			PageID: pageID,
			ItemID: itemID,
		},
	}
}

/* 该扫描器用于扫描堆页面中的所有 entry */
type Scanner struct {
	oid  metadata.OID
	desc *common.EntryDesc
	smgr *smngr.StorageManager
	keys []*common.ScanKey

	pageNum page.PageNumber   /* 表的页面总数 */
	pageID  page.PageNumber   /* 下一个扫描的页面号 */
	itemID  page.OffsetNumber /* 下一个扫描的 entry */
	buf     *buffer.Buffer    /* 当前持有的缓冲区页面 */

	copy bool /* 读取的所有 entry 是否需要拷贝 */
}

func NewScanner(oid metadata.OID, desc *common.EntryDesc, smgr *smngr.StorageManager, keys []*common.ScanKey) *Scanner {
	return &Scanner{
		oid:  oid,
		desc: desc,
		smgr: smgr,
		keys: keys,
	}
}

/* 扫描操作开始前, 调用该接口进行初始化操作 */
func (sc *Scanner) BeginScan(copy bool) error {
	var err error

	/* 页面总数 */
	sc.pageNum, err = sc.smgr.PageNum()
	if err != nil {
		return err
	}

	sc.copy = copy

	/* 初始化页面 */
	if sc.pageID >= sc.pageNum {
		return nil
	}
	return sc.moveNextPage()
}

/* 扫描操作结束后, 调用该接口进行清理操作 */
func (sc *Scanner) EndScan() {
	if sc.buf != nil {
		sc.buf.Release()
	}
}

/* 获取下一个 entry */
func (sc *Scanner) GetNextEntry() (*HeapEntry, error) {
	for {
		/* 扫描结束 */
		if sc.buf == nil {
			return nil, nil
		}

		/* 扫描该页的所有 entry */
		sc.buf.Mu.RLock()
		p := page.Page(sc.buf.Data)
		max := p.GetMaxItemID()
		for sc.itemID < max {
			item := p.GetItem(sc.itemID)
			sc.itemID++
			if !item.IsLive() {
				continue
			}

			e := p.GetEntry(item)
			if sc.copy {
				new := make(page.Entry, len(e))
				copy(new, e)
				e = new
			}

			res, err := sc.check(e)
			if err != nil {
				return nil, err
			}
			if !res {
				continue
			}

			sc.buf.Mu.RUnlock()
			return NewHeapEntry(e, sc.pageID-1, sc.itemID-1), nil
		}
		sc.buf.Mu.RUnlock()

		/* 所有页面扫描结束 */
		if sc.pageID >= sc.pageNum {
			sc.buf.Release()
			sc.buf = nil
			return nil, nil
		}

		/* 移动至下一页 */
		if err := sc.moveNextPage(); err != nil {
			return nil, err
		}
	}
}

/* 移动至下一个页面, 这里会对 buffer 加引用数 */
func (sc *Scanner) moveNextPage() error {
	buf, err := buffer.BufferPool.GetBuffer(buffer.NewPageIdentifier(sc.oid, sc.pageID))
	if err != nil {
		return err
	}
	sc.buf = buf
	sc.pageID++
	return nil
}

/* 检查 entry 是否符合条件 */
func (sc *Scanner) check(e page.Entry) (bool, error) {
	for _, sk := range sc.keys {
		value := e.GetColumnValue(sc.desc.Cols, sk.ColNo)
		res, err := fmngr.Call(sk.CmpFuncID, value, sk.Argument)
		if err != nil || !datum.DatumGetValue[bool](res) {
			return false, err
		}
	}
	return true, nil
}
