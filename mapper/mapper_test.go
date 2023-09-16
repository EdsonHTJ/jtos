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
