package httputils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/silversixx/s3-go/pkg/logger"
)

func ResponseJson(w http.ResponseWriter, status int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)

}

func ResponseJsonError(w http.ResponseWriter, status int, respStr string, err error) {
	if err != nil {
		logger.Error(fmt.Sprintf("Error: %s", err.Error()))
	}
	ResponseJson(w, status, map[string]string{"error": respStr})
}

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}

func GetErrorMsg(err error) string {
	var msg string
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	serr, serr_ok := err.(validator.ValidationErrors)

	switch {
	case errors.As(err, &syntaxError):
		msg = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)

	case errors.Is(err, io.ErrUnexpectedEOF):
		msg = fmt.Sprintf("Request body contains badly-formed JSON")

	case errors.As(err, &unmarshalTypeError):
		msg = fmt.Sprintf("Request body contains an invalid value for the %q field", unmarshalTypeError.Field)

	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		msg = fmt.Sprintf("Request body contains unknown field %s", fieldName)

	case serr_ok:
		msg = fmt.Sprintf("Request body contains invalid value for the %q field", serr[0].Field())
	case errors.Is(err, io.EOF):
		msg = "Request body must not be empty"

	case err.Error() == "http: request body too large":
		msg = "Request body must not be larger than 1MB"

	default:
		msg = "Error parsing request " + err.Error()
	}
	logger.Error(msg)

	return msg
}