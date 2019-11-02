package lexer

import (
	"fmt"
	"github.com/fatih/color"
	"unicode/utf8"
	"unicode"
	"strings"
)

type Lexer struct {
	Name          string
	Line          int
	PosToLine     int
	Input         string
	Tokens        chan LexToken
	State         LexFn
	NextState     LexFn
	BracketsCount int
	Error         *LexingError
	LastRuneSize  int

	Start int
	Pos   int
}

type LexingError struct {
	Lexer    *Lexer
	Got      string
	Expected string
	Line     int
	Pos      int
}

var boldBlack *color.Color = color.New(color.Bold, color.FgBlack)
var boldRed *color.Color = color.New(color.Bold, color.FgRed)

func (this *LexingError) Error() string {
	return fmt.Sprintf("%s %s Something when wrong while processing input data, got %s when expecting %s.\n",
		boldBlack.Sprint(fmt.Sprintf("%s:%d:%d:", this.Lexer.Name, this.Line, this.Pos)),
		boldRed.Sprint("LexingError:"), boldBlack.Sprint(this.Got),
		boldBlack.Sprint(this.Expected))
}

func (this *Lexer) PosInLine() int {
	return this.Start - this.PosToLine
}

/*
Pushes a token onto the channel, from previous pos (Start) to current pos (Pos)
Takes a TokenType as parameter
*/
func (this *Lexer) Emit(tokenType TokenType) {
	if Debug {
		println("\tEmit")
	}
	t := LexToken{Type: tokenType, Value: this.Input[this.Start:this.Pos]}
	fmt.Printf("\t\tToken %d: %s\n", t.Type, t.Value)
	this.Tokens <- t
	this.Start = this.Pos
}

/*
Returns current rune at Pos
*/
func (this *Lexer) Peek() (r rune) {
	if Debug {
		println("\tPeek")
	}
	r, _ = utf8.DecodeRuneInString(this.InputToEnd())
	return
}

func (this *Lexer) Next() rune {
	if Debug {
		println("\tNext")
	}
	r, size := utf8.DecodeRuneInString(this.InputToEnd())
	this.Pos += size
	this.LastRuneSize = size
	// if this.Start+this.Pos >= len(this.Input) {
	// 	this.Emit(TOKEN_EOF)
	// }
	return r
}

func (this *Lexer) Jump(r rune) {
	if Debug {
		println("\tJump")
	}
	this.Pos += len(fmt.Sprintf("%c", r))
	this.Start = this.Pos
	return
}

/*
Increment Pos of the size of the next rune
Pushes EOF token to the channel if we reached end of input
*/
func (this *Lexer) Inc() {
	if Debug {
		println("\tInc")
	}
	_, size := utf8.DecodeRuneInString(this.InputToEnd())
	this.Pos += size
	// if this.Start+this.Pos >= len(this.Input) {
	// 	this.Emit(TOKEN_EOF)
	// }
}

func (this *Lexer) Back() {
	if Debug {
		println("\tBack")
	}
	if this.LastRuneSize == 0 {
		panic("lexer.LastRuneSize is 0, keep in mind that the lexer.Back() call works only for one rune.")
	}
	this.Pos -= this.LastRuneSize
	this.LastRuneSize = 0
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
		r := this.Peek()

		if r != '\t' && r != ' ' {
			break
		}
		this.Jump(r)
	}
}

func (this *Lexer) ParseKey() bool {
	r := this.Peek()
	if strings.ContainsRune(KEYS, r) {
		this.Inc()
		for unicode.IsLower(this.Peek()) {
			this.Pos += utf8.RuneLen(r)
		}
		return true
	}
	return false
}

/*
Start a new lexer with a given input string. This returns the
instance of the lexer and a channel of tokens. Reading this stream
is the way to parse a given input and perform processing.
BUFF_SIZE should be > 1 to be buffered and initializable at init stage
For memory footprint considerations, BUFF_SIZE should be kept as small as possible.
Hence 2 is the king choice and can be hardcoded.
*/
func BeginLexing(input string, name string) *Lexer {
	l := &Lexer{
		Name:   name,
		Input:  input,
		State:  LexBegin,
		Tokens: make(chan LexToken, 2),
	}
	go l.run()
	return l
}

/*
Returns next item if there is one, otherwise move one step ahead
*/
func (this *Lexer) run() {
	if Debug {
		println("Start go routine")
	}
	for this.State != nil {
		this.State = this.State(this)
	}
	//panic("WTF?! Out of infinit loop.")
}
