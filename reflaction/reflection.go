package main

import (
	"fmt"
	"os"
	"reflect"
)

type TypeOne struct {
	X int
	Y float64
	Z string
}

type TypeSecond struct {
	F int
	G int
	H string
	I float64
}

func main() {
	x := 100
	xRefl := reflect.ValueOf(&x).Elem()
	xType := xRefl.Type()
	fmt.Printf("The type of x is %s.\n", xType)
	//	fmt.Printf("The type of x is %s.\n", reflect.TypeOf(x))

	A := TypeOne{100, 200.12, "Struct TypeOne"}
	B := TypeSecond{1, 2, "Struct TypeSecond", -1.2}
	var r reflect.Value

	arguments := os.Args
	if len(arguments) == 1 {
		r = reflect.ValueOf(&A).Elem()
	} else {
		r = reflect.ValueOf(&B).Elem()
	}

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are:\n", r.NumField(), iType)

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Field name: %s ", iType.Field(i).Name)
		fmt.Printf("with type: %s ", r.Field(i).Type())
		fmt.Printf("and value %v\n", r.Field(i).Interface())
	}
}
