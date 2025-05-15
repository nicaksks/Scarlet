package errors

type ErrorResponse struct {
	Code    int `json:"Cod"`
	Message any `json:"Msg"`
}

func ResponseError(code int, message any) ErrorResponse {
	return ErrorResponse{Code: code, Message: message}
}
