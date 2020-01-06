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

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent  takes in string ident
// and checks if we have the value of
// ident stored in our 'keywords' map
// if it exists in our 'keywords' map,
// then we look up the keyword and
// return the stored TokenType.
// Otherwise, we'll treat it as an identifier
// and return a TokenType of IDENT.
func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}
