package catalog

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/access/heapam"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/storage/page"
)

type Meta struct {
	Values []datum.Datum  /* meta 解析后的值 */
	Pos    *page.EntryPos /* meta 的位置 */
}

func NewScanner(cat *Catalog, keys []*common.ScanKey) *heapam.Scanner {
	return heapam.NewScanner(cat.OID, cat.Desc, cat.Smgr, keys)
}

/* 搜索系统表, 找到一个 meta 或者返回空 */
func SearchCatalog(oid metadata.OID, keys []*common.ScanKey) (*Meta, error) {
	metas, err := SearchCatalogList(oid, keys)
	if err != nil {
		return nil, err
	}

	if len(metas) > 1 {
		return nil, errlog.New("too many catalog meta")
	}

	if metas == nil {
		return nil, nil
	}

	return metas[0], nil
}

/* 搜索系统表, 只有找到一个 meta 才成功 */
func SearchCatalogOne(oid metadata.OID, keys []*common.ScanKey) (*Meta, error) {
	metas, err := SearchCatalogList(oid, keys)
	if err != nil {
		return nil, err
	}

	if len(metas) > 1 {
		return nil, errlog.New("too many catalog meta")
	}

	if metas == nil {
		return nil, errlog.New("not find catalog meta")
	}

	return metas[0], nil
}

/* 搜索系统表, 返回所有找到的 meta */
func SearchCatalogList(oid metadata.OID, keys []*common.ScanKey) ([]*Meta, error) {
	cat, err := open(oid)
	if err != nil {
		return nil, err
	}
	defer cat.Smgr.Close()

	hes, err := searchAll(cat, keys)
	if err != nil {
		return nil, err
	}

	var metas []*Meta
	for _, he := range hes {
		metas = append(metas, &Meta{
			Values: he.Entry.GetAllColumnsValue(cat.Desc.Cols),
			Pos:    he.Pos,
		})
	}

	return metas, nil
}

/* 搜索系统表中所有符合条件的对象 */
func searchAll(cat *Catalog, keys []*common.ScanKey) ([]*heapam.HeapEntry, error) {
	sc := NewScanner(cat, keys)
	if err := sc.BeginScan(true); err != nil {
		return nil, err
	}
	defer sc.EndScan()

	var hes []*heapam.HeapEntry
	for {
		he, err := sc.GetNextEntry()
		if err != nil {
			return nil, err
		}
		if he == nil {
			break
		}
		hes = append(hes, he)
	}
	return hes, nil
}
