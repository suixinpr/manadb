package ast

/*
 * node is the struct implements Node interface.
 * Node implementations should embed it in.
 *
 * nodeImpl implements Node interface.
 */
type node struct{}

func (n *node) nodeImpl() {}

/*
 * stmtNode implements StmtNode interface.
 * Statement implementations should embed it in.
 *
 * stmtNodeImpl implements StmtNode interface.
 */
type stmtNode struct {
	node
}

func (n *stmtNode) stmtNodeImpl() {}

/*
 * exprNode is the struct implements ExprNode interface.
 * Expression implementations should embed it in.
 *
 * exprNodeImpl implements ExprNode interface.
 */
type exprNode struct {
	node
}

func (n *exprNode) exprNodeImpl() {}
