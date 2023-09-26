package gogen

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"sort"

	"github.com/EdsonHTJ/jtos/domain"
)

const (
	GO_TYPE_INT       GoPrimitives = iota
	GO_TYPE_STRING    GoPrimitives = iota
	GO_TYPE_FLOAT     GoPrimitives = iota
	GO_TYPE_BOOL      GoPrimitives = iota
	GO_TYPE_INTERFACE GoPrimitives = iota
)

type GoPrimitives uint8

func (g GoPrimitives) getString() string {
	switch g {
	case GO_TYPE_INT:
		return "int32"
	case GO_TYPE_STRING:
		return "string"
	case GO_TYPE_FLOAT:
		return "float64"
	case GO_TYPE_BOOL:
		return "bool"
	case GO_TYPE_INTERFACE:
		return "interface{}"
	default:
		return "interface{}"
	}
}

// GoField represents a struct field in Go
type GoField struct {
	Name        string
	Type        GoType
	JsonName    string
	IsPrimitive bool
}

// GoType represents a type in Go
// It can be a primitive type or a complex type
type GoType struct {
	isArray       bool
	IsPrimitive   bool
	PrimitiveType GoPrimitives
	CustomType    string
}

// getString returns the string representation of the type
// For example for an it array it returns an "[]int32"
func (g *GoType) getString() string {
	out := ""
	if g.isArray {
		out += "[]"
	}

	if g.IsPrimitive {
		out += g.PrimitiveType.getString()
	} else {
		out += g.CustomType
	}

	return out
}

// GoStruct represents a struct in Go
type GoStruct struct {
	Name   string
	Fields []GoField
}

// CalcHash calculates the hash of the struct
// It is used to check if the struct already exists inside the structs map
func (g *GoStruct) CalcHash() string {
	input := ""
	for _, field := range g.Fields {
		input += "\\name: " + field.Name + "\\type: " + field.Type.getString()
	}

	sha256 := sha256.Sum256([]byte(input))
	return hex.EncodeToString(sha256[:])
}

// GoGenObject represents an object on the json
type GoGenObject struct {
	Name string
	Obj  domain.Value
}

type Output string

func (r *Output) insertHeader(packageName string) {
	*r = Output("package " + packageName + "\n\n")
}

func (r *Output) insertStruct(goStruct GoStruct) {
	toAdd := "type " + goStruct.Name + " struct {\n"
	for _, field := range goStruct.Fields {
		toAdd += "\t" + field.Name + " " +
			field.Type.getString() + " `json:\"" + field.JsonName + "\"`\n"
	}

	toAdd += "}\n\n"
	*r += Output(toAdd)
}

// GoGen represents the generator for Go
// It contains the maps different structs and the package name
type GoGen struct {
	Structures map[string]GoStruct
}

func New() *GoGen {
	return &GoGen{Structures: map[string]GoStruct{}}
}

func (g *GoGen) Generate(packageName string) string {
	structs := make([]GoStruct, 0)
	for _, goStruct := range g.Structures {
		structs = append(structs, goStruct)
	}

	sort.Slice(structs, func(i, j int) bool {
		return len(structs[i].Fields) > (len(structs[j].Fields))
	})

	var output Output = ""
	output.insertHeader(packageName)
	for _, goStruct := range structs {
		output.insertStruct(goStruct)
	}

	return string(output)
}

func (g *GoGen) GetOutPath(packageName string) string {
	return packageName + string(os.PathSeparator) + packageName + ".go"
}

func (g *GoGen) InsertObject(name string, object domain.Object) {
	g.ParseObject(name, object)
}

func (g *GoGen) ParseObject(name string, object domain.Object) GoType {
	newName := toCamelCase(name)
	goStruct := GoStruct{Name: newName, Fields: []GoField{}}

	// Convert map to slice to allow sorting
	objects := make([]GoGenObject, 0)
	for k, v := range object {
		objects = append(objects, GoGenObject{Name: k, Obj: v})
	}

	// Sort the slice of objects by lenght
	// This sort is needed because different fields can have the same structure signature
	// For example a field with name Transactions and PendingTransactions, both this fields should have the
	// same structure signature of Transactions: []Transaction
	sort.Slice(objects, func(i, j int) bool {
		if len(objects[i].Name) != len(objects[j].Name) {
			return len(objects[i].Name) < len(objects[j].Name)
		}

		return objects[i].Name < objects[j].Name
	})

	for _, obj := range objects {
		if IsPrimitiveValue(obj.Obj) {
			goStruct.Fields = append(goStruct.Fields, GoField{IsPrimitive: true, Name: toCamelCase(obj.Name), Type: ParsePrimitiveValue(obj.Obj), JsonName: obj.Name})
		} else {
			goStruct.Fields = append(goStruct.Fields, GoField{IsPrimitive: false, Name: toCamelCase(obj.Name), Type: g.ParseNonPrimitiveValue(obj.Name, obj.Obj), JsonName: obj.Name})
		}
	}

	// Sort the slice of fields by name and primitive type
	// The primitive types are sorted first
	sort.Slice(goStruct.Fields, func(i, j int) bool {
		if goStruct.Fields[i].IsPrimitive == goStruct.Fields[j].IsPrimitive {
			return goStruct.Fields[i].Name < goStruct.Fields[j].Name
		}

		return goStruct.Fields[i].IsPrimitive
	})

	// Checks if the struct already exists
	structure, ok := g.Structures[goStruct.CalcHash()]
	if !ok {
		g.Structures[goStruct.CalcHash()] = goStruct
	} else {
		newName = structure.Name
	}

	return GoType{isArray: false, IsPrimitive: false, CustomType: newName}
}

// ParseNonPrimitiveValue parses a non primitive value
func (g *GoGen) ParseNonPrimitiveValue(key string, value domain.Value) (gtype GoType) {
	switch value.Type {
	case domain.VALUE_OBJECT:
		return g.ParseObject(key, value.Data.(domain.Object))
	case domain.VALUE_ARRAY_OBJ:
		valueArray := value.Data.([]domain.Object)
		if len(valueArray) > 0 {
			gtype = g.ParseObject(key, valueArray[0])
			gtype.isArray = true
			return
		} else {
			gtype = GoType{isArray: false, IsPrimitive: true}
			return
		}
	default:
		gtype = GoType{isArray: false, IsPrimitive: true}
		return
	}
}

func ParsePrimitiveValue(value domain.Value) (gtype GoType) {
	gtype.IsPrimitive = true
	switch value.Type {
	case domain.VALUE_INTEGER:
		gtype.PrimitiveType = GO_TYPE_INT
		return
	case domain.VALUE_STRING:
		gtype.PrimitiveType = GO_TYPE_STRING
		return
	case domain.VALUE_FLOAT:
		gtype.PrimitiveType = GO_TYPE_FLOAT
		return
	case domain.VALUE_BOOL:
		gtype.PrimitiveType = GO_TYPE_BOOL
		return
	case domain.VALUE_NULL:
		gtype.PrimitiveType = GO_TYPE_INTERFACE
		return
	case domain.VALUE_ARRAY_INT:
		gtype.isArray = true
		gtype.PrimitiveType = GO_TYPE_INT
		return
	case domain.VALUE_ARRAY_STR:
		gtype.isArray = true
		gtype.PrimitiveType = GO_TYPE_STRING
		return
	case domain.VALUE_ARRAY_FLOAT:
		gtype.isArray = true
		gtype.PrimitiveType = GO_TYPE_FLOAT
		return
	default:
		gtype.PrimitiveType = GO_TYPE_INTERFACE
		return
	}
}
