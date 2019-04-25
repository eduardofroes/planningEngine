package syntactical

import (
	"planningEngine/lexical"
)

type SyntacticalAnalyzer struct {
	InputTokens         *[]lexical.Token
	SyntacticalSequence []Sequence
	LastAnalyze         bool
}

type Sequence struct {
	Name        string
	TokensCheck [][]lexical.Token
}

// GetAllSequence is responsible to gets all sequence which is allowed in syntactical analyzer.
func GetAllSequence() []Sequence {
	return []Sequence{
		{
			Name: lexical.If.Name,
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
		},
		{
			Name: lexical.Propose.Name,
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
		},
		{
			Name: lexical.Apply.Name,
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
		},
		{
			Name: lexical.Identifier.Name,
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
		},
		{
			Name: "Assignment",
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
					lexical.DoubleQuotes,
					lexical.Entity,
					lexical.DoubleQuotes,
				},
			},
		},
	}
}

func (s SyntacticalAnalyzer) SyntacticalAnalyzer() (bool, error) {

	return true, nil
}
