package language

import "fmt"

type i interface {
	sayHi()
}
type t struct {
	tt string
}

func (t *t) sayHi() {
	fmt.Printf("say hi %s", t.tt)
}

func TransferSolution() {
	assertTransfer()
}

func assertTransfer() {

	//var a uint = 2342 // not right ,a must be interface type
	var a interface{} = 2342
	if s, ok := a.(int64); ok {
		print(s)
	} else {
		print("not int64")
	}

	var b interface{} = t{
		"tt",
	}
	if ss, ok := b.(i); ok {
		ss.sayHi()
	}

}
