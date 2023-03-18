package generic

import "fmt"

type MyFunc[T, R any] func(str T) R

func GenericSolution() {

	var realFun = func(str int) string {
		return fmt.Sprintf("str=%d", str)
	}

	var fun MyFunc[int, string] = realFun

	result := call(fun, 12)

	fmt.Println(result)
}

func call[T, R any](myFunc MyFunc[T, R], t T) R {
	return myFunc(t)
}
