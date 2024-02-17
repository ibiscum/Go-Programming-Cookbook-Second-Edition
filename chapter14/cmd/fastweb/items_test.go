package main

import (
	"testing"
)

func TestAddItem(t *testing.T) {
	type args struct {
		item string
	}
	tests := []struct {
		name string
		args args
	}{
		{"base-case", args{"test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddItem(tt.args.item)
		})
	}
}

func TestReadItems(t *testing.T) {
	// type args struct {
	// 	item string
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want []string
	// }{
	// 	{"base-case", args{"test"}, []string{"test"}},
	// }
	// for _, tt := range tests {
	// 	AddItem(tt.args.item)
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := ReadItems(); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("ReadItems() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}
