package lexer

import (
	"fmt"
	"strings"
)

type LexFn func(*Lexer) LexFn

const RuneError = '\uFFFD' // the "error" Rune or "Unicode replacement character"
const Debug = false

/*
Jumps all spaces before to transit to return the next state passed as parameter
*/
func LexFnSpacesJumpWrapper(this *Lexer, fn LexFn) LexFn {
	this.SkipWhitespace()
	r := this.Peek()
	if r == '#' {
		for ; r != '\n' && r != RuneError; r = this.Peek() {
			this.Jump(r)
		}
	}
	return fn
}

/*
This lexer function starts everything off. It determines which state
should be next considering the first runes in the string.
Does not jumps Spaces, need to be wrapped in LexFnSpacesJumpWrapper at first launch
*/
func LexBegin(this *Lexer) LexFn {
	if Debug {
		println("Start LexBegin")
	}

	str := this.InputToEnd()
	if strings.HasPrefix(str, LEFT_BRACKET) {
		return LexFnSpacesJumpWrapper(this, LexLeftBracket)
	} else if !this.Facts && strings.HasPrefix(str, EQUALS) {
		this.Facts = true
		return LexFnSpacesJumpWrapper(this, LexEquals)
	} else if !this.Query && strings.HasPrefix(str, QUERY) {
		this.Query = true
		return LexFnSpacesJumpWrapper(this, LexQuery)
	} else if strings.HasPrefix(str, "\n") {
		this.Jump('\n')
		this.NewLine()
		return LexFnSpacesJumpWrapper(this, LexBegin)
	} else {
		return LexFnSpacesJumpWrapper(this, LexKey)
	}
}

/*
Emits a FALSE token then returns LexKey, to determine what is false.
*/
func LexFalse(this *Lexer) LexFn {
	if Debug {
		println("Start LexFalse")
	}
	this.Pos += len(FALSE)
	this.Emit(TOKEN_FALSE)
	return LexFnSpacesJumpWrapper(this, LexKey)
}

/*
Emits QUERY token then transit to LexKeyQuery state
*/
func LexQuery(this *Lexer) LexFn {
	if Debug {
		println("Start LexQuery")
	}
	this.Pos += len(QUERY)
	this.Emit(TOKEN_QUERY)
	return LexFnSpacesJumpWrapper(this, LexKeyQuery)
}

/*
Inside a QUERY line, emits a KEY token if there is then recursively transits or LexEnd
*/
func LexKeyQuery(this *Lexer) LexFn {
	if Debug {
		println("Start LexKeyQuery")
	}
	if this.ParseKey() {
		this.Emit(TOKEN_KEY)
		return LexFnSpacesJumpWrapper(this, LexKeyQuery)
	}
	return LexFnSpacesJumpWrapper(this, LexEndLine) // To check, bad shit could be afterwards
}

/*
This lexer function emits a TOKEN_LEFT_BRACKET then returns
the lexer for a key.
*/
func LexLeftBracket(this *Lexer) LexFn {
	if Debug {
		println("Start LexLeftBracket")
	}
	this.Pos += len(LEFT_BRACKET)
	this.BracketsCount += 1
	this.Emit(TOKEN_LEFT_BRACKET)
	return LexFnSpacesJumpWrapper(this, LexKey)
}

/*
This lexer function emits a TOKEN_RIGHT_BRACKET then returns
the lexer for a key.
*/
func LexRightBracket(this *Lexer) LexFn {
	if Debug {
		println("Start LexRightBracket")
	}
	this.Pos += len(RIGHT_BRACKET)
	this.BracketsCount -= 1
	if this.BracketsCount < 0 {
		this.Error = &LexingError{this, fmt.Sprintf("'%c'", this.Peek()), "capital letter or '('", this.Line, this.PosInLine()}
		return LexError
	}
	this.Emit(TOKEN_RIGHT_BRACKET)
	return LexFnSpacesJumpWrapper(this, LexSymbol)
}

/*
Main Key Lexer, transits to False if needed, then emits
either Key (to transit to symbol) or transits to LeftBracket
*/
func LexKey(this *Lexer) LexFn {
	if Debug {
		println("Start LexKey")
	}
	str := this.InputToEnd()
	if strings.HasPrefix(str, FALSE) {
		return LexFalse
	} else if this.ParseKey() {
		this.Emit(TOKEN_KEY)
		return LexFnSpacesJumpWrapper(this, LexSymbol)
	} else if strings.HasPrefix(str, LEFT_BRACKET) {
		return LexLeftBracket
	} else {
		this.Error = LexingError{this, fmt.Sprintf("'%c'", this.Peek()), "capital letter or '('"}
		return LexError
	}
}

