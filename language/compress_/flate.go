package compress_

import (
	"bytes"
	"compress/flate"
	"io"
)

type FlateCompress struct{}

func init() {
	register(new(FlateCompress))
}

func (FlateCompress) name() string {
	return "flate"
}

func (FlateCompress) compress(data []byte) ([]byte, error) {
	var bf bytes.Buffer
	w, _ := flate.NewWriter(&bf, flate.BestCompression)
	w.Write(data)
	w.Close()
	return bf.Bytes(), nil
}

func (FlateCompress) deCompress(data []byte) ([]byte, error) {
	r := flate.NewReader(bytes.NewReader(data))
	return io.ReadAll(r)
}
