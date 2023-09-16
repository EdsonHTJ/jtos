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

	token, err := m.GetNextToken()
	if err != nil {
		return domain.Object{}, err
	}

	if token.Type != domain.TYPE_OBJECT_OPEN {
		return domain.Object{}, fmt.Errorf("expected {, got %s", token.Value)
	}

	for {
		token, err = m.GetNextToken()
		if err != nil {
			return domain.Object{}, err
		}

		if token.Type != domain.TYPE_KEY_STRING {
			return domain.Object{}, fmt.Errorf("expected key, got %s", token.Value)
		}

		keyName := token.Value

		token, err = m.GetNextToken()
		if err != nil {
			return domain.Object{}, err
		}

		if token.Type != domain.TYPE_COLON {
			return domain.Object{}, fmt.Errorf("expected :, got %s", token.Value)
		}

		token, err = m.GetNextToken()
		if err != nil {
			return domain.Object{}, err
		}

		switch token.Type {
		case domain.TYPE_INTEGER:
			integer, err := parseInteger(token)
			if err != nil {
				return domain.Object{}, err
			}

			object[keyName] = integer

		case domain.TYPE_STRING | domain.TYPE_KEY_STRING:
			object[keyName] = domain.Value{
				Type: domain.VALUE_STRING,
				Data: token.Value,
			}
		case domain.TYPE_FLOAT:
			float, err := parseFloat(token)
			if err != nil {
				return domain.Object{}, err
			}

			object[keyName] = float
		default:
			return domain.Object{}, fmt.Errorf("expected integer, string or float, got %s", token.Value)
		}

		token, err = m.GetNextToken()
		if err != nil {
			return domain.Object{}, err
		}

		switch token.Type {
		case domain.TYPE_COMMA:
			continue
		case domain.TYPE_OBJECT_CLOSE:
			return object, nil
		default:
			return domain.Object{}, fmt.Errorf("expected , or }, got %s", token.Value)
		}
	}
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
