package lexical

//Token represents the string that will be matched by regex when tokenizer method is running.
type Token struct {
	Name       string
	Expression string
	ValueFound string
}

var Apply = Token{Name: "Apply", Expression: `\W+apply\W+`}
var Propose = Token{Name: "Propose", Expression: `\W+propose\W+`}
var If = Token{Name: "If", Expression: `\W+if\W*`}
var Else = Token{Name: "Else", Expression: `\W*else\W*`}
var Save = Token{Name: "Save", Expression: `\W+save\W+`}
var Delete = Token{Name: "Delete", Expression: `\W+delete\W+`}
var Fetch = Token{Name: "Fetch", Expression: `\W+fetch\W+`}
var Where = Token{Name: "Where", Expression: `\W+where\W+`}
var Null = Token{Name: "Null", Expression: `\W+null\W+`}
var Return = Token{Name: "Return", Expression: `\W+return\W+`}
var In = Token{Name: "In", Expression: `\W+in\W+`}
var Identifier = Token{Name: "Identifier", Expression: `\W+Identifier\W+`}
var Attribute = Token{Name: "Attribute", Expression: `\W+Attribute\W+`}
var Value = Token{Name: "Value", Expression: `\W+Value\W+`}

var OpenParentheses = Token{Name: "OpenParentheses", Expression: `\(`}
var CloseParentheses = Token{Name: "CloseParentheses", Expression: `\)`}
var Equal = Token{Name: "Equal", Expression: `\=`}
var Less = Token{Name: "Less", Expression: `\<`}
var Great = Token{Name: "Great", Expression: `\>`}
var Not = Token{Name: "Not", Expression: `\!`}
var Comma = Token{Name: "Comma", Expression: `\,`}
var OpenBrace = Token{Name: "OpenBrace", Expression: `\{`}
var CloseBrace = Token{Name: "CloseBrace", Expression: `\}`}
var DoubleQuotes = Token{Name: "DoubleQuotes", Expression: `\"`}
var Comparator = Token{Name: "Comparator", Expression: `\s*(\&\&|\|\|\|)\s*`}

var Entity = Token{Name: "Entity", Expression: `\w+(\.\w+)*`}
var IntValue = Token{Name: "IntValue", Expression: `\s+[0-9]+\s+`}
var FloatValue = Token{Name: "FloatValue", Expression: `\s+[0-9]+.[0-9+]\s+`}
var StringValue = Token{Name: "StringValue", Expression: `\"(.*?)\"\s*`}
var Operator = Token{Name: "Operator", Expression: `\s*(\+|\-|\*|/|\%)\s*`}

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
		Comparator,
		Entity,
		Identifier,
		Attribute,
		Value,
		DoubleQuotes,
		IntValue,
		FloatValue,
		Operator,
		Save,
		Delete,
		Fetch,
		Apply,
		Propose,
		Where,
		Return,
		Null,
		In,
		Err,
		SubSequence,
	}
}
