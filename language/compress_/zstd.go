package compress_

import (
	"bytes"
	"github.com/valyala/gozstd"
	"io"
)

type ZstdCompress struct{}

func init() {
	register(new(ZstdCompress))
}
func (ZstdCompress) name() string {
	return "zstd"
}

func (ZstdCompress) compress(data []byte) ([]byte, error) {
	var bf bytes.Buffer
	w := gozstd.NewWriter(&bf)
	w.Write(data)
	w.Close()
	return bf.Bytes(), nil
}

func (ZstdCompress) deCompress(data []byte) ([]byte, error) {
	r := gozstd.NewReader(bytes.NewReader(data))
	return io.ReadAll(r)

}
