package parser

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func isIdentHead(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_' || ch >= utf8.MaxRune
}

func isIdentChar(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch >= utf8.MaxRune
}

/*
 * 使用字典树来存储字符的匹配规则
 */
type trieNode struct {
	childs [127]*trieNode
	token  int
}

var charTable trieNode

func init() {
	/* single character*/
	initTokenByte('+', int('+'))
	initTokenByte('-', int('-'))
	initTokenByte('*', int('*'))
	initTokenByte('/', int('/'))
	initTokenByte('(', int('('))
	initTokenByte(')', int(')'))
	initTokenByte('[', int('['))
	initTokenByte(']', int(']'))
	initTokenByte(';', int(';'))
	initTokenByte(',', int(','))
	initTokenByte('.', int('.'))
	initTokenByte('>', int('>'))
	initTokenByte('<', int('<'))
	initTokenByte('=', int('='))

	/* multi character*/
	initTokenString("<=", LE)
	initTokenString(">=", GE)
	initTokenString("<>", NE)
	initTokenString("!=", NE)
}

func initTokenByte(ch rune, tok int) {
	if charTable.childs[ch] == nil {
		charTable.childs[ch] = &trieNode{}
	}
	node := charTable.childs[ch]
	if node.token != 0 {
		panic(fmt.Sprintf("repeated char rule: %v", ch))
	}
	node.token = tok
}

func initTokenString(str string, tok int) {
	node := &charTable
	for _, ch := range str {
		if node.childs[ch] == nil {
			node.childs[ch] = &trieNode{}
		}
		node = node.childs[ch]
	}
	if node.token != 0 {
		panic(fmt.Sprintf("repeated char rule: %v", str))
	}
	node.token = tok
}

/* ReservedKeyword and UnReservedKeyword */
var keywords = map[string]int{
	"AND":     AND,
	"AS":      AS,
	"BOOLEAN": BOOLEAN,
	"BYTE":    BYTE,
	"CHAR":    CHAR,
	"CREATE":  CREATE,
	"CROSS":   CROSS,
	"DELETE":  DELETE,
	"DOUBLE":  DOUBLE,
	"DROP":    DROP,
	"EXPLAIN": EXPLAIN,
	"FALSE":   FALSE,
	"FLOAT":   FLOAT,
	"FLOAT32": FLOAT32,
	"FLOAT64": FLOAT64,
	"FROM":    FROM,
	"FULL":    FULL,
	"INNER":   INNER,
	"INSERT":  INSERT,
	"INT":     INT,
	"INT8":    INT8,
	"INT16":   INT16,
	"INT32":   INT32,
	"INT64":   INT64,
	"INTO":    INTO,
	"JOIN":    JOIN,
	"LEFT":    LEFT,
	"LIMIT":   LIMIT,
	"NATURAL": NATURAL,
	"NOT":     NOT,
	"NULL":    NULL,
	"OFFSET":  OFFSET,
	"ON":      ON,
	"OR":      OR,
	"OUTER":   OUTER,
	"RIGHT":   RIGHT,
	"SELECT":  SELECT,
	"SET":     SET,
	"TABLE":   TABLE,
	"TEXT":    TEXT,
	"TRUE":    TRUE,
	"UINT8":   UINT8,
	"UINT16":  UINT16,
	"UINT32":  UINT32,
	"UINT64":  UINT64,
	"UPDATE":  UPDATE,
	"UNKNOWN": UNKNOWN,
	"VALUES":  VALUES,
	"VARCHAR": VARCHAR,
	"WHERE":   WHERE,
}
