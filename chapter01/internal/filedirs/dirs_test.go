package filedirs

import (
	"os"
	"testing"
)

func TestOperate(t *testing.T) {
	os.RemoveAll("example_dir")

	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Operate(); (err != nil) != tt.wantErr {
				t.Errorf("Operate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOperateMkDir(t *testing.T) {
	os.Mkdir("example_dir", os.FileMode(0755))

	tests := []struct {
		name    string
		wantErr bool
	}{
		{"base-case", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Operate(); (err != nil) != tt.wantErr {
				t.Errorf("Operate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
