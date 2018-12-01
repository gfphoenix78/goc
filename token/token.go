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

	INT_LIT		// 34
	FLOAT_LIT // 11.23
	BOOL_LIT // true/false

	// keyword. not package/import now

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

	// operator
	ADD   // +
	SUB   // -
	MUL   // *
	DIV   // /
	REM   // %
	// bit ops
	AND   // &
	OR    // |
	NOT   // ~
	XOR   // ^
	SHL // <<
	SHR // >>
	// logic ops
	LAND   // &&
	LOR    // ||
	LNOT	// !
	// conditional ops
	EQ     // ==
	LT     // <
	GT     // >
	ASSIGN // =
	NE     // !=
	LE     // <=
	GE     // >=
	

	ADD_ASSIGN   // +=
	SUB_ASSIGN   // -=
	MUL_ASSIGN   // *=
	DIV_ASSIGN   // /=
	REM_ASSIGN   // %=
	AND_ASSIGN   // &=
	OR_ASSIGN    // |=
	XOR_ASSIGN   // ^=
	RIGHT_ASSIGN // >>=
	LEFT_ASSIGN  // <<=


	ARROW  // ->
	
	INC    // ++
	DEC    // --

QUE   // ?
	LPAREN   // (
	LBRACK   // [
	LBRACE   // {
	COMMA    // ,
	PERIOD   // .
	ELLIPSIS // ...

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :

	// special tokens
	EOF
	COMMENT // /* or //
	UNKNOWN
)

func (k TokenKind) String() string {
	return [...]string{
		IDENT: "IDENT",

		INT_CONST:    "INT_CONST",
		FLOAT_CONST:  "FLOAT_CONST",
		STRING_CONST: "STRING_CONST",
		CHAR_CONST:   "CHAR_CONST",

		INT:    "INT",
		VOID:   "VOID",
		CHAR:   "CHAR",
		FLOAT:  "FLOAT",
		LONG:   "LONG",
		SHORT:  "SHORT",
		DOUBLE: "DOUBLE",

		DO:       "DO",
		WHILE:    "WHILE",
		IF:       "IF",
		ELSE:     "ELSE",
		FOR:      "FOR",
		AUTO:     "AUTO",
		RETURN:   "RETURN",
		SWITCH:   "SWITCH",
		CASE:     "CASE",
		DEFAULT:  "DEFAULT",
		CONTINUE: "CONTINUE",
		BREAK:    "BREAK",
		GOTO:     "GOTO",
		CONST:    "CONST",
		EXTERN:   "EXTERN",
		REGISTER: "REGISTER",
		SIGNED:   "SIGNED",
		UNSIGNED: "UNSIGNED",
		SIZEOF:   "SIZEOF",
		STATIC:   "STATIC",
		STRUCT:   "STRUCT",
		TYPEDEF:  "TYPEDEF",
		UNION:    "UNION",
		VOLATILE: "VOLATILE",
		ENUM:     "ENUM",

		ADD:   "ADD",
		SUB:   "SUB",
		MUL:   "MUL",
		DIV:   "DIV",
		REM:   "REM",
		AND:   "AND",
		OR:    "OR",
		QUE:   "QUE",
		XOR:   "XOR",
		TILDE: "TILDE",

		ADD_ASSIGN:   "ADD_ASSIGN",
		SUB_ASSIGN:   "SUB_ASSIGN",
		MUL_ASSIGN:   "MUL_ASSIGN",
		DIV_ASSIGN:   "DIV_ASSIGN",
		REM_ASSIGN:   "REM_ASSIGN",
		RIGHT_ASSIGN: "RIGHT_ASSIGN",
		LEFT_ASSIGN:  "LEFT_ASSIGN",
		AND_ASSIGN:   "AND_ASSIGN",
		OR_ASSIGN:    "OR_ASSIGN",
		XOR_ASSIGN:   "XOR_ASSIGN",

		SHL: "SHL",
		SHR: "SHR",
		ARROW:  "ARROW",
		LAND:   "LAND",
		LOR:    "LOR",
		INC:    "INC",
		DEC:    "DEC",
		EQ:     "EQ",
		LT:     "LT",
		GT:     "GT",
		ASSIGN: "ASSIGN",
		NOT:    "NOT",
		NE:     "NE",
		LE:     "LE",
		GE:     "GE",

		LPAREN: "LPAREN",
		LBRACK: "LBRACK",
		LBRACE: "LBRACE",
		COMMA:  "COMMA",
		PERIOD: "PERIOD",

		RPAREN:    "RPAREN",
		RBRACK:    "RBRACK",
		RBRACE:    "RBRACE",
		SEMICOLON: "SEMICOLON",
		COLON:     "COLON",
		ELLIPSIS:  "ELLIPSIS",

		EOF:     "EOF",
		COMMENT: "COMMENT",
		UNKNOWN: "UNKNOWN",
	}[k]
}
