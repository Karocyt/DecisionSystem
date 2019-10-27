package lexer

import (
	"strings"
)

type LexFn func(*Lexer) LexFn

/*
This lexer function starts everything off. It determines if we are
beginning with a bracket, a key or facts.
*/
func LexBegin(lexer *Lexer) LexFn {
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

func LexQuery(lexer *Lexer) LexFn {
	this.Next()
	this.Emit(TOKEN_QUERY)
	return LexKeyQuery
}

func LexKeyQuery(lexer *Lexer) LexFn {
	this.Next()
	if strings.ContainsRune(KEYS, this.Input[this.Start]) {
		this.Emit(TOKEN_KEY)
		return LexKeyQuery
	} else if this.Input[this.Start] == `/n` {
		return LexError
	}
	return LexEnd // To check, bad shit could be afterwards
}

/*
This lexer function emits a TOKEN_LEFT_BRACKET then returns
the lexer for a key.
*/
func LexLeftBracket(lexer *Lexer) LexFn {
  this.Pos += len(LEFT_BRACKET)
  this.BracketCount += 1
  this.Emit(TOKEN_LEFT_BRACKET)
  return LexKey
}

func LexRightBracket(lexer *Lexer) LexFn {
  this.Pos += len(RIGHT_BRACKET)
  this.BracketCount -= 1
  if this.BracketCount < 0 {
  	return LexError
  }
  this.Emit(TOKEN_RIGHT_BRACKET)
  return LexSymbol
}

func LexKey(lexer *Lexer) LexFn {
	if strings.ContainsRune(KEYS, this.Input[this.Pos]) {
		this.Next()
		this.Emit(TOKEN_KEY)
		return LexSymbol
	} else if strings.HasPrefix(this.InputToEnd(), LEFT_BRACKET) {
    	return LexLeftBracket
    } else {
    	return LexError
    }
}

func LexSymbol(lexer *Lexer) LexFn {
	if strings.HasPrefix(this.InputToEnd(), IMPLIES) {
    	return LexImplies
    } else if strings.HasPrefix(this.InputToEnd(), IF_ONLY_IF) {
    	return LexIfOnlyIf
    } else if strings.HasPrefix(this.InputToEnd(), IF_ONLY_IF) {
    	return LexOperator
    }
}

func LexImplies(lexer *Lexer) LexFn {
	this.Pos += len(IMPLIES)
	this.Emit(lexertoker.TOKEN_IMPLIES)
	return LexResult
}

func LexIfOnlyIf(lexer *Lexer) LexFn {
	this.Pos += len(IF_ONLY_IF)
	this.Emit(lexertoker.TOKEN_IF_ONLY_IF)
	return LexResult
}

func LexOperator(lexer *Lexer) LexFn {
	this.Next()
	if strings.HasPrefix(OPERATORS, this.Input[this.Start]) {
		this.Emit(TOKEN_OPERATOR)
		return LexKey
	}
	return LexError
}

func LexResult(lexer *Lexer) LexFn {
	this.Next()
	if strings.ContainsRune(KEYS, this.Input[this.Start]) {
		this.Emit(TOKEN_KEY)
	}
	if strings.ContainsRune(OPERATORS, this.Input[this.Pos]) {
		this.Next()
		this.Emit(TOKEN_OPERATOR)
		this.Next()
		if strings.ContainsRune(KEYS, this.Input[this.Start]) {
			this.Emit(TOKEN_KEY)
		} else {
			return LexError
		}
	}
	this.Next()
	if this.Input[this.Start] == `/n` {
		this.Emit(TOKEN_EOL)
		return LexBegin
	}
	return LexError
}

func LexEquals(lexer *Lexer) LexFn {
  	this.Pos += len(EQUALS)
  	this.Emit(TOKEN_EQUALS)
  	return LexFact
}

func LexFact(lexer *Lexer) LexFn {
	if strings.ContainsRune(KEYS, this.Input[this.Pos]) {
		this.Pos += 1
		this.Emit(TOKEN_KEY)
		return LexFact
	} else {
		return LexEnd
	}
}
