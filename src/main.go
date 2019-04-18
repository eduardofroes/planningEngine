package main

import (
	"fmt"
	"io/ioutil"
	"planningEngine/src/lexical"
)

func main() {

	fileBytes, err := ioutil.ReadFile("//Users/du/go/src/planningEngine/resource/rulesExemples/rules.per")

	if err != nil {
		fmt.Print(err)
	}

	lex := lexical.LexicalAnalyzer{InputRule: string(fileBytes)}
	lex.Tokenize()
}
