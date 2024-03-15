package parser

import (
	"testing"
)

func isEqualTokens(tokens1 []int, tokens2 []int) bool {
	if len(tokens1) != len(tokens2) {
		return false
	}

	for i, v := range tokens1 {
		if v != tokens2[i] {
			return false
		}
	}

	return true
}

func TestCaseLex(t *testing.T) {
	tests := []struct {
		input  string
		tokens []int
	}{
		{
			input:  "CREATE TABLE TEST (A INT);",
			tokens: []int{CREATE, TABLE, IDENT, '(', IDENT, INT, ')', ';', 0},
		},
		{
			input:  "create table test (a int);",
			tokens: []int{CREATE, TABLE, IDENT, '(', IDENT, INT, ')', ';', 0},
		},
		{
			input:  "CREATE TABLE TEST (A FLOAT, B INT);",
			tokens: []int{CREATE, TABLE, IDENT, '(', IDENT, FLOAT, ',', IDENT, INT, ')', ';', 0},
		},
		{
			input:  "create table test (a float, b int);",
			tokens: []int{CREATE, TABLE, IDENT, '(', IDENT, FLOAT, ',', IDENT, INT, ')', ';', 0},
		},
		{
			input:  "DROP TABLE test;",
			tokens: []int{DROP, TABLE, IDENT, ';', 0},
		},
		{
			input:  "drop table test;",
			tokens: []int{DROP, TABLE, IDENT, ';', 0},
		},
		{
			input:  "INSERT INTO TEST VALUES (100, 3.14);",
			tokens: []int{INSERT, INTO, IDENT, VALUES, '(', INT_LITERAL, ',', FLOAT_LITERAL, ')', ';', 0},
		},
		{
			input:  "insert into test values (100, 3.14);",
			tokens: []int{INSERT, INTO, IDENT, VALUES, '(', INT_LITERAL, ',', FLOAT_LITERAL, ')', ';', 0},
		},
		{
			input:  "SELECT A, B FROM TEST;",
			tokens: []int{SELECT, IDENT, ',', IDENT, FROM, IDENT, ';', 0},
		},
		{
			input:  "select a, b from test;",
			tokens: []int{SELECT, IDENT, ',', IDENT, FROM, IDENT, ';', 0},
		},
	}

	for _, tt := range tests {
		var sym yySymType
		var tokens []int
		scan := NewScanner(tt.input)
		for {
			token := scan.Lex(&sym)
			tokens = append(tokens, token)
			if token == 0 {
				break
			}
		}
		if !isEqualTokens(tokens, tt.tokens) {
			t.Errorf("error: got = %v, expected = %v", tokens, tt.tokens)
		}
	}
}
