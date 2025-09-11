package lexer

import "strings"

var keywords = map[string]TokenType{
	"IDENTIFICATION":  KEYWORD,
	"DIVISION":        KEYWORD,
	"PROGRAM-ID":      KEYWORD,
	"INSTALLATION":    KEYWORD,
	"DATE-WRITTEN":    KEYWORD,
	"ENVIRONMENT":     KEYWORD,
	"CONFIGURATION":   KEYWORD,
	"SPECIAL-NAMES":   KEYWORD,
	"INPUT-OUTPUT":    KEYWORD,
	"SECTION":         KEYWORD,
	"FILE-CONTROL":    KEYWORD,
	"DATA":            KEYWORD,
	"WORKING-STORAGE": KEYWORD,
	"PIC":             KEYWORD,
	"VALUE":           KEYWORD,
	"PROCEDURE":       KEYWORD,
	"DISPLAY":         KEYWORD,
	"EXIT":            KEYWORD,
}

func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[strings.ToUpper(identifier)]; ok {
		return tok
	}
	return IDENTIFIER
}
