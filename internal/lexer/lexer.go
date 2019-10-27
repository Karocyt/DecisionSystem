package lexer

import (
  "unicode"
	"unicode/utf8"
)

type Lexer struct {
	Error			     string
	Input			     string
	Tokens			   chan LexToken
	State			     LexFn
	BracketsCount	 int

	Start			int
	Pos				int
	Width			int
}

/*
Pushes a token onto the channel, from previous pos (Start) to current pos (Pos)
Takes a TokenType as parameter  
*/
func (this *Lexer) Emit(tokenType TokenType) {
  this.Tokens <- LexToken{Type: tokenType, Value: this.Input[this.Start:this.Pos]}
  this.Start = this.Pos
}

/*
Returns current rune at Pos
*/
func (this *Lexer) Next() (r rune) {
  r, _ = utf8.DecodeRuneInString(this.InputToEnd)
  return
}

/*
Increment Pos of the size of the next rune
Pushes EOF token to the channel if we reached end of input
*/
func (this *Lexer) Inc() {
  _, size := utf8.DecodeRuneInString(this.InputToEnd)
  lexer.Pos += size
  if this.Start + this.Pos >= len(this.Input) {
    this.Emit(TOKEN_EOF)
  }
}

/*
Return a slice of the input from the current pos to the end of the input string.
*/
func (this *Lexer) InputToEnd() string {
  return this.Input[this.Pos:]
}

/*
Skips whitespace in infinite loop until we get something else or EOF.
(Could use this.Next() ?)
*/
func (this *Lexer) SkipWhitespace() {
  for {
    r := this.Next()

    if !unicode.IsSpace(ch) {
      break
    }
    this.Inc()
  }
}

/*
Start a new lexer with a given input string. This returns the
instance of the lexer and a channel of tokens. Reading this stream
is the way to parse a given input and perform processing.
BUFF_SIZE should be > 1 to be buffered and initializable at init stage
For memory footprint considerations, BUFF_SIZE should be kept as small as possible.
Hence 2 is the king choice and can be hardcoded.
*/
func BeginLexing(filename string, input string) *Lexer {
  l := &Lexer{
    Input:  input,
    State:  LexBegin,
    Tokens: make(chan LexToken, 2),
  }

  return l
}