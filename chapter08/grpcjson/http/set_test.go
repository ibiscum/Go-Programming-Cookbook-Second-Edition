package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter08/internal/grpcjson/kvintern"
)

func TestController_SetHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		c    *Controller
		args args
	}{
		{"base-case", &Controller{KeyValue: kvintern.NewKeyValue()}, args{httptest.NewRecorder(), httptest.NewRequest("GET", "/", bytes.NewBufferString(`{"key":"key", "value":"value"}`))}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.SetHandler(tt.args.w, tt.args.r)
		})
	}
}
