%{
/*
 * cd .\internal\parser\
 * goyacc -o gram.go gram.y
 */
package parser

import (
	"github/suixinpr/manadb/internal/parser/ast"
	"github/suixinpr/manadb/internal/parser/types"
)
%}

%union {
	ident string
	item  any
	node  ast.Node
	expr  ast.ExprNode
	stmt  ast.StmtNode
}

/* keywords */
%token <ident>
	/* identifier */
	IDENT  /* identifier */

	/* special token  */
	ERROR  /* error */

	/* ReservedKeyword */
	AND         /* and */
	AS          /* as */
	CREATE      /* create */
	CROSS       /* cross */
	DELETE      /* delete */
	DROP        /* drop */
	EXPLAIN     /* explain */
	FALSE       /* false */
	FROM        /* from */
	FULL        /* full */
	INNER       /* inner */
	INSERT      /* insert */
	INTO        /* into */
	JOIN        /* join */
	LEFT        /* left */
	LIMIT       /* limit */
	NATURAL     /* natural */
	NOT         /* not */
	NULL        /* null */
	OFFSET      /* offset */
	ON          /* on */
	OR          /* or */
	OUTER       /* outer */
	RIGHT       /* right */
	SELECT      /* select */
	SET         /* set */
	UPDATE      /* update */
	TABLE       /* table */
	TRUE        /* true */
	VALUES      /* values */
	WHERE       /* where */

	/* UnReservedKeyword */
	BOOLEAN     /* boolean */
	BYTE        /* type */
	CHAR        /* char */
	DOUBLE      /* double */
	FLOAT       /* float */
	FLOAT32     /* float32 */
	FLOAT64     /* float64 */
	INT         /* int */
	INT8        /* int8 */
	INT16       /* int16 */
	INT32       /* int32 */
	INT64       /* int64 */
	TEXT        /* text */
	UINT8       /* uint8 */
	UINT16      /* uint16 */
	UINT32      /* uint32 */
	UINT64      /* uint64 */
	VARCHAR     /* varchar */
	UNKNOWN     /* unknown */

/* literal */
%token <item>
	FLOAT_LITERAL    /* floating-point literal */
	INT_LITERAL      /* integer literal */
	STRING_LITERAL   /* string literal */

/* multi-character operators */
%token
	LE      /* <= */
	GE      /* >= */
	NE      /* <> != */

%type<stmt>
	Stmt
	CreateTableStmt
	DeleteStmt
	DropTableStmt
	ExplainableStmt
	ExplainStmt
	InsertStmt
	SelectStmt
	UpdateStmt

%type<expr>
	ColumnRef
	Const
	Expression
	FuncCall
	SimpleExpr
	WhereClauseOpt

%type<node>
	Assignment
	BuiltinType
	ColumnType
	Field
	FromClauseOpt
	JoinTable
	LimitClauseOpt
	SpecTable
	TableElement
	TableRef
	TableRefs

%type<item>
	AssignmentList
	ExpressionList
	ExpressionListOpt
	FieldList
	InsertColumns
	InsertColumnsOpt
	InsertValues
	JoinType
	SpecTables
	TableElementList
	TableElementListOpt
	TypeLen
	ValuesList

%type<ident>
	ColumnName
	FunctionName
	Identifier
	OuterOpt
	TableName
	UnReservedKeyword

/*
 * 操作符优先级
 */
%right '('
%left ')'
%left OR
%left AND
%left '=' '>' '<' LE GE NE
%left '+' '-'
%left '*' '/'
%right NOT

/*
 * JOIN 相关优先级
 */
%left JOIN CROSS INNER LEFT RIGHT FULL NATURAL

%%

/* 解析入口 */
Top:
	Stmt ';'
	{
		yylex.(*Scanner).res = $1
	}

Stmt:
	CreateTableStmt
|	DeleteStmt
|	DropTableStmt
|	ExplainStmt
|	InsertStmt
|	SelectStmt
|	UpdateStmt

/********************************************************************************
 *
 *  CREATE TABLE
 *
 ********************************************************************************/

CreateTableStmt:
	CREATE TABLE TableName TableElementListOpt
	{
		$$ = &ast.CreateTableStmt{
			TableName: $3,
			TableElementList: $4.([]*ast.TableElement),
		}
	}

TableElementListOpt:
	'(' TableElementList ')'
	{
		$$ = $2
	}

TableElementList:
	TableElement
	{
		$$ = []*ast.TableElement{$1.(*ast.TableElement)}
	}
|	TableElementList ',' TableElement
	{
		$$ = append($1.([]*ast.TableElement), $3.(*ast.TableElement))
	}

