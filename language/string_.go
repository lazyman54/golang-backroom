package language

import (
	"fmt"
	"strings"
)

func StringSolution() {
	//traversalString()
	trimSpace()
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
