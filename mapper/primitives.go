package mapper

import (
	"fmt"
	"strconv"

	"github.com/EdsonHTJ/jtos/domain"
)

//This file defines the functions that parse the json primitives
//and return the corresponding domain.Value
//The expects functions checks if the next token is the expected one
//and returns an error if it is not

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

	if token.Type != domain.TOKEN_SIMPLE_STRING {
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

	if (token.Type != domain.TOKEN_SPECIAL_STRING) && (token.Type != domain.TOKEN_SIMPLE_STRING) {
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

	value, err := parseReservedWord(token)
	if err != nil {
		return false, err
	}

	if value.Type != domain.VALUE_BOOL {
		return false, fmt.Errorf("expected true or false, got %s", token.Value)
	}

	return value.Data.(bool), nil
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
