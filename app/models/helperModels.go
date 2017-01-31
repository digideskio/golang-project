package models

type (

	Response struct {
		Status    string `json:"status"`
		Message   string `json:"message"`
		Result    Result `json:"result"`
	}

	Result struct {
		DataType  string `json:"dataType"`
		Data      interface{} `json:"data"`
	}	
)
