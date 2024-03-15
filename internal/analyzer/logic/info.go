package logic

import (
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/mana/metadata"
)

type TableInfo struct {
	/* 表的 OID, 如果这不是一个真实的表, 则为 InvalidOID */
	TableID metadata.OID

	/* 表的名字 */
	TableName string

	/* 表的所有列信息 */
	Columns []*ColumnInfo
}

type ColumnInfo struct {
	Name    string
	Mod     int16
	TypeID  metadata.OID
	TypeLen int16
}

type TargetInfo struct {
	Name string /* 投影列的名字(别名) */
	Expr expression.Expression
}

type SetInfo struct {
	ColNo int
	Expr  expression.Expression
}
