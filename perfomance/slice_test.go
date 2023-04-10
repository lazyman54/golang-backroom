package perfomance

import "testing"

func BenchmarkSliceNoCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sli := make([]int, 0, 4)
		sli = append(sli, 1, 2, 3, 4)
	}
}

func BenchmarkSliceCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sli := make([]int, 0, 2)
		sli = append(sli, 1, 2, 3, 4)
	}
}
