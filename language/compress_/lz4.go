package compress_

import (
	"bytes"
	"github.com/pierrec/lz4"
	"io"
)

type Lz4Compress struct{}

func init() {
	register(new(Lz4Compress))
}

func (Lz4Compress) name() string {
	return "lz4"
}

func (Lz4Compress) compress(data []byte) ([]byte, error) {
	//var b = make([]byte, len(data))
	//size, _ := lz4.CompressBlock(data, b, []int{1})
	//fmt.Println("real size:", size)
	//return b[:size], nil
	var bf bytes.Buffer
	w := lz4.NewWriter(&bf)
	w.Write(data)
	w.Close()
	return bf.Bytes(), nil
}

func (Lz4Compress) deCompress(data []byte) ([]byte, error) {
	//var b []byte
	//lz4.UncompressBlock(data, b)
	//fmt.Println("originData:", string(b))
	//return b, nil

	r := lz4.NewReader(bytes.NewReader(data))
	return io.ReadAll(r)

}
