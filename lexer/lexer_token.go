package lexer

import . "gocc/token"

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
