package biz

import (
	"fmt"
)

func FinanceSolution() {

	var carAmount float64 = 349000
	var otherAmount float64 = 8000

	var payRate float64 = 0.3
	var rate float64 = 0
	var timeForMonth float64 = 36

	firstPay := float64(carAmount)*payRate + float64(otherAmount)

	monthPay := ((carAmount - firstPay) * (1 + rate)) / timeForMonth

	fmt.Printf("firstPay:%f, pay for month:%f", firstPay, monthPay)

}
