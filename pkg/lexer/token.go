package lexer

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

type TokenType string

const (
	// Special
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers + keywords
	IDENTIFIER TokenType = "IDENT"
	KEYWORD    TokenType = "KEYWORD"

	// Literals
	NUMBER TokenType = "NUMBER"
	STRING TokenType = "STRING"

	// Symbols
	PERIOD TokenType = "."
	COMMA  TokenType = ","
	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	PLUS   TokenType = "+"
	MINUS  TokenType = "-"
	TIMES  TokenType = "*"
	DIVIDE TokenType = "/"

	// Comments
	COMMENT TokenType = "COMMENT"
)
