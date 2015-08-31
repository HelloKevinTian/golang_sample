/*
	在Golang中用名字调用函数，函数作为第一类值，接口的方案更为通用，与反射直接挂钩
	Call函数式网上大神所写
*/
package main

import (
	// "errors"
	"fmt"
	"reflect"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Exception:", r)
		}
	}()
	//第一种方式
	// funcs := map[string]func(){"foo": foo}
	// funcs["foo"]()

	//第二种方式
	funcs := map[string]interface{}{"foo": foo, "bar": bar}
	// funcs["foo"]() //不能这样用
	Call(funcs, "foo")
	Call(funcs, "bar", 1, 2, 3)

	fmt.Println("main over...")
}

func foo() {
	fmt.Println("foo ...")
}

func bar(a, b, c int) {
	fmt.Println("bar:", a, b, c)
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		panic("The number of params is not adapted.")
		// err := errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}
