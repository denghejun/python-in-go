//package go_python
// main.go
package main

import "github.com/sbinet/go-python"

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
	name := python.PyString_FromString("Tom")
	python.PyTuple_SetItem(args, 0, name)

	helloFunc.Call(args, python.PyDict_New())
}
