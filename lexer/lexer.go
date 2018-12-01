package lexer

import "gocc/token"

type Lexer struct {
	scanner *Scanner
}

func NewLexer(source []byte) *Lexer {
	return &Lexer{scanner: NewScanner(source)}
}

func (l *Lexer) Pos() token.Position {
	return l.scanner.Pos()
}

func (l *Lexer) Next() *token.Token {
	t := &token.Token{}
	if l.scanner.IsEnd() {
		t.Kind = token.EOF
		t.Pos = l.scanner.Pos()
		return t
	}

	c := l.skipSpace()
	pos := l.scanner.Pos()

	switch  {
		case isAlpha(c) || c=='_' : l.parseAlpha(t)
		case isDigit(c) : l.parseNumber(t)
		case c=='"': l.parseString(t)
		case c=='\'': l.parseChar(t)
		default: l.parseOther(t)
	}
	if t.Kind == token.COMMENT {
		return l.Next()
	}
	t.Pos = pos
	return t
}

func (l *Lexer) Reset(pos token.Position) {
	l.scanner.Reset(pos)
}

func (l *Lexer) consume() (byte, bool) {
	l.scanner.Step()
	if l.scanner.IsEnd() {
		return 0, false
	}
	return l.scanner.Get(), true
}

func (l *Lexer) parseAlpha(t *token.Token) {
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
	if s=="true" || s=="false" {
		return token.BOOL_LIT
	}
	return token.IDENT
}
// only support int now
func (l *Lexer) parseNumber(t *token.Token) {
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

func (l *Lexer) skipSpace() byte {
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
	return (c == ' ' || c == '\t' || c == '\f' || c == '\r')
}

func isReturn(c byte) bool {
	return c == '\n'
}

func isAlpha(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

func isDigit(c byte) bool {
	return ('0' <= c && c <= '9')
}

func (c byte) (token.TokenKind, bool) {
	if v, ok := singleTokens[c]; ok {
		return v, ok
	}
	return token.EOF, false
}

func (l *Lexer) parseChar(t *token.Token) {
	panic("not implement yet")
}

func (l *Lexer) parseString(t *token.Token) {
	t.Str = string(l.readString())
	t.Kind = token.STRING_LIT
}
// no escape char yet
func (l *Lexer) readString() []byte {
	ok := false
	c := l.scanner.Get()

	var s []byte
	for c,ok=l.consume(); ok && c!='"'; c,ok=l.consume() {
		s = append(s, l.scanner.Get())
	}
	if c!='"' {
		panic("not complete string")
	}
	l.consume()
	return s
}
func (l *Lexer)checkComment(t *token.Token) {
	if t.Kind == COMMENT {
		var comments []byte
		var ok bool
		for c:=l.scanner.Get(); c != '\n' && !l.scanner.IsEnd(); {
			comments = append(comments, c)
			c, _ = l.consume()
		}
		t.Str = string(comments)
	}
}
func (l *Lexer)parseOther(t *token.Token) {
	c := l.scanner.Get()
	switch c {
	case '+': l.parseType3(t, ADD, INC, ADD_ASSIGN)
	case '-': l.parseType3(t, SUB, DEC, SUB_ASSIGN)
	case '!': l.parseType2(t, LNOT, NE)
	case '%': l.parseType2(t, REM, REM_ASSIGN)
	case '^': l.parseType2(t, XOR, XOR_ASSIGN)
	case '&': l.parseType4(t, AND, LAND, AND_ASSIGN, LAND_ASSIGN)
	case '|': l.parseType4(t, OR, LOR, OR_ASSIGN, LOR_ASSIGN)
	case '*': l.parseType2(t, MUL, MUL_ASSIGN)
	case '=': l.parseType2(t, ASSIGN, EQ)
	case '<': l.parseType4(t, LT, SHL, LE, SHL_ASSIGN)
	case '>': l.parseType4(t, GT, SHR, GE, SHR_ASSIGN)
	case '/': l.parseType3(t, DIV, COMMENT, DIV_ASSIGN)
				l.checkComment(t)
	default:
		l.parseOther2(t)
	}
}
func (l *Lexer)parseOther2(t *token.Token) {
	c := l.scanner.Get()
	if v, ok := singleTokens[c]; ok {
		l.consume()
		t.Kind = v
		t.Str = string(c)
		return
	}else {
		panic("unknown char:"+string(c))
	}
}
// 'x' 'x='
func (l *Lexer)parseType2(t *token.Token, tk, tke token.TokenKind) {
	old = l.scanner.Get()
	c, ok := l.consume()
	if ok && c=='=' {
		t.Kind = tke
		t.Str = string([]byte{old,c})
		l.consume()
	}else {
		t.Kind = tk
		t.Str = string(old)
	}
}
'x' 'xx' 'x='
func (l *Lexer)parseType3(t *token.Token, tk, tkd, tke token.TokenKind) {
	c1 := l.scanner.Get()
	c2, ok := l.consume()
	switch {
	case ok && c2 == c1:
		t.Kind = tkd
		t.Str = string([]byte{c1,c2})
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
func (l *Lexer)parseType4(t *token.Token, tk, tkd, tke , tkde token.TokenKind) {
	var c1, c2, c3 byte
	var ok bool
	c1 = l.scanner.Get()
	c2, ok = l.consume()
	switch {
	case ok && c2=='=':
		t.Kind = tke
		t.Str = string([]byte{c1, c2})
		l.consume()
	case ok && c2==c1: // tkd or tkde
		c3, ok = l.consume()
		if ok && c3=='=' {
			t.Kind = tkde
			t.Str = string([]byte{c1,c2,c3})
			l.consume()
		}else {
			t.Kind = tkd
			t.Str = string([]byte{c1,c2})
		}
	default:
		t.Kind = tk
		t.Str = string(c1)
	}	
}