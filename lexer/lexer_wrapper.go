package lexer

import "github.com/gfphoenix78/goc/token"

type Lexer2 struct {
	*BaseLexer
	last   *token.Token // last consumed token
	enable bool         // if enabled, the next token is `last`, because the last token is a fake ';'
}

func NewLexer2(source []byte) *Lexer2 {
	return &Lexer2{
		BaseLexer: NewLexer(source),
	}
}
func semi(tk *token.Token) *token.Token {
	tk.Pos.Column += len(tk.Str) + 1
	return &token.Token{
		Kind: token.SEMICOLON,
		Pos:  tk.Pos,
		Str:  ";",
	}
}
func (l *Lexer2) Next() *token.Token {
	if l.enable {
		tk := l.last
		l.last = nil
		l.enable = false
		return tk
	}
	tk := l.BaseLexer.Next()
	if l.last == nil || l.last.Pos.Line == tk.Pos.Line {
		l.last = tk
		return tk
	}
	// if to a new line
	switch l.last.Kind {
	case token.IDENT,
		token.INT_LIT,
		token.BOOL_LIT,
		token.STRING_LIT,
		token.CHAR_LIT,
		// key words
		token.RBRACE,
		token.RPAREN,
		token.RBRACK,
		token.INC,
		token.DEC:
	default:
		l.last = tk
		return tk
	}
	last := l.last
	l.last = tk
	l.enable = true
	return semi(last)
}
