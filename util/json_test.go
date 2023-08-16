package util

import "testing"

func TestJsonSolution(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JsonSolution()
		})
	}
}

func Test_jsonUnmarshal(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "json unmarshal"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jsonUnmarshal()
		})
	}
}
