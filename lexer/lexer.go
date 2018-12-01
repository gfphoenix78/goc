package lexer

import (
	"goc/token"
)

type BaseLexer struct {
	scanner *Scanner
}

func NewLexer(source []byte) *BaseLexer {
	return &BaseLexer{scanner: NewScanner(source)}
}

func (l *BaseLexer) Pos() token.Position {
	return l.scanner.Pos()
}

func (l *BaseLexer) Next() *token.Token {
	if l.scanner.IsEnd() {
		return &token.Token{
			Kind: token.EOF,
			Pos:  l.scanner.Pos(),
		}
	}
	t := &token.Token{}
	c := l.skipSpace()
	t.Pos = l.scanner.Pos()
	switch {
	case isAlpha(c) || c == '_':
		l.parseAlpha(t)
	case isDigit(c):
		l.parseNumber(t)
	case c == '"':
		l.parseString(t)
	case c == '\'':
		l.parseChar(t)
	default:
		l.parseOther(t)
	}
	if t.Kind == token.COMMENT {
		return l.Next()
	}
	return t
}

func (l *BaseLexer) Reset(pos token.Position) {
	l.scanner.Reset(pos)
}

func (l *BaseLexer) consume() (byte, bool) {
	l.scanner.Step()
	if l.scanner.IsEnd() {
		return 0, false
	}
	return l.scanner.Get(), true
}

func (l *BaseLexer) parseAlpha(t *token.Token) {
	var s []byte
	var ok bool
	c := l.scanner.Get()
	for isAlpha(c) || isDigit(c) || c == '_' {
		s = append(s, c)
		if c, ok = l.consume(); !ok {
			break
		}
	}
	t.Str = string(s)
	t.Kind = checkIdent(t.Str)
}

func checkIdent(s string) token.TokenKind {
	if v, ok := keywords[s]; ok {
		return v
	}
	return token.IDENT
}

// only support int now
func (l *BaseLexer) parseNumber(t *token.Token) {
	var s []byte
	c := l.scanner.Get()
	var ok bool
	for isDigit(c) {
		s = append(s, c)
		if c, ok = l.consume(); !ok {
			break
		}
	}
	t.Str = string(s)
	t.Kind = token.INT_LIT
}

func (l *BaseLexer) skipSpace() byte {
	c := l.scanner.Get()
	ok := false
	for isWhitespace(c) || isReturn(c) {
		if c, ok = l.consume(); !ok {
			break
		}
	}
	return c
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\f' || c == '\r'
}

func isReturn(c byte) bool {
	return c == '\n'
}

