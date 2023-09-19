package mapper

import (
	"fmt"
	"strconv"

	"github.com/EdsonHTJ/jtos/domain"
)

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
