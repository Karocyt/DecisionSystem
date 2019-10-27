package lexer

import (
	"unicode"
)

type Lexer struct {
	Name			     string
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
Returns current rune and increment pos by one rune
Pushes EOF token to the channel if we reached end of input
*/
func (this *Lexer) Next() {
  r, size := utf8.DecodeRuneInString(lexer.Input[lexer.Start])
  lexer.Pos += size
  if this.Start + this.Pos >= len(this.Input) {
    this.Emit(TOKEN_EOF)
  }
  return r
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
    ch := this.Next()

    if !unicode.IsSpace(ch) {
      this.Dec()
      break
    }

    if ch == EOF {
      this.Emit(TOKEN_EOF)
      break
    }
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
  l := &lexer.Lexer{
    Name:   name,
    Input:  input,
    State:  lexer.LexBegin,
    Tokens: make(chan LexToken, 2),
  }

  return l
}