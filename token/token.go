package token

type Position struct {
	Line   int
	Column int
	Offset int
}

type TokenKind int

type Token struct {
	Kind TokenKind
	Str  string
	Pos  Position
}

func (t *Token) String() string {
	return t.Str
}

const (
	IDENT TokenKind = iota

	INT_LIT    // 34
	FLOAT_LIT  // 11.23
	STRING_LIT // "hi"
	CHAR_LIT   // 'c'

	// keyword. not package/import now
	begin_keyword
	IF
	ELSE
	FOR
	RETURN
	SWITCH
	CASE
	DEFAULT
	CONTINUE
	BREAK
	FALLTHROUGH
	GOTO
	FUNC
	VAR
	STRUCT
	TYPE
	TRUE
	FALSE
	end_keyword

	// operator
	ADD // +
	SUB // -
	MUL // *
	DIV // /
	REM // %
	// bit ops
	AND // &
	OR  // |
	NOT // ~
	XOR // ^
	SHL // <<
	SHR // >>
	// logic ops
	LAND // &&
	LOR  // ||
	LNOT // !
	// conditional ops
	EQ     // ==
	LT     // <
	GT     // >
	ASSIGN // =
	NE     // !=
	LE     // <=
	GE     // >=

	ADD_ASSIGN  // +=
	SUB_ASSIGN  // -=
	MUL_ASSIGN  // *=
	DIV_ASSIGN  // /=
	REM_ASSIGN  // %=
	AND_ASSIGN  // &=
	OR_ASSIGN   // |=
	XOR_ASSIGN  // ^=
	SHL_ASSIGN  // <<=
	SHR_ASSIGN  // >>=
	LAND_ASSIGN // &&=
	LOR_ASSIGN  // ||=

	INC // ++
	DEC // --

	QMARK    // ?
	LPAREN   // (
	LBRACK   // [
	LBRACE   // {
	RPAREN   // )
	RBRACK   // ]
	RBRACE   // }
	COMMA    // ,
	AT       // @
	SHARP    // #
	PERIOD   // .
	ELLIPSIS // ...

	SEMICOLON // ;
	COLON     // :

	// special tokens
	EOF
	COMMENT
	UNKNOWN
)

func (k TokenKind) IsKeyword() bool {
	return k > begin_keyword && k < end_keyword
}

var str_map = [...]string{
	IDENT: "IDENT",

	INT_LIT:    "INT_LIT",
	FLOAT_LIT:  "FLOAT_LIT",
	STRING_LIT: "STRING_LIT",
	CHAR_LIT:   "CHAR_LIT",

	IF:          "IF",
	ELSE:        "ELSE",
	FOR:         "FOR",
	RETURN:      "RETURN",
	SWITCH:      "SWITCH",
	CASE:        "CASE",
	DEFAULT:     "DEFAULT",
	CONTINUE:    "CONTINUE",
	BREAK:       "BREAK",
	FALLTHROUGH: "FALLTHROUGH",
	GOTO:        "GOTO",
	FUNC:        "FUNC",
	VAR:         "VAR",
	STRUCT:      "STRUCT",
	TYPE:        "TYPE",
	TRUE:        "TRUE",
	FALSE:       "FALSE",

	ADD: "ADD",
	SUB: "SUB",
	MUL: "MUL",
	DIV: "DIV",
	REM: "REM",

	AND: "AND",
	OR:  "OR",
	NOT: "NOT",
	XOR: "XOR",
	SHL: "SHL",
	SHR: "SHR",

	LAND:   "LAND",
	LOR:    "LOR",
	LNOT:   "LNOT",
	EQ:     "EQ",
	LT:     "LT",
	GT:     "GT",
	ASSIGN: "ASSIGN",
	NE:     "NE",
	LE:     "LE",
	GE:     "GE",

	ADD_ASSIGN:  "ADD_ASSIGN",
	SUB_ASSIGN:  "SUB_ASSIGN",
	MUL_ASSIGN:  "MUL_ASSIGN",
	DIV_ASSIGN:  "DIV_ASSIGN",
	REM_ASSIGN:  "REM_ASSIGN",
	AND_ASSIGN:  "AND_ASSIGN",
	OR_ASSIGN:   "OR_ASSIGN",
	XOR_ASSIGN:  "XOR_ASSIGN",
	SHL_ASSIGN:  "SHL_ASSIGN",
	SHR_ASSIGN:  "SHR_ASSIGN",
	LAND_ASSIGN: "LAND_ASSIGN",
	LOR_ASSIGN:  "LOR_ASSIGN",

	INC:       "INC",
	DEC:       "DEC",
	QMARK:     "QMARK",
	LPAREN:    "LPAREN",
	LBRACK:    "LBRACK",
	LBRACE:    "LBRACE",
	RPAREN:    "RPAREN",
	RBRACK:    "RBRACK",
	RBRACE:    "RBRACE",
	AT:        "AT",
	SHARP:     "SHARP",
	COMMA:     "COMMA",
	PERIOD:    "PERIOD",
	ELLIPSIS:  "ELLIPSIS",
	SEMICOLON: "SEMICOLON",
	COLON:     "COLON",

	EOF: "EOF",
}

func (k TokenKind) String() string {
	return str_map[k]
}
