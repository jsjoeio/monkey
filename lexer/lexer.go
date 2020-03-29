package lexer

import "monkey/token"

type Lexer struct {
	input string
	position int // current position in input (points to current char)
	readPosition int // current reading position input (after current char)
	ch byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

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
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// purpose of this function is to give us the next character and
// advance our position in the input string
func (l *Lexer) readChar() {
	// check if we have reached the end of input
	if l.readPosition >= len(l.input) {
		// ASCII code for the "NUL" is 0
		// signifies either "we haven't read anything yet" or "end of file"
		l.ch = 0
	} else {
		// otherwise we set it to next character
		l.ch = l.input[l.readPosition]
	}
	// set position
	l.position = l.readPosition
	// increment so we point to next position
	l.readPosition += 1
}