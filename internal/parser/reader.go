package parser

import (
	"unicode/utf8"
)

/*
 * 读取器实现
 */
type Pos struct {
	Line   int /* 当前所在行 */
	Col    int /* 当前行所在列 */
	Offset int
}

type reader struct {
	s string
	p Pos
	l int
}

/* 结束符 */
const eof rune = 0

func (r *reader) eof() bool {
	return r.p.Offset >= r.l
}

/* 查看字符 */
func (r *reader) peek() rune {
	if r.eof() {
		return eof
	}
	ch, _ := utf8.DecodeRuneInString(r.s[r.p.Offset:])
	return ch
}

func (r *reader) inc(ch rune) {
	if ch == '\n' {
		r.p.Line++
		r.p.Col = 0
	}
	n := utf8.RuneLen(ch)
	r.p.Offset += n
	r.p.Col += n
}

/*
 * 移动至最后满足条件的, 并返回下一个字符(第一个未满足条件的字符)
 */
func (r *reader) incAsLongAs(fn func(b rune) bool) rune {
	for {
		ch := r.peek()
		if !fn(ch) {
			return ch
		}
		if r.eof() {
			return 0
		}
		r.inc(ch)
	}
}

func (r *reader) data(from Pos) string {
	return r.s[from.Offset:r.p.Offset]
}

func (r *reader) pos() Pos {
	return r.p
}
