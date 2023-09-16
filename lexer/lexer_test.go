package lexer_test

import (
	"testing"

	"github.com/EdsonHTJ/jtos/domain"
	"github.com/EdsonHTJ/jtos/lexer"
	"github.com/stretchr/testify/require"
)

func TestLexer(t *testing.T) {
	jsonstr := `{"name": "Thomas", "age": 25, "height": 1.75, "weight": 70.5,}`

	l := lexer.New()
	tokens, err := l.GetTokens(jsonstr)
	require.NoError(t, err)

	require.Equal(t, "{", tokens[0].Value)
	require.Equal(t, domain.TOKEN_OBJECT_OPEN, tokens[0].Type)

	require.Equal(t, "\"name\"", tokens[1].Value)
	require.Equal(t, domain.TOKEN_SIMPLE_STRING, tokens[1].Type)

	require.Equal(t, ":", tokens[2].Value)
	require.Equal(t, domain.TOKEN_COLON, tokens[2].Type)

	require.Equal(t, "\"Thomas\"", tokens[3].Value)
	require.Equal(t, domain.TOKEN_SIMPLE_STRING, tokens[3].Type)

	require.Equal(t, ",", tokens[4].Value)
	require.Equal(t, domain.TOKEN_COMMA, tokens[4].Type)

	require.Equal(t, "\"age\"", tokens[5].Value)
	require.Equal(t, domain.TOKEN_SIMPLE_STRING, tokens[5].Type)

	require.Equal(t, ":", tokens[6].Value)
	require.Equal(t, domain.TOKEN_COLON, tokens[6].Type)

	require.Equal(t, "25", tokens[7].Value)
	require.Equal(t, domain.TOKEN_INTEGER, tokens[7].Type)

	require.Equal(t, ",", tokens[8].Value)
	require.Equal(t, domain.TOKEN_COMMA, tokens[8].Type)

	require.Equal(t, "\"height\"", tokens[9].Value)
	require.Equal(t, domain.TOKEN_SIMPLE_STRING, tokens[9].Type)

	require.Equal(t, ":", tokens[10].Value)
	require.Equal(t, domain.TOKEN_COLON, tokens[10].Type)

	require.Equal(t, "1.75", tokens[11].Value)
	require.Equal(t, domain.TOKEN_FLOAT, tokens[11].Type)

	require.Equal(t, ",", tokens[12].Value)
	require.Equal(t, domain.TOKEN_COMMA, tokens[12].Type)

	require.Equal(t, "\"weight\"", tokens[13].Value)
	require.Equal(t, domain.TOKEN_SIMPLE_STRING, tokens[13].Type)

	require.Equal(t, ":", tokens[14].Value)
	require.Equal(t, domain.TOKEN_COLON, tokens[14].Type)

	require.Equal(t, "70.5", tokens[15].Value)
	require.Equal(t, domain.TOKEN_FLOAT, tokens[15].Type)

	require.Equal(t, ",", tokens[16].Value)
	require.Equal(t, domain.TOKEN_COMMA, tokens[16].Type)

	require.Equal(t, "}", tokens[17].Value)
	require.Equal(t, domain.TOKEN_OBJECT_CLOSE, tokens[17].Type)
}
