package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/EdsonHTJ/jtos"
	"github.com/EdsonHTJ/jtos/wasm/dto"
)

func ParseJsonToLang(this js.Value, args []js.Value) interface{} {
	parseRequest := dto.ParseRequest{}
	response := dto.ParseResponse{}

	err := json.Unmarshal([]byte(args[0].String()), &parseRequest)
	if err != nil {
		response.Error = err.Error()
		return js.ValueOf(response)
	}

	parsed, err := jtos.ParseJsonFile(parseRequest.PackageName, parseRequest.MainStruct, parseRequest.Json, jtos.GOLANG_GENERATOR)
	if err != nil {
		response.Error = err.Error()
		return js.ValueOf(response)
	}

	response.Output = parsed.Output
	response.RecomendedPath = parsed.RecomendedPath

	return js.ValueOf(response)
}

func main() {
	js.Global().Set("parseJsonToLang", js.FuncOf(ParseJsonToLang))
}
