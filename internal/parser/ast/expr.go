package ast

import (
	"github/suixinpr/manadb/internal/parser/types"
)

type ExprNode interface {
	Node
	exprNodeImpl()
}

var (
	_ ExprNode = &OperationExpr{}
	_ ExprNode = &ConstExpr{}
	_ ExprNode = &ColumnRefExpr{}
	_ ExprNode = &FuncCallExpr{}
)

/* Operation, 一个表达式运算 */
type OperationExpr struct {
	exprNode

	OprName string
	L       ExprNode
	R       ExprNode
}

/* Const, 常量 */
type ConstExpr struct {
	exprNode

	Type types.ConstType
	B    bool
	I    int64
	U    uint64
	F    float64
	S    string
}

func MakeConstExpr(val any) *ConstExpr {
	switch val := val.(type) {
	case nil:
		return &ConstExpr{Type: types.ConstNull}
	case bool:
		return &ConstExpr{Type: types.ConstBoolean, B: val}
	case int64:
		return &ConstExpr{Type: types.ConstInt, I: val}
	case uint64:
		return &ConstExpr{Type: types.ConstUint, U: val}
	case float64:
		return &ConstExpr{Type: types.ConstFloat, F: val}
	case string:
		return &ConstExpr{Type: types.ConstString, S: val}
	default:
		return nil
	}
}

/* ColumnRef, 引用的列 */
type ColumnRefExpr struct {
	exprNode

	TableName  string
	ColumnName string
}

/* FuncCall, 函数调用 */
type FuncCallExpr struct {
	exprNode

	FuncName string
	Args     []ExprNode
}
