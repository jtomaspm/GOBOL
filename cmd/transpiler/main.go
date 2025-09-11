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
	for _, token := range l.GetTokens() {
		fmt.Printf("%+v\n", token)
	}
}
