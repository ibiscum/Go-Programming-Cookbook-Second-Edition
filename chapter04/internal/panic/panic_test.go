package panic

import (
	"testing"
)

func TestPanic(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"base-case", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("the code did not panic")
				}
			}()

			Panic()
		})
	}
}

func TestCatcher(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"base-case", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Catcher()
		})
	}
}
