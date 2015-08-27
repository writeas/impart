package impart

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type (
	// envelope contains metadata and optional data for a response object.
	envelope struct {
		MetaData metadata    `json:"meta"`
		Data     interface{} `json:"data,omitempty"`
	}

	// metadata contains more information about the response
	metadata struct {
		Code         int    `json:"code"`
		ErrorType    string `json:"error_type,omitempty"`
		ErrorMessage string `json:"error_message,omitempty"`
	}
)

func writeBody(w http.ResponseWriter, body []byte, status int, contentType string) error {
	w.Header().Set("Content-Type", contentType+"; charset=UTF-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.WriteHeader(status)
	_, err := w.Write(body)
	return err
}

func renderJSON(w http.ResponseWriter, value interface{}, status int) error {
	body, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return writeBody(w, body, status, "application/json")
}

func renderString(w http.ResponseWriter, status int, msg string) error {
	return writeBody(w, []byte(msg), status, "text/plain")
}

func WriteSuccess(w http.ResponseWriter, data interface{}, status int) error {
	env := &envelope{
		MetaData: metadata{
			Code: status,
		},
		Data: data,
	}
	return renderJSON(w, env, status)
}

func WriteError(w http.ResponseWriter, e HTTPError) error {
	env := &envelope{
		MetaData: metadata{
			Code:         e.Status,
			ErrorMessage: e.Message,
		},
	}
	return renderJSON(w, env, e.Status)
}
