// Code generated by goyacc -o gram.go gram.y. DO NOT EDIT.

//line gram.y:2
/*
 * cd .\internal\parser\
 * goyacc -o gram.go gram.y
 */
package parser

import __yyfmt__ "fmt"

//line gram.y:6

import (
	"github/suixinpr/manadb/internal/parser/ast"
	"github/suixinpr/manadb/internal/parser/types"
)

//line gram.y:14
type yySymType struct {
	yys   int
	ident string
	item  any
	node  ast.Node
	expr  ast.ExprNode
	stmt  ast.StmtNode
}

const IDENT = 57346
const ERROR = 57347
const AND = 57348
const AS = 57349
const CREATE = 57350
const CROSS = 57351
const DELETE = 57352
const DROP = 57353
const EXPLAIN = 57354
const FALSE = 57355
const FROM = 57356
const FULL = 57357
const INNER = 57358
const INSERT = 57359
const INTO = 57360
const JOIN = 57361
const LEFT = 57362
const LIMIT = 57363
const NATURAL = 57364
const NOT = 57365
const NULL = 57366
const OFFSET = 57367
const ON = 57368
const OR = 57369
const OUTER = 57370
const RIGHT = 57371
const SELECT = 57372
const SET = 57373
const UPDATE = 57374
const TABLE = 57375
const TRUE = 57376
const VALUES = 57377
const WHERE = 57378
const BOOLEAN = 57379
const BYTE = 57380
const CHAR = 57381
const DOUBLE = 57382
const FLOAT = 57383
const FLOAT32 = 57384
const FLOAT64 = 57385
const INT = 57386
const INT8 = 57387
const INT16 = 57388
const INT32 = 57389
const INT64 = 57390
const TEXT = 57391
const UINT8 = 57392
const UINT16 = 57393
const UINT32 = 57394
const UINT64 = 57395
const VARCHAR = 57396
const UNKNOWN = 57397
const FLOAT_LITERAL = 57398
const INT_LITERAL = 57399
const STRING_LITERAL = 57400
const LE = 57401
const GE = 57402
const NE = 57403

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENT",
	"ERROR",
	"AND",
	"AS",
	"CREATE",
	"CROSS",
	"DELETE",
	"DROP",
	"EXPLAIN",
	"FALSE",
	"FROM",
	"FULL",
	"INNER",
	"INSERT",
	"INTO",
	"JOIN",
	"LEFT",
	"LIMIT",
	"NATURAL",
	"NOT",
	"NULL",
	"OFFSET",
	"ON",
	"OR",
	"OUTER",
	"RIGHT",
	"SELECT",
	"SET",
	"UPDATE",
	"TABLE",
	"TRUE",
	"VALUES",
	"WHERE",
	"BOOLEAN",
	"BYTE",
	"CHAR",
	"DOUBLE",
	"FLOAT",
	"FLOAT32",
	"FLOAT64",
	"INT",
	"INT8",
	"INT16",
	"INT32",
	"INT64",
	"TEXT",
	"UINT8",
	"UINT16",
	"UINT32",
	"UINT64",
	"VARCHAR",
	"UNKNOWN",
	"FLOAT_LITERAL",
	"INT_LITERAL",
	"STRING_LITERAL",
	"LE",
	"GE",
	"NE",
	"'('",
	"')'",
	"'='",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"';'",
	"','",
	"'.'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line gram.y:847

//line yacctab:1
var yyExca = [...]int16{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 32,
	62, 124,
	73, 122,
	-2, 123,
	-1, 143,
	73, 122,
	-2, 123,
}

const yyPrivate = 57344

const yyLast = 485

