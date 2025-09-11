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
	IDENTIFIER TokenType = "IDENTIFIER"
	KEYWORD    TokenType = "KEYWORD"

	// Literals
	NUMBER TokenType = "NUMBER"
	STRING TokenType = "STRING"

	// Symbols
	PERIOD TokenType = "PERIOD" // .
	COMMA  TokenType = "COMMA"  // ,
	LPAREN TokenType = "LPAREN" // (
	RPAREN TokenType = "RPAREN" // )
	PLUS   TokenType = "PLUS"   // +
	MINUS  TokenType = "MINUS"  // -
	TIMES  TokenType = "TIMES"  // *
	DIVIDE TokenType = "DIVIDE" // /

	// Comments
	COMMENT TokenType = "COMMENT"
)
