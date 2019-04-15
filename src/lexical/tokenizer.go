package lexical

import (
	"fmt"
	"regexp"
)

// LexicalAnalyzer Structure used to analyze rules statements.
type LexicalAnalyzer struct {
	InputRule string
	TokensFound []Token
}

func (l LexicalAnalyzer) Tokenize(ruleString string) ([]Token, error) {
	tokensFound := []Token{}
	word := ""
	for _, c := range ruleString {
		character := string(c)
		tokens := GetAllTokens()
		word = word + character
		for _, token := range tokens {
			matched, err := regexp.Match(fmt.Sprintf(`%s`, token.Expression), []byte(word))
			if err != nil {
				return nil, err
			}
			if matched {
				tokensFound = append(tokensFound, Token{Name:token.Name, Expression:token.Expression, ValueFound:word})
				word = ""
			}
		}
	}

	fmt.Println(fmt.Sprintf("Test: %v", tokensFound))

	return tokensFound, nil
}
