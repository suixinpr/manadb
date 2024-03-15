package parser

import "github/suixinpr/manadb/internal/parser/ast"

func Parse(input string) (ast.StmtNode, error) {
	scan := NewScanner(input)
	yyParse(scan)
	return scan.res, scan.err
}
