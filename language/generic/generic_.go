package generic

import (
	"fmt"
	"reflect"
)

type data struct {
	value string
}
type MyFunc[T, R any] func(str T) R

func GenericSolution() {

	var realFun = func(str int) *data {
		d := new(data)
		d.value = fmt.Sprintf("str=%d", str)
		return nil
	}

	var fun MyFunc[int, *data] = realFun

	result := call(fun, 12)

	fmt.Println(result)
}

func call[T, R any](myFunc MyFunc[T, R], t T) R {
	var my = new(R)

	var mydata = data{}
	var mydataPoint = *my
	pv1 := reflect.ValueOf(my)
	fmt.Println(pv1.Kind() == reflect.Ptr)
	fmt.Printf("my data:%v", mydata)
	fmt.Printf("data:%v", my)
	fmt.Printf("data:%v", mydataPoint)
	fmt.Println(my)
	return myFunc(t)
}