var yyAct = [...]uint8{
	139, 138, 84, 37, 210, 116, 141, 145, 79, 201,
	119, 85, 8, 192, 69, 91, 31, 166, 30, 72,
	32, 74, 72, 72, 33, 108, 24, 81, 117, 72,
	28, 157, 110, 228, 223, 17, 92, 73, 169, 73,
	73, 73, 166, 166, 103, 71, 81, 73, 75, 77,
	227, 96, 97, 98, 82, 78, 94, 95, 93, 87,
	88, 89, 90, 165, 99, 100, 101, 102, 96, 97,
	98, 194, 136, 94, 95, 93, 87, 88, 89, 90,
	195, 31, 213, 30, 72, 168, 89, 90, 124, 125,
	126, 127, 128, 129, 130, 131, 132, 133, 134, 135,
	121, 123, 73, 170, 80, 113, 122, 122, 82, 142,
	147, 114, 171, 72, 211, 193, 154, 87, 88, 89,
	90, 112, 146, 107, 150, 104, 143, 221, 122, 109,
	15, 73, 20, 143, 121, 152, 18, 105, 202, 148,
	122, 158, 10, 156, 11, 12, 13, 164, 161, 167,
	200, 14, 162, 26, 159, 199, 198, 196, 226, 19,
	72, 163, 11, 197, 15, 1, 16, 205, 35, 14,
	208, 9, 142, 203, 204, 207, 91, 86, 73, 209,
	7, 70, 15, 151, 16, 25, 146, 106, 144, 143,
	76, 160, 122, 206, 23, 214, 212, 92, 4, 215,
	149, 72, 72, 72, 218, 219, 220, 111, 153, 27,
	137, 140, 22, 34, 115, 222, 143, 224, 225, 73,
	73, 73, 65, 155, 118, 172, 173, 229, 38, 96,
	97, 98, 42, 63, 94, 95, 93, 87, 88, 89,
	90, 36, 6, 64, 21, 5, 44, 45, 46, 47,
	48, 49, 50, 51, 52, 53, 54, 55, 56, 57,
	58, 59, 60, 61, 62, 66, 67, 68, 3, 2,
	0, 43, 34, 0, 0, 0, 39, 40, 83, 41,
	0, 65, 0, 158, 0, 0, 91, 0, 0, 164,
	161, 42, 63, 0, 162, 0, 159, 0, 0, 0,
	0, 0, 64, 163, 0, 44, 45, 46, 47, 48,
	49, 50, 51, 52, 53, 54, 55, 56, 57, 58,
	59, 60, 61, 62, 66, 67, 68, 34, 0, 0,
	43, 0, 0, 0, 0, 39, 40, 29, 41, 96,
	97, 98, 0, 0, 94, 95, 93, 87, 88, 89,
	90, 0, 0, 0, 0, 0, 0, 0, 0, 91,
	44, 45, 46, 47, 48, 49, 50, 51, 52, 53,
	54, 55, 56, 57, 58, 59, 60, 61, 62, 91,
	92, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 120, 0, 0, 0, 0, 0, 217, 0,
	92, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 96, 97, 98, 34, 0, 94, 95, 93,
	87, 88, 89, 90, 0, 0, 0, 0, 0, 0,
	0, 0, 96, 97, 98, 0, 0, 94, 95, 93,
	87, 88, 89, 90, 0, 216, 0, 0, 44, 45,
	46, 47, 48, 49, 50, 51, 52, 53, 54, 55,
	56, 57, 58, 59, 60, 61, 62, 174, 175, 176,
	177, 178, 179, 180, 181, 182, 183, 184, 185, 186,
	187, 188, 189, 190, 191,
}

var yyPact = [...]int16{
	134, -1000, -36, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	103, 145, 99, 152, 135, 268, 411, -1000, 411, 411,
	411, -1000, -1000, -1000, -1000, -1000, 411, 32, -1000, 209,
	-62, 170, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 209,
	209, 209, 209, 209, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	63, 106, -1000, -1000, 61, 93, -40, -1000, 59, 93,
	268, 411, -1000, 209, -63, 323, 411, 209, 209, 209,
	209, 209, 209, 209, 209, 209, 209, 209, 209, -1000,
	-1000, -1000, -1000, 9, 209, 411, -1000, 411, -1000, 209,
	411, 100, 411, 122, -1000, -41, 274, -1000, -1000, 411,
	-1000, -1000, -1000, -1000, 17, 17, -1000, -1000, -8, 280,
	50, 50, 50, 50, 50, 50, -1000, 0, -55, 353,
	13, -1000, -26, -1000, 40, -1000, 430, 353, -1000, -1000,
	-1000, -59, 53, 8, -1000, -1000, 209, 411, 137, 136,
	131, -1000, 110, 110, 110, -1000, 209, 93, 411, 209,
	-1000, 411, -1000, -1000, -1000, -1000, 52, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 52, 20, 209, -1000, 411, 373, 274, 411, 411,
	411, -1000, -1000, -1000, -1000, 353, -1000, -1000, 353, -1000,
	-1000, 70, -1000, 209, -29, -1000, 209, 209, -1000, -1000,
	132, -13, -30, -1000, 353, 353, 209, -1000, -1000, 353,
}

