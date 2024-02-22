package main

import (
	"encoding/json"
	"net/http"

	"github.com/apex/log"
	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter08/internal/grpcjson/keyvalue"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetHandler wraps our RPC Get call
func (c *Controller) GetHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	kv := keyvalue.GetKeyValueRequest{Key: key}

	gresp, err := c.Get(r.Context(), &kv)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		log.Errorf("failed to get: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	resp, err := json.Marshal(gresp)
	if err != nil {
		log.Errorf("failed to marshal: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		panic(err)
	}
}
