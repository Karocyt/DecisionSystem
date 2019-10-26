package lexer

type TokenType int

const (
	LEFT_BRACKET 	string = "("
	RIGHT_BRACKET 	string = ")"
	IMPLIES 		string = "=>"
	IF_ONLY_IF 		string = "<=>"
	EQUALS 			string = "="
	QUERY 			string = "?"
	KEYS 			string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	OPERATORS		string = "+!|^"

  	TOKEN_ERROR 	TokenType = iota
  	TOKEN_EOL
  	TOKEN_EOF

  	TOKEN_LEFT_BRACKET
  	TOKEN_RIGHT_BRACKET

  	TOKEN_KEY
  	TOKEN_EQUALS
  	TOKEN_OPERATOR
  	TOKEN_IMPLIES
  	TOKEN_IF_ONLY_IF
)

type Token struct {
  	Type  TokenType
  	Value string
}

