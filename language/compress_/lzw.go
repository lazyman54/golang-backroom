package compress_

import (
	"bytes"
	"compress/lzw"
	"io"
)

type LzwCompress struct{}

func init() {
	register(new(LzwCompress))
}

func (LzwCompress) name() string {
	return "lzw"
}

func (LzwCompress) compress(data []byte) ([]byte, error) {
	var bf bytes.Buffer
	w := lzw.NewWriter(&bf, lzw.LSB, 8)
	w.Write(data)
	w.Close()
	return bf.Bytes(), nil
}

func (LzwCompress) deCompress(data []byte) ([]byte, error) {
	r := lzw.NewReader(bytes.NewReader(data), lzw.LSB, 8)
	return io.ReadAll(r)
}
