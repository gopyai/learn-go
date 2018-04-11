package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 1. Reflection goes from interface value to reflection object.
	r1a()
	r1b()
	r1c()

	// 2. Reflection goes from reflection object to interface value.
	r2a()

	// 3. To modify a reflection object, the value must be settable.
	r3a()

	// Struct
	rs()
}

func r1a() {
	fmt.Println("\n### 1A ###")
	var x float64 = 3.4
	fmt.Println("typeof(x):", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))

	v := reflect.ValueOf(x)
	fmt.Println("typeof(v):", reflect.TypeOf(v))
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
	fmt.Println("value:", v)
}

func r1b() {
	fmt.Println("\n### 1B ###")
	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8)

	y := v.Uint() //Uint64
	fmt.Println("type:", reflect.TypeOf(y))
}

func r1c() {
	fmt.Println("\n### 1C ###")
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("type is main.MyInt:", v.Type().String() == "main.MyInt")
	fmt.Println("kind:", v.Kind())

	fmt.Println(v.Interface().(MyInt))
	//fmt.Println(v.Interface().(int)) // will error conversion
}

func r2a() {
	fmt.Println("\n### 2A ###")

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())

	y := v.Interface().(float64)
	fmt.Println(y)
}

func r3a() {
	fmt.Println("\n### 3A ###")

	var x float64 = 3.4
	v := reflect.ValueOf(x)
	//v.SetFloat(7.1) // Error: will panic.
	fmt.Println("settability of v:", v.CanSet())

	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	v = p.Elem()
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}

func rs() {
	fmt.Println("\n### Struct ###")
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	//

	fmt.Println("t is", t)
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
}
