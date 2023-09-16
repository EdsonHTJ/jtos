package domain

const (
	TYPE_OBJECT_OPEN   = iota
	TYPE_OBJECT_CLOSE  = iota
	TYPE_ARRAY_OPEN    = iota
	TYPE_ARRAY_CLOSE   = iota
	TYPE_INTEGER       = iota
	TYPE_FLOAT         = iota
	TYPE_STRING        = iota
	TYPE_KEY_STRING    = iota
	TYPE_COLON         = iota
	TYPE_COMMA         = iota
	TYPE_RESERVED_WORD = iota
)

const (
	VALUE_OBJECT      = iota
	VALUE_INTEGER     = iota
	VALUE_STRING      = iota
	VALUE_FLOAT       = iota
	VALUE_ARRAY_INT   = iota
	VALUE_ARRAY_STR   = iota
	VALUE_ARRAY_FLOAT = iota
)

type TokenType uint8
type TokenList []Token

type Token struct {
	Type  TokenType
	Value string
}

type ValueType uint8

type Value struct {
	Type ValueType
	Data interface{}
}

type Object map[string]Value
