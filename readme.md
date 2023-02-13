data

```go
%v 值的本来输出
%+v 字段名+值的输出
%#v go语言语法格式的值
%c 输出单个字符
%T go语言语法格式的类型和值
%%输出%本体
%b 整型以二进制方式显示
%o 整型以八进制方式显示
%d 整型以十进制方式显示
%x 整型以十六进制方式显示
%X整型以十六进制、字母大写方式显示
%U Unicode字符
%f 浮点数
%p 指针，十六进制方式显示
%3.4f 浮点数，3代表宽度，4代表小数点后的位数
```
close	主要用来关闭channel
len	用来求长度，比如string、array、slice、map、channel
new	用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
make	用来分配内存，主要用来分配引用类型，比如chan、map、slice
append	用来追加元素到数组、slice中
panic和recover	用来做错误处理，recover必须于defer一起使用


```go
uint8    (0 到 255)
uint16   (0 到 65535) 2^16-1
uint32   (0 到 4294967295)
uint64   (0 到 18446744073709551615)
int8     (-128 到 127)
int16    (-32768 到 32767)
int32    (-2147483648 到 2147483647)
int64    (-9223372036854775808 到 9223372036854775807)
```
uint
int
在32位系统中是uint/int32,在64位操作系统中是：uint/int64
uintptr 用于存放一个指针

len(str)	求长度
+或fmt.Sprintf	拼接字符串
strings.Split	分割
strings.contains	判断是否包含
strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
strings.Index(),strings.LastIndex()	子串出现的位置
strings.Join(a[]string, sep string)	join操作

uint8类型，或者叫 byte 型，代表一个ASCII码字符。
rune类型，代表一个 UTF-8字符（实际是int32）
（全英文，用byte或者rune没什么区别，如果有中文，则只能用rune了。rune等价于java的char）

**什么时候应该使用指针类型接收者？**
1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

package管理
命令	                介绍
go mod init	    初始化项目依赖，生成go.mod文件
go mod download	根据go.mod文件下载依赖
go mod tidy  	比对项目文件中引入的依赖与go.mod进行比对
go mod graph 	输出依赖关系图
go mod edit	    编辑go.mod文件
go mod vendor	将项目的所有依赖导出至vendor目录
go mod verify	检验一个依赖包是否被篡改过
go mod why	    解释为什么需要某个依赖

example
```shell
go mod init holiday
```
module holiday //定义当前项目的导入路径
go 1.17 //定义golang的版本号
replace location-a => location-b 代码里使用的是location-a，实际去location-b查找这个包

使用git打tag标签
```shell
 git tag -a v0.1.0 -m "release version v0.1.0"
 git push origin v0.1.0
```

类型断言：
X.(T)
X 当前的值，T假设的类型
```go
if value,ok:=X.(T);ok{
    //do something
}
switch v:=x.(type){
case string:
	//do something
	case int:
		//do something
}
```

```go
string 和 []byte 的互相转换
var byteData []byte = []byte(str)
var str string = string(byteData[:])
```