var yyPgo = [...]int16{
	0, 269, 268, 198, 245, 244, 242, 180, 12, 171,
	3, 241, 0, 228, 24, 25, 6, 226, 225, 30,
	8, 224, 223, 28, 7, 5, 214, 211, 1, 210,
	209, 208, 207, 200, 191, 190, 188, 187, 4, 183,
	14, 181, 20, 9, 2, 168, 165,
}

var yyR1 = [...]int8{
	0, 46, 1, 1, 1, 1, 1, 1, 1, 2,
	37, 36, 36, 24, 18, 3, 4, 6, 5, 5,
	5, 5, 7, 32, 32, 31, 31, 33, 33, 39,
	39, 8, 30, 30, 19, 19, 19, 19, 9, 27,
	27, 16, 20, 20, 15, 15, 22, 22, 22, 22,
	26, 26, 35, 35, 25, 25, 23, 21, 21, 21,
	34, 34, 34, 34, 43, 43, 42, 42, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	45, 45, 45, 45, 45, 45, 45, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 11,
	11, 11, 11, 11, 11, 10, 10, 13, 29, 29,
	28, 28, 44, 40, 41, 17, 17, 17, 17, 17,
	17, 17, 17, 17, 17, 17, 17, 17, 17, 17,
	17, 17, 17, 38,
}

var yyR2 = [...]int8{
	0, 2, 1, 1, 1, 1, 1, 1, 1, 4,
	3, 1, 3, 2, 1, 4, 3, 2, 1, 1,
	1, 1, 5, 0, 3, 1, 3, 1, 1, 4,
	5, 5, 1, 3, 1, 3, 1, 3, 6, 1,
	3, 3, 0, 2, 0, 2, 0, 2, 4, 4,
	1, 3, 1, 3, 1, 1, 1, 4, 4, 6,
	1, 2, 2, 2, 0, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	1, 1, 1, 2, 2, 2, 2, 2, 3, 1,
	1, 1, 1, 1, 1, 1, 3, 4, 0, 1,
	1, 3, 1, 1, 1, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 3,
}

var yyChk = [...]int16{
	-1000, -46, -1, -2, -3, -4, -6, -7, -8, -9,
	8, 10, 11, 12, 17, 30, 32, 71, 33, 14,
	33, -5, -3, -7, -8, -9, 18, -30, -19, 69,
	-44, -12, -42, -14, 4, -45, -11, -10, -13, 67,
	68, 70, 23, 62, 37, 38, 39, 40, 41, 42,
	43, 44, 45, 46, 47, 48, 49, 50, 51, 52,
	53, 54, 55, 24, 34, 13, 56, 57, 58, -40,
	-41, -23, -44, -42, -44, -23, -35, -23, -23, -20,
	72, 14, -14, 69, -44, 73, 7, 67, 68, 69,
	70, 6, 27, 66, 64, 65, 59, 60, 61, -14,
	-14, -14, -14, -12, 62, 31, -37, 62, -15, 36,
	72, -32, 62, -15, -19, -26, -25, -23, -21, 73,
	69, -40, -42, -40, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, 63, -29, -28, -12,
	-27, -16, -10, -42, -36, -24, -40, -12, -23, -33,
	-8, -39, 35, -31, -10, -22, 21, 72, 9, 22,
	-34, 16, 20, 29, 15, 63, 72, -20, 72, 64,
	63, 72, -18, -17, 37, 38, 39, 40, 41, 42,
	43, 44, 45, 46, 47, 48, 49, 50, 51, 52,
	53, 54, 72, 62, 63, 72, -12, -25, 19, 19,
	19, -43, 28, -43, -43, -12, -15, -16, -12, -24,
	-38, 62, -38, 62, -28, -10, 72, 25, -25, -25,
	-25, 57, -28, 63, -12, -12, 26, 63, 63, -12,
}

