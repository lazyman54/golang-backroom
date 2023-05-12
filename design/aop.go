package design

import "fmt"

type Hello struct {
	SayHello func() string
	SayWord  func() string
}

type HelloWord struct {
	word string
}

func (cls *HelloWord) SayHello() string {
	fmt.Println("hello")
	return "hello"
}
func (cls *HelloWord) SayWord() string {
	fmt.Println(cls.word)
	return cls.word
}
