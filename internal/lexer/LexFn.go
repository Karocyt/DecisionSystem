package lexer

import (
	"fmt"
	"strings"
)

type LexFn func(*Lexer) LexFn

const RuneError = '\uFFFD' // the "error" Rune or "Unicode replacement character"
const Debug = false

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
This lexer function starts everything off. It determines if we are
beginning with a bracket, a key or facts.
*/
func LexBegin(this *Lexer) LexFn {
	if Debug {
		println("Start LexBegin")
	}

	str := this.InputToEnd()
	if strings.HasPrefix(str, LEFT_BRACKET) {
		return LexFnSpacesJumpWrapper(this, LexLeftBracket)
	} else if strings.HasPrefix(str, EQUALS) {
		return LexFnSpacesJumpWrapper(this, LexEquals)
	} else if strings.HasPrefix(str, QUERY) {
		return LexFnSpacesJumpWrapper(this, LexQuery)
	} else if strings.HasPrefix(str, "\n") {
		this.Jump('\n')
		this.NewLine()
		return LexFnSpacesJumpWrapper(this, LexBegin)
	} else {
		return LexFnSpacesJumpWrapper(this, LexKey)
	}
}

func LexFalse(this *Lexer) LexFn {
	if Debug {
		println("Start LexFalse")
	}
	this.Pos += len(FALSE)
	this.Emit(TOKEN_FALSE)
	return LexFnSpacesJumpWrapper(this, LexKey)
}

func LexQuery(this *Lexer) LexFn {
	if Debug {
		println("Start LexQuery")
	}
	this.Pos += len(QUERY)
	this.Emit(TOKEN_QUERY)
	return LexFnSpacesJumpWrapper(this, LexKeyQuery)
}

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

func LexRightBracket(this *Lexer) LexFn {
	if Debug {
		println("Start LexRightBracket")
	}
	this.Pos += len(RIGHT_BRACKET)
	this.BracketsCount -= 1
	if this.BracketsCount < 0 {
		return LexError
	}
	this.Emit(TOKEN_RIGHT_BRACKET)
	return LexFnSpacesJumpWrapper(this, LexSymbol)
}

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
		this.Error = &LexingError{this, fmt.Sprintf("'%c'", this.Peek()), "capital letter or '('"}
		return LexError
	}
}

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
	this.Pos += len(IMPLIES)
	this.Emit(TOKEN_IMPLIES)
	return LexFnSpacesJumpWrapper(this, LexResult)
}

func LexIfOnlyIf(this *Lexer) LexFn {
	if Debug {
		println("Start LexIfOnlyIf")
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
	this.Error = &LexingError{this, fmt.Sprintf("'%c'", r), "operator"}
	return LexError
}

func LexResult(this *Lexer) LexFn {
	if Debug {
		println("Start LexResult")
	}
	if this.ParseKey() {
		this.Emit(TOKEN_KEY)
	} else {
		this.Error = &LexingError{this, fmt.Sprintf("'%c'", this.Peek()), "capital letter"}
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
			this.Error = &LexingError{this, fmt.Sprintf("'%c'", this.Peek()), "capital letter"}
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

func LexError(this *Lexer) LexFn {
	if Debug {
		println("Start LexError")
	}
	close(this.Tokens)
	return nil
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
		close(this.Tokens)
		return nil
	}
	this.Error = &LexingError{this, fmt.Sprintf("'%c'", this.Peek()), "newline or EOF"}
	return LexError
}
