package lexer

import (
  "unicode"
	"unicode/utf8"
  "fmt"
)

type Lexer struct {
	Name			     string
  Line           int
  PosToLine      int
	Input			     string
	Tokens			   chan LexToken
	State			     LexFn
	BracketsCount	 int
  Error          *LexingError

	Start			int
	Pos				int
}

type LexingError struct {
  lexer    *Lexer
  Expected string
  Got      string
  Line     int
  Pos     int
}

func (this *LexingError) Error() string {
  return fmt.Sprintf("LexingError in %s:l%d:c%d Something when wrong while processing input data, got %s when expecting %s./n",
                    this.lexer.Name, this.Line, this.Pos, this.Got, this.Expected)
}

func (this *Lexer) PosInLine() int {
  return this.Start - this.PosToLine
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
  r, _ = utf8.DecodeRuneInString(this.InputToEnd())
  return
}

/*
Increment Pos of the size of the next rune
Pushes EOF token to the channel if we reached end of input
*/
func (this *Lexer) Inc() {
  _, size := utf8.DecodeRuneInString(this.InputToEnd())
  this.Pos += size
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

    if !unicode.IsSpace(r) {
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
func BeginLexing(input string) *Lexer {
  l := &Lexer{
    Input:  input,
    State:  LexBegin,
    Tokens: make(chan LexToken, 2),
  }

  return l
}

/*
Returns next item if there is one, otherwise move one step ahead
*/
func (this *Lexer) NextToken() LexToken {
  for {
    select {
    case t := <-this.Tokens:
      return t
    default:
      this.State = this.State(this)
    }
  }
  panic("WTF?! Out of infinit loop.")
}