package pipeline

import (
	"context"
	"testing"
)

func TestWorker_Print(t *testing.T) {
	type fields struct {
		in  chan string
		out chan string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Worker{
				in:  tt.fields.in,
				out: tt.fields.out,
			}
			w.Print(tt.args.ctx)
		})
	}
}
