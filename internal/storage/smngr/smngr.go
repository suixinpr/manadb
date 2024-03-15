package smngr

import (
	"fmt"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/storage/page"
	"io"
	"log"
	"os"
)

func Init(path string) {
	err := os.Chdir(path)
	if err != nil {
		log.Panic(err)
	}
}

/*
 * 某个对象的存储管理器.
 */
type StorageManager struct {
	OID  metadata.OID
	path string   /* 文件路径 */
	file *os.File /* 文件描述符 */
}

/* 打开一个管理器，没有做其余多余的任何事情 */
func OpenStorageManager(oid metadata.OID) *StorageManager {
	return &StorageManager{
		OID:  oid,
		path: fmt.Sprint(oid),
		file: nil,
	}
}

/* 创建一个新的文件, 并打开它 */
func (smgr *StorageManager) Create() error {
	var err error
	smgr.file, err = os.OpenFile(smgr.path, os.O_CREATE, os.ModePerm)
	return errlog.Err(err)
}

/* 删除文件 */
func (smgr *StorageManager) Remove() error {
	err := os.Remove(smgr.path)
	return errlog.Err(err)
}

/* 打开对应路径的文件 */
func (smgr *StorageManager) Open() error {
	var err error
	smgr.file, err = os.OpenFile(smgr.path, os.O_RDWR, os.ModePerm)
	return errlog.Err(err)
}

/* 关闭文件 */
func (smgr *StorageManager) Close() error {
	err := smgr.file.Close()
	return errlog.Err(err)
}

/* 页面数 */
func (smgr *StorageManager) PageNum() (page.PageNumber, error) {
	fi, err := os.Stat(smgr.path)
	if err != nil {
		return 0, errlog.Err(err)
	}
	return page.PageNumber(fi.Size() / int64(page.PageSize)), nil
}

/* 读取文件中的某个页面 */
func (smgr *StorageManager) Read(buf []byte, pageId page.PageNumber) error {
	offset := int64(pageId) * int64(page.PageSize)
	n, err := smgr.file.ReadAt(buf, offset)

	/* 读取失败 */
	if err != nil {
		return errlog.Err(err)
	}

	/* 读取数据长度不对 */
	if n != page.PageSize {
		return errlog.Err(io.ErrUnexpectedEOF)
	}

	return nil
}

/* 将某个页面写入文件中 */
func (smgr *StorageManager) Write(buf []byte, pageId page.PageNumber) error {
	offset := int64(pageId) * int64(page.PageSize)
	n, err := smgr.file.WriteAt(buf, offset)

	/* 写入失败 */
	if err != nil {
		return errlog.Err(err)
	}

	/* 写入数据长度不对 */
	if n != page.PageSize {
		return errlog.Err(io.ErrShortWrite)
	}

	return nil
}

func (smgr *StorageManager) Extend() error {
	n, err := smgr.PageNum()
	if err != nil {
		return errlog.Err(err)
	}

	buf := page.NewPage()
	buf.Init()
	return smgr.Write(buf, n)
}
