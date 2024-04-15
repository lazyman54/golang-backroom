package language

import (
	"fmt"
	"strings"
)

func StringSolution() {
	//traversalString()
	stringToRune()
}

func trimSpace() {

	s := "   "
	s = strings.TrimSpace(s)
	fmt.Printf(s)
	fmt.Printf("\n %v", s == "")
}

func traversalString() {
	s := "hello沙河"

	fmt.Println(len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c)", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v(%c)", r, r)
	}

}

func stringToRune() {
	strObj := new(strStruct)
	strObj.name = PtrString("I am a string")
	fmt.Printf("print:%v", strObj)
}

type strStruct struct {
	name *string
}

func PtrString(arg string) *string {
	return &arg
}
