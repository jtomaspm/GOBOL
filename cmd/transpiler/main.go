package main

import (
	"fmt"
	"os"

	"github.com/jtomaspm/GOBOL/pkg/lexer"
)

func main() {
	f, err := os.Open("/home/pop/Code/GOBOL/resources/cobol/TEST0001.PCOD")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	l := lexer.New(f)
	for {
		tok := l.NextToken()
		fmt.Printf("%+v\n", tok)
		if tok.Type == lexer.EOF {
			break
		}
	}
}
