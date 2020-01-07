package lexer

import (
	"strconv"

	"github.com/zeddee/learn-write-interpreter-in-go/token"
)

// Lexer tracks our progress while lexing the input string
type Lexer struct {
	input        string
	position     int  // index of current char in input
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New initialises a new lexer
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken returns the next token in the input string
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
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
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// add default case to catch cases
		// where we need to lex a sequence of chars
		// e.g. for any identifiers
		if isLetter(l.ch) {
			tok.Literal = l.readLiteralSequence(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			// we're exiting early here because
			// we've already called l.readChar()
			// (which we _must_ call in order to get
			// the complete identifer string)
			// inside our readIdentifier() call.
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readLiteralSequence(isDigit)
			tok.Type = token.INT
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
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

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func checkCharTokenType(ch byte) token.TokenType {
	switch {
	// ch is a letter
	case 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_':
		return token.IDENT
	// ch is an int
	case func() bool {
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
func (l *Lexer) readLiteralSequence(check checkCharType) string {
	// store the current value of l.position
	// so we know where our character
	// sequence starts in l.input
	position := l.position
	for check(l.ch) {
		// read characters in 'input',
		// check if character fulfils the
		// type check implemented by checkCharType()
		// and advance l.position
		// once we hit checkCharType() == false,
		// we stop and return the string
		// of characters that we've read
		l.readChar()
	}

	return l.input[position:l.position]
}

// isLetter checks if ch is a letter.
// the scope of this checker will determine
// what characters will be legal in keywords
// and identifiers.
func isLetter(ch byte) bool {
	if 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' {
		return true
	}
	return false
}

func isDigit(ch byte) bool {
	// reimplemented as strconv.Atoi check instead
	// of literal comparisons.
	// hopefully more robust
	if _, err := strconv.Atoi(string(ch)); err != nil {
		return false
	}
	return true
}

// skipWhitespace skips whitespace that does
// not qualify as EOF
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}