TableElement:
	ColumnName ColumnType
	{
		$$ = &ast.TableElement{
			ColumnName: $1,
			ColumnType: $2.(*ast.TypeName),
		}
	}

ColumnType:
	BuiltinType

/********************************************************************************
 *
 *  DELETE
 *
 ********************************************************************************/

DeleteStmt:
	DELETE FROM SpecTable WhereClauseOpt
	{
		$$ = &ast.DeleteStmt{
			From: $3.(*ast.SpecTable),
			WhereClause: $4,
		}
	}

/********************************************************************************
 *
 *  DROP TABLE
 *
 ********************************************************************************/

DropTableStmt:
	DROP TABLE SpecTables
	{
		$$ = &ast.DropTableStmt{
			Tables: $3.([]*ast.SpecTable),
		}
	}

/********************************************************************************
 *
 *  EXPLAIN
 *
 ********************************************************************************/

ExplainStmt:
	EXPLAIN ExplainableStmt
	{
		$$ = &ast.ExplainStmt{
			Stmt: $2,
		}
	}

ExplainableStmt:
	DeleteStmt
|	InsertStmt
|	SelectStmt
|	UpdateStmt

/********************************************************************************
 *
 *  INSERT
 *
 ********************************************************************************/

InsertStmt:
	INSERT INTO SpecTable InsertColumnsOpt InsertValues
	{
		stmt := $5.(*ast.InsertStmt)
		stmt.Table = $3.(*ast.SpecTable)
		if ($4 != nil) {
			stmt.Columns = $4.([]ast.Node)
		}
		$$ = stmt
	}

InsertColumnsOpt:
	/* empty */
	{
		$$ = nil
	}
|	'(' InsertColumns ')'
	{
		$$ = $2
	}

InsertColumns:
	ColumnRef
	{
		$$ = []ast.Node{$1.(ast.Node)}
	}
|	InsertColumns ',' ColumnRef
	{
		$$ = append($1.([]ast.Node), $3)
	}

InsertValues:
	SelectStmt
	{
		$$ = &ast.InsertStmt{Select: $1.(*ast.SelectStmt)}
	}
|	ValuesList
	{
		$$ = &ast.InsertStmt{ValuesList: $1.([][]ast.ExprNode)}
	}

ValuesList:
	VALUES '(' ExpressionList ')'
	{
		$$ = [][]ast.ExprNode{$3.([]ast.ExprNode)}
	}
|	ValuesList ',' '(' ExpressionList ')'
	{
		$$ = append($1.([][]ast.ExprNode), $4.([]ast.ExprNode))
	}

/********************************************************************************
 *
 *  SELECT
 *
 ********************************************************************************/

SelectStmt:
	SELECT FieldList FromClauseOpt WhereClauseOpt LimitClauseOpt
	{
		stmt := &ast.SelectStmt{
			Fields: $2.([]ast.Node),
			FromClause: $3,
			WhereClause: $4,
		}
		if $5 != nil {
			stmt.LimitClause = $5.(*ast.LimitClause)
		}
		$$ = stmt
	}

FieldList:
	Field
	{
		$$ = []ast.Node{$1.(ast.Node)}
	}
|	FieldList ',' Field
	{
		$$ = append($1.([]ast.Node), $3.(ast.Node))
	}

Field:
	'*'
	{
		$$ = &ast.FieldStar{}
	}
|	TableName '.' '*'
	{
		$$ = &ast.FieldStar{TableName: $1}
	}
|	Expression
	{
		$$ = &ast.FieldColumn{Expr: $1}
	}
|	Expression AS ColumnName
	{
		$$ = &ast.FieldColumn{Expr: $1, Name: $3}
	}

/********************************************************************************
 *
 *  UPDATE
 *
 ********************************************************************************/

UpdateStmt:
	UPDATE SpecTable SET AssignmentList FromClauseOpt WhereClauseOpt
	{
		stmt := &ast.UpdateStmt{
			Table: $2.(*ast.SpecTable),
			SetList: $4.([]*ast.Assignment),
			WhereClause: $6,
		}
		$$ = stmt
	}

AssignmentList:
	Assignment
	{
		$$ = []*ast.Assignment{$1.(*ast.Assignment)}
	}
|	AssignmentList ',' Assignment
	{
		$$ = append($1.([]*ast.Assignment), $3.(*ast.Assignment))
	}

Assignment:
	ColumnRef '=' Expression
	{
		$$ = &ast.Assignment{ColumnRef: $1.(*ast.ColumnRefExpr), Expr: $3}
	}

/********************************************************************************
 *
 *  Clause
 *
 ********************************************************************************/

