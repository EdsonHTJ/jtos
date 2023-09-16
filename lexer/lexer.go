package lexer

import (
	"fmt"
	"strings"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits   = "0123456789"
)

const (
	OPEN_BRACE  = '{'
	CLOSE_BRACE = '}'
	OPEN_BRACK  = '['
	CLOSE_BRACK = ']'
	COLON       = ':'
	COMMA       = ','
	PERIOD      = '.'
	QUOTE       = '"'
)

const (
	TYPE_OBJECT_OPEN   = iota
	TYPE_OBJECT_CLOSE  = iota
	TYPE_ARRAY_OPEN    = iota
	TYPE_ARRAY_CLOSE   = iota
	TYPE_INTEGER       = iota
	TYPE_FLOAT         = iota
	TYPE_STRING        = iota
	TYPE_KEY_STRING    = iota
	TYPE_COLON         = iota
	TYPE_COMMA         = iota
	TYPE_RESERVED_WORD = iota
)

const (
	STATE_PARSING_INITIAL       = iota
	STATE_PARSING_STRING        = iota
	STATE_PARSING_INTEGER       = iota
	STATE_PARSING_FLOAT         = iota
	STATE_PARSING_RESERVED_WORD = iota
)

type State uint8
type TokenType uint8
type TokenList []Token

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	State        State
	CurrentToken Token
	Tokens       TokenList
	Line         uint16
}

func New() Lexer {
	return Lexer{
		State: STATE_PARSING_INITIAL,
	}
}

func (l *Lexer) GetTokens(str string) (TokenList, error) {
	for _, c := range str {
		err := l.parse(c)
		if err != nil {
			return nil, err
		}
	}

	return l.Tokens, nil
}

func (l *Lexer) parse(c rune) error {

	switch l.State {
	case STATE_PARSING_INITIAL:
		return l.parseObject(c)
	case STATE_PARSING_STRING:
		return l.parseString(c)
	case STATE_PARSING_INTEGER:
		return l.parseInteger(c)
	case STATE_PARSING_FLOAT:
		return l.parseFloat(c)
	case STATE_PARSING_RESERVED_WORD:
		return l.parseReservedWord(c)
	}

	return nil
}

func (l *Lexer) appendToken(token Token) {
	l.Tokens = append(l.Tokens, token)
	l.CurrentToken = Token{}
	l.State = STATE_PARSING_INITIAL
}

func (l *Lexer) parseObject(c rune) error {
	switch {
	case isWhitespace(c):
	case isNewLine(c):
		l.Line++
	case isDigit(c):
		l.State = STATE_PARSING_INTEGER
		l.CurrentToken.Type = TYPE_INTEGER
		l.CurrentToken.Value = string(c)
	case isLetter(c):
		l.CurrentToken.Type = TYPE_RESERVED_WORD
		l.CurrentToken.Value = string(c)
		l.State = STATE_PARSING_RESERVED_WORD
	case c == QUOTE:
		l.State = STATE_PARSING_STRING
		l.CurrentToken.Type = TYPE_KEY_STRING
		l.CurrentToken.Value = string(c)
	case c == OPEN_BRACE:
		l.appendToken(Token{Type: TYPE_OBJECT_OPEN, Value: string(c)})
	case c == CLOSE_BRACE:
		l.appendToken(Token{Type: TYPE_OBJECT_CLOSE, Value: string(c)})
	case c == OPEN_BRACK:
		l.appendToken(Token{Type: TYPE_ARRAY_OPEN, Value: string(c)})
	case c == CLOSE_BRACK:
		l.appendToken(Token{Type: TYPE_ARRAY_CLOSE, Value: string(c)})
	case c == COLON:
		l.appendToken(Token{Type: TYPE_COLON, Value: string(c)})
	case c == COMMA:
		l.appendToken(Token{Type: TYPE_COMMA, Value: string(c)})
	default:
		return fmt.Errorf("Unexpected character '%c' at line %d", c, l.Line)
	}

	return nil
}

func (l *Lexer) parseReservedWord(c rune) error {
	switch {
	case isLetter(c):
		l.CurrentToken.Value += string(c)
	default:
		l.appendToken(l.CurrentToken)
		l.State = STATE_PARSING_INITIAL
		l.parse(c)
	}

	return nil
}

func (l *Lexer) parseString(c rune) error {
	switch {
	case c == QUOTE:
		l.CurrentToken.Value += string(c)
		l.appendToken(l.CurrentToken)
	case isNonAlphanumeric(c):
		l.CurrentToken.Type = TYPE_STRING
		l.CurrentToken.Value += string(c)
	default:
		l.CurrentToken.Value += string(c)
	}

	return nil
}

func (l *Lexer) parseInteger(c rune) error {
	switch {
	case isDigit(c):
		l.CurrentToken.Value += string(c)
	case c == PERIOD:
		l.State = STATE_PARSING_FLOAT
		l.CurrentToken.Type = TYPE_FLOAT
		l.CurrentToken.Value += string(c)
	default:
		l.appendToken(l.CurrentToken)
		l.State = STATE_PARSING_INITIAL
		l.parse(c)
	}

	return nil
}

func (l *Lexer) parseFloat(c rune) error {
	switch {
	case isDigit(c):
		l.CurrentToken.Value += string(c)
	default:
		l.appendToken(l.CurrentToken)
		l.State = STATE_PARSING_INITIAL
		l.parse(c)
	}

	return nil
}

func isNewLine(c rune) bool {
	return c == '\n'
}

func isWhitespace(c rune) bool {
	return c == ' ' || c == '\t'
}

func isNonAlphanumeric(c rune) bool {
	return !isLetter(c) && !isDigit(c)
}

func isLetter(c rune) bool {
	return strings.Contains(alphabet, string(c))
}

func isDigit(c rune) bool {
	return strings.Contains(digits, string(c))
}
