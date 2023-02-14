package bytes_

import (
	"testing"
)

func TestBytes(t *testing.T) {
	buf := make([]byte, 16, 16)
	for i := 0; i < cap(buf); i++ {
		buf[i] = byte(i)
	}
	t.Logf("before buf len:%d, cap:%d", len(buf), cap(buf))
	buf = buf[:8]
	t.Logf("buf len:%d, cap:%d", len(buf), cap(buf))
	for i := 0; i < cap(buf); i++ {
		t.Log(buf[i])
	}
}
