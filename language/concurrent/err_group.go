package concurrent

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"time"
)

func ErrGroupSolution() {

	ctx := context.Background()
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(taskExe)
	eg.Go(taskWithErr)
	eg.Go(taskWithDelay)
	if err := eg.Wait(); err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Println("done")
}

func taskWithErr() error {

	randomInt := rand.Intn(2000)
	time.Sleep(time.Duration(randomInt) * time.Millisecond)
	fmt.Printf("task with err exe, wait time:%d ms\n", randomInt)
	return errors.New("taskWithErr")
}
func taskWithDelay() error {

	randomInt := rand.Intn(2000)
	time.Sleep(time.Duration(randomInt) * time.Millisecond)
	fmt.Printf("task with delay exe, wait time:%d ms\n", randomInt)
	return nil
}

func taskExe() error {
	fmt.Printf("task exe\n")
	return nil
}

type data struct {
	name string
	age  int
}

func ErrGroupSolutionWithData() {
	ctx := context.Background()
	var dataList []*data
	for i := 0; i < 2; i++ {
		dataList = append(dataList, &data{
			name: fmt.Sprintf("name%d", i),
			age:  i,
		})
	}
	eg, ctx := errgroup.WithContext(ctx)

	for _, da := range dataList {
		var da1 = da
		eg.Go(func() error {
			da1.name = fmt.Sprintf("%s_new", da1.name)
			fmt.Printf("da:%p\n", da1)
			time.Sleep(time.Duration(2) * time.Second)
			return nil
		})
		eg.Go(func() error {
			da1.age = da1.age + 1
			fmt.Printf("da:%p\n", da1)
			time.Sleep(time.Duration(2) * time.Second)
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		for _, da := range dataList {
			fmt.Printf("data:%+v \n", *da)
		}
	}
}
