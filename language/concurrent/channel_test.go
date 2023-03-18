package concurrent

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var wg1 sync.WaitGroup

func ChannelSolution() {
	//recv_()
	select_()
}
func recv_() {

	wg1.Add(1)
	ch := make(chan int, 1)
	go recv(ch)
	ch <- 10
	close(ch)
	fmt.Println("send success")
	wg1.Wait()
}
func recv(c <-chan int) {
	time.Sleep(2 * time.Second)
	//ret := <-c
	//fmt.Println("receive", ret)
	//ret = <-c
	//fmt.Println("receive", ret)
	/*当通道c关闭时，range函数会自动退出*/
	for v := range c {
		fmt.Println("receive: ", v)
	}
	wg1.Done()
}

func select_() {
	ch := make(chan int, 2)
	for i := 1; i <= 10; i++ {
		fmt.Printf("the %d time: ", i)
		select {
		case v := <-ch:
			fmt.Println(v)
		case ch <- i:
			fmt.Println()
		}
	}
}

func doCheckClose(taskCh chan int) {
	for {
		select {
		case t, beforeClosed := <-taskCh:
			if !beforeClosed {
				fmt.Println("taskCh has been closed")
				return
			}
			time.Sleep(time.Millisecond)
			fmt.Printf("task %d is done\n", t)
		}
	}
}

func sendTasksCheckClose() {
	taskCh := make(chan int, 10)
	go doCheckClose(taskCh)
	for i := 0; i < 1000; i++ {
		taskCh <- i
	}
	close(taskCh)
}

func TestDoCheckClose(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	sendTasksCheckClose()
	time.Sleep(time.Second)
	runtime.GC()
	t.Log(runtime.NumGoroutine())
}
