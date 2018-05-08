package main

import (
	"reflect"
	"fmt"
)

type (
	A struct {
		Name string "Namamu"
		age  int    "Umurmu"
	}
	B struct {
		A          "Huray"
		Kids   int "Woi"
		isGood bool
	}
)

func desc(v reflect.Value) {
	t := v.Type()
	k := v.Kind()

	fmt.Println("Type:", t, "Kind:", k)

	switch k := v.Kind(); k {
	case reflect.Ptr:
		fmt.Println("Ptr")
		desc(v.Elem())
	case reflect.Struct:
		fmt.Println("Struct")
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println("Field:", f.Name, "Tag:", f.Tag, "Type:", f.Type)
			desc(v.Field(i))
		}
	case reflect.String:
		fmt.Println("String")
	case reflect.Int:
		fmt.Println("Int")
	case reflect.Bool:
		fmt.Println("Bool")
	default:
		fmt.Println("NOT FOUND")
	}
}

func main() {
	b := &B{Kids: 10}
	v := reflect.ValueOf(b)
	desc(v)
}
