package lexer

import (
	"strings"
	"fmt"
)

type LexFn func(*Lexer) LexFn

/*
This lexer function starts everything off. It determines if we are
beginning with a bracket, a key or facts.
*/
func LexBegin(this *Lexer) LexFn {
  this.SkipWhitespace()

  if strings.HasPrefix(this.InputToEnd(), LEFT_BRACKET) {
    return LexLeftBracket
  } else if strings.HasPrefix(this.InputToEnd(), EQUALS) {
    return LexEquals
  } else if strings.HasPrefix(this.InputToEnd(), QUERY) {
    return LexQuery
  } else {
  	return LexKey
  }
}

func LexQuery(this *Lexer) LexFn {
	this.Pos += len(QUERY)
	this.Emit(TOKEN_QUERY)
	return LexKeyQuery
}

func LexKeyQuery(this *Lexer) LexFn {
	this.Inc()
	r := this.Next()
	if strings.ContainsRune(KEYS, r) {
		this.Emit(TOKEN_KEY)
		return LexKeyQuery
	} else if r == rune(10) {
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
  this.Pos += len(LEFT_BRACKET)
  this.BracketsCount += 1
  this.Emit(TOKEN_LEFT_BRACKET)
  return LexKey
}

func LexRightBracket(this *Lexer) LexFn {
  this.Pos += len(RIGHT_BRACKET)
  this.BracketsCount -= 1
  if this.BracketsCount < 0 {
  	return LexError
  }
  this.Emit(TOKEN_RIGHT_BRACKET)
  return LexSymbol
}

func LexKey(this *Lexer) LexFn {
	r := this.Next()
	if strings.ContainsRune(KEYS, r) {
		this.Inc()
		this.Emit(TOKEN_KEY)
		return LexSymbol
	} else if strings.HasPrefix(this.InputToEnd(), LEFT_BRACKET) {
    	return LexLeftBracket
    } else {
    	this.Error = &LexingError{this, fmt.Sprintf("'%c'", r), "capital letter or '('", this.Line, this.PosInLine()}
    	return LexError
    }
}

func LexSymbol(this *Lexer) LexFn {
	if strings.HasPrefix(this.InputToEnd(), IMPLIES) {
    	return LexImplies
    } else if strings.HasPrefix(this.InputToEnd(), IF_ONLY_IF) {
    	return LexIfOnlyIf
    }
    return LexOperator
}

func LexImplies(this *Lexer) LexFn {
	this.Pos += len(IMPLIES)
	this.Emit(TOKEN_IMPLIES)
	return LexResult
}

func LexIfOnlyIf(this *Lexer) LexFn {
	this.Pos += len(IF_ONLY_IF)
	this.Emit(TOKEN_IF_ONLY_IF)
	return LexResult
}

func LexOperator(this *Lexer) LexFn {
	r := this.Next()
	if strings.ContainsRune(OPERATORS, r) {
		this.Emit(TOKEN_OPERATOR)
		return LexKey
	}
	this.Error = &LexingError{this, fmt.Sprintf("'%c'", r), "operator", this.Line, this.PosInLine()}
	return LexError
}

func LexResult(this *Lexer) LexFn {
	r := this.Next()
	if strings.ContainsRune(KEYS, r) {
		this.Emit(TOKEN_KEY)
	}
	if strings.ContainsRune(OPERATORS, r) {
		this.Inc()
		this.Emit(TOKEN_OPERATOR)
		r = this.Next()
		if strings.ContainsRune(KEYS, r) {
			this.Emit(TOKEN_KEY)
		} else {
			this.Error = &LexingError{this, fmt.Sprintf("'%c'", r), "capital letter", this.Line, this.PosInLine()}
			return LexError
		}
	}
	r = this.Next()
	if r == rune(10) {
		this.Emit(TOKEN_EOL)
		return LexBegin
	}
	this.Error = &LexingError{this, fmt.Sprintf("'%c'", r), "newline or EOF", this.Line, this.PosInLine()}
	return LexError
}

func LexEquals(this *Lexer) LexFn {
  	this.Pos += len(EQUALS)
  	this.Emit(TOKEN_EQUALS)
  	return LexFact
}

func LexFact(this *Lexer) LexFn {
	if strings.ContainsRune(KEYS, this.Next()) {
		this.Pos += 1
		this.Emit(TOKEN_KEY)
		return LexFact
	} else {
		return LexEnd
	}
}

func LexError(this *Lexer) LexFn {
	close(this.Tokens)
	return nil
}

func LexEnd(this *Lexer) LexFn {
	close(this.Tokens)
	return nil
}

