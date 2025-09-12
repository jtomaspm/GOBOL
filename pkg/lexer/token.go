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
	//Arithmetic operators
	ADDITION       TokenType = "ADDITION"       // +
	SUBTRACTION    TokenType = "SUBTRACTION"    // -
	DIVISION       TokenType = "DIVISION"       // /
	MULTIPLICATION TokenType = "MULTIPLICATION" // *
	EXPONENTIATION TokenType = "EXPONENTIATION" // **
	//Relational operators
	GREATER_THAN       TokenType = "GREATER_THAN"       // >
	LESS_THAN          TokenType = "LESS_THAN"          // <
	EQUAL              TokenType = "EQUAL"              // =
	PSEUDO_EQUAL       TokenType = "PSEUDO_EQUAL"       // ==
	GREATER_THAN_EQUAL TokenType = "GREATER_THAN_EQUAL" // >=
	LESS_THAN_EQUAL    TokenType = "LESS_THAN_EQUAL"    // <=
	NOT_EQUAL          TokenType = "NOT_EQUAL"          // <>
	// Other symbols
	COMMENT_INDICATOR  TokenType = "COMMENT_INDICATOR"  // *>
	COMPILER_DIRECTIVE TokenType = "COMPILER_DIRECTIVE" // >>
	PERIOD             TokenType = "PERIOD"             // .
	COMMA              TokenType = "COMMA"              // ,
	LPAREN             TokenType = "LPAREN"             // (
	RPAREN             TokenType = "RPAREN"             // )

	// Comments
	COMMENT TokenType = "COMMENT"
)
