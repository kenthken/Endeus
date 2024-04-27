package jsonresponse

import (
	"encoding/json"
	"endeus/api/utils"
	"errors"
	"net/http"
)

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) error {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	if err != nil {
		return errors.New(utils.JsonError)
	}
	return nil
}
