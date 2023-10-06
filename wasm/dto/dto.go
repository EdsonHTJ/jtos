package dto

type ParseRequest struct {
	PackageName string `json:"packageName"`
	MainStruct  string `json:"mainStruct"`
	Json        string `json:"json"`
}

type ParseResponse struct {
	Output         string `json:"output"`
	RecomendedPath string `json:"recomendedPath"`
	Error          string `json:"error"`
}
