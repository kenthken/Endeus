package exceptions

import (
	jsonresponse "endeus/api/helper/json/json-response"
	"endeus/api/utils"
	"net/http"

	"github.com/sirupsen/logrus"
)

// BaseErrorResponse defines the general error response structure
type BaseErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
	Err        error  `json:"-"`
}

func NewBadRequestException(writer http.ResponseWriter, request *http.Request, err *BaseErrorResponse) {
	statusCode := http.StatusBadRequest

	if err.Message == "" {
		err.Message = utils.BadRequestError
	}
	if err.Err != nil {
		logrus.Info(err)
		res := &BaseErrorResponse{
			StatusCode: statusCode,
			Message:    err.Err.Error(),
			//Data:       err,
		}

		writer.WriteHeader(statusCode)
		jsonresponse.WriteToResponseBody(writer, res)
		return
	}
}

func NewNotFoundException(writer http.ResponseWriter, request *http.Request, err *BaseErrorResponse) {
	statusCode := http.StatusNotFound
	if err.Message == "" {
		err.Message = utils.GetDataNotFound
	}
	if err.Err != nil {
		logrus.Info(err)
		res := &BaseErrorResponse{
			StatusCode: statusCode,
			Message:    err.Message,
			//Data:       err,
		}

		writer.WriteHeader(statusCode)
		jsonresponse.WriteToResponseBody(writer, res)
		return
	}
}
