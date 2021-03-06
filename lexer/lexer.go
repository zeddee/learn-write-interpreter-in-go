// Package lexer contains
// code that reads and tokenises
// incoming strings that it ingests,
// primarily through Lexer.NextToken() calls.
package lexer

import (
	"strconv"

	"github.com/zeddee/learn-write-interpreter-in-go/token"
)

// Lexer is a data structure that contains
// the string we need to lex,
// and our lexing progress through it.
type Lexer struct {
	input        string
	position     int  // index of current char in input
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// NewLexer initialises a new lexer
// by taking an input string
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // Initialises the rest of the Lexer struct
	// by attempting to read the first byte
	return l
}

// readChar reads a character in our input string
// and updates the lexer.position and lexer.readPosition
// fields in our lexer
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // sets current char to ASCII for NULL
	} else {
		l.ch = l.input[l.readPosition] // set current char to index of the position to read
	}
	l.position = l.readPosition // updates current position
	l.readPosition++            // updates position to read on the next readChar call
}

// peekChar reads the byte ahead of current l.position,
// but doesn't move it i.e. no read progress is made.
// This is for a case where we just want to see what's ahead,
// but don't want to lex it yet
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// newToken returns a new token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// NextToken returns the next token in the input string
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		// comparison operators are a bit more
		// complex to check for.
		// we need to check if we have
		// a special sequence of characters
		// e.g. two operators put together
		// that would constitute a comparison
		// operator.
		// then we can lex and tokenise
		// that sequence of characters
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			// manually create a token.Token
			// because our newToken function
			// only takes in a single byte.
			tok = token.Token{
				Type:    token.EQ,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{
				Type:    token.NEQ,
				Literal: string(ch) + string(l.ch),
			}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.RT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// add default case to catch cases
		// where we need to lex a sequence of chars
		// e.g. for any identifiers

		// switch in a switch ...
		switch checkCharTokenType(l.ch) {
		case token.IDENT:
			tok.Literal = l.readLiteralSequence(token.IDENT)
			tok.Type = token.LookupIdent(tok.Literal)
			// we're exiting early here because
			// we've already called l.readChar()
			// (which we _must_ call in order to get
			// the complete identifer string)
			// inside our readLiteralSequence() call.
			return tok
		case token.INT:
			tok.Literal = l.readLiteralSequence(token.INT)
			tok.Type = token.INT
			return tok
		case token.ILLEGAL:
			fallthrough
		default:
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// checkCharTokenType reads a byte and
// returns its token type.
func checkCharTokenType(ch byte) token.TokenType {
	switch {
	// ch is a letter
	case 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_':
		return token.IDENT
	// ch is an int
	case func() bool {
		// a bit of an unweildy anonymous function
		// to handle the truthiness of strconv.Atoi(ch)
		if _, err := strconv.Atoi(string(ch)); err != nil {
			return false
		}
		return true
	}():
		return token.INT
	}
	return token.ILLEGAL
}

type checkCharType func(ch byte) bool

// readLiteralSequence reads a sequence of characters
// of the same type (str, int, etc.)
// and returns that sequence as a string
func (l *Lexer) readLiteralSequence(expectedToken token.TokenType) string {
	// store the current value of l.position
	// so we know where our character
	// sequence starts in l.input
	position := l.position
	for checkCharTokenType(l.ch) == expectedToken {
		// read characters in 'input',
		// check if character fulfils the
		// type check implemented by checkCharTokenType()
		// and advance l.position
		// once we hit checkCharTokenType() != expectedToken,
		// we stop and return the string
		// of characters that we've read
		l.readChar()
	}

	return l.input[position:l.position]
}

// skipWhitespace skips whitespace that does
// not qualify as EOF
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
