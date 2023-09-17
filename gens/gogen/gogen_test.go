package gogen_test

import (
	"os"
	"testing"

	"github.com/EdsonHTJ/jtos/gens/gogen"
	"github.com/EdsonHTJ/jtos/lexer"
	"github.com/EdsonHTJ/jtos/mapper"
	"github.com/stretchr/testify/require"
)

func TestGogenParse(t *testing.T) {
	jsonstr :=
		`{"name": "Thom$$$--as", "age": 25, "height": 1.75, "weight": 70.5,
	 "isMarried": true, "children": ["John", "Mary"], "car": {"model": "Mustang",
	 "year": 1964}, "secondcar": {"model": "Mustang",
	 "year": 1964}, "thirdcar": {"model": "Mustang",
	 "year": 1964}, "fourthcar": {"model": "Mustang",
	 "year": 1964}}`

	lexer := lexer.New()
	tokens, err := lexer.GetTokens(jsonstr)
	require.NoError(t, err)

	mapper := mapper.New(tokens)
	object, err := mapper.ParseObject()
	require.NoError(t, err)

	gogen := gogen.New()
	gogen.ParseObject("Person", object)

	output := gogen.Generate("test")

	os.WriteFile("test/test.go", []byte(output), 0644)
}
