package lexical

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

// LexicalAnalyzer Structure used to analyze rules statements.
type LexicalAnalyzer struct {
	InputRule   string
	TokensFound []Token
}

func (l LexicalAnalyzer) Tokenize() error {
	tokensFound := []Token{}
	word := ""
	for i, c := range l.InputRule {
		word = word + string(c)
		tokensCandidate, err := l.GetTokensCandidate(word)

		if err != nil {
			logrus.Error("Error in match tokens in Lexical Analyzer:%v", err)
			return err
		}

		if len(tokensCandidate) > 0 {
			for _, tokenCandidate := range tokensCandidate {
				if tokenCandidate.Name != Entity.Name {
					tokensFound = append(
						tokensFound,
						Token{
							Name:       tokenCandidate.Name,
							Expression: tokenCandidate.Expression,
							ValueFound: strings.TrimSpace(word),
						},
					)
					word = ""
				} else {
					if i+1 < len(l.InputRule) {
						nextTokensCandidate, err := l.GetTokensCandidate(string(l.InputRule[i+1]))
						if err != nil {
							logrus.Error("Error in match tokens in Lexical Analyzer:%v", err)
							return err
						}
						for _, nextTokenCandidate := range nextTokensCandidate {
							if nextTokenCandidate.Name != Entity.Name && nextTokenCandidate.Name != IntValue.Name {
								tokensFound = append(
									tokensFound,
									Token{
										Name:       tokenCandidate.Name,
										Expression: tokenCandidate.Expression,
										ValueFound: strings.TrimSpace(word),
									},
								)
								word = ""
							} else {
								break
							}
						}
					} else {
						tokensFound = append(tokensFound, Err)
					}
				}
			}
		}
	}

	l.TokensFound = tokensFound

	for _, f := range tokensFound {
		logrus.Info(fmt.Sprintf("%s : %s", f.Name, f.ValueFound))
	}

	return nil
}

func (l LexicalAnalyzer) GetTokensCandidate(Word string) ([]Token, error) {
	tokens := GetAllTokens()
	tokensCandidate := []Token{}
	for _, token := range tokens {
		if token.Expression != "" {
			matched, err := regexp.Match(fmt.Sprintf(`%s`, token.Expression), []byte(Word))
			if err != nil {
				return nil, err
			}
			if matched {
				tokensCandidate = append(
					tokensCandidate,
					Token{
						Name:       token.Name,
						Expression: token.Expression,
						ValueFound: Word,
					},
				)
			}
		}
	}

	return tokensCandidate, nil
}

func TokenContains(tokens []Token, targetToken Token) bool {
	for _, token := range tokens {
		if token.Name == targetToken.Name {
			return true
		}
	}
	return false
}
