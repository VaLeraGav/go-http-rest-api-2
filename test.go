package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct{}

func (m MyStruct) Hello(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func dynamicCall(obj interface{}, methodName string, params ...interface{}) {
	v := reflect.ValueOf(obj)
	method := v.MethodByName(methodName)

	if !method.IsValid() {
		fmt.Printf("Method %s not found\n", methodName)
		return
	}

	// Подготовка аргументов
	args := make([]reflect.Value, len(params))
	for i, param := range params {
		args[i] = reflect.ValueOf(param)
	}

	method.Call(args)
}

func main() {
	myObj := MyStruct{}
	dynamicCall(myObj, "Hello", "World")
}