func isAlpha(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// escape: '\\' '\n' '\r' '\'' '\"' '\t'
// the current byte is '\\'
// after calling this, the current char is the last valid escape char
func (l *BaseLexer) escape() byte {
	var ch byte
	c, ok := l.consume()
	if !ok {
		panic("unexpected EOF")
	}
	switch c {
	case '\\':
		ch = '\\'
	case 'n':
		ch = '\n'
	case 'r':
		ch = '\r'
	case 't':
		ch = '\t'
	case '\'':
		ch = '\''
	case '"':
		ch = '"'
	default:
		panic("unexpected escape char = '" + string([]byte{c, '\''}))
	}
	return ch
}
func (l *BaseLexer) parseChar(t *token.Token) {
	c, ok := l.consume()
	if !ok {
		panic("incomplete char")
	}
	if c != '\\' {
		t.Str = string(c)
	} else {
		t.Str = string(l.escape())
	}
	if c, ok = l.consume(); !ok || c != '\'' {
		panic("expected end of '")
	}
	l.consume()
	t.Kind = token.CHAR_LIT
}

func (l *BaseLexer) parseString(t *token.Token) {
	t.Str = string(l.readString())
	t.Kind = token.STRING_LIT
}

// no escape char yet
func (l *BaseLexer) readString() []byte {
	ok := false
	c := l.scanner.Get()

	var s []byte
	for c, ok = l.consume(); ok && c != '"'; c, ok = l.consume() {
		if c != '\\' {
			s = append(s, c)
		} else {
			s = append(s, l.escape())
		}
	}
	if c != '"' {
		panic("not complete string")
	}
	l.consume()
	return s
}
func (l *BaseLexer) checkComment(t *token.Token) {
	if t.Kind == token.COMMENT {
		comments := []byte{'/', '/'}
		for c := l.scanner.Get(); c != '\n' && !l.scanner.IsEnd(); {
			comments = append(comments, c)
			c, _ = l.consume()
		}
		t.Str = string(comments)
	}
}
func (l *BaseLexer) parseOther(t *token.Token) {
	c := l.scanner.Get()
	switch c {
	case '+':
		l.parseType3(t, token.ADD, token.INC, token.ADD_ASSIGN)
	case '-':
		l.parseType3(t, token.SUB, token.DEC, token.SUB_ASSIGN)
	case '!':
		l.parseType2(t, token.LNOT, token.NE)
	case '%':
		l.parseType2(t, token.REM, token.REM_ASSIGN)
	case '^':
		l.parseType2(t, token.XOR, token.XOR_ASSIGN)
	case '&':
		l.parseType4(t, token.AND, token.LAND, token.AND_ASSIGN, token.LAND_ASSIGN)
	case '|':
		l.parseType4(t, token.OR, token.LOR, token.OR_ASSIGN, token.LOR_ASSIGN)
	case '*':
		l.parseType2(t, token.MUL, token.MUL_ASSIGN)
	case '=':
		l.parseType2(t, token.ASSIGN, token.EQ)
	case '<':
		l.parseType4(t, token.LT, token.SHL, token.LE, token.SHL_ASSIGN)
	case '>':
		l.parseType4(t, token.GT, token.SHR, token.GE, token.SHR_ASSIGN)
	case '/':
		l.parseType3(t, token.DIV, token.COMMENT, token.DIV_ASSIGN)
		l.checkComment(t)
	default:
		l.parseOther2(t)
	}
}
func (l *BaseLexer) parseOther2(t *token.Token) {
	c := l.scanner.Get()
	if v, ok := singleTokens[c]; ok {
		l.consume()
		t.Kind = v
		t.Str = string(c)
		return
	} else {
		panic("unknown char:" + string(c))
	}
}

// 'x' 'x='
func (l *BaseLexer) parseType2(t *token.Token, tk, tke token.TokenKind) {
	old := l.scanner.Get()
	c, ok := l.consume()
	if ok && c == '=' {
		t.Kind = tke
		t.Str = string([]byte{old, c})
		l.consume()
	} else {
		t.Kind = tk
		t.Str = string(old)
	}
}

// 'x' 'xx' 'x='
func (l *BaseLexer) parseType3(t *token.Token, tk, tkd, tke token.TokenKind) {
	c1 := l.scanner.Get()
	c2, ok := l.consume()
	switch {
	case ok && c2 == c1:
		t.Kind = tkd
		t.Str = string([]byte{c1, c2})
		l.consume()
	case ok && c2 == '=':
		t.Kind = tke
		t.Str = string([]byte{c1, c2})
		l.consume()
	default:
		t.Kind = tk
		t.Str = string(c1)
	}
}

// 'x' 'xx' 'x=' 'xx='
func (l *BaseLexer) parseType4(t *token.Token, tk, tkd, tke, tkde token.TokenKind) {
	var c1, c2, c3 byte
	var ok bool
	c1 = l.scanner.Get()
	c2, ok = l.consume()
	switch {
	case ok && c2 == '=':
		t.Kind = tke
		t.Str = string([]byte{c1, c2})
		l.consume()
	case ok && c2 == c1: // tkd or tkde
		c3, ok = l.consume()
		if ok && c3 == '=' {
			t.Kind = tkde
			t.Str = string([]byte{c1, c2, c3})
			l.consume()
		} else {
			t.Kind = tkd
			t.Str = string([]byte{c1, c2})
		}
	default:
		t.Kind = tk
		t.Str = string(c1)
	}
}
