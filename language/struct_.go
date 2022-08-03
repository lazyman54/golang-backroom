package language

import "fmt"

func StructSolution() {

	m := make(map[string]*student)
	stus := []student{
		{name: "xiao wang zi", age: 18},
		{name: "na zha", age: 23},
		{name: "da wang ba", age: 9000},
	}
	/*rang的做法是，新建一个变量，来接收值，每一次循环把取到的值放到变量中，变量的地址不变*/
	for _, stu := range stus {
		fmt.Printf("stu prt:%p, and stu:%v\n", &stu, stu)
		m[stu.name] = &stu
	}
	for i := 0; i < len(stus); i++ {
		fmt.Printf("stu ptr:%p, and stu:%v\n", &stus[i], stus[i])
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
		fmt.Printf("k ptr:%p, v ptr:%p", &k, &v)
	}
}

type student struct {
	name string
	age  int
}
