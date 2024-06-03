package dto

import "net/http"

type ResponseResult struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Result  interface{} `json:"result"`
}

func SetResponseData(Result interface{}) ResponseResult {
	return ResponseResult{Code: http.StatusOK, Message: "success", Result: Result}
}

func SetResponseFailure(message interface{}) ResponseResult {
	return ResponseResult{Code: http.StatusBadRequest, Message: message, Result: map[string]interface{}{}}
}

func SetResponseSuccess(Result interface{}) ResponseResult {
	return ResponseResult{Code: http.StatusOK, Message: "success", Result: Result}
}