FromClauseOpt:
	/* empty */
	{
		$$ = nil
	}
|	FROM TableRefs
	{
		$$ = $2
	}

WhereClauseOpt:
	/* empty */
	{
		$$ = nil
	}
|	WHERE Expression
	{
		$$ = $2
	}

LimitClauseOpt:
	/* empty */
	{
		$$ = nil
	}
|	LIMIT Expression
	{
		$$ = &ast.LimitClause{
			Count: $2,
		}
	}
|	LIMIT Expression ',' Expression
	{
		$$ = &ast.LimitClause{
			Count: $4,
			Offset: $2,
		}
	}
|	LIMIT Expression OFFSET Expression
	{
		$$ = &ast.LimitClause{
			Count: $2,
			Offset: $4,
		}
	}

TableRefs:
	TableRef
	{
		$$ = $1
	}
|	TableRefs ',' TableRef
	{
		$$ = &ast.JoinTable{
			Left: $1,
			Right: $3,
		}
	}

SpecTables:
	SpecTable
	{
		$$ = []*ast.SpecTable{$1.(*ast.SpecTable)}
	}
|	SpecTables ',' SpecTable
	{
		$$ = append($1.([]*ast.SpecTable), $3.(*ast.SpecTable))
	}

/*
 * TableRef 包含了所有的逻辑上的表结构
 */
TableRef:
	SpecTable
	{
		$$ = $1
	}
|	JoinTable
	{
		$$ = $1
	}

/*
 * SpecTable 包含了所有的系统表中存在的表
 */
SpecTable:
	TableName
	{
		$$ = &ast.SpecTable{TableName: $1}
	}

/*
 * JoinTable 包含了所有 JOIN 后的表
 */
JoinTable:
	TableRef CROSS JOIN TableRef
	{
		$$ = &ast.JoinTable{Type: types.CrossJoin, Left: $1, Right: $4}
	}
|	TableRef NATURAL JOIN TableRef
	{
		$$ = &ast.JoinTable{Type: types.NaturalJoin, Left: $1, Right: $4}
	}
|	TableRef JoinType JOIN TableRef ON Expression
	{
		$$ = &ast.JoinTable{Type: $2.(types.JoinType), Left: $1, Right: $4, On: $6}
	}

JoinType:
	INNER
	{
		$$ = types.InnerJoin
	}
|	LEFT OuterOpt
	{
		$$ = types.LeftJoin
	}
|	RIGHT OuterOpt
	{
		$$ = types.RightJoin
	}
|	FULL OuterOpt
	{
		$$ = types.FullJoin
	}

OuterOpt:
	/* empty */
	{}
|	OUTER

/********************************************************************************
 *
 *  Identifier
 *
 ********************************************************************************/

/* 标识符 */
Identifier:
	IDENT
|	UnReservedKeyword

UnReservedKeyword:
	BOOLEAN
|	BYTE
|	CHAR
|	DOUBLE
|	FLOAT
|	FLOAT32
|	FLOAT64
|	INT
|	INT8
|	INT16
|	INT32
|	INT64
|	TEXT
|	UINT8
|	UINT16
|	UINT32
|	UINT64
|	VARCHAR
|	UNKNOWN

/********************************************************************************
 *
 *  Expr
 *
 *  语义规则的书写从上往下优先级递增
 *
 ********************************************************************************/

/* 表达式 */
Expression:
	SimpleExpr
|	Expression '+' Expression
	{
		$$ = &ast.OperationExpr{OprName: "+", L: $1, R: $3}
	}
|	Expression '-' Expression
	{
		$$ = &ast.OperationExpr{OprName: "-", L: $1, R: $3}
	}
|	Expression '*' Expression
	{
		$$ = &ast.OperationExpr{OprName: "*", L: $1, R: $3}
	}
|	Expression '/' Expression
	{
		$$ = &ast.OperationExpr{OprName: "/", L: $1, R: $3}
	}
|	Expression AND Expression
	{
		$$ = &ast.OperationExpr{OprName: "and", L: $1, R: $3}
	}
|	Expression OR Expression
	{
		$$ = &ast.OperationExpr{OprName: "or", L: $1, R: $3}
	}
|	Expression '<' Expression
	{
		$$ = &ast.OperationExpr{OprName: "<", L: $1, R: $3}
	}
|	Expression '=' Expression
	{
		$$ = &ast.OperationExpr{OprName: "=", L: $1, R: $3}
	}
|	Expression '>' Expression
	{
		$$ = &ast.OperationExpr{OprName: ">", L: $1, R: $3}
	}
|	Expression LE Expression
	{
		$$ = &ast.OperationExpr{OprName: "<=", L: $1, R: $3}
	}
