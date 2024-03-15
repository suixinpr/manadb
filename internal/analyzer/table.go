package analyzer

import (
	"github/suixinpr/manadb/internal/access/common"
	"github/suixinpr/manadb/internal/analyzer/logic"
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
	"github/suixinpr/manadb/internal/parser/ast"
	"github/suixinpr/manadb/internal/relation/catalog"
)

type parseNamespace struct {
	table *logic.TableInfo
	index int
}

func (ana *Analyzer) transformSpecTable(n *ast.SpecTable) (*logic.TableObject, error) {
	table := new(logic.TableInfo)

	/* 获取表 oid */
	meta, err := catalog.SearchCatalogOne(metadata.ManaTablesID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaTablesTblName, metadata.TextID, datum.StringGetDatum(n.TableName)),
	})
	if err != nil {
		return nil, err
	}
	table.TableID = datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaTablesTblID])
	table.TableName = datum.DatumGetString(meta.Values[metadata.ManaTablesTblName])

	/* 获取表的所有列 */
	metas, err := catalog.SearchCatalogList(metadata.ManaColumnsID, []*common.ScanKey{
		catalog.NewScanKey(metadata.ManaColumnsColTblID, metadata.OIDID, datum.ValueGetDatum(table.TableID)),
	})
	if err != nil {
		return nil, err
	}
	for _, meta := range metas {
		table.Columns = append(table.Columns, &logic.ColumnInfo{
			Name:    datum.DatumGetString(meta.Values[metadata.ManaColumnsColName]),
			TypeID:  datum.DatumGetValue[metadata.OID](meta.Values[metadata.ManaColumnsColTypeID]),
			TypeLen: datum.DatumGetValue[int16](meta.Values[metadata.ManaColumnsColTypeLen]),
		})
	}

	idx := len(ana.tables)
	ana.namespace = append(ana.namespace, &parseNamespace{
		table: table,
		index: idx,
	})
	ana.tables = append(ana.tables, table)
	return &logic.TableObject{Index: idx}, nil
}

func (ana *Analyzer) transformJoinTable(n *ast.JoinTable) (*logic.JoinObject, error) {
	left, err := ana.transformFromClauseItem(n.Left)
	if err != nil {
		return nil, err
	}

	right, err := ana.transformFromClauseItem(n.Right)
	if err != nil {
		return nil, err
	}

	on, err := ana.transformExprNode(n.On)
	if err != nil {
		return nil, err
	}

	table := new(logic.TableInfo)
	table.Columns = append(table.Columns, ana.tables[left.ObjIndex()].Columns...)
	table.Columns = append(table.Columns, ana.tables[right.ObjIndex()].Columns...)

	ana.tables = append(ana.tables, table)
	return &logic.JoinObject{
		Index: len(ana.tables) - 1,
		Type:  n.Type,
		Left:  left,
		Right: right,
		On:    on,
	}, nil
}
