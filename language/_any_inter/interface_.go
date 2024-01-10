package _any_inter

import "fmt"

type InterfaceFunc[A any] func(input string) A

func InterfaceFuncExec() {

	var realFunc InterfaceFunc[string] = func(input string) string {
		result := fmt.Sprintf("%s_func", input)
		return result
	}

	var result = realFunc("eric")
	fmt.Println(result)
}
