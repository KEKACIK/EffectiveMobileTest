package api

import (
	"TestTask/pkg/logging"
	"encoding/json"
	"fmt"
	"net/http"
)

func debug(l *logging.Logger, url, method, host string) {
	l.Debug(fmt.Sprintf("%s %s %s", url, method, host))
}

func ErrorResponse(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "error",
		"message": err.Error(),
	})
}
