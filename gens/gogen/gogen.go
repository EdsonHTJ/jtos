package gogen

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/EdsonHTJ/jtos/domain"
)

type GoField struct {
	Name string
	Type string
}

type GoStruct struct {
	Name   string
	Fields []GoField
}

func (g *GoStruct) CalcHash() string {
	input := ""
	for _, field := range g.Fields {
		input += "\\name: " + field.Name + "\\type: " + field.Type
	}

	sha256 := sha256.Sum256([]byte(input))
	return hex.EncodeToString(sha256[:])
}

type GoGen struct {
	Output     string
	Structures map[string]GoStruct
}

func New() *GoGen {
	return &GoGen{Output: "", Structures: map[string]GoStruct{}}
}

func (g *GoGen) ParseObject(name string, object domain.Object) {
	// //goStruct := GoStruct{Name: name, Fields: []GoField{}}
	// for k, v := range object {
	// 	goStruct.Fields = append(goStruct.Fields, GoField{Name: k, Type: g.parseValueType(v.Type)})
	// }

}

func ParsePrimitiveValue(value domain.Value) string {
	switch value.Type {
	case domain.VALUE_INTEGER:
		return "int32"
	case domain.VALUE_STRING:
		return "string"
	case domain.VALUE_FLOAT:
		return "float64"
	case domain.VALUE_BOOL:
		return "bool"
	case domain.VALUE_NULL:
		return "interface{}"
	default:
		return "interface{}"
	}
}
