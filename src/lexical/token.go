package lexical

type Token string

const (
	If              Token = "\bif\b"
	Else            Token = "\belse\b"
	OpenParenthese  Token = `\(`
	CloseParenthese Token = `\)`
	Equality        Token = "=="
	Inequality      Token = "!="
	Less            Token = "<"
	LessEqual       Token = "<="
	Great           Token = ">"
	GreatEqual      Token = ">="
	Not             Token = "!"
	OpenBrace       Token = `\{`
	CloseBrace      Token = `\}`
	Dot             Token = "."
	And             Token = "&&"
	Or              Token = "||"
	Identifier      Token = "\bIdentifier\b"
	IdentifierName  Token = "[a-zA-Z]\\w*"
	Attribute       Token = "\bAttribute\b"
	AttributeName   Token = "^[a-zA-Z]\\w*"
	Value           Token = "\bValue\b"
	IntValue        Token = "[0-9]+"
	FloatValue      Token = "[0-9]+.[0-9+]"
	StringValue     Token = "\"[a-zA-Z]\\w*\""
	Plus            Token = "+"
	Minus           Token = "-"
	Multiply        Token = "*"
	Divide          Token = "/"
	Rest            Token = "%"
	Save            Token = "\bsave\b"
	Delete          Token = "\bdelete\b"
	Get             Token = "\bget\b"
	Return          Token = "\breturn\b"
	Input           Token = "\binput.[a-zA-Z]+"
	Memory          Token = "\bmemory.[a-zA-Z]+"
)

func GetAllTokens() []Token {
	tokens := []Token{
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
		IdentifierName,
		Attribute,
		AttributeName,
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
