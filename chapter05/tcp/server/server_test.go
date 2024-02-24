package main

import (
	"net"
	"testing"
)

func Test_echoBackCapitalized(t *testing.T) {
	type args struct {
		conn net.Conn
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			echoBackCapitalized(tt.args.conn)
		})
	}
}
