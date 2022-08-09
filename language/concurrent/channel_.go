package concurrent

import (
	"fmt"
	"sync"
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
