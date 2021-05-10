//package go_python
// main.go
package main

import (
	"fmt"
	"github.com/sbinet/go-python"
)

func main() {
	python.Initialize()
	defer python.Finalize()

	fooModule := python.PyImport_ImportModule("foo")
	if fooModule == nil {
		panic("Error importing module")
	}

	helloFunc := fooModule.GetAttrString("hello")
	if helloFunc == nil {
		panic("Error importing function")
	}

	// init args, if no any args needed, please give a "python.PyTuple_New(0)"
	args := python.PyTuple_New(1) // 1 arg
	name_mock := python.PyString_FromString("Tom")
	python.PyTuple_SetItem(args, 0, name_mock)

	// get the return value
	result := helloFunc.Call(args, python.PyDict_New())
	name := python.PyDict_GetItemString(result, "name")
	age := python.PyDict_GetItemString(result, "age")
	//x := python.PyDict_GetItem(result, (python.PyList_GetItem(python.PyDict_Keys(result),0)));
	fmt.Printf("The result from Python is : { name: %v, age: %v}", name, age)
}
