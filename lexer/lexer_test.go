package lexer

import (
	"testing"

	"github.com/zeddee/learn-write-interpreter-in-go/token"
)

func TestNextToken(t *testing.T) {
	type expectedTokens struct {
		expectedType    token.TokenType
		expectedLiteral string
	}

	tester := func(l *Lexer, tokens []expectedTokens) {
		for i, tt := range tokens {
			tok := l.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf(
					"tests[%d] — tokentype wrong, expected=%q, got  %q",
					i, tt.expectedType, tok.Type)
			}

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("tests[%d] — literal wrong. expected %q, got %q",
					i, tt.expectedLiteral, tok.Literal)
			}
		}
	}

	t.Run("=+(){},;", func(t *testing.T) {
		l := NewLexer(`=+(){},;`)

		simpleCase := []expectedTokens{
			{token.ASSIGN, "="},
			{token.PLUS, "+"},
			{token.LPAREN, "("},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.RBRACE, "}"},
			{token.COMMA, ","},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}

		tester(l, simpleCase)
	})
	t.Run("add function", func(t *testing.T) {
		l := NewLexer(`let five = 5;
		let ten = 10;
		let add = fn(x, y){
			x + y;
		};
		let result = add(five + ten);
		!-/*5;
		5 < 10 > 5;
		`)

		expected := []expectedTokens{
			{token.LET, "let"},
			{token.IDENT, "five"},
			{token.ASSIGN, "="},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "ten"},
			{token.ASSIGN, "="},
			{token.INT, "10"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "add"},
			{token.ASSIGN, "="},
			{token.FUNCTION, "fn"},
			{token.LPAREN, "("},
			{token.IDENT, "x"},
			{token.COMMA, ","},
			{token.IDENT, "y"},
			{token.RPAREN, ")"},
			{token.LBRACE, "{"},
			{token.IDENT, "x"},
			{token.PLUS, "+"},
			{token.IDENT, "y"},
			{token.SEMICOLON, ";"},
			{token.RBRACE, "}"},
			{token.SEMICOLON, ";"},
			{token.LET, "let"},
			{token.IDENT, "result"},
			{token.ASSIGN, "="},
			{token.IDENT, "add"},
			{token.LPAREN, "("},
			{token.IDENT, "five"},
			{token.PLUS, "+"},
			{token.IDENT, "ten"},
			{token.RPAREN, ")"},
			{token.SEMICOLON, ";"},
			{token.BANG, "!"},
			{token.MINUS, "-"},
			{token.SLASH, "/"},
			{token.ASTERISK, "*"},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.INT, "5"},
			{token.LT, "<"},
			{token.INT, "10"},
			{token.RT, ">"},
			{token.INT, "5"},
			{token.SEMICOLON, ";"},
			{token.EOF, ""},
		}
		tester(l, expected)
	})
}
