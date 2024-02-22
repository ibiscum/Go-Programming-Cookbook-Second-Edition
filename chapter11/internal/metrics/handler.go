package metrics

import (
	"net/http"
	"time"

	"github.com/rcrowley/go-metrics"
)

// CounterHandler will update a counter each time it's called
func CounterHandler(w http.ResponseWriter, r *http.Request) {
	c := metrics.GetOrRegisterCounter("counterhandler.counter", nil)
	c.Inc(1)

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("success"))
	if err != nil {
		panic(err)
	}
}

// TimerHandler records the duration required to compelete
func TimerHandler(w http.ResponseWriter, r *http.Request) {
	currt := time.Now()
	t := metrics.GetOrRegisterTimer("timerhandler.timer", nil)

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("success"))
	if err != nil {
		panic(err)
	}
	t.UpdateSince(currt)
}
