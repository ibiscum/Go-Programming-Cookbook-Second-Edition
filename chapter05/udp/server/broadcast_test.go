package main

import (
	"net"
	"testing"
)

func Test_broadcast(t *testing.T) {
	type args struct {
		conn  *net.UDPConn
		conns *connections
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			broadcast(tt.args.conn, tt.args.conns)
		})
	}
}
