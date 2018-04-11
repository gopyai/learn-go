// main
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type (
	V struct {
		Name  string
		Value []float64
	}
)

func main() {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	enc.Encode(V{Name: "Arief", Value: []float64{1, 2, 3, 4}})

	fmt.Println(b.Bytes())

	var v V
	dec := gob.NewDecoder(&b)
	dec.Decode(&v)

	fmt.Println(v)
}
