package main

import (
	"fmt"
	"math"
)

func main() {
	var g, m, d, t float64
	g = 0.01
	m = 0
	d = 0.999

	t = 0
	for i := 0; i < 100; i++ {
		t++
		m = d*m + (1-d)*g
		mx := m / (1 - math.Pow(d, t))
		fmt.Printf("%f <= %f\n", mx, m)
	}
}