/*

*/
func LexSymbol(this *Lexer) LexFn {
	if Debug {
		println("Start LexSymbol")
	}
	str := this.InputToEnd()
	if strings.HasPrefix(str, IMPLIES) {
		return LexImplies
	} else if strings.HasPrefix(str, IF_ONLY_IF) {
		return LexIfOnlyIf
	} else if strings.HasPrefix(str, RIGHT_BRACKET) {
		return LexRightBracket
	}
	return LexOperator
}

func LexImplies(this *Lexer) LexFn {
	if Debug {
		println("Start LexImplies")
	}
	if this.BracketsCount != 0 {
		this.Error = &LexingError{this, IMPLIES, RIGHT_BRACKET, this.Line, this.PosInLine()}
	}
	this.Pos += len(IMPLIES)
	this.Emit(TOKEN_IMPLIES)
	return LexFnSpacesJumpWrapper(this, LexResult)
}

func LexIfOnlyIf(this *Lexer) LexFn {
	if Debug {
		println("Start LexIfOnlyIf")
	}
	if this.BracketsCount != 0 {
		this.Error = &LexingError{this, IF_ONLY_IF, RIGHT_BRACKET, this.Line, this.PosInLine()}
	}
	this.Pos += len(IF_ONLY_IF)
	this.Emit(TOKEN_IF_ONLY_IF)
	return LexFnSpacesJumpWrapper(this, LexResult)
}

func LexOperator(this *Lexer) LexFn {
	if Debug {
		println("Start LexOperator")
	}
	r := this.Next()
	if strings.ContainsRune(OPERATORS, r) {
		this.Emit(TOKEN_OPERATOR)
		return LexFnSpacesJumpWrapper(this, LexKey)
	}
	this.Error = LexingError{this, fmt.Sprintf("'%c'", r), "operator"}
	return LexError
}

func LexResult(this *Lexer) LexFn {
	if Debug {
		println("Start LexResult")
	}
	if this.ParseKey() {
		this.Emit(TOKEN_KEY)
	} else {
		this.Error = LexingError{this, fmt.Sprintf("'%c'", this.Peek()), "capital letter"}
		return LexError
	}
	this.SkipWhitespace()
	if strings.ContainsRune(OPERATORS, this.Peek()) {
		if Debug {
			println("\tOperator found")
		}
		this.Inc()
		this.Emit(TOKEN_OPERATOR)
		this.SkipWhitespace()
		if this.ParseKey() {
			if Debug {
				println("\tKey found")
			}
			this.Emit(TOKEN_KEY)
		} else {
			this.Error = LexingError{this, fmt.Sprintf("'%c'", this.Peek()), "capital letter"}
			return LexError
		}
	}
	return LexFnSpacesJumpWrapper(this, LexEndLine)
}

func LexEquals(this *Lexer) LexFn {
	if Debug {
		println("Start LexEquals")
	}
	this.Pos += len(EQUALS)
	this.Emit(TOKEN_EQUALS)
	return LexFnSpacesJumpWrapper(this, LexFact)
}

func LexFact(this *Lexer) LexFn {
	if Debug {
		println("Start LexFact")
	}
	if this.ParseKey() {
		this.Emit(TOKEN_KEY)
		return LexFnSpacesJumpWrapper(this, LexFact)
	} else {
		return LexEndLine
	}
}

func LexEndLine(this *Lexer) LexFn {
	if Debug {
		println("Start LexEndLine")
	}
	if this.Peek() == '\n' {
		if Debug {
			println("\tNewline")
		}
		this.Inc()	
		this.Emit(TOKEN_EOL)
		this.NewLine()
		return LexFnSpacesJumpWrapper(this, LexBegin)
	} else if this.Peek() == RuneError {
		if Debug {
			println("EOF")
		}
		return LexEnd
	}
	this.Error = LexingError{this, fmt.Sprintf("'%c'", this.Peek()), "newline or EOF"}
	return LexError
}

func LexError(this *Lexer) LexFn {
	if Debug {
		println("Start LexError")
	}
	close(this.Tokens)
	return nil
}

func LexEnd(this *Lexer) LexFn {
	if this.Query && this.Facts {
		close(this.Tokens)
		return nil
	} else if !this.Facts {
		this.Error = &LexingError{this, "end of file", "Facts", this.Line, this.PosInLine()}
	} else {
		this.Error = &LexingError{this, "end of file", "Query", this.Line, this.PosInLine()}
	}
	return LexError
}