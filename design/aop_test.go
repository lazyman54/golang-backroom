package design

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAopProxy(t *testing.T) {

	proxy := InvocationProxy.NewProxyInstance(&Hello{}, func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value {
		return []reflect.Value{reflect.ValueOf("this is a proxy function")}
	}).(*Hello)
	fmt.Println(proxy.SayHello())
}

func TestAopProxyWithClass(t *testing.T) {
	impl := &HelloWord{word: "aop word"}
	proxy := InvocationProxy.NewProxyInstance(&Hello{}, func(obj any, method InvocationMethod, args []reflect.Value) []reflect.Value {
		return method.Invoke(impl, args)
	}).(*Hello)

	proxy.SayWord()
	fmt.Println(proxy.SayHello())
}
