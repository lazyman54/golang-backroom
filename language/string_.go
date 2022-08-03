package language

import "fmt"

func StringSolution() {
	traversalString()
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