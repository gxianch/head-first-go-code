package main

import (
	"fmt"
	"mypkg"
)

func main() {
	var value mypkg.MyInterface
	value = mypkg.MyType(1)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	//value.MethodNotInInterface()  //error
	fmt.Println(value.MethodWithReturnValue())

	var value1 = mypkg.MyType(1)
	value1.MethodNotInInterface()
}
