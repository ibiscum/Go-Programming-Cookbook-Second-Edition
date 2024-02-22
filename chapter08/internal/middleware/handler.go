package middleware

import (
	"net/http"
)

// Handler is very basic
func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("success"))
	if err != nil {
		panic(err)
	}
}
