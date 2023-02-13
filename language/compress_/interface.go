package compress_

type ICompress interface {
	name() string
	compress(data []byte) ([]byte, error)
	deCompress(data []byte) ([]byte, error)
}
