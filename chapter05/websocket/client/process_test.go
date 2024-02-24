package main

import (
	"testing"

	"github.com/gorilla/websocket"
)

func Test_process(t *testing.T) {
	type args struct {
		c *websocket.Conn
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			process(tt.args.c)
		})
	}
}
