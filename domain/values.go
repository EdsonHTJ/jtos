package domain

// List of possible values
const (
	VALUE_OBJECT      ValueType = iota
	VALUE_INTEGER     ValueType = iota
	VALUE_STRING      ValueType = iota
	VALUE_FLOAT       ValueType = iota
	VALUE_BOOL        ValueType = iota
	VALUE_NULL        ValueType = iota
	VALUE_ARRAY_INT   ValueType = iota
	VALUE_ARRAY_STR   ValueType = iota
	VALUE_ARRAY_FLOAT ValueType = iota
	VALUE_ARRAY_BOOL  ValueType = iota
	VALUE_ARRAY_OBJ   ValueType = iota
)

type ValueType uint8

type Value struct {
	Type ValueType
	Data interface{}
}

type Object map[string]Value
