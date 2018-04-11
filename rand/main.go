// main
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var chances float64 = 1 / 100.0
	fmt.Println(chances)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		if rand.Float64() < chances {
			fmt.Println("Wow!!!")
		}
	}
}
