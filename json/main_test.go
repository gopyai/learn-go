package main_test

import (
	"encoding/json"
	"fmt"
)

type (
	data struct {
		Name string
		Age  int
	}
)

func Example() {
	var d data
	fmt.Println("Before:", d)
	unmarshal(marshal(), &d)
	fmt.Println("After:", d)
	// Output:
	// Before: { 0}
	// Marshal: {"Name":"Arief","Age":17}
	// Unmarshal: &{Arief 17}
	// After: {Arief 17}
}

func marshal() []byte {
	b, err := json.Marshal(data{"Arief", 17})
	panicIf(err)
	fmt.Println("Marshal:", string(b))
	return b
}

func unmarshal(b []byte, d interface{}) {
	err := json.Unmarshal(b, d)
	panicIf(err)
	fmt.Println("Unmarshal:", d)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
