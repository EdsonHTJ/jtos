package domain

// List of possible tokens of a json file
const (
	TOKEN_OBJECT_OPEN    TokenType = iota
	TOKEN_OBJECT_CLOSE   TokenType = iota
	TOKEN_ARRAY_OPEN     TokenType = iota
	TOKEN_ARRAY_CLOSE    TokenType = iota
	TOKEN_INTEGER        TokenType = iota
	TOKEN_FLOAT          TokenType = iota
	TOKEN_SPECIAL_STRING TokenType = iota
	TOKEN_SIMPLE_STRING  TokenType = iota
	TOKEN_COLON          TokenType = iota
	TOKEN_COMMA          TokenType = iota
	TOKEN_RESERVED_WORD  TokenType = iota
)

type TokenType uint8
type TokenList []Token

type Token struct {
	Type  TokenType
	Value string
}
