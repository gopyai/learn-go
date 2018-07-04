package _1

import (
	"fmt"
	"reflect"
)

// 1. Reflection goes from interface value to reflection object.
func Example01A() {
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

	// Output:
	// typeof(x): float64
	// value: 3.4
	// typeof(v): reflect.Value
	// type: float64
	// kind: float64
	// kind is float64: true
	// value: 3.4
	// value: 3.4
}

// 1. Reflection goes from interface value to reflection object.
func Example01B() {
	var x uint8 = 'x'
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is uint8: ", v.Kind() == reflect.Uint8)

	y := v.Uint() // Uint64
	fmt.Println("type:", reflect.TypeOf(y))

	// Output:
	// type: uint8
	// kind is uint8:  true
	// type: uint64
}

// 1. Reflection goes from interface value to reflection object.
func Example01C() {
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("type is main.MyInt:", v.Type().String() == "main.MyInt")
	fmt.Println("kind:", v.Kind())

	fmt.Println(v.Interface().(MyInt))
	// fmt.Println(v.Interface().(int)) // will error conversion

	// Output:
	// type: main.MyInt
	// type is main.MyInt: true
	// kind: int
	// 7
}

// 2. Reflection goes from reflection object to interface value.
func Example02A() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())

	y := v.Interface().(float64)
	fmt.Println(y)

	// Output:
	// type: float64
	// kind: float64
	// 3.4
}

// 3. To modify a reflection object, the value must be settable.
func Example03A() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// v.SetFloat(7.1) // Error: will panic.
	fmt.Println("settability of v:", v.CanSet())

	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	v = p.Elem()
	fmt.Println("settability of v:", v.CanSet())
	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)

	// Output:
	// settability of v: false
	// type of p: *float64
	// settability of p: false
	// settability of v: true
	// 7.1
	// 7.1
}

func ExampleStruct() {
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

	// Output:
	// 0: A int = 23
	// 1: B string = skidoo
	// t is {23 skidoo}
	// t is now {77 Sunset Strip}
}
