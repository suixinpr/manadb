package heapam

import (
	"github/suixinpr/manadb/internal/storage/buffer"
	"github/suixinpr/manadb/internal/storage/page"
	"github/suixinpr/manadb/internal/storage/smngr"
)

/* 向目标存储对象中插入一行记录 */
func InsertEntry(smgr *smngr.StorageManager, e page.Entry) error {
	n, err := smgr.PageNum()
	if err != nil {
		return err
	}

	/* 没有页面 */
	if n == 0 {
		err := smgr.Extend()
		if err != nil {
			return err
		}
		n++
	}

	buf, err := buffer.BufferPool.GetBuffer(buffer.NewPageIdentifier(smgr.OID, n-1))
	if err != nil {
		return err
	}
	buf.Mu.Lock()
	p := page.Page(buf.Data)

	if !p.CanInsertEntry(e) {
		buf.Mu.Unlock()

		err := smgr.Extend()
		if err != nil {
			return err
		}
		n++

		buf, err = buffer.BufferPool.GetBuffer(buffer.NewPageIdentifier(smgr.OID, n-1))
		if err != nil {
			return err
		}
		buf.Mu.Lock()
		p = page.Page(buf.Data)
	}

	err = p.InsertEntry(e)
	if err != nil {
		return err
	}

	buf.Mu.Unlock()
	return nil
}

/* 从目标存储对象中删除一行记录 */
func DeleteEntry(smgr *smngr.StorageManager, pos *page.EntryPos) error {
	buf, err := buffer.BufferPool.GetBuffer(buffer.NewPageIdentifier(smgr.OID, pos.PageID))
	if err != nil {
		return err
	}

	buf.Mu.Lock()
	p := page.Page(buf.Data)
	err = p.DeleteEntry(pos.ItemID)
	if err != nil {
		return err
	}

	buf.Mu.Unlock()
	return nil
}

/* 更新目标存储对象中的一行记录 */
func UpdataEntry(smgr *smngr.StorageManager, pos *page.EntryPos, e page.Entry) error {
	buf, err := buffer.BufferPool.GetBuffer(buffer.NewPageIdentifier(smgr.OID, pos.PageID))
	if err != nil {
		return err
	}

	buf.Mu.Lock()
	p := page.Page(buf.Data)
	if p.TryUpdateEntry(pos.ItemID, e) {
		buf.Mu.Unlock()
		return nil
	}

	buf.Mu.Unlock()

	/* delete old entry, insert new entry. */
	err = DeleteEntry(smgr, pos)
	if err != nil {
		return err
	}
	return InsertEntry(smgr, e)
}
