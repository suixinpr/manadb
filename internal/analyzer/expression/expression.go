package expression

import (
	"github/suixinpr/manadb/internal/mana/datum"
	"github/suixinpr/manadb/internal/mana/metadata"
)

type Expression interface {
	/* 表达式的类型. */
	TypeID() metadata.OID

	/* 遍历所有表达式, 该函数主要设计于收集信息. */
	Walker(walk Walker)

	/* 遍历所有表达式, 该函数主要设计于新建表达式. */
	Mutator(mutate Mutator) Expression

	/* 执行表达式, 并将值作为 Datum 返回. */
	Evaluate(outer, inner []datum.Datum) (datum.Datum, error)

	/* 将表达式转化为字符串 */
	ToString() (string, error)
}

type Walker func(expr Expression)

type Mutator func(expr Expression) Expression
