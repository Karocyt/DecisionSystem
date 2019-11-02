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
	if r == '\n' {
		if Debug {
			println("\tNewline")
		}
		this.Pos++
		this.Emit(TOKEN_EOL)
		return LexBegin
	} else if r == RuneError {
		if Debug {
			println("EOF")
		}
		this.Pos++
		return LexEnd
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
	this.SkipWhitespace()

	if strings.HasPrefix(this.InputToEnd(), LEFT_BRACKET) {
		return LexLeftBracket
	} else if strings.HasPrefix(this.InputToEnd(), EQUALS) {
		return LexEquals
	} else if strings.HasPrefix(this.InputToEnd(), QUERY) {
		return LexQuery
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
	// r := this.Peek()
	// if !strings.ContainsRune(KEYS, r) && !strings.HasPrefix(this.InputToEnd(), LEFT_BRACKET) {
	// 	this.Error = &LexingError{this, fmt.Sprintf("%c", r), "capital letter or '('", this.Line, this.PosInLine()}
	// }
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
	} else if this.Peek() == rune(10) {
		this.Error = &LexingError{this, "newline", "capital letter or EOF", this.Line, this.PosInLine()}
		return LexError
	}
	return LexEnd // To check, bad shit could be afterwards
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
		this.Error = &LexingError{this, fmt.Sprintf("'%c'", this.Peek()), "capital letter or '('", this.Line, this.PosInLine()}
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
	this.Error = &LexingError{this, fmt.Sprintf("'%c'", r), "operator", this.Line, this.PosInLine()}
	return LexError
}

func LexResult(this *Lexer) LexFn {
	if Debug {
		println("Start LexResult")
	}
	if this.ParseKey() {
		this.Emit(TOKEN_KEY)
	} else {
		this.Error = &LexingError{this, fmt.Sprintf("'%c'", this.Peek()), "capital letter", this.Line, this.PosInLine()}
		return LexError
	}
	this.SkipWhitespace()
	r := this.Next()
	if strings.ContainsRune(OPERATORS, r) {
		if Debug {
			println("\tOperator found")
		}
		this.Emit(TOKEN_OPERATOR)
		this.SkipWhitespace()
		if this.ParseKey() {
			if Debug {
				println("\tKey found")
			}
			this.Emit(TOKEN_KEY)
		} else {
			this.Error = &LexingError{this, fmt.Sprintf("'%c'", this.Peek()), "capital letter", this.Line, this.PosInLine()}
			return LexError
		}
		r = this.Next()
	}
	if r == '\n' {
		if Debug {
			println("\tNewline")
		}
		this.Emit(TOKEN_EOL)
		return LexBegin
	} else if r == RuneError {
		if Debug {
			println("EOF")
		}
		return LexEnd
	}
	this.Error = &LexingError{this, fmt.Sprintf("'%c'", r), "newline or EOF", this.Line, this.PosInLine()}
	return LexError
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
		return LexEnd
	}
}

func LexError(this *Lexer) LexFn {
	if Debug {
		println("Start LexError")
	}
	close(this.Tokens)
	return nil
}

func LexEnd(this *Lexer) LexFn {
	if Debug {
		println("Start LexEnd")
	}
	close(this.Tokens)
	return nil
}
