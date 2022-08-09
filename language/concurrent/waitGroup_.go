package concurrent

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	fmt.Println("hello")
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("i", i)
			wg.Done()
		}()
	}
	wg.Done()
}

func WaitGroupSolution() {
	wg.Add(6)
	go hello()
	fmt.Println("你好")
	wg.Wait()
}
