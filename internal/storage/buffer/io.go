package buffer

import (
	"fmt"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/storage/page"
	"github/suixinpr/manadb/internal/storage/smngr"
)

/* 页面标识符, 确定一个页面所在的文件和文件内位置 */
type PageIdentifier struct {
	oid    metadata.OID    /* 对象 */
	pageID page.PageNumber /* 页面编号 */
}

func NewPageIdentifier(oid metadata.OID, pageID page.PageNumber) *PageIdentifier {
	return &PageIdentifier{
		oid:    oid,
		pageID: pageID,
	}
}

func (pi *PageIdentifier) hash() string {
	return fmt.Sprint(pi)
}

func writePage(data page.Page, pi *PageIdentifier) error {
	smgr := smngr.OpenStorageManager(pi.oid)
	if err := smgr.Open(); err != nil {
		return err
	}
	if err := smgr.Write(data, pi.pageID); err != nil {
		return err
	}
	return smgr.Close()
}

func readPage(data page.Page, pi *PageIdentifier) error {
	smgr := smngr.OpenStorageManager(pi.oid)
	if err := smgr.Open(); err != nil {
		return err
	}
	if err := smgr.Read(data, pi.pageID); err != nil {
		return err
	}
	return smgr.Close()
}
