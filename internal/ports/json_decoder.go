package ports

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type malformedPayload struct {
	status int
	msg    string
}

func (mr *malformedPayload) Error() string {
	return mr.msg
}

func decodeJSONBody(b io.Reader, dst interface{}) error {
	if b == nil {
		return &malformedPayload{
			status: http.StatusBadRequest,
			msg:    "Invalid: Empty body",
		}
	}

	err := json.NewDecoder(b).Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.Is(err, io.EOF):
			msg := "Invalid: Empty body"
			return &malformedPayload{status: http.StatusBadRequest, msg: msg}
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains invalid JSON (at position %d)", syntaxError.Offset)
			return &malformedPayload{status: http.StatusBadRequest, msg: msg}
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf(
				"Request body contains invalid value for the %q field (at position %d)",
				unmarshalTypeError.Field,
				unmarshalTypeError.Offset)
			return &malformedPayload{status: http.StatusBadRequest, msg: msg}
		}
		return err
	}

	return nil
}
