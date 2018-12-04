package lexer

import . "github.com/gfphoenix78/goc/token"

var (
	keywords = map[string]TokenKind{
		"if":          IF,
		"else":        ELSE,
		"for":         FOR,
		"return":      RETURN,
		"switch":      SWITCH,
		"case":        CASE,
		"default":     DEFAULT,
		"continue":    CONTINUE,
		"break":       BREAK,
		"fallthrough": FALLTHROUGH,
		"goto":        GOTO,
		"func":        FUNC,
		"var":         VAR,
		"struct":      STRUCT,
		"type":        TYPE,
		"true":        TRUE,
		"false":       FALSE,
	}

	singleTokens = map[byte]TokenKind{
		'?': QMARK,
		'(': LPAREN,
		')': RPAREN,
		'{': LBRACE,
		'}': RBRACE,
		'[': LBRACK,
		']': RBRACK,
		';': SEMICOLON,
		',': COMMA,
		':': COLON,
		'.': PERIOD,
		'~': NOT,
		'@': AT,
		'#': SHARP,
	}
)
