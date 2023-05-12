package generic

type genData[T any] struct {
	data int64
}

type normalData struct {
	data string
}

type genInterface[T any] interface {
	getGenData() *T
}

func register(obj genInterface[any]) {

}

type DefaultGenStruct[T any] struct {
	data genData[T]
}

func NewDefaultGenStruct[T any]() *DefaultGenStruct[T] {
	return &DefaultGenStruct[T]{data: genData[T]{data: 1}}
}

func (i *DefaultGenStruct[T]) getGenData() *T {
	return nil
}

type MyGenStruct[T any] struct {
	DefaultGenStruct[T]
}

func NewMyGenStruct[T any]() *MyGenStruct[T] {
	return &MyGenStruct[T]{}
}
func (m *MyGenStruct[T]) getGenData() *T {
	return nil
}

func init() {
	//register(NewDefaultGenStruct[int64]())
	register(NewMyGenStruct[int64]())
}
