package domain

// List of possible values
const (
	VALUE_OBJECT      = iota
	VALUE_INTEGER     = iota
	VALUE_STRING      = iota
	VALUE_FLOAT       = iota
	VALUE_BOOL        = iota
	VALUE_NULL        = iota
	VALUE_ARRAY_INT   = iota
	VALUE_ARRAY_STR   = iota
	VALUE_ARRAY_FLOAT = iota
	VALUE_ARRAY_BOOL  = iota
)

type ValueType uint8

type Value struct {
	Type ValueType
	Data interface{}
}

type Object map[string]Value
