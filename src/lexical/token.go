package lexical

import (
	"fmt"
)

type Token struct {
	Name       string
	Expression string
	ValueFound string
}

var Function = Token{Name: "Func", Expression: `\s*(func)\s+`}
var If = Token{Name: "If", Expression: fmt.Sprintf(`\s+(if)\s*%s`, OpenParenthese.Expression)}
var Else = Token{Name: "Else", Expression: `\s*(else)\s*\(`}
var OpenParenthese = Token{Name: "OpenParenthese", Expression: `\(`}
var CloseParenthese = Token{Name: "CloseParenthese", Expression: `\)`}
var Equality = Token{Name: "Equality", Expression: `\=\=`}
var Inequality = Token{Name: "Inequality", Expression: `\!\=`}
var Less = Token{Name: "Less", Expression: `\<`}
var LessEqual = Token{Name: "LessEqual", Expression: `\<\=`}
var Great = Token{Name: "Great", Expression: `\>`}
var GreatEqual = Token{Name: "GreatEqual", Expression: `\>\=`}
var Not = Token{Name: "Not", Expression: `\!`}
var OpenBrace = Token{Name: "OpenBrace", Expression: `\{`}
var CloseBrace = Token{Name: "CloseBrace", Expression: `\}`}
var Dot = Token{Name: "Dot", Expression: `\.`}
var And = Token{Name: "And", Expression: `\s+\&\&\s+`}
var Or = Token{Name: "Or", Expression: `\s+\|\|\s+`}
var Identifier = Token{Name: "Identifier", Expression: `Identifier\s+`}

//var IdentifierName = Token{Name: "IdentifierName", Expression: "[a-zA-Z]\\w*"}
var Attribute = Token{Name: "Attribute", Expression: `Attribute\s+`}

//var AttributeName = Token{Name: "AttributeName", Expression: "[a-zA-Z]\\w*"}
var Value = Token{Name: "Value", Expression: `Value\s+`}
var IntValue = Token{Name: "IntValue", Expression: `\s+[0-9]+`}
var FloatValue = Token{Name: "FloatValue", Expression: `\s+[0-9]+.[0-9+]`}
var StringValue = Token{Name: "StringValue", Expression: `\"\w*\"`}
var Plus = Token{Name: "Plus", Expression: `\+`}
var Minus = Token{Name: "Minus", Expression: `\-`}
var Multiply = Token{Name: "Multiply", Expression: `\*`}
var Divide = Token{Name: "Divide", Expression: `\/`}
var Rest = Token{Name: "Rest", Expression: `\%`}
var Save = Token{Name: "Save", Expression: `save\s+`}
var Delete = Token{Name: "Delete", Expression: `delete\s+`}
var Get = Token{Name: "Get", Expression: `get\s+`}
var Return = Token{Name: "Return", Expression: `return\s+`}
var Input = Token{Name: "Input", Expression: `input\.[a-zA-Z]+\s+`}
var Memory = Token{Name: "Memory", Expression: `memory\.[a-zA-Z]+\s+`}

func GetAllTokens() []Token {
	tokens := []Token{
		Function,
		If,
		Else,
		OpenParenthese,
		CloseParenthese,
		Equality,
		Inequality,
		Less,
		LessEqual,
		Great,
		GreatEqual,
		Not,
		OpenBrace,
		CloseBrace,
		And,
		Or,
		Identifier,
		//	IdentifierName,
		Attribute,
		//	AttributeName,
		Value,
		IntValue,
		FloatValue,
		StringValue,
		Plus,
		Minus,
		Multiply,
		Divide,
		Rest,
		Save,
		Delete,
		Get,
		Return,
		Input,
		Memory,
	}
	return tokens
}
