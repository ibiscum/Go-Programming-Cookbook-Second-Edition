package crypto

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// GuessHandler checks if ?message=password
func GuessHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		// if we can't parse the form
		// we'll assume it is malformed
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("error reading guess"))
		if err != nil {
			panic(err)
		}
		return
	}

	msg := r.FormValue("message")

	// "password"
	real := []byte("$2a$10$2ovnPWuIjMx2S0HvCxP/mutzdsGhyt8rq/JqnJg/6OyC3B0APMGlK")

	if err := bcrypt.CompareHashAndPassword(real, []byte(msg)); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("try again"))
		if err != nil {
			panic(err)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("you got it"))
	if err != nil {
		panic(err)
	}
}
