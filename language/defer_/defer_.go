package defer_

import (
	"errors"
	"fmt"
	"os"
)

func DeferSolution() {
	//main()
	//main2()
	//main4()
	//main5()
	main6()
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}
func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	//预计：
	/*
		5
		6
		5
		5
	*/
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main2() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
	/*
		预计：
			B ,10, 20, 30
			BB, 10,30,40
			A,10,20,30
			AA,10,30,40
		实际：
			A 1 2 3
			B 10 2 12
			BB 10 12 22
			AA 1 3 4
	*/
}

type T int

func (t T) M(n int) T {
	print(n)
	return t
}
func main3() {
	var t T
	defer t.M(1).M(2)
	t.M(3).M(4)
	//1342
}

type Add func(int, int) int

func main4() {
	var f Add
	defer f(1, 2)
	fmt.Println("end")
	//end; panic
}

func test1() {
	fmt.Println("test")
}
func main5() {
	defer test1()
	os.Exit(0)
	//empty print
}

func e1() {
	var err error
	defer fmt.Println(err)
	err = errors.New("e1 defer err")
	// nil
}
func e2() {
	var err error
	defer func() {
		fmt.Println(err)
	}()
	err = errors.New("e2 defer err")
	// e2 defer err
}
func e3() {
	var err error
	defer func(err error) {
		fmt.Println(err)
	}(err)
	err = errors.New("e3 defer err")
	//nil
}
func main6() {
	e1()
	e2()
	e3()
}
