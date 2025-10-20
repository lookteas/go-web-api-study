package handler

import (
	"encoding/json"
	"net/http"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse 成功响应
func SuccessResponse(w http.ResponseWriter, data interface{}) {
	response := Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// ErrorResponse 错误响应
func ErrorResponse(w http.ResponseWriter, code int, message string) {
	response := Response{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

// HelloHandler 示例处理器
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"greeting": "Hello from Go Web API!",
		"version":  "1.0.0",
		"method":   r.Method,
		"path":     r.URL.Path,
	}
	SuccessResponse(w, data)
}