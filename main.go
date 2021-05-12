package main

import (
	"fmt"
	"github.com/DataDog/go-python3"
)

type User struct {
	name string
	age  int
}

func main() {
	// 1. Init python and load module 'foo'
	python3.Py_Initialize()
	defer python3.Py_Finalize()
	fooModule := python3.PyImport_ImportModule("package_a.foo") // "{package.module} = {directroy.py_filename}"
	if fooModule == nil {
		panic("Error importing module")
	}

	// 2. Get the callable function 'hello' from python module 'foo'
	helloFunc := fooModule.GetAttrString("hello")

	if helloFunc == nil {
		panic("Error importing function")
	}

	// 3. Prepare parameters (if no any args needed, please give a "python3.PyTuple_New(0)")
	args := python3.PyTuple_New(2) // 2 args
	paramName := python3.PyUnicode_FromString("Leo")
	paramAge := python3.PyLong_FromGoInt(99)
	python3.PyTuple_SetItem(args, 0, paramName)
	python3.PyTuple_SetItem(args, 1, paramAge)

	// 4. Call Python function
	result := helloFunc.Call(args, python3.Py_None)

	// 5. Receive the result to Golang
	user := User{
		name: python3.PyUnicode_AsUTF8(python3.PyDict_GetItemString(result, "name")),
		age:  python3.PyLong_AsLong(python3.PyDict_GetItemString(result, "age")),
	}

	fmt.Printf("The result from Python is : { name: %v, age: %v}", user.name, user.age)
}
