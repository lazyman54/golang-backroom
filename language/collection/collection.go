package collection

import "fmt"

func SliceStudy() {
	//array()
	fmt.Println("now try list")
	list()

}

/*
*
数组是值类型，通过方法修改，只是修改副本
*/
func array() {

	var array [5]int           // 默认都是0
	var _ = [5]int{1, 2}       //结果是：1，2，0，0，0
	var _ = [...]int{1, 3, 5}  // 结果是：1，3，5
	var _ = [5]int{0: 5, 4: 3} //结果：5，0，0，0，3
	_ = [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	} //多维数组的创建
	_ = [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
	}

	for i := 0; i < len(array); i++ {
		array[i] = i + 1
	}
	for index, value := range array {
		fmt.Printf("%d: %d\n", index, value)
	}

	//out of bound
	//fmt.Println(array[5])
}

/*
切片是引用类型，拷贝
*/
func list() {
	var array = [5]int{2, 4, 6, 8, 10}
	//定义
	var _ []int         //==nil
	var _ = []int{}     // !=nil
	var list = []int{1} //定义和初始化切片,len=1,cap==1
	fmt.Printf("len:%d, cap:%d \n", len(list), cap(list))
	list = array[1:4] // _: 4,6,8;len=3, cap=4
	fmt.Printf("len:%d, cap:%d \n", len(list), cap(list))
	var _ = make([]int, 5, 9) // len=5, cap=9

	// s2 = s1; s1[0] = newValue => s2[0]==newValue ture

	list = append(list, 4, 5)
	for index, value := range list {
		fmt.Printf("%d: %d\n", index, value)
	}

	/*拷贝*/
	_ = list //s2 指向同一个内存区域
	s2 := make([]int, len(list), cap(list))
	copy(s2, list) // 深度拷贝

	/*删除元素*/
	list = append(list[0:2], list[3:]...)

}

func map_() {

	//定义
	var _ map[string]string
	var myMap = map[string]string{
		"a": "aa",
		"b": "bb",
	}
	var _ = make(map[string]string, 8)

	//	判断key是否存在
	if value, ok := myMap["a"]; ok {
		fmt.Printf("value:%s \n", value)
	} else {
		fmt.Printf("there is not key\n")
	}
	//	删除
	delete(myMap, "c")
	//	map自己无法排序，需要通过外置的slice来排序
}
