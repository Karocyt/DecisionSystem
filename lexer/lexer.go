package lexer

import (
	"utf8"
	"unicode"
)

type Lexer struct {
	Name			string
	Input			string
	Tokens			chan lexertoken.Token
	State			LexFn
	BracketsCount	int

	Start			int
	Pos				int
	Width			int
}

/*
Pushes a token onto the channel, from previous pos (Start) to current pos (Pos)
Takes a TokenType as parameter  
*/
func (this *Lexer) Emit(tokenType lexertoken.TokenType) {
  this.Tokens <- lexertoken.Token{Type: tokenType, Value: this.Input[this.Start:this.Pos]}
  this.Start = this.Pos
}

/*
Returns current rune and increment pos by one rune
*/
func (this *Lexer) Next() {
  r := lexer.Input[lexer.Start]
  lexer.Inc()
  return r
}

/*
Increment current pos of one rune
Pushes EOF token to the channel if we reached RuneCountInString
*/
func (this *Lexer) Inc() {
  this.Pos++
  if this.Pos >= utf8.RuneCountInString(this.Input) {
    this.Emit(lexertoken.TOKEN_EOF)
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
(Could use this.Inc() ?)
*/
func (this *Lexer) SkipWhitespace() {
  for {
    ch := this.Next()

    if !unicode.IsSpace(ch) {
      this.Dec()
      break
    }

    if ch == lexertoken.EOF {
      this.Emit(lexertoken.TOKEN_EOF)
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
func BeginLexing(filename string, input string) *lexer.Lexer {
  l := &lexer.Lexer{
    Name:   name,
    Input:  input,
    State:  lexer.LexBegin,
    Tokens: make(chan lexertoken.Token, 2),
  }

  return l
}