|	Expression GE Expression
	{
		$$ = &ast.OperationExpr{OprName: ">=", L: $1, R: $3}
	}
|	Expression NE Expression
	{
		$$ = &ast.OperationExpr{OprName: "<>", L: $1, R: $3}
	}

/* 简单表达式, 不存在二元及以上运算 */
SimpleExpr:
	Const
|	ColumnRef
|	FuncCall
|	'+' SimpleExpr
	{
		$$ = &ast.OperationExpr{OprName: "+", L: nil, R: $2}
	}
|	'-' SimpleExpr
	{
		$$ = &ast.OperationExpr{OprName: "-", L: nil, R: $2}
	}
|	'*' SimpleExpr
	{
		$$ = &ast.OperationExpr{OprName: "*", L: nil, R: $2}
	}
|	'/' SimpleExpr
	{
		$$ = &ast.OperationExpr{OprName: "/", L: nil, R: $2}
	}
|	NOT SimpleExpr
	{
		$$ = &ast.OperationExpr{OprName: "not", L: nil, R: $2}
	}
|	'(' Expression ')'
	{
		$$ = $2
	}

Const:
	NULL
	{
		$$ = ast.MakeConstExpr(nil)
	}
|	TRUE
	{
		$$ = ast.MakeConstExpr(true)
	}
|	FALSE
	{
		$$ = ast.MakeConstExpr(false)
	}
|	FLOAT_LITERAL
	{
		$$ = ast.MakeConstExpr($1)
	}
|	INT_LITERAL
	{
		$$ = ast.MakeConstExpr($1)
	}
|	STRING_LITERAL
	{
		$$ = ast.MakeConstExpr($1)
	}

/* 表达式中的列名 */
ColumnRef:
	ColumnName
	{
		$$ = &ast.ColumnRefExpr{
			TableName: "",
			ColumnName: $1,
		}
	}
|	TableName '.' ColumnName
	{
		$$ = &ast.ColumnRefExpr{
			TableName: $1,
			ColumnName: $3,
		}
	}

FuncCall:
	FunctionName '(' ExpressionListOpt ')'
	{
		$$ = &ast.FuncCallExpr{
			FuncName: $1,
			Args: $3.([]ast.ExprNode),
		}
	}

ExpressionListOpt:
	/* empty */
	{
		$$ = []ast.ExprNode{}
	}
|	ExpressionList

/* 表达式列表 */
ExpressionList:
	Expression
	{
		$$ = []ast.ExprNode{$1}
	}
|	ExpressionList ',' Expression
	{
		$$ = append($1.([]ast.ExprNode), $3)
	}

/* 表名 */
TableName:
	Identifier

/* 列名 */
ColumnName:
	Identifier

/* 函数名 */
FunctionName:
	Identifier

/********************************************************************************
 *
 *  Type
 *
 ********************************************************************************/

BuiltinType:
	BOOLEAN
	{
		$$ = ast.MakeTypeName("boolean")
	}
|	BYTE
	{
		$$ = ast.MakeTypeName("byte")
	}
|	CHAR TypeLen
	{
		tp := ast.MakeTypeName("char")
		tp.SetLen($2.(int))
		$$ = tp
	}
|	DOUBLE
	{
		$$ = ast.MakeTypeName("float64")
	}
|	FLOAT
	{
		$$ = ast.MakeTypeName("float32")
	}
|	FLOAT32
	{
		$$ = ast.MakeTypeName("float32")
	}
|	FLOAT64
	{
		$$ = ast.MakeTypeName("float64")
	}
|	INT
	{
		$$ = ast.MakeTypeName("int32")
	}
|	INT8
	{
		$$ = ast.MakeTypeName("int8")
	}
|	INT16
	{
		$$ = ast.MakeTypeName("int16")
	}
|	INT32
	{
		$$ = ast.MakeTypeName("int32")
	}
|	INT64
	{
		$$ = ast.MakeTypeName("int64")
	}
|	TEXT
	{
		$$ = ast.MakeTypeName("text")
	}
|	UINT8
	{
		$$ = ast.MakeTypeName("uint8")
	}
|	UINT16
	{
		$$ = ast.MakeTypeName("uint16")
	}
|	UINT32
	{
		$$ = ast.MakeTypeName("uint32")
	}
|	UINT64
	{
		$$ = ast.MakeTypeName("uint64")
	}
|	VARCHAR TypeLen
	{
		tp := ast.MakeTypeName("varchar")
		tp.SetLen($2.(int))
		$$ = tp
	}

TypeLen:
	'(' INT_LITERAL ')'
	{
		$$ = int($2.(uint64))
	}

%%