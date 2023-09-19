package gogen

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
	"unicode"

	"github.com/EdsonHTJ/jtos/domain"
)

type GoField struct {
	Name        string
	Type        string
	JsonName    string
	IsPrimitive bool
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
	Structures map[string]GoStruct
}

type GoGenObject struct {
	Name string
	Obj  domain.Value
}

func New() *GoGen {
	return &GoGen{Structures: map[string]GoStruct{}}
}

func (g *GoGen) ParseObject(name string, object domain.Object) string {
	newName := toCamelCase(name)
	goStruct := GoStruct{Name: newName, Fields: []GoField{}}

	objects := make([]GoGenObject, 0)
	for k, v := range object {
		objects = append(objects, GoGenObject{Name: k, Obj: v})
	}

	sort.Slice(objects, func(i, j int) bool {
		return objects[i].Name < objects[j].Name
	})

	for _, obj := range objects {
		if IsPrimitiveValue(obj.Obj) {
			goStruct.Fields = append(goStruct.Fields, GoField{IsPrimitive: true, Name: toCamelCase(obj.Name), Type: ParsePrimitiveValue(obj.Obj), JsonName: obj.Name})
		} else {
			goStruct.Fields = append(goStruct.Fields, GoField{IsPrimitive: false, Name: toCamelCase(obj.Name), Type: g.ParseNonPrimitiveValue(obj.Name, obj.Obj), JsonName: obj.Name})
		}
	}

	sort.Slice(goStruct.Fields, func(i, j int) bool {
		if goStruct.Fields[i].IsPrimitive == goStruct.Fields[j].IsPrimitive {
			return goStruct.Fields[i].Name < goStruct.Fields[j].Name
		}

		return goStruct.Fields[i].IsPrimitive
	})

	structure, ok := g.Structures[goStruct.CalcHash()]
	if !ok {
		g.Structures[goStruct.CalcHash()] = goStruct
	} else {
		newName = structure.Name
	}

	return newName
}

func (g *GoGen) ParseNonPrimitiveValue(key string, value domain.Value) string {
	switch value.Type {
	case domain.VALUE_OBJECT:
		return g.ParseObject(key, value.Data.(domain.Object))
	case domain.VALUE_ARRAY_OBJ:
		valueArray := value.Data.([]domain.Object)
		if len(valueArray) > 0 {
			return "[]" + g.ParseObject(key, valueArray[0])
		} else {
			return "interface{}"
		}
	default:
		return "interface{}"
	}
}

func (g *GoGen) Generate(packageName string) string {
	result := ""

	structs := make([]GoStruct, 0)
	for _, goStruct := range g.Structures {
		structs = append(structs, goStruct)
	}

	sort.Slice(structs, func(i, j int) bool {
		return len(structs[i].Fields) > (len(structs[j].Fields))
	})

	result += "package " + packageName + "\n\n"
	for _, goStruct := range structs {
		result += "type " + goStruct.Name + " struct {\n"
		for _, field := range goStruct.Fields {
			result += "\t" + field.Name + " " +
				field.Type + " `json:\"" + field.JsonName + ",omitempty\"`\n"
		}
		result += "}\n\n"
	}

	return result
}

func (g *GoGen) GetOutPath(packageName string) string {
	return packageName + "/" + packageName + ".go"
}

func IsPrimitiveValue(value domain.Value) bool {
	switch value.Type {
	case domain.VALUE_INTEGER, domain.VALUE_STRING, domain.VALUE_FLOAT, domain.VALUE_BOOL,
		domain.VALUE_ARRAY_BOOL, domain.VALUE_ARRAY_FLOAT, domain.VALUE_ARRAY_INT, domain.VALUE_ARRAY_STR, domain.VALUE_NULL:
		return true
	default:
		return false
	}
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
	case domain.VALUE_ARRAY_INT:
		return "[]int32"
	case domain.VALUE_ARRAY_STR:
		return "[]string"
	case domain.VALUE_ARRAY_FLOAT:
		return "[]float64"
	default:
		return "interface{}"
	}
}

func toCamelCase(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for i, word := range words {
		words[i] = strings.Title(word)
	}

	return strings.Join(words, "")
}
