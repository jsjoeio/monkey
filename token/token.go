package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
		ILLEGAL = "ILLEGAL"
		EOF = "EOF"

		// Identifers + literals
		IDENT = "IDENT" // add, helloWorld, convertThing, x, y
		INT = "INT" // 14324

		// Operators
		ASSIGN = "="
		PLUS = "+"

		// Delimiters
		COMMA = ","
		SEMICOLON = ";"
		LPAREN = "("
		RPAREN = ")"
		LBRACE = "{"
		RBRACE = "}"

		// Keywords
		FUNCTION = "FUNCTION"
		LET = "LET"
)