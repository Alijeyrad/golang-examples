package models

type Request struct {
	Route   string      `json:"route"`
	Content interface{} `json:"content"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}
