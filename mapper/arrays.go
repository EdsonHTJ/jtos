package mapper

import (
	"fmt"

	"github.com/EdsonHTJ/jtos/domain"
)

//This file defines the functions that parse the json arrays
//and return the corresponding domain.Value
//The expects functions checks if the next token is the expected one
//and returns an error if it is not

func (m *Mapper) parseArray() (domain.Value, error) {
	token, err := m.GetNextToken()
	if err != nil {
		return domain.Value{}, err
	}

	switch token.Type {
	case domain.TOKEN_INTEGER:
		// The parse array expects and initial '[' token
		// so we need to go back one token
		m.TokenIndex--
		return m.parseIntArray()

	case (domain.TOKEN_STRING | domain.TOKEN_KEY_STRING):
		m.TokenIndex--
		return m.parseStringArray()

	case domain.TOKEN_FLOAT:
		m.TokenIndex--
		return m.parseFloatArray()

	default:
		return domain.Value{}, fmt.Errorf("expected integer, string or float, got %s", token.Value)
	}
}

func (m *Mapper) parseIntArray() (domain.Value, error) {
	arr := make([]int, 0)
	for {
		intval, err := m.expectInteger()
		if err != nil {
			return domain.Value{}, err
		}

		arr = append(arr, intval)

		token, err := m.GetNextToken()
		if err != nil {
			return domain.Value{}, err
		}

		switch token.Type {
		case domain.TOKEN_COMMA:
			continue

		case domain.TOKEN_ARRAY_CLOSE:
			return domain.Value{
				Type: domain.VALUE_ARRAY_INT,
				Data: arr,
			}, nil
		default:
			return domain.Value{}, fmt.Errorf("expected , or ], got %s", token.Value)
		}
	}
}

func (m *Mapper) parseFloatArray() (domain.Value, error) {
	arr := make([]float64, 0)
	for {
		val, err := m.expectFloat()
		if err != nil {
			return domain.Value{}, err
		}

		arr = append(arr, val)

		token, err := m.GetNextToken()
		if err != nil {
			return domain.Value{}, err
		}

		switch token.Type {
		case domain.TOKEN_COMMA:
			continue

		case domain.TOKEN_ARRAY_CLOSE:
			return domain.Value{
				Type: domain.VALUE_ARRAY_FLOAT,
				Data: arr,
			}, nil
		default:
			return domain.Value{}, fmt.Errorf("expected , or ], got %s", token.Value)
		}
	}
}

func (m *Mapper) parseStringArray() (domain.Value, error) {
	arr := make([]string, 0)
	for {
		val, err := m.expectString()
		if err != nil {
			return domain.Value{}, err
		}

		arr = append(arr, val)

		token, err := m.GetNextToken()
		if err != nil {
			return domain.Value{}, err
		}

		switch token.Type {
		case domain.TOKEN_COMMA:
			continue

		case domain.TOKEN_ARRAY_CLOSE:
			return domain.Value{
				Type: domain.VALUE_ARRAY_STR,
				Data: arr,
			}, nil
		default:
			return domain.Value{}, fmt.Errorf("expected , or ], got %s", token.Value)
		}
	}
}
