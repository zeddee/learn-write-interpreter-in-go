package token

// TokenType is a string that describes type of the token
type TokenType string

// Token describes each token our lexer can parse
type Token struct {
	Type    TokenType
	Literal string
}

// define all possible token types
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + Literals
	IDENT = "IDENT" // userspace names e.g. add, foobar, x, y, ...
	INT   = "INT"   // whole no.s

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords. These are reserved,
	// and should not overlap with Identifiers
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
