package lexer

import "github.com/zeddee/learn-write-interpreter-in-go/token"

// Lexer tracks our progress while lexing the input string
type Lexer struct {
	input        string
	position     int  // index of current char in input
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New initialises a new lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

// NextToken returns the next token in the input string
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

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
	case '0':
		tok.Literal = ""
		tok.Type = token.EOF
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
