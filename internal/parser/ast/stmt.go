package ast

type StmtNode interface {
	Node
	stmtNodeImpl()
}

var (
	_ StmtNode = &CreateTableStmt{}
	_ StmtNode = &DeleteStmt{}
	_ StmtNode = &DropTableStmt{}
	_ StmtNode = &InsertStmt{}
	_ StmtNode = &SelectStmt{}
	_ StmtNode = &UpdateStmt{}
)

type CreateTableStmt struct {
	stmtNode

	TableName        string
	TableElementList []*TableElement
}

type DeleteStmt struct {
	stmtNode

	From        *SpecTable
	WhereClause ExprNode
}

type DropTableStmt struct {
	stmtNode

	Tables []*SpecTable
}

type ExplainStmt struct {
	stmtNode

	Stmt StmtNode
}

type InsertStmt struct {
	stmtNode

	Table   *SpecTable
	Columns []Node

	Select     *SelectStmt
	ValuesList [][]ExprNode
}

type UpdateStmt struct {
	stmtNode

	Table       *SpecTable
	SetList     []*Assignment
	WhereClause ExprNode
}

type SelectStmt struct {
	stmtNode

	Fields      []Node
	FromClause  Node
	WhereClause ExprNode
	LimitClause *LimitClause
}
