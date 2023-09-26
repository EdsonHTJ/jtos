package jtos

import (
	"fmt"

	"github.com/EdsonHTJ/jtos/domain"
	"github.com/EdsonHTJ/jtos/gens"
	"github.com/EdsonHTJ/jtos/gens/gogen"
	"github.com/EdsonHTJ/jtos/lexer"
	"github.com/EdsonHTJ/jtos/mapper"
)

var Lib = 0

const (
	GOLANG_GENERATOR GeneratorCode = iota
)

type GeneratorCode uint8

type ParseResponse struct {
	Output         string
	RecomendedPath string
}

var generators = map[GeneratorCode]gens.Generator{
	GOLANG_GENERATOR: gogen.New(),
}

func Unmarshal(json string) (domain.Object, error) {
	tokens, err := lexer.GetTokens(json)
	if err != nil {
		return nil, err
	}

	return mapper.MapTokensToObject(tokens)
}

//TODO: add marshal

func ParseJsonFile(packageName string, mainStruct string, json string, generatorCode GeneratorCode) (ParseResponse, error) {
	object, err := Unmarshal(json)
	if err != nil {
		return ParseResponse{}, err
	}

	gen, ok := generators[generatorCode]
	if !ok {
		return ParseResponse{}, fmt.Errorf("invalid generator %d", generatorCode)
	}

	response := ParseResponse{}
	gen.InsertObject(mainStruct, object)
	response.Output = gen.Generate(packageName)
	response.RecomendedPath = gen.GetOutPath(packageName)

	return response, nil
}
