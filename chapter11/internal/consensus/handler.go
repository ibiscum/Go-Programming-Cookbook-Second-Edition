package consensus

import (
	"net/http"
	"time"
)

// Handler grabs the get param ?next= and tries
// to transition to the state contained there
func Handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	state := r.FormValue("next")
	for address, raft := range rafts {
		if address != raft.Leader() {
			continue
		}
		result := raft.Apply([]byte(state), 1*time.Second)
		if result.Error() != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newState, ok := result.Response().(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if newState != state {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte("invalid transition"))
			if err != nil {
				panic(err)
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(newState))
		if err != nil {
			panic(err)
		}
		return
	}
}
