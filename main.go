package main

import (
	"fmt"
	"os"

	"github.com/EdsonHTJ/jtos/lexer"
)

func main() {
	jsrt, err := os.ReadFile("./example.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsrt))
	lexer := lexer.New()
	tokens, err := lexer.GetTokens(string(jsrt))
	if err != nil {
		panic(err)
	}

	for _, token := range tokens {
		fmt.Println("Token: ", token.Value)
		fmt.Println("Type: ", token.Type)
		fmt.Println("====================================")

	}
}
