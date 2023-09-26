package gogen

import (
	"strings"
	"unicode"

	"github.com/EdsonHTJ/jtos/domain"
)

// IsPrimitiveValue checks if the value is a primitive value
func IsPrimitiveValue(value domain.Value) bool {
	switch value.Type {
	case domain.VALUE_INTEGER, domain.VALUE_STRING, domain.VALUE_FLOAT, domain.VALUE_BOOL,
		domain.VALUE_ARRAY_BOOL, domain.VALUE_ARRAY_FLOAT, domain.VALUE_ARRAY_INT, domain.VALUE_ARRAY_STR, domain.VALUE_NULL:
		return true
	default:
		return false
	}
}

func toCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for i, word := range words {
		words[i] = strings.Title(word)
	}

	return strings.Join(words, "")
}
