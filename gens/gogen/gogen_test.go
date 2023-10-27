package gogen_test

import (
	"os"
	"strings"
	"testing"

	"github.com/EdsonHTJ/jtos/gens/gogen"
	"github.com/EdsonHTJ/jtos/lexer"
	"github.com/EdsonHTJ/jtos/mapper"
	"github.com/stretchr/testify/require"
)

//TODO: Improve generation tests

func TestGogenParse(t *testing.T) {
	jsonstr :=
		`{"name": "Thomas", "age": 25, "height": 1.75, "weight": 70.5,
	 "isMarried": true, "children": ["John", "Mary"], "car": {"model": "Mustang",
	 "year": 1964}, "secondcar": {"model": "Mustang",
	 "year": 1964}, "thirdcar": {"model": "Mustang",
	 "year": 1964}, "fourthcar": {"model": "Mustang",
	 "year": 1964}}`

	tokens, err := lexer.GetTokens(jsonstr)
	require.NoError(t, err)

	object, err := mapper.MapTokensToObject(tokens)
	require.NoError(t, err)

	gogen := gogen.New()
	gogen.ParseObject("Person", object)

	output := gogen.Generate("test")

	//Test if the package is correct
	require.Equal(t, strings.Index(output, "package test"), 0)

	//Test if the PersonStructureIsCorrect

	require.True(t, strings.Contains(output, "type Person struct {"))
	require.True(t, strings.Contains(output, "Age int32 `json:\"age\"`"))
	require.True(t, strings.Contains(output, "Children []string `json:\"children\"`"))
	require.True(t, strings.Contains(output, "Height float64 `json:\"height\"`"))
	require.True(t, strings.Contains(output, "IsMarried bool `json:\"isMarried\"`"))
	require.True(t, strings.Contains(output, "Name string `json:\"name\"`"))
	require.True(t, strings.Contains(output, "Weight float64 `json:\"weight\"`"))
	require.True(t, strings.Contains(output, "Car Car `json:\"car\"`"))
	require.True(t, strings.Contains(output, "Fourthcar Car `json:\"fourthcar\"`"))
	require.True(t, strings.Contains(output, "Secondcar Car `json:\"secondcar\"`"))
	require.True(t, strings.Contains(output, "Thirdcar Car `json:\"thirdcar\"`"))

	//Test if the CarStructureIsCorrect
	require.True(t, strings.Contains(output, "type Car struct {"))
	require.True(t, strings.Contains(output, "Model string `json:\"model\"`"))
	require.True(t, strings.Contains(output, "Year int32 `json:\"year\"`"))
}

func TestGogenWithFile(t *testing.T) {
	bt, err := os.ReadFile("../../data.json")
	require.NoError(t, err)

	jsonstr := string(bt)

	tokens, err := lexer.GetTokens(jsonstr)
	require.NoError(t, err)

	object, err := mapper.MapTokensToObject(tokens)
	require.NoError(t, err)

	gogen := gogen.New()
	gogen.ParseObject("Data", object)

	output := gogen.Generate("test")
	os.WriteFile("../../test.go", []byte(output), 0644)
}
