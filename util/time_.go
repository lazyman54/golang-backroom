package util

import (
	"fmt"
	carbon "github.com/golang-module/carbon"
	"time"
)

func TimeSolution() {

	timeUtilAndSince()

}

func timeUtilAndSince() {

	now := time.Now()
	time.Sleep(time.Duration(4) * time.Second)
	fmt.Println(time.Since(now).Milliseconds())

}

func timeCarbon() {
	i := carbon.Now().AddDays(5).DayOfWeek()
	fmt.Println(i)
	day := carbon.Now().SetWeekStartsAt(carbon.Monday).StartOfWeek()
	lastDay := carbon.Now().SetWeekStartsAt(carbon.Monday).EndOfWeek()
	fmt.Println(day)
	fmt.Println(lastDay)

	now := time.Now()
	fmt.Println(now)
	before := now.Add(time.Duration(-7 * time.Hour * 24))
	fmt.Println(before)
}
