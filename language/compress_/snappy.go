package compress_

import (
	"github.com/golang/snappy"
)

type SnappyCompress struct{}

func init() {
	register(new(SnappyCompress))
}

func (SnappyCompress) name() string {
	return "snappy"
}

func (SnappyCompress) compress(data []byte) ([]byte, error) {
	return snappy.Encode(nil, data), nil
}

func (SnappyCompress) deCompress(data []byte) ([]byte, error) {
	return snappy.Decode(nil, data)
}
