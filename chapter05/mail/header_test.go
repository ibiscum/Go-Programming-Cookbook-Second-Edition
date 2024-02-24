package main

import (
	"net/mail"
	"testing"
)

func Test_printHeaderInfo(t *testing.T) {
	type args struct {
		header mail.Header
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printHeaderInfo(tt.args.header)
		})
	}
}
