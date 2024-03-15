package ast

import "github/suixinpr/manadb/internal/parser/types"

/*
 * Node is the basic element of the AST.
 */
type Node interface {
	nodeImpl()
}

var (
	_ Node = &TypeName{}
	_ Node = &SpecTable{}
	_ Node = &TableElement{}
)

/*
 * 一个类型, 在语法分析时会碰到类型时会创建一个该对象
 */
type TypeName struct {
	node

	Name string
	Len  int
}

func MakeTypeName(name string) *TypeName {
	return &TypeName{Name: name}
}

func (node *TypeName) SetLen(n int) {
	node.Len = n
}

/*
 * From 子句中的对象, 表示一个具体的表
 */
type SpecTable struct {
	node

	TableName string
}

/*
 * From 子句中的对象, 表示一个 Join 连接
 */
type JoinTable struct {
	node

	Type  types.JoinType
	Left  Node
	Right Node
	On    ExprNode
}

type TableElement struct {
	node

	ColumnName string
	ColumnType *TypeName
}

/*
 * 表示 SELECT TableName.*, 如果未指定 TableName 则其为空串.
 */
type FieldStar struct {
	node

	TableName string
}

type FieldColumn struct {
	node

	Expr ExprNode
	Name string
}

type Assignment struct {
	node

	ColumnRef *ColumnRefExpr
	Expr      ExprNode
}

type LimitClause struct {
	node

	Count  ExprNode
	Offset ExprNode
}
