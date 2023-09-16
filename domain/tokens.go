package domain

const (
	TOKEN_OBJECT_OPEN   = iota
	TOKEN_OBJECT_CLOSE  = iota
	TOKEN_ARRAY_OPEN    = iota
	TOKEN_ARRAY_CLOSE   = iota
	TOKEN_INTEGER       = iota
	TOKEN_FLOAT         = iota
	TOKEN_STRING        = iota
	TOKEN_KEY_STRING    = iota
	TOKEN_COLON         = iota
	TOKEN_COMMA         = iota
	TOKEN_RESERVED_WORD = iota
)

type TokenType uint8
type TokenList []Token

type Token struct {
	Type  TokenType
	Value string
}
