package compress_

import (
	"bytes"
	"compress/zlib"
	"io"
)

type ZlibCompress struct{}

func init() {
	register(new(ZlibCompress))
}
func (ZlibCompress) name() string {
	return "zlib"
}

func (ZlibCompress) compress(data []byte) ([]byte, error) {
	var bf bytes.Buffer
	w := zlib.NewWriter(&bf)
	w.Write(data)
	w.Close()
	return bf.Bytes(), nil
}

func (ZlibCompress) deCompress(data []byte) ([]byte, error) {
	r, _ := zlib.NewReader(bytes.NewReader(data))
	return io.ReadAll(r)
}

