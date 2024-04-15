package concurrent

import "testing"

func TestWaitGroupSolution(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			WaitGroupSolution()
		})
	}
}
