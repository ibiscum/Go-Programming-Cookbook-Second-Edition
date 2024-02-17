package firebase

import (
	"context"
	"os"
	"testing"
)

func TestAuthenticate(t *testing.T) {
	if os.Getenv("TEST_WIP") != "" {
		t.Skip("Skipping not finished test")
	}
	type args struct {
		ctx        context.Context
		collection string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"base-case", args{context.Background(), "test"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Authenticate(tt.args.ctx, tt.args.collection)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
