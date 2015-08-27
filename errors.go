package impart

import (
	"net/http"
)

type HTTPError struct {
	Status  int
	Message string
}

func (h HTTPError) Error() string {
	if h.Message == "" {
		return http.StatusText(h.Status)
	}
	return h.Message
}
