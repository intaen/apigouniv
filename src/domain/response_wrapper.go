package domain

type ResponseWrapper struct {
	StatusCode string      `json:"code"`
	Message    string      `json:"message"`
	Result     interface{} `json:"result"`
}
