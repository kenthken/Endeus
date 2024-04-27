package helper_test

import (
	exceptions "endeus/api/expections"
	"net/http"
)

func ReturnError(writer http.ResponseWriter, request *http.Request, err *exceptions.BaseErrorResponse) {
	if err.StatusCode == http.StatusBadRequest {
		exceptions.NewBadRequestException(writer, request, err)
		return
	} else if err.StatusCode == http.StatusNotFound {
		exceptions.NewNotFoundException(writer, request, err)
		return
	}
}
