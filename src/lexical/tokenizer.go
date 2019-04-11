package lexical

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

// LexicalAnalyzer Structure used to analyze rules statements.
type LexicalAnalyzer struct {
	InputRule string
}

func (l LexicalAnalyzer) Analyze(ruleString string) []Token {

	wordsRule := strings.Split(ruleString, " ")

	for _, frase := range wordsRule {
		word := ""
		for i, c := range frase {
			character := string(c)
			tokens := GetAllTokens()
			logrus.Info(tokens)

			word = word + character

			if len(ruleString) >= i+1 {
				logrus.Info(word)
			}
		}
	}

	return []Token{Token(0)}
}

func (l LexicalAnalyzer) Tokenize(ruleString string) {

	for _, token := range GetAllTokens() {
		sent := regexp.MustCompile(fmt.Sprintf(`\A(#{%s})`, token))
		indexes := sent.FindIndex([]byte(ruleString))
		//matched, err := regexp.Match(fmt.Sprintf(`\A(#{%s})/`, token), []byte(ruleString))
		fmt.Println(indexes)
	}
}
