package storage

import (
	"testing"
)

func TestExec(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Exec(); (err != nil) != tt.wantErr {
				t.Errorf("Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPerformOperations(t *testing.T) {
	type args struct {
		s Storage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PerformOperations(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("PerformOperations() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
