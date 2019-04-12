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
				for _, token := range GetAllTokens() {
					sent := regexp.MustCompile(fmt.Sprintf(`%s`, token.Expression))
					indexes := sent.FindAllIndex([]byte(ruleString), 2)
					//matched, err := regexp.Match(fmt.Sprintf(`\A(#{%s})/`, token), []byte(ruleString))
					fmt.Println(fmt.Sprintf("%s - %v", token, indexes))
				}
			}
		}
	}

	return []Token{}
}

func (l LexicalAnalyzer) Tokenize(ruleString string) {

	//tokens := GetAllTokens()

	word := ""
	for _, c := range ruleString {
		word = word + string(c)

	}

	// for _, token := range GetAllTokens() {
	// 	sent := regexp.MustCompile(fmt.Sprintf(`%s`, token.Expression))
	// 	indexes := sent.FindAllIndex([]byte(ruleString), 2)
	// 	//matched, err := regexp.Match(fmt.Sprintf(`\A(#{%s})/`, token), []byte(ruleString))
	// 	fmt.Println(fmt.Sprintf("%s - %v", token, indexes))
	// }
}
