package pipeline

import (
	"context"
	"reflect"
	"testing"
)

func TestNewPipeline(t *testing.T) {
	type args struct {
		ctx         context.Context
		numEncoders int
		numPrinters int
	}
	tests := []struct {
		name  string
		args  args
		want  chan string
		want1 chan string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewPipeline(tt.args.ctx, tt.args.numEncoders, tt.args.numPrinters)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPipeline() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewPipeline() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
