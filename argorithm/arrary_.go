package argorithm

import (
	"fmt"
	"math"
	"sort"
)

/*
标题
两个自然数数组差值绝对值最小值

题目描述
给定两个自然数数组A和B，从A中任取一个元素a，从B中任取一个元素b。d=绝对值(a-b)求d的最小值。A和B的长度都不为0
举例: A=[ 1,3,2,10]，B = [7,13，9,13]，输出1 (即10-9)给出时间复杂度和空间复杂度
*/
func Solution() {
	arr1 := []int{1, 3, 2, 11}
	arr2 := []int{7, 13, 9, 13}
	sort.Ints(arr1)
	sort.Ints(arr2)
	index1 := 0
	index2 := 0
	minNum := math.MaxInt
	for index1 < len(arr1) && index2 < len(arr2) {
		num1 := arr1[index1]
		num2 := arr2[index2]

		val := int(math.Abs(float64(num1 - num2)))
		if val < minNum {
			minNum = val
		}
		if num1 <= num2 {
			index1++
		} else {
			index2++
		}
	}
	fmt.Println(minNum)
}
