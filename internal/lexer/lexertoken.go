package lexer

type TokenType int

const (
	LEFT_BRACKET  string = "("
	RIGHT_BRACKET string = ")"
	IMPLIES       string = "=>"
	IF_ONLY_IF    string = "<=>"
	EQUALS        string = "="
	QUERY         string = "?"
	FALSE         string = "!"
	KEYS          string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	OPERATORS     string = "+|^"

	TOKEN_ERROR TokenType = iota
	TOKEN_EOL             // 10
	TOKEN_EOF             // 11

	TOKEN_LEFT_BRACKET  // 12
	TOKEN_RIGHT_BRACKET // 13

	TOKEN_KEY        // 14
	TOKEN_FALSE      // 15
	TOKEN_EQUALS     // 16
	TOKEN_OPERATOR   // 17
	TOKEN_IMPLIES    // 18
	TOKEN_IF_ONLY_IF // 19
	TOKEN_QUERY      // 20
)

type LexToken struct {
	Type  TokenType
	Value string
}

func (t LexToken) String() string {
        return t.Value
}
