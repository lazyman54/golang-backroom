package concurrent

import (
	"fmt"
	"sync"
)

var m sync.Mutex
var x int

func MutexSolution() {
	var addFun = func() {
		for i := 0; i < 5000; i++ {
			m.Lock()
			x += 1
			m.Unlock()
		}
		wg.Done()
	}
	wg.Add(2)
	go addFun()
	go addFun()
	wg.Wait()
	fmt.Println(x)

}
