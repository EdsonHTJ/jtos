package mapper_test

import (
	"testing"

	"github.com/EdsonHTJ/jtos/domain"
	"github.com/EdsonHTJ/jtos/lexer"
	"github.com/EdsonHTJ/jtos/mapper"
	"github.com/stretchr/testify/require"
)

func TestMapperComplexData(t *testing.T) {
	jsonstr :=
		`{"name": "Thomas", "age": 25, "height": 1.75, "weight": 70.5,
	 "isMarried": true, "children": ["John", "Mary"], "car": {"model": "Mustang",
	 "year": 1964}}`

	lexer := lexer.New()
	tokens, err := lexer.GetTokens(jsonstr)
	require.NoError(t, err)

	mapper := mapper.New(tokens)
	object, err := mapper.ParseObject()

	require.NoError(t, err)
	require.Equal(t, "Thomas", object["name"].Data)
	require.Equal(t, 25, object["age"].Data)
	require.Equal(t, 1.75, object["height"].Data)
	require.Equal(t, 70.5, object["weight"].Data)
	require.Equal(t, true, object["isMarried"].Data)
	require.Equal(t, "John", object["children"].Data.([]string)[0])
	require.Equal(t, "Mary", object["children"].Data.([]string)[1])
	require.Equal(t, "Mustang", object["car"].Data.(domain.Object)["model"].Data)
	require.Equal(t, 1964, object["car"].Data.(domain.Object)["year"].Data)

}

func TestMapperArrays(t *testing.T) {
	jsonStr := `{"intArr": [1,2,3], "floatArr": [1.1, 2.2, 3.3], "stringArr": ["a", "b", "c"], "boolArr": [true, false, true]}`

	lexer := lexer.New()
	tokens, err := lexer.GetTokens(jsonStr)
	require.NoError(t, err)

	mapper := mapper.New(tokens)
	object, err := mapper.ParseObject()
	require.NoError(t, err)

	require.Equal(t, 1, object["intArr"].Data.([]int)[0])
	require.Equal(t, 2, object["intArr"].Data.([]int)[1])
	require.Equal(t, 3, object["intArr"].Data.([]int)[2])

	require.Equal(t, 1.1, object["floatArr"].Data.([]float64)[0])
	require.Equal(t, 2.2, object["floatArr"].Data.([]float64)[1])
	require.Equal(t, 3.3, object["floatArr"].Data.([]float64)[2])

	require.Equal(t, "a", object["stringArr"].Data.([]string)[0])
	require.Equal(t, "b", object["stringArr"].Data.([]string)[1])
	require.Equal(t, "c", object["stringArr"].Data.([]string)[2])

	require.Equal(t, true, object["boolArr"].Data.([]bool)[0])
	require.Equal(t, false, object["boolArr"].Data.([]bool)[1])
	require.Equal(t, true, object["boolArr"].Data.([]bool)[2])
}

func TestMapperArrayOfObject(t *testing.T) {
	jsonstr := `{"data":[{"first": "Thomas", "age": -25}, {"name": "Mary", "age": 20}]}`

	lexer := lexer.New()
	tokens, err := lexer.GetTokens(jsonstr)
	require.NoError(t, err)

	mapper := mapper.New(tokens)
	object, err := mapper.ParseObject()
	require.NoError(t, err)

	require.Equal(t, "Thomas", object["data"].Data.([]domain.Object)[0]["first"].Data)
	require.Equal(t, -25, object["data"].Data.([]domain.Object)[0]["age"].Data)
	require.Equal(t, "Mary", object["data"].Data.([]domain.Object)[1]["name"].Data)
	require.Equal(t, 20, object["data"].Data.([]domain.Object)[1]["age"].Data)
}
