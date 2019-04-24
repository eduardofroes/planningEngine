package lexical

//Token represents the string that will be matched by regex when tokenizer method is running.
type Token struct {
	Name       string
	Expression string
	ValueFound string
}

var Apply = Token{Name: "Apply", Expression: `\s*apply\s+`}
var Propose = Token{Name: "Propose", Expression: `\s*propose\s+`}
var If = Token{Name: "If", Expression: `s+if\s+`}
var Else = Token{Name: "Else", Expression: `\s*else\s*`}
var OpenParentheses = Token{Name: "OpenParentheses", Expression: `\(`}
var CloseParentheses = Token{Name: "CloseParentheses", Expression: `\)`}
var Equal = Token{Name: "Equal", Expression: `\=`}
var Less = Token{Name: "Less", Expression: `\<`}
var Great = Token{Name: "Great", Expression: `\>`}
var Not = Token{Name: "Not", Expression: `\!`}
var Comma = Token{Name: "Comma", Expression: `\,`}
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
var StringValue = Token{Name: "StringValue", Expression: `\"(.*?)\"\s+`}
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
var Output = Token{Name: "Input", Expression: `\s*output(\.[a-zA-Z0-9]+)*\s+`}
var Memory = Token{Name: "Memory", Expression: `\s*memory(\.[a-zA-Z0-9]+)*\s+`}
var In = Token{Name: "In", Expression: `\s+in\s+`}
var Err = Token{Name: "Error"}
var SubSequence = Token{Name: "SubSequence"}

func GetAllTokens() []Token {
	return []Token{
		If,
		Else,
		OpenParentheses,
		CloseParentheses,
		Equal,
		Less,
		Great,
		Not,
		Comma,
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
		StringValue,
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
		Output,
		Memory,
		In,
		Err,
		SubSequence,
	}
}
