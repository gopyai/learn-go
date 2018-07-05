package main

import "fmt"

func main() {
	data := []byte{1, 2, 3}
	fmt.Println(data)
	f(data)
	fmt.Println(data)
}

func f(data []byte) {
	data[0] = 10
}
