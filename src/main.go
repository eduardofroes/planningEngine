package main

import (
	"fmt"
	"io/ioutil"
	"planningEngine/lexical"
)

func main() {

	fileBytes, err := ioutil.ReadFile("/opt/repository/planningEngine/resource/rulesExemples/rules.per")
	if err != nil {
		fmt.Print(err)
	}

	lex := lexical.LexicalAnalyzer{InputRule: string(fileBytes)}
	lex.Analyze(lex.InputRule)
}
