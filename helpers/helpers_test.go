package helpers

import (
	"testing"
)

func TestHash(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash(tt.args.data); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsHashProofed(t *testing.T) {
	type args struct {
		hash       string
		difficulty int
		prefix     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHashProofed(tt.args.hash, tt.args.difficulty, tt.args.prefix); got != tt.want {
				t.Errorf("IsHashProofed() = %v, want %v", got, tt.want)
			}
		})
	}
}
