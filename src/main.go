package main

import (
	"fmt"
	"io/ioutil"
	"planningEngine/src/lexical"

	"github.com/sirupsen/logrus"
)

func main() {

	//fileBytes, err := ioutil.ReadFile("/opt/repository/planningEngine/resource/rulesExemples/rules.per")
	fileBytes, err := ioutil.ReadFile("/Users/du/go/src/planningEngine/resource/rulesExemples/rules.per")

	if err != nil {
		fmt.Print(err)
	}

	lex := lexical.LexicalAnalyzer{InputRule: string(fileBytes)}
	lex.Tokenize()

	for _, f := range lex.TokensFound {
		logrus.Info(fmt.Sprintf("%s : %s", f.Name, f.ValueFound))
	}
}
