package str

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func StrPerformanceSolution() {

}

func BenchmarkPlusConcat(b *testing.B) {
	benchmark(b, plusConcat)
}
func BenchmarkSprintfConcat(b *testing.B) {
	benchmark(b, sprintfConcat)
}
func BenchmarkBuilderConcat(b *testing.B) {
	benchmark(b, builderConcat)
}
func BenchmarkBufferConcat(b *testing.B) {
	benchmark(b, bufferConcat)
}
func BenchmarkByteConcat(b *testing.B) {
	benchmark(b, byteConcat)
}

func benchmark(b *testing.B, f func(n int, s string) string) {

	originStr := randomString(10)
	for i := 0; i < b.N; i++ {
		f(500000, originStr)
	}

}

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}
func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}
func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
func bufferConcat(n int, s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}
func byteConcat(n int, s string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, s...)
	}
	return string(buf)
}
