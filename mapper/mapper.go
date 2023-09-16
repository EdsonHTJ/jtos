package mapper

import (
	"fmt"

	"github.com/EdsonHTJ/jtos/domain"
)

type Mapper struct {
	TokenList  domain.TokenList
	TokenIndex uint32
}

func New(tokenList domain.TokenList) Mapper {
	return Mapper{
		TokenList:  tokenList,
		TokenIndex: 0,
	}
}

// List of reserved values
const (
	BOOLEAN_TRUE  = "true"
	BOOLEAN_FALSE = "false"
	NULL_VALUE    = "null"
)

// GetNextToken returns the next token in the list
func (m *Mapper) GetNextToken() (domain.Token, error) {
	if m.TokenIndex < uint32(len(m.TokenList)) {
		token := m.TokenList[m.TokenIndex]
		m.TokenIndex++
		return token, nil
	}

	return domain.Token{}, fmt.Errorf("no more tokens")
}

// ParseObject parses a json object
func (m *Mapper) ParseObject() (domain.Object, error) {

	object := domain.Object{}

	err := m.expectObjectOpen()
	if err != nil {
		return domain.Object{}, err
	}

	for {
		keyName, err := m.expectKeyedString()
		if err != nil {
			return domain.Object{}, err
		}

		err = m.expectColon()
		if err != nil {
			return domain.Object{}, err
		}

		token, err := m.GetNextToken()
		if err != nil {
			return domain.Object{}, err
		}

		innerValue, err := m.parseInnerValue(token)
		if err != nil {
			return domain.Object{}, err
		}

		if innerValue.Type != domain.VALUE_NULL {
			object[keyName] = innerValue
		}

		token, err = m.GetNextToken()
		if err != nil {
			return domain.Object{}, err
		}

		switch token.Type {
		case domain.TOKEN_COMMA:
			continue
		case domain.TOKEN_OBJECT_CLOSE:
			return object, nil
		default:
			return domain.Object{}, fmt.Errorf("expected , or }, got %s", token.Value)
		}
	}
}

// Parses an inner valuer of a json key-value pair
func (m *Mapper) parseInnerValue(token domain.Token) (domain.Value, error) {
	switch token.Type {
	case domain.TOKEN_INTEGER:
		integer, err := parseInteger(token)
		if err != nil {
			return domain.Value{}, err
		}

		return integer, nil

	case domain.TOKEN_STRING | domain.TOKEN_KEY_STRING:
		return parseString(token), nil

	case domain.TOKEN_RESERVED_WORD:
		return parseReservedWord(token)

	case domain.TOKEN_FLOAT:
		float, err := parseFloat(token)
		if err != nil {
			return domain.Value{}, err
		}

		return float, nil

	case domain.TOKEN_ARRAY_OPEN:
		return m.parseArray()

	case domain.TOKEN_OBJECT_OPEN:
		// The parse object expects and initial '{' token
		// so we need to go back one token
		m.TokenIndex--
		innerObj, err := m.ParseObject()
		if err != nil {
			return domain.Value{}, err
		}

		return domain.Value{
			Type: domain.VALUE_OBJECT,
			Data: innerObj,
		}, nil

	default:
		return domain.Value{}, fmt.Errorf("expected integer, string or float, got %s", token.Value)
	}
}
