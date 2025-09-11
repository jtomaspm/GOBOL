package lexer

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

type Lexer struct {
	r    *bufio.Reader
	line int
	col  int
	done bool
}

func New(r io.Reader) *Lexer {
	return &Lexer{r: bufio.NewReader(r), line: 0, col: 0, done: false}
}

func (l *Lexer) GetTokens() []Token {
	tokens := []Token{}
	for {
		tokens = append(tokens, l.nextLine()...)
		if l.done {
			break
		}
	}
	return tokens
}

func (l *Lexer) nextLine() []Token {
	for {
		line, err := l.r.ReadString('\n')
		if err == io.EOF && line == "" {
			l.done = true
			return []Token{Token{Type: EOF, Literal: "", Line: l.line, Column: l.col}}
		}
		l.line++

		// Handle fixed form: strip cols 1–6 and cols 73+, check col 7
		clean := l.stripLine(line)

		// If comment
		if strings.HasPrefix(clean, "*") {
			return []Token{Token{Type: COMMENT, Literal: strings.TrimSpace(clean), Line: l.line, Column: 7}}
		}

		// Tokenize remaining
		tokens := l.tokenize(clean)
		if len(tokens) == 0 {
			continue
		}
		return tokens
	}
}

// Strip COBOL line formatting (seq + indicator, limit to col 72)
func (l *Lexer) stripLine(line string) string {
	// Ensure at least 8 chars for safety
	if len(line) < 8 {
		return ""
	}
	end := min(72, len(line))

	// Check column 7 for comment indicator
	indicator := line[6] // index 6 == col 7
	if indicator == '*' {
		return "*" + strings.TrimRight(line[7:end], "\r\n")
	}

	// Extract columns 8–72 (index 7..72)
	return strings.TrimRight(line[7:end], "\r\n")
}

// Tokenize a single line into tokens
func (l *Lexer) tokenize(s string) []Token {
	var tokens []Token
	col := 8 // actual COBOL text starts at col 8

	reader := strings.NewReader(s)
	for {
		ch, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}

		// Skip whitespace
		if unicode.IsSpace(ch) {
			col++
			continue
		}

		startCol := col

		// Identifiers / keywords
		if unicode.IsLetter(ch) {
			ident := string(ch)
			for {
				r, _, err := reader.ReadRune()
				if err == io.EOF || !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-') {
					if err == nil {
						reader.UnreadRune()
					}
					break
				}
				ident += string(r)
			}
			tokens = append(tokens, Token{Type: LookupIdentifier(ident), Literal: ident, Line: l.line, Column: startCol})
			continue
		}

		// Numbers
		if unicode.IsDigit(ch) {
			num := string(ch)
			for {
				r, _, err := reader.ReadRune()
				if err == io.EOF || !(unicode.IsDigit(r)) {
					if err == nil {
						reader.UnreadRune()
					}
					break
				}
				num += string(r)
			}
			tokens = append(tokens, Token{Type: NUMBER, Literal: num, Line: l.line, Column: startCol})
			continue
		}

		// Strings (single or double quotes)
		if ch == '"' || ch == '\'' {
			quote := ch
			lit := ""
			for {
				r, _, err := reader.ReadRune()
				if err == io.EOF || r == quote {
					break
				}
				lit += string(r)
			}
			tokens = append(tokens, Token{Type: STRING, Literal: lit, Line: l.line, Column: startCol})
			continue
		}

		// Symbols
		switch ch {
		case '.':
			tokens = append(tokens, Token{Type: PERIOD, Literal: ".", Line: l.line, Column: startCol})
		case ',':
			tokens = append(tokens, Token{Type: COMMA, Literal: ",", Line: l.line, Column: startCol})
		case '(':
			tokens = append(tokens, Token{Type: LPAREN, Literal: "(", Line: l.line, Column: startCol})
		case ')':
			tokens = append(tokens, Token{Type: RPAREN, Literal: ")", Line: l.line, Column: startCol})
		case '+':
			tokens = append(tokens, Token{Type: PLUS, Literal: "+", Line: l.line, Column: startCol})
		case '-':
			tokens = append(tokens, Token{Type: MINUS, Literal: "-", Line: l.line, Column: startCol})
		case '*':
			tokens = append(tokens, Token{Type: TIMES, Literal: "*", Line: l.line, Column: startCol})
		case '/':
			tokens = append(tokens, Token{Type: DIVIDE, Literal: "/", Line: l.line, Column: startCol})
		default:
			tokens = append(tokens, Token{Type: ILLEGAL, Literal: string(ch), Line: l.line, Column: startCol})
		}
	}

	return tokens
}
