package lexical

//Token represents the string that will be matched by regex when tokenizer method is running.
type Token struct {
	Name       string
	Expression string
	ValueFound string
}

var Apply = Token{Name: "Apply", Expression: `\s*apply\s+`}
var Propose = Token{Name: "Propose", Expression: `\s*propose\s+`}
var If = Token{Name: "If", Expression: `if\s+\(*`}
var Else = Token{Name: "Else", Expression: `\s*else\s*`}
var OpenParenthese = Token{Name: "OpenParenthese", Expression: `\(`}
var CloseParenthese = Token{Name: "CloseParenthese", Expression: `\)`}
var Equal = Token{Name: "Equal", Expression: `\=`}
var Less = Token{Name: "Less", Expression: `\<`}
var Great = Token{Name: "Great", Expression: `\>`}
var Not = Token{Name: "Not", Expression: `\!`}
var DoubleQuotes = Token{Name: "DoubleQuotes", Expression: `\"`}
var OpenBrace = Token{Name: "OpenBrace", Expression: `\{`}
var CloseBrace = Token{Name: "CloseBrace", Expression: `\}`}
var And = Token{Name: "And", Expression: `\s+\&\&\s+`}
var Or = Token{Name: "Or", Expression: `\s+\|\|\s+`}
var Identifier = Token{Name: "Identifier", Expression: `Identifier\s+`}
var Entity = Token{Name: "Entity", Expression: `[a-zA-Z0-9]+(\.[a-zA-Z0-9]+)*`}
var Attribute = Token{Name: "Attribute", Expression: `Attribute\s+`}
var Value = Token{Name: "Value", Expression: `Value\s+`}
var IntValue = Token{Name: "IntValue", Expression: `(\=\=|\=|\s+)[0-9]+\s+`}
var FloatValue = Token{Name: "FloatValue", Expression: `(\=\=|\=|\s+)[0-9]+.[0-9+]`}
var Plus = Token{Name: "Plus", Expression: `\+`}
var Minus = Token{Name: "Minus", Expression: `\-`}
var Multiply = Token{Name: "Multiply", Expression: `\*`}
var Divide = Token{Name: "Divide", Expression: `\/`}
var Rest = Token{Name: "Rest", Expression: `\%`}
var Save = Token{Name: "Save", Expression: `\s*save\s+`}
var Delete = Token{Name: "Delete", Expression: `\s*delete\s+`}
var Fetch = Token{Name: "Fetch", Expression: `\s*fetch\s+`}
var Where = Token{Name: "Where", Expression: `\s*where\s+`}
var Null = Token{Name: "Null", Expression: `\s*null\s+`}
var Return = Token{Name: "Return", Expression: `\s*return\s+`}
var Input = Token{Name: "Input", Expression: `\s*input(\.[a-zA-Z0-9]+)*\s+`}
var Memory = Token{Name: "Memory", Expression: `\s*memory(\.[a-zA-Z0-9]+)*\s+`}
var In = Token{Name: "In", Expression: `\s+(in)\s+`}
var Err = Token{Name: "Error"}

func GetAllTokens() []Token {
	tokens := []Token{
		If,
		Else,
		OpenParenthese,
		CloseParenthese,
		Equal,
		Less,
		Great,
		Not,
		DoubleQuotes,
		OpenBrace,
		CloseBrace,
		And,
		Or,
		Entity,
		Identifier,
		Attribute,
		Value,
		IntValue,
		FloatValue,
		Plus,
		Minus,
		Multiply,
		Divide,
		Rest,
		Save,
		Delete,
		Fetch,
		Apply,
		Propose,
		Where,
		Return,
		Null,
		Input,
		Memory,
		Err,
		In,
	}
	return tokens
}
