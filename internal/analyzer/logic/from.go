package logic

import (
	"github/suixinpr/manadb/internal/analyzer/expression"
	"github/suixinpr/manadb/internal/parser/types"
)

type FromObject interface {
	ObjIndex() int
}

type TableObject struct {
	Index int
}

func (obj *TableObject) ObjIndex() int {
	return obj.Index
}

type JoinObject struct {
	Index int
	Type  types.JoinType
	Left  FromObject
	Right FromObject
	On    expression.Expression /* join ... on 条件 */
}

func (obj *JoinObject) ObjIndex() int {
	return obj.Index
}
