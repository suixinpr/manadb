package parser

import (
	"fmt"
	"github/suixinpr/manadb/internal/mana/errlog"
	"github/suixinpr/manadb/internal/parser/ast"
	"math"
	"strconv"
	"strings"
	"unicode"
)

type Scanner struct {
	r *reader

	pos Pos
	err error
	res ast.StmtNode
}

func NewScanner(s string) *Scanner {
	r := &reader{
		s: s,
		l: len(s),
	}
	return &Scanner{r: r}
}

func (s *Scanner) Error(str string) {
	val := s.r.s[s.pos.Offset:]
	s.err = fmt.Errorf("line %d column %d near %s, %s",
		s.r.p.Line, s.r.p.Col, val, str)
}

func (s *Scanner) Lex(v *yySymType) int {
	tok, pos, lit := s.scan()
	s.pos = pos
	switch tok {
	case IDENT:
		return s.handleIdent(v, lit)
	case INT_LITERAL:
		return s.handleInt(v, lit)
	case FLOAT_LITERAL:
		return s.handleFloat(v, lit)
	case STRING_LITERAL:
		return s.handleString(v, lit)
	}
	return tok
}

func (s *Scanner) scan() (tok int, pos Pos, lit string) {
	ch := s.r.peek()
	if unicode.IsSpace(ch) {
		ch = s.skipSpace()
	}
	pos = s.r.pos()
	if ch == eof {
		return 0, pos, ""
	}
	if isIdentHead(ch) {
		return s.scanIdentifier()
	}
	if unicode.IsDigit(ch) {
		return s.scanNumber()
	}
	if ch == '\'' {
		return s.scanString()
	}
	return s.scanCharacters()
}

/* 跳过所有的空格, 不做记录 */
func (s *Scanner) skipSpace() rune {
	return s.r.incAsLongAs(func(b rune) bool {
		return unicode.IsSpace(b)
	})
}

/* 标识符和关键字 */
func (s *Scanner) scanIdentifier() (tok int, pos Pos, lit string) {
	pos = s.r.pos()
	_ = s.r.incAsLongAs(isIdentChar)
	return IDENT, pos, s.r.data(pos)
}

/* 扫描数字 */
func (s *Scanner) scanNumber() (tok int, pos Pos, lit string) {
	pos = s.r.pos()
	ch0 := s.skipDigit()
	if ch0 == '.' {
		s.r.inc(ch0)
		ch1 := s.r.peek()
		/* 3.141 is float, 3. is int */
		if unicode.IsDigit(ch1) {
			_ = s.skipDigit()
			return FLOAT_LITERAL, pos, s.r.data(pos)
		}
	}
	return INT_LITERAL, pos, s.r.data(pos)
}

func (s *Scanner) skipDigit() rune {
	return s.r.incAsLongAs(func(b rune) bool {
		return unicode.IsDigit(b)
	})
}

/* 字符串 */
func (s *Scanner) scanString() (tok int, pos Pos, lit string) {
	pos = s.r.pos()
	s.r.inc('\'')
	ch := s.r.incAsLongAs(func(b rune) bool {
		return b != '\''
	})
	if ch == 0 {
		return ERROR, pos, ""
	}
	s.r.inc('\'')
	return STRING_LITERAL, pos, s.r.data(pos)
}

func (s *Scanner) scanCharacters() (tok int, pos Pos, lit string) {
	pos = s.r.pos()
	ch := s.r.peek()
	node := &charTable
	for node.childs[ch] != nil {
		node = node.childs[ch]
		s.r.inc(ch)
		ch = s.r.peek()
	}
	tok = node.token
	if tok == 0 {
		s.err = errlog.New("unexpected rune")
	}
	return tok, pos, s.r.data(pos)
}

func (s *Scanner) handleIdent(v *yySymType, lit string) int {
	tok, ok := keywords[strings.ToUpper(lit)]
	if ok {
		v.ident = strings.ToLower(lit)
		return tok
	}
	v.ident = strings.ToLower(lit)
	return IDENT
}

func (s *Scanner) handleInt(v *yySymType, lit string) int {
	i, err := strconv.ParseUint(lit, 10, 64)
	if err != nil {
		return ERROR
	}
	switch {
	case i <= math.MaxInt64:
		v.item = int64(i)
	default:
		v.item = i
	}
	return INT_LITERAL
}

func (s *Scanner) handleFloat(v *yySymType, lit string) int {
	f, err := strconv.ParseFloat(lit, 64)
	if err != nil {
		return ERROR
	}
	v.item = f
	return FLOAT_LITERAL
}

func (s *Scanner) handleString(v *yySymType, lit string) int {
	v.item = lit[1 : len(lit)-1] /* 去掉首尾的单引号 */
	return STRING_LITERAL
}
