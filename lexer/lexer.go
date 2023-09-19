package lexer

import (
	"fmt"
	"strings"

	"github.com/EdsonHTJ/jtos/domain"
)

// Alphabet of the lexer
const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	digits   = "0123456789"
)

// Reserved Tokens
const (
	OPEN_BRACE  = '{'
	CLOSE_BRACE = '}'
	OPEN_BRACK  = '['
	CLOSE_BRACK = ']'
	COLON       = ':'
	COMMA       = ','
	PERIOD      = '.'
	QUOTE       = '"'
	MINUS       = '-'
)

// State of the lexer
const (
	STATE_PARSING_INITIAL       = iota
	STATE_PARSING_STRING        = iota
	STATE_PARSING_INTEGER       = iota
	STATE_PARSING_FLOAT         = iota
	STATE_PARSING_RESERVED_WORD = iota
)

type State uint8

type Lexer struct {
	State        State
	CurrentToken domain.Token
	Tokens       domain.TokenList
	Line         uint16
}

// New creates a new lexer
func New() Lexer {
	return Lexer{
		State: STATE_PARSING_INITIAL,
	}
}

// GetTokens parse a json string to find and label the tokens
func (l *Lexer) GetTokens(str string) (domain.TokenList, error) {
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
		return l.parseFull(c)
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

func (l *Lexer) appendToken(token domain.Token) {
	l.Tokens = append(l.Tokens, token)
	l.CurrentToken = domain.Token{}
	l.State = STATE_PARSING_INITIAL
}

// parseFull the higher level state machine,
// it parses the full json file and redirects the flow to the specifics states
func (l *Lexer) parseFull(c rune) error {
	switch {
	case isWhitespace(c):
	case isNewLine(c):
		l.Line++
	case (isDigit(c) || c == MINUS):
		l.State = STATE_PARSING_INTEGER
		l.CurrentToken.Type = domain.TOKEN_INTEGER
		l.CurrentToken.Value = string(c)
	case isLetter(c):
		l.CurrentToken.Type = domain.TOKEN_RESERVED_WORD
		l.CurrentToken.Value = string(c)
		l.State = STATE_PARSING_RESERVED_WORD
	case c == QUOTE:
		l.State = STATE_PARSING_STRING
		l.CurrentToken.Type = domain.TOKEN_SIMPLE_STRING
		l.CurrentToken.Value = string(c)
	case c == OPEN_BRACE:
		l.appendToken(domain.Token{Type: domain.TOKEN_OBJECT_OPEN, Value: string(c)})
	case c == CLOSE_BRACE:
		l.appendToken(domain.Token{Type: domain.TOKEN_OBJECT_CLOSE, Value: string(c)})
	case c == OPEN_BRACK:
		l.appendToken(domain.Token{Type: domain.TOKEN_ARRAY_OPEN, Value: string(c)})
	case c == CLOSE_BRACK:
		l.appendToken(domain.Token{Type: domain.TOKEN_ARRAY_CLOSE, Value: string(c)})
	case c == COLON:
		l.appendToken(domain.Token{Type: domain.TOKEN_COLON, Value: string(c)})
	case c == COMMA:
		l.appendToken(domain.Token{Type: domain.TOKEN_COMMA, Value: string(c)})
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
		l.CurrentToken.Type = domain.TOKEN_SPECIAL_STRING
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
		l.CurrentToken.Type = domain.TOKEN_FLOAT
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
	return c == ' ' || c == '\t' || c == '\r'
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
