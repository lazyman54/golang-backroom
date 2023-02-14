package perfomance

import "testing"

type Item struct {
	id  int
	val [4096]byte
}

func BenchmarkRangePointer(b *testing.B) {
	var persons [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, value := range persons {
			tmp = value.id
		}
		_ = tmp
	}

}

func BenchmarkForPointer(b *testing.B) {

	var persons [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for i := 0; i < len(persons); i++ {
			tmp = persons[i].id

		}
		_ = tmp
	}
}
