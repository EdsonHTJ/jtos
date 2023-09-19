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

	case domain.TOKEN_SPECIAL_STRING, domain.TOKEN_SIMPLE_STRING:
		m.TokenIndex--
		return m.parseStringArray()

	case domain.TOKEN_FLOAT:
		m.TokenIndex--
		return m.parseFloatArray()

	case domain.TOKEN_OBJECT_OPEN:
		m.TokenIndex--
		return m.parseObjectArray()

	case domain.TOKEN_RESERVED_WORD:
		m.TokenIndex--
		if token.Value == "true" || token.Value == "false" {
			return m.parseBooleanArray()
		} else {
			return domain.Value{}, fmt.Errorf("expected boolean, got %s", token.Value)
		}

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

func (m *Mapper) parseBooleanArray() (domain.Value, error) {
	arr := make([]bool, 0)
	for {
		val, err := m.expectBool()
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

func (m *Mapper) parseObjectArray() (domain.Value, error) {
	arr := make([]domain.Object, 0)

	objFields := domain.Object{}
	for {
		val, err := m.ParseObject()
		if err != nil {
			return domain.Value{}, err
		}

		for k, v := range val {
			objFields[k] = domain.Value{
				Type: v.Type,
				Data: nil,
			}
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
			ajustedArr := make([]domain.Object, 0)
			for _, itr := range arr {
				ajustedObj := domain.Object{}

				for k, v := range objFields {
					ajustedObj[k] = v
				}

				for k, v := range itr {
					ajustedObj[k] = v
				}
				ajustedArr = append(ajustedArr, ajustedObj)
			}

			return domain.Value{
				Type: domain.VALUE_ARRAY_OBJ,
				Data: ajustedArr,
			}, nil
		default:
			return domain.Value{}, fmt.Errorf("expected , or ], got %s", token.Value)
		}
	}
}
