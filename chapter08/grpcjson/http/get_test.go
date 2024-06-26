package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter08/internal/grpcjson/keyvalue"
	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter08/internal/grpcjson/kvintern"
)

func TestController_GetHandler(t *testing.T) {
	k := kvintern.NewKeyValue()
	_, err := k.Set(context.Background(), &keyvalue.SetKeyValueRequest{Key: "test", Value: "value"})
	if err != nil {
		panic(err)
	}

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name string
		c    *Controller
		args args
	}{
		{"no key", &Controller{KeyValue: kvintern.NewKeyValue()}, args{httptest.NewRecorder(), httptest.NewRequest("GET", "/?key=test", nil)}},
		{"key", &Controller{KeyValue: k}, args{httptest.NewRecorder(), httptest.NewRequest("GET", "/?key=test", nil)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.GetHandler(tt.args.w, tt.args.r)
		})
	}
}
