package syntactical

import (
	"planningEngine/lexical"

	"github.com/sirupsen/logrus"
)

type SyntacticalAnalyzer struct {
	InputTokens         *[]lexical.Token
	SyntacticalSequence []Sequence
	LastAnalyze         bool
}

type Sequence struct {
	TokensCheck [][]lexical.Token
}

//SyntacticValidation is responsible for validate whether the rules are syntactically correct.
func (s *SyntacticalAnalyzer) SyntacticValidation(tokens []lexical.Token) bool {

	sequenceMap := GetAllSequence()
	for _, token := range tokens {
		tokensMatrix := sequenceMap[token.Name]
		logrus.Info("%v", tokensMatrix)
	}

	return true
}

// GetAllSequence is responsible to gets all sequence which is allowed in syntactical analyzer.
func GetAllSequence() map[string]Sequence {

	sequenceMap := make(map[string]Sequence)

	sequenceMap[lexical.If.Name] = Sequence{
		TokensCheck: [][]lexical.Token{
			{
				lexical.If,
				lexical.OpenParentheses,
				lexical.SubSequence,
				lexical.CloseParentheses,
				lexical.OpenBrace,
				lexical.SubSequence,
				lexical.CloseBrace,
			},
			{
				lexical.If,
				lexical.SubSequence,
				lexical.OpenBrace,
				lexical.SubSequence,
				lexical.CloseBrace,
			},
		},
	}

	sequenceMap[lexical.Propose.Name] = Sequence{
		TokensCheck: [][]lexical.Token{
			{
				lexical.Propose,
				lexical.Entity,
				lexical.OpenParentheses,
				lexical.SubSequence,
				lexical.CloseParentheses,
				lexical.OpenBrace,
				lexical.SubSequence,
				lexical.CloseBrace,
			},
		},
	}

	sequenceMap[lexical.Apply.Name] = Sequence{
		TokensCheck: [][]lexical.Token{
			{
				lexical.Apply,
				lexical.Entity,
				lexical.OpenParentheses,
				lexical.SubSequence,
				lexical.CloseParentheses,
				lexical.OpenBrace,
				lexical.SubSequence,
				lexical.CloseBrace,
			},
		},
	}

	sequenceMap[lexical.Identifier.Name] = Sequence{
		TokensCheck: [][]lexical.Token{
			{
				lexical.Identifier,
				lexical.Entity,
			},
			{
				lexical.Identifier,
				lexical.Entity,
				lexical.Comma,
			},
		},
	}

	sequenceMap["Assingment"] = Sequence{
		TokensCheck: [][]lexical.Token{
			{
				lexical.Entity,
				lexical.Equal,
				lexical.FloatValue,
			},
			{
				lexical.Entity,
				lexical.Equal,
				lexical.IntValue,
			},
			{
				lexical.Entity,
				lexical.Equal,
				lexical.StringValue,
			},
		},
	}

	return sequenceMap
}

func (s SyntacticalAnalyzer) SyntacticalAnalyzer() (bool, error) {

	return true, nil
}
