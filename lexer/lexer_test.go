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
}
