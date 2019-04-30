package lexical

import (
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

// LexicalAnalyzer Structure used to analyze rules statements.
type LexicalAnalyzer struct {
	FilesPath   []string
	InputRule   string
	TokensFound []Token
}

// Tokenize method is responsible for tokenize string into tokens.
func (l *LexicalAnalyzer) Tokenize() error {
	tokensFound := []Token{}
	word := " "
	for i, c := range wordProcessing(l.InputRule) {
		word = word + string(c)
		if len(strings.TrimSpace(word)) == 0 {
			continue
		}

		tokensCandidate, err := l.GetTokensCandidate(GetAllTokens(), word)
		if err != nil {
			logrus.Fatal("Error in match tokens in Lexical Analyzer: %s", err.Error())
			return err
		}

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
				word = " "
			} else {
				if i+1 < len(l.InputRule) {
					if len(strings.TrimSpace(word+" ")) == 0 {
						continue
					}

					if string(l.InputRule[i+1]) == " " {
						nextTokensCandidate, err := l.GetTokensCandidate(GetAllTokens(), word+" ")
						if err != nil {
							logrus.Fatal("Error in match tokens in Lexical Analyzer: %s", err.Error())
							return err
						}

						if len(nextTokensCandidate) == 1 {
							tokensFound = append(
								tokensFound,
								Token{
									Name:       tokenCandidate.Name,
									Expression: tokenCandidate.Expression,
									ValueFound: strings.TrimSpace(word),
								},
							)
							word = " "
						}
					} else {
						nextTokensCandidate, err := l.GetTokensCandidate(GetAllTokens(), string(l.InputRule[i+1]))
						if err != nil {
							logrus.Fatal("Error in match tokens in Lexical Analyzer: %s", err.Error())
							return err
						}

						for _, nextTokenCandidate := range nextTokensCandidate {
							if nextTokenCandidate.Name != Entity.Name &&
								nextTokenCandidate.Name != IntValue.Name {
								tokensFound = append(
									tokensFound,
									Token{
										Name:       tokenCandidate.Name,
										Expression: tokenCandidate.Expression,
										ValueFound: strings.TrimSpace(word),
									},
								)
								word = " "
							}
						}
					}
				}
			}
		}
	}

	l.TokensFound = tokensProcessing(&tokensFound)

	return nil
}

// GetTokensCandidate is a method that gets all tokens candidate which was matched by an regex.
func (l LexicalAnalyzer) GetTokensCandidate(Tokens []Token, Word string) ([]Token, error) {
	tokensCandidate := []Token{}

	if Tokens == nil {
		Tokens = GetAllTokens()
	}

	for _, token := range Tokens {
		if token.Expression != "" {
			matched, err := regexp.Match(token.Expression, []byte(Word))
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

func tokensProcessing(tokensFound *[]Token) []Token {
	processedTokens := replaceSequence(tokensFound, []Token{DoubleQuotes, Entity, DoubleQuotes}, StringValue)
	processedTokens = replaceSequence(&processedTokens, []Token{Equal, Equal}, ComparableSignal)
	processedTokens = replaceSequence(&processedTokens, []Token{Less, Equal}, ComparableSignal)
	processedTokens = replaceSequence(&processedTokens, []Token{Great, Equal}, ComparableSignal)
	processedTokens = replaceSequence(&processedTokens, []Token{Not, Equal}, ComparableSignal)

	return processedTokens
}

func replaceSequence(tokensFound *[]Token, tokenSequence []Token, tokenApply Token) []Token {
	tokensCounter := 0
	processedTokens := []Token{}
	previousTokens := []Token{}

	for _, tokenFound := range *tokensFound {
		if tokenFound.Name == tokenSequence[tokensCounter].Name {
			tokensCounter++
			if tokensCounter == len(tokenSequence) {
				previousTokens = append(previousTokens, tokenFound)
				newTokenApply := Token{Name: tokenApply.Name, Expression: tokenApply.Expression}

				for _, previousToken := range previousTokens {
					newTokenApply.ValueFound += previousToken.ValueFound
				}

				processedTokens = append(processedTokens, newTokenApply)
				previousTokens = []Token{}
				tokensCounter = 0
			} else {
				previousTokens = append(previousTokens, tokenFound)
			}
		} else {
			processedTokens = append(processedTokens, tokenFound)
			previousTokens = []Token{}
			tokensCounter = 0
		}
	}

	return processedTokens
}

func wordProcessing(word string) string {
	return strings.Replace(
		strings.Replace(
			strings.Replace(word, "\n", " ", -1),
			"\t",
			" ",
			-1,
		),
		"\r",
		" ",
		-1,
	)
}

//TokenContains verifies if there is a specific token in slice.
func TokenContains(tokens []Token, targetToken Token) bool {
	for _, token := range tokens {
		if token.Name == targetToken.Name {
			return true
		}
	}
	return false
}
