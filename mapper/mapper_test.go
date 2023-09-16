package mapper_test

import (
	"testing"

	"github.com/EdsonHTJ/jtos/lexer"
	"github.com/EdsonHTJ/jtos/mapper"
	"github.com/stretchr/testify/require"
)

func TestMapper(t *testing.T) {
	jsonstr := `{"name": "Thomas", "age": 25, "height": 1.75, "weight": 70.5}`
	lexer := lexer.New()
	tokens, err := lexer.GetTokens(jsonstr)
	require.NoError(t, err)

	mapper := mapper.New(tokens)
	object, err := mapper.ParseObject()
	require.NoError(t, err)
	require.Equal(t, "Thomas", object["name"].Data)
	require.Equal(t, 25, object["age"].Data)
}
