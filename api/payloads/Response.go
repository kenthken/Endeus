package payloads

import (
	"encoding/json"
	jsonresponse "endeus/api/helper/json/json-response"
	"net/http"
)

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ResponsePagination struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Page       int         `json:"page"`
	Limit      int         `json:"page_limit"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
}

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// NewHandleError creates and returns a new error response
func NewHandleError(writer http.ResponseWriter, errorMessage string, statusCode int) {
	response := ErrorResponse{
		Error: errorMessage,
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		http.Error(writer, "Failed to encode error response", http.StatusInternalServerError)
		return
	}
}

func NewHandleSuccess(writer http.ResponseWriter, data interface{}, message string, status int) {
	res := Response{
		StatusCode: status,
		Message:    message,
		Data:       data,
	}

	jsonresponse.WriteToResponseBody(writer, res)
}

func NewHandleSuccessPagination(writer http.ResponseWriter, data interface{}, message string, status int, limit int, page int, totalRows int64, totalPages int) {
	res := ResponsePagination{
		StatusCode: status,
		Message:    message,
		Page:       page,
		Limit:      limit,
		TotalRows:  totalRows,
		TotalPages: totalPages,
		Data:       data,
	}

	jsonresponse.WriteToResponseBody(writer, res)
}
