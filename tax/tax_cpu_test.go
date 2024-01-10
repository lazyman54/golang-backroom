package tax

import "testing"

func TestCalcAndPrint(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "my"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CalcAndPrint()
		})
	}
}
