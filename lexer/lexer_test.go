package lexer_test

import (
	"testing"

	"github.com/EdsonHTJ/jtos/domain"
	"github.com/EdsonHTJ/jtos/lexer"
	"github.com/stretchr/testify/require"
)

func TestLexer(t *testing.T) {
	jsonstr := `{"name": "Thomas d'ante", "age": 25, "height": 1.75, "weight": 70.5, "money": -1.45,
	 "children": [{"name":"jhon"}, {"name": "mary"}], "married": true}`

	tokens, err := lexer.GetTokens(jsonstr)
	require.NoError(t, err)

	require.Equal(t, "{", tokens[0].Value)
	require.Equal(t, domain.TOKEN_OBJECT_OPEN, tokens[0].Type)

	require.Equal(t, "\"name\"", tokens[1].Value)
	require.Equal(t, domain.TOKEN_SIMPLE_STRING, tokens[1].Type)

	require.Equal(t, ":", tokens[2].Value)
	require.Equal(t, domain.TOKEN_COLON, tokens[2].Type)

	require.Equal(t, "\"Thomas d'ante\"", tokens[3].Value)
	require.Equal(t, domain.TOKEN_SPECIAL_STRING, tokens[3].Type)

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

	require.Equal(t, "\"money\"", tokens[17].Value)
	require.Equal(t, domain.TOKEN_SIMPLE_STRING, tokens[17].Type)

	require.Equal(t, ":", tokens[18].Value)
	require.Equal(t, domain.TOKEN_COLON, tokens[18].Type)

	require.Equal(t, "-1.45", tokens[19].Value)
	require.Equal(t, domain.TOKEN_FLOAT, tokens[19].Type)

	require.Equal(t, ",", tokens[20].Value)
	require.Equal(t, domain.TOKEN_COMMA, tokens[20].Type)

	require.Equal(t, "\"children\"", tokens[21].Value)
	require.Equal(t, domain.TOKEN_SIMPLE_STRING, tokens[21].Type)

	require.Equal(t, ":", tokens[22].Value)
	require.Equal(t, domain.TOKEN_COLON, tokens[22].Type)

	require.Equal(t, "[", tokens[23].Value)
	require.Equal(t, domain.TOKEN_ARRAY_OPEN, tokens[23].Type)

	require.Equal(t, "{", tokens[24].Value)
	require.Equal(t, domain.TOKEN_OBJECT_OPEN, tokens[24].Type)

	require.Equal(t, "\"name\"", tokens[25].Value)
	require.Equal(t, domain.TOKEN_SIMPLE_STRING, tokens[25].Type)

	require.Equal(t, ":", tokens[26].Value)
	require.Equal(t, domain.TOKEN_COLON, tokens[26].Type)

	require.Equal(t, "\"jhon\"", tokens[27].Value)
	require.Equal(t, domain.TOKEN_SIMPLE_STRING, tokens[27].Type)

	require.Equal(t, "}", tokens[28].Value)
	require.Equal(t, domain.TOKEN_OBJECT_CLOSE, tokens[28].Type)

	require.Equal(t, ",", tokens[29].Value)
	require.Equal(t, domain.TOKEN_COMMA, tokens[29].Type)

	require.Equal(t, "{", tokens[30].Value)
	require.Equal(t, domain.TOKEN_OBJECT_OPEN, tokens[30].Type)

	require.Equal(t, "\"name\"", tokens[31].Value)
	require.Equal(t, domain.TOKEN_SIMPLE_STRING, tokens[31].Type)

	require.Equal(t, ":", tokens[32].Value)
	require.Equal(t, domain.TOKEN_COLON, tokens[32].Type)

	require.Equal(t, "\"mary\"", tokens[33].Value)
	require.Equal(t, domain.TOKEN_SIMPLE_STRING, tokens[33].Type)

	require.Equal(t, "}", tokens[34].Value)
	require.Equal(t, domain.TOKEN_OBJECT_CLOSE, tokens[34].Type)

	require.Equal(t, "]", tokens[35].Value)
	require.Equal(t, domain.TOKEN_ARRAY_CLOSE, tokens[35].Type)

	require.Equal(t, ",", tokens[36].Value)
	require.Equal(t, domain.TOKEN_COMMA, tokens[36].Type)

	require.Equal(t, "\"married\"", tokens[37].Value)
	require.Equal(t, domain.TOKEN_SIMPLE_STRING, tokens[37].Type)

	require.Equal(t, ":", tokens[38].Value)
	require.Equal(t, domain.TOKEN_COLON, tokens[38].Type)

	require.Equal(t, "true", tokens[39].Value)
	require.Equal(t, domain.TOKEN_RESERVED_WORD, tokens[39].Type)

	require.Equal(t, "}", tokens[40].Value)
	require.Equal(t, domain.TOKEN_OBJECT_CLOSE, tokens[40].Type)

}
