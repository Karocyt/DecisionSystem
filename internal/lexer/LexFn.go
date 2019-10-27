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
  lexer.SkipWhitespace()

  if strings.HasPrefix(lexer.InputToEnd(), LEFT_BRACKET) {
    return LexLeftBracket
  } else if strings.HasPrefix(lexer.InputToEnd(), EQUALS) {
    return LexEquals
  } else if strings.HasPrefix(lexer.InputToEnd(), QUERY) {
    return LexQuery
  } else {
  	return LexKey
  }
}

func LexQuery(lexer *Lexer) LexFn {
	lexer.Pos += len(QUERY)
	lexer.Emit(TOKEN_QUERY)
	return LexKeyQuery
}

func LexKeyQuery(lexer *Lexer) LexFn {
	lexer.Inc()
	if strings.ContainsRune(KEYS, lexer.Input[lexer.Start]) {
		lexer.Emit(TOKEN_KEY)
		return LexKeyQuery
	} else if lexer.Input[lexer.Start] == `/n` {
		return LexError
	}
	return LexEnd // To check, bad shit could be afterwards
}

/*
This lexer function emits a TOKEN_LEFT_BRACKET then returns
the lexer for a key.
*/
func LexLeftBracket(lexer *Lexer) LexFn {
  lexer.Pos += len(LEFT_BRACKET)
  lexer.BracketCount += 1
  lexer.Emit(TOKEN_LEFT_BRACKET)
  return LexKey
}

func LexRightBracket(lexer *Lexer) LexFn {
  lexer.Pos += len(RIGHT_BRACKET)
  lexer.BracketCount -= 1
  if lexer.BracketCount < 0 {
  	return LexError
  }
  lexer.Emit(TOKEN_RIGHT_BRACKET)
  return LexSymbol
}

func LexKey(lexer *Lexer) LexFn {
	if strings.ContainsRune(KEYS, lexer.Input[lexer.Pos]) {
		lexer.Inc()
		lexer.Emit(TOKEN_KEY)
		return LexSymbol
	} else if strings.HasPrefix(lexer.InputToEnd(), LEFT_BRACKET) {
    	return LexLeftBracket
    } else {
    	return LexError
    }
}

func LexSymbol(lexer *Lexer) LexFn {
	if strings.HasPrefix(lexer.InputToEnd(), IMPLIES) {
    	return LexImplies
    } else if strings.HasPrefix(lexer.InputToEnd(), IF_ONLY_IF) {
    	return LexIfOnlyIf
    } else if strings.HasPrefix(lexer.InputToEnd(), IF_ONLY_IF) {
    	return LexOperator
    }
}

func LexImplies(lexer *Lexer) LexFn {
	lexer.Pos += len(IMPLIES)
	lexer.Emit(lexertoker.TOKEN_IMPLIES)
	return LexResult
}

func LexIfOnlyIf(lexer *Lexer) LexFn {
	lexer.Pos += len(IF_ONLY_IF)
	lexer.Emit(lexertoker.TOKEN_IF_ONLY_IF)
	return LexResult
}

func LexOperator(lexer *Lexer) LexFn {
	lexer.Inc()
	if strings.HasPrefix(OPERATORS, lexer.Input[lexer.Start]) {
		lexer.Emit(TOKEN_OPERATOR)
		return LexKey
	}
	return LexError
}

func LexResult(lexer *Lexer) LexFn {
	lexer.Inc()
	if strings.ContainsRune(KEYS, lexer.Input[lexer.Start]) {
		lexer.Emit(TOKEN_KEY)
	}
	if strings.ContainsRune(OPERATORS, lexer.Input[lexer.Pos]) {
		lexer.Inc()
		lexer.Emit(TOKEN_OPERATOR)
		lexer.Inc()
		if strings.ContainsRune(KEYS, lexer.Input[lexer.Start]) {
			lexer.Emit(TOKEN_KEY)
		} else {
			return LexError
		}
	}
	lexer.Inc()
	if lexer.Input[lexer.Start] == `/n` {
		lexer.Emit(TOKEN_EOL)
		return LexBegin
	}
	return LexError
}

func LexEquals(lexer *Lexer) LexFn {
  	lexer.Pos += len(EQUALS)
  	lexer.Emit(TOKEN_EQUALS)
  	return LexFact
}

func LexFact(lexer *Lexer) LexFn {
	if strings.ContainsRune(KEYS, lexer.Input[lexer.Pos]) {
		lexer.Pos += 1
		lexer.Emit(TOKEN_KEY)
		return LexFact
	} else {
		return LexEnd
	}
}
