package parser

import (
	"encoding/json"
	"testing"
)

func TestCaseParse(t *testing.T) {
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
		{
			input:  "SELECT TEST.A FROM TEST;",
			tokens: []int{SELECT, IDENT, ',', IDENT, FROM, IDENT, ';', 0},
		},
		{
			input:  "select test.a from test;",
			tokens: []int{SELECT, IDENT, ',', IDENT, FROM, IDENT, ';', 0},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result, _ := Parse(tt.input)
			b, _ := json.Marshal(result)
			t.Errorf(string(b))
		})
	}
}
