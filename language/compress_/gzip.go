package compress_

import (
	"bytes"
	"compress/gzip"
	"io"
)

type GzipCompress struct{}

func init() {
	register(new(GzipCompress))
}

func (GzipCompress) name() string {
	return "gzip"
}

func (GzipCompress) compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(data); err != nil {
		return nil, err
	}
	if err := gz.Flush(); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (GzipCompress) deCompress(data []byte) ([]byte, error) {
	gz, _ := gzip.NewReader(bytes.NewReader(data))
	defer gz.Close()
	return io.ReadAll(gz)
}
