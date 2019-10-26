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

  if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LEFT_BRACKET) {
    return LexLeftBracket
  } else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.EQUALS) {
    return LexEquals
  } else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.QUERY) {
    return LexQuery
  } else {
  	return LexKey
  }
}

func LexQuery(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.QUERY)
	lexer.Emit(lexertoken.TOKEN_QUERY)
	return LexKeyQuery
}

func LexKeyQuery(lexer *Lexer) LexFn {
	lexer.Inc()
	if strings.ContainsRune(lexertoken.KEYS, lexer.Input[lexer.Start]) {
		lexer.Emit(lexertoken.TOKEN_KEY)
		return LexKeyQuery
	} else if lexer.Input[lexer.Start] == '/n' {
		return LexError
	}
	return LexEnd // To check, bad shit could be afterwards
}

/*
This lexer function emits a TOKEN_LEFT_BRACKET then returns
the lexer for a key.
*/
func LexLeftBracket(lexer *Lexer) LexFn {
  lexer.Pos += len(lexertoken.LEFT_BRACKET)
  lexer.BracketCount += 1
  lexer.Emit(lexertoken.TOKEN_LEFT_BRACKET)
  return LexKey
}

func LexRightBracket(lexer *Lexer) LexFn {
  lexer.Pos += len(lexertoken.RIGHT_BRACKET)
  lexer.BracketCount -= 1
  if lexer.BracketCount < 0 {
  	return LexError
  }
  lexer.Emit(lexertoken.TOKEN_RIGHT_BRACKET)
  return LexSymbol
}

func LexKey(lexer *Lexer) LexFn {
	if strings.ContainsRune(lexertoken.KEYS, lexer.Input[lexer.Pos]) {
		lexer.Inc()
		lexer.Emit(lexertoken.TOKEN_KEY)
		return LexSymbol
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LEFT_BRACKET) {
    	return LexLeftBracket
    } else {
    	return LexError
    }
}

func LexSymbol(lexer *Lexer) LexFn {
	if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IMPLIES) {
    	return LexImplies
    } else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IF_ONLY_IF) {
    	return LexIfOnlyIf
    } else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IF_ONLY_IF) {
    	return LexOperator
    }
}

func LexImplies(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.IMPLIES)
	lexer.Emit(lexertoker.TOKEN_IMPLIES)
	return LexResult
}

func LexIfOnlyIf(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.IF_ONLY_IF)
	lexer.Emit(lexertoker.TOKEN_IF_ONLY_IF)
	return LexResult
}

func LexOperator(lexer *Lexer) LexFn {
	lexer.Inc()
	if strings.HasPrefix(lexertoken.OPERATORS, lexer.Input[lexer.Start]) {
		lexer.Emit(lexertoken.TOKEN_OPERATOR)
		return LexKey
	}
	return LexError
}

func LexResult(lexer *Lexer) LexFn {
	lexer.Inc()
	if strings.ContainsRune(lexertoken.KEYS, lexer.Input[lexer.Start]) {
		lexer.Emit(lexertoken.TOKEN_KEY)
	}
	if strings.ContainsRune(lexertoken.OPERATORS, lexer.Input[lexer.Pos]) {
		lexer.Inc()
		lexer.Emit(lexertoken.TOKEN_OPERATOR)
		lexer.Inc()
		if strings.ContainsRune(lexertoken.KEYS, lexer.Input[lexer.Start]) {
			lexer.Emit(lexertoken.TOKEN_KEY)
		} else {
			return LexError
		}
	}
	lexer.Inc()
	if lexer.Input[lexer.Start] == '/n'
		lexer.Emit(lexertoken.TOKEN_EOL)
		return LexBegin
	}
	return LexError
}

func LexEquals(lexer *Lexer) LexFn {
  	lexer.Pos += len(lexertoken.EQUALS)
  	lexer.Emit(lexertoken.TOKEN_EQUALS)
  	return LexFact
}

func LexFact(lexer *Lexer) LexFn {
	if strings.ContainsRune(lexertoken.KEYS, lexer.Input[lexer.Pos]) {
		lexer.Pos += 1
		lexer.Emit(lexertoken.TOKEN_KEY)
		return LexFact
	} else {
		return LexEnd
	}
}
