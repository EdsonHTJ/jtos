package mapper

import (
	"fmt"
	"strconv"

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

const (
	BOOLEAN_TRUE  = "true"
	BOOLEAN_FALSE = "false"
	NULL_VALUE    = "null"
)

func (m *Mapper) GetNextToken() (domain.Token, error) {
	if m.TokenIndex < uint32(len(m.TokenList)) {
		token := m.TokenList[m.TokenIndex]
		m.TokenIndex++
		return token, nil
	}

	return domain.Token{}, fmt.Errorf("no more tokens")
}

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
	default:
		return domain.Value{}, fmt.Errorf("expected integer, string or float, got %s", token.Value)
	}
}

func (m *Mapper) parseArray() (domain.Value, error) {
	token, err := m.GetNextToken()
	if err != nil {
		return domain.Value{}, err
	}

	switch token.Type {
	case domain.TOKEN_INTEGER:
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

func (m *Mapper) expectBoolArray() (domain.Value, error) {
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
				Type: domain.VALUE_ARRAY_BOOL,
				Data: arr,
			}, nil
		default:
			return domain.Value{}, fmt.Errorf("expected , or ], got %s", token.Value)
		}
	}
}

func (m *Mapper) expectObjectOpen() error {
	token, err := m.GetNextToken()
	if err != nil {
		return err
	}

	if token.Type != domain.TOKEN_OBJECT_OPEN {
		return fmt.Errorf("expected {, got %s", token.Value)
	}

	return nil
}

func (m *Mapper) expectKeyedString() (string, error) {
	token, err := m.GetNextToken()
	if err != nil {
		return "", err
	}

	if token.Type != domain.TOKEN_KEY_STRING {
		return "", fmt.Errorf("expected key, got %s", token.Value)
	}

	return parseRawString(token.Value), nil
}

func (m *Mapper) expectInteger() (int, error) {
	token, err := m.GetNextToken()
	if err != nil {
		return 0, err
	}

	if token.Type != domain.TOKEN_INTEGER {
		return 0, fmt.Errorf("expected integer, got %s", token.Value)
	}

	return strconv.Atoi(token.Value)
}

func (m *Mapper) expectFloat() (float64, error) {
	token, err := m.GetNextToken()
	if err != nil {
		return 0, err
	}

	if token.Type != domain.TOKEN_FLOAT {
		return 0, fmt.Errorf("expected float, got %s", token.Value)
	}

	return strconv.ParseFloat(token.Value, 64)
}

func (m *Mapper) expectString() (string, error) {
	token, err := m.GetNextToken()
	if err != nil {
		return "", err
	}

	if (token.Type != domain.TOKEN_STRING) && (token.Type != domain.TOKEN_KEY_STRING) {
		return "", fmt.Errorf("expected string, got %s", token.Value)
	}

	return parseRawString(token.Value), nil
}

func (m *Mapper) expectBool() (bool, error) {
	token, err := m.GetNextToken()
	if err != nil {
		return false, err
	}

	if token.Type != domain.TOKEN_RESERVED_WORD {
		return false, fmt.Errorf("expected true or false, got %s", token.Value)
	}

	switch token.Value {
	case BOOLEAN_TRUE:
		return true, nil
	case BOOLEAN_FALSE:
		return false, nil
	default:
		return false, fmt.Errorf("expected true or false, got %s", token.Value)
	}
}

func (m *Mapper) expectColon() error {
	token, err := m.GetNextToken()
	if err != nil {
		return err
	}

	if token.Type != domain.TOKEN_COLON {
		return fmt.Errorf("expected :, got %s", token.Value)
	}

	return nil
}

func parseInteger(t domain.Token) (domain.Value, error) {
	value, err := strconv.Atoi(t.Value)
	if err != nil {
		return domain.Value{}, err
	}

	return domain.Value{
		Type: domain.VALUE_INTEGER,
		Data: value,
	}, nil
}

func parseFloat(t domain.Token) (domain.Value, error) {
	value, err := strconv.ParseFloat(t.Value, 64)
	if err != nil {
		return domain.Value{}, err
	}

	return domain.Value{
		Type: domain.VALUE_FLOAT,
		Data: value,
	}, nil
}

func parseString(t domain.Token) domain.Value {
	return domain.Value{
		Type: domain.VALUE_STRING,
		Data: parseRawString(t.Value),
	}
}

func parseRawString(rawString string) string {
	return rawString[1 : len(rawString)-1]
}

func parseReservedWord(token domain.Token) (domain.Value, error) {
	switch token.Value {
	case BOOLEAN_TRUE:
		return domain.Value{
			Type: domain.VALUE_BOOL,
			Data: true,
		}, nil
	case BOOLEAN_FALSE:
		return domain.Value{
			Type: domain.VALUE_BOOL,
			Data: false,
		}, nil
	case NULL_VALUE:
		return domain.Value{
			Type: domain.VALUE_NULL,
			Data: nil,
		}, nil
	default:
		return domain.Value{}, fmt.Errorf("expected true, false or null, got %s", token.Value)
	}
}