var yyDef = [...]int16{
	0, -2, 0, 2, 3, 4, 5, 6, 7, 8,
	0, 0, 0, 0, 0, 0, 0, 1, 0, 0,
	0, 17, 18, 19, 20, 21, 0, 42, 32, 34,
	0, 36, -2, 87, 66, 67, 100, 101, 102, 0,
	0, 0, 0, 0, 68, 69, 70, 71, 72, 73,
	74, 75, 76, 77, 78, 79, 80, 81, 82, 83,
	84, 85, 86, 109, 110, 111, 112, 113, 114, 115,
	0, 0, 56, 122, 0, 44, 16, 52, 23, 44,
	0, 0, 105, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 103,
	104, 106, 107, 0, 118, 0, 9, 0, 15, 0,
	0, 0, 0, 46, 33, 43, 50, 54, 55, 0,
	35, 116, 123, 37, 88, 89, 90, 91, 92, 93,
	94, 95, 96, 97, 98, 99, 108, 0, 119, 120,
	42, 39, 0, -2, 0, 11, 0, 45, 53, 22,
	27, 28, 0, 0, 25, 31, 0, 0, 0, 0,
	0, 60, 64, 64, 64, 117, 0, 44, 0, 0,
	10, 0, 13, 14, 125, 126, 0, 128, 129, 130,
	131, 132, 133, 134, 135, 136, 137, 138, 139, 140,
	141, 0, 0, 0, 24, 0, 47, 51, 0, 0,
	0, 61, 65, 62, 63, 121, 38, 40, 41, 12,
	127, 0, 142, 0, 0, 26, 0, 0, 57, 58,
	0, 0, 0, 29, 48, 49, 0, 143, 30, 59,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	62, 63, 69, 67, 72, 68, 73, 70, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 71,
	66, 64, 65,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:173
		{
			yylex.(*Scanner).res = yyDollar[1].stmt
		}
	case 9:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:194
		{
			yyVAL.stmt = &ast.CreateTableStmt{
				TableName:        yyDollar[3].ident,
				TableElementList: yyDollar[4].item.([]*ast.TableElement),
			}
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:203
		{
			yyVAL.item = yyDollar[2].item
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:209
		{
			yyVAL.item = []*ast.TableElement{yyDollar[1].node.(*ast.TableElement)}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:213
		{
			yyVAL.item = append(yyDollar[1].item.([]*ast.TableElement), yyDollar[3].node.(*ast.TableElement))
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:219
		{
			yyVAL.node = &ast.TableElement{
				ColumnName: yyDollar[1].ident,
				ColumnType: yyDollar[2].node.(*ast.TypeName),
			}
		}
	case 15:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:237
		{
			yyVAL.stmt = &ast.DeleteStmt{
				From:        yyDollar[3].node.(*ast.SpecTable),
				WhereClause: yyDollar[4].expr,
			}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:252
		{
			yyVAL.stmt = &ast.DropTableStmt{
				Tables: yyDollar[3].item.([]*ast.SpecTable),
			}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:266
		{
			yyVAL.stmt = &ast.ExplainStmt{
				Stmt: yyDollar[2].stmt,
			}
		}
	case 22:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:286
		{
			stmt := yyDollar[5].item.(*ast.InsertStmt)
			stmt.Table = yyDollar[3].node.(*ast.SpecTable)
			if yyDollar[4].item != nil {
				stmt.Columns = yyDollar[4].item.([]ast.Node)
			}
			yyVAL.stmt = stmt
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:297
		{
			yyVAL.item = nil
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:301
		{
			yyVAL.item = yyDollar[2].item
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:307
		{
			yyVAL.item = []ast.Node{yyDollar[1].expr.(ast.Node)}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:311
		{
			yyVAL.item = append(yyDollar[1].item.([]ast.Node), yyDollar[3].expr)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:317
		{
			yyVAL.item = &ast.InsertStmt{Select: yyDollar[1].stmt.(*ast.SelectStmt)}
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:321
		{
			yyVAL.item = &ast.InsertStmt{ValuesList: yyDollar[1].item.([][]ast.ExprNode)}
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:327
		{
			yyVAL.item = [][]ast.ExprNode{yyDollar[3].item.([]ast.ExprNode)}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:331
		{
			yyVAL.item = append(yyDollar[1].item.([][]ast.ExprNode), yyDollar[4].item.([]ast.ExprNode))
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
//line gram.y:343
		{
			stmt := &ast.SelectStmt{
				Fields:      yyDollar[2].item.([]ast.Node),
				FromClause:  yyDollar[3].node,
				WhereClause: yyDollar[4].expr,
			}
			if yyDollar[5].node != nil {
				stmt.LimitClause = yyDollar[5].node.(*ast.LimitClause)
			}
			yyVAL.stmt = stmt
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:357
		{
			yyVAL.item = []ast.Node{yyDollar[1].node.(ast.Node)}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:361
		{
			yyVAL.item = append(yyDollar[1].item.([]ast.Node), yyDollar[3].node.(ast.Node))
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:367
		{
			yyVAL.node = &ast.FieldStar{}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:371
		{
			yyVAL.node = &ast.FieldStar{TableName: yyDollar[1].ident}
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:375
		{
			yyVAL.node = &ast.FieldColumn{Expr: yyDollar[1].expr}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:379
		{
			yyVAL.node = &ast.FieldColumn{Expr: yyDollar[1].expr, Name: yyDollar[3].ident}
		}
	case 38:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:391
		{
			stmt := &ast.UpdateStmt{
				Table:       yyDollar[2].node.(*ast.SpecTable),
				SetList:     yyDollar[4].item.([]*ast.Assignment),
				WhereClause: yyDollar[6].expr,
			}
			yyVAL.stmt = stmt
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:402
		{
			yyVAL.item = []*ast.Assignment{yyDollar[1].node.(*ast.Assignment)}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:406
		{
			yyVAL.item = append(yyDollar[1].item.([]*ast.Assignment), yyDollar[3].node.(*ast.Assignment))
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:412
		{
			yyVAL.node = &ast.Assignment{ColumnRef: yyDollar[1].expr.(*ast.ColumnRefExpr), Expr: yyDollar[3].expr}
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:424
		{
			yyVAL.node = nil
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:428
		{
			yyVAL.node = yyDollar[2].node
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:434
		{
			yyVAL.expr = nil
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:438
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:444
		{
			yyVAL.node = nil
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:448
		{
			yyVAL.node = &ast.LimitClause{
				Count: yyDollar[2].expr,
			}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:454
		{
			yyVAL.node = &ast.LimitClause{
				Count:  yyDollar[4].expr,
				Offset: yyDollar[2].expr,
			}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:461
		{
			yyVAL.node = &ast.LimitClause{
				Count:  yyDollar[2].expr,
				Offset: yyDollar[4].expr,
			}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:470
		{
			yyVAL.node = yyDollar[1].node
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:474
		{
			yyVAL.node = &ast.JoinTable{
				Left:  yyDollar[1].node,
				Right: yyDollar[3].node,
			}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:483
		{
			yyVAL.item = []*ast.SpecTable{yyDollar[1].node.(*ast.SpecTable)}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:487
		{
			yyVAL.item = append(yyDollar[1].item.([]*ast.SpecTable), yyDollar[3].node.(*ast.SpecTable))
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:496
		{
			yyVAL.node = yyDollar[1].node
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:500
		{
			yyVAL.node = yyDollar[1].node
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:509
		{
			yyVAL.node = &ast.SpecTable{TableName: yyDollar[1].ident}
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:518
		{
			yyVAL.node = &ast.JoinTable{Type: types.CrossJoin, Left: yyDollar[1].node, Right: yyDollar[4].node}
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:522
		{
			yyVAL.node = &ast.JoinTable{Type: types.NaturalJoin, Left: yyDollar[1].node, Right: yyDollar[4].node}
		}
	case 59:
		yyDollar = yyS[yypt-6 : yypt+1]
//line gram.y:526
		{
			yyVAL.node = &ast.JoinTable{Type: yyDollar[2].item.(types.JoinType), Left: yyDollar[1].node, Right: yyDollar[4].node, On: yyDollar[6].expr}
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:532
		{
			yyVAL.item = types.InnerJoin
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:536
		{
			yyVAL.item = types.LeftJoin
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:540
		{
			yyVAL.item = types.RightJoin
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:544
		{
			yyVAL.item = types.FullJoin
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:550
		{
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:597
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "+", L: yyDollar[1].expr, R: yyDollar[3].expr}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:601
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "-", L: yyDollar[1].expr, R: yyDollar[3].expr}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:605
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "*", L: yyDollar[1].expr, R: yyDollar[3].expr}
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:609
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "/", L: yyDollar[1].expr, R: yyDollar[3].expr}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:613
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "and", L: yyDollar[1].expr, R: yyDollar[3].expr}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:617
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "or", L: yyDollar[1].expr, R: yyDollar[3].expr}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:621
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "<", L: yyDollar[1].expr, R: yyDollar[3].expr}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:625
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "=", L: yyDollar[1].expr, R: yyDollar[3].expr}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:629
		{
			yyVAL.expr = &ast.OperationExpr{OprName: ">", L: yyDollar[1].expr, R: yyDollar[3].expr}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:633
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "<=", L: yyDollar[1].expr, R: yyDollar[3].expr}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:637
		{
			yyVAL.expr = &ast.OperationExpr{OprName: ">=", L: yyDollar[1].expr, R: yyDollar[3].expr}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:641
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "<>", L: yyDollar[1].expr, R: yyDollar[3].expr}
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:651
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "+", L: nil, R: yyDollar[2].expr}
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:655
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "-", L: nil, R: yyDollar[2].expr}
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:659
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "*", L: nil, R: yyDollar[2].expr}
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:663
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "/", L: nil, R: yyDollar[2].expr}
		}
	case 107:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:667
		{
			yyVAL.expr = &ast.OperationExpr{OprName: "not", L: nil, R: yyDollar[2].expr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:671
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:677
		{
			yyVAL.expr = ast.MakeConstExpr(nil)
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:681
		{
			yyVAL.expr = ast.MakeConstExpr(true)
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:685
		{
			yyVAL.expr = ast.MakeConstExpr(false)
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:689
		{
			yyVAL.expr = ast.MakeConstExpr(yyDollar[1].item)
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:693
		{
			yyVAL.expr = ast.MakeConstExpr(yyDollar[1].item)
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:697
		{
			yyVAL.expr = ast.MakeConstExpr(yyDollar[1].item)
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:704
		{
			yyVAL.expr = &ast.ColumnRefExpr{
				TableName:  "",
				ColumnName: yyDollar[1].ident,
			}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:711
		{
			yyVAL.expr = &ast.ColumnRefExpr{
				TableName:  yyDollar[1].ident,
				ColumnName: yyDollar[3].ident,
			}
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
//line gram.y:720
		{
			yyVAL.expr = &ast.FuncCallExpr{
				FuncName: yyDollar[1].ident,
				Args:     yyDollar[3].item.([]ast.ExprNode),
			}
		}
	case 118:
		yyDollar = yyS[yypt-0 : yypt+1]
//line gram.y:729
		{
			yyVAL.item = []ast.ExprNode{}
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:737
		{
			yyVAL.item = []ast.ExprNode{yyDollar[1].expr}
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:741
		{
			yyVAL.item = append(yyDollar[1].item.([]ast.ExprNode), yyDollar[3].expr)
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:765
		{
			yyVAL.node = ast.MakeTypeName("boolean")
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:769
		{
			yyVAL.node = ast.MakeTypeName("byte")
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:773
		{
			tp := ast.MakeTypeName("char")
			tp.SetLen(yyDollar[2].item.(int))
			yyVAL.node = tp
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:779
		{
			yyVAL.node = ast.MakeTypeName("float64")
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:783
		{
			yyVAL.node = ast.MakeTypeName("float32")
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:787
		{
			yyVAL.node = ast.MakeTypeName("float32")
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:791
		{
			yyVAL.node = ast.MakeTypeName("float64")
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:795
		{
			yyVAL.node = ast.MakeTypeName("int32")
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:799
		{
			yyVAL.node = ast.MakeTypeName("int8")
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:803
		{
			yyVAL.node = ast.MakeTypeName("int16")
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:807
		{
			yyVAL.node = ast.MakeTypeName("int32")
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:811
		{
			yyVAL.node = ast.MakeTypeName("int64")
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:815
		{
			yyVAL.node = ast.MakeTypeName("text")
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:819
		{
			yyVAL.node = ast.MakeTypeName("uint8")
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:823
		{
			yyVAL.node = ast.MakeTypeName("uint16")
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:827
		{
			yyVAL.node = ast.MakeTypeName("uint32")
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
//line gram.y:831
		{
			yyVAL.node = ast.MakeTypeName("uint64")
		}
	case 142:
		yyDollar = yyS[yypt-2 : yypt+1]
//line gram.y:835
		{
			tp := ast.MakeTypeName("varchar")
			tp.SetLen(yyDollar[2].item.(int))
			yyVAL.node = tp
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
//line gram.y:843
		{
			yyVAL.item = int(yyDollar[2].item.(uint64))
		}
	}
	goto yystack /* stack new state and value */
}
