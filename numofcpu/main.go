// main
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	count = 1000000
)

func main() {
	a := make([]float64, count)
	b := make([]float64, count)
	c := make([]float64, count)

	for i := 0; i < count; i++ {
		a[i] = 10
		b[i] = 2
	}

	// Option 1
	t := time.Now()
	for i := 0; i < count; i++ {
		c[i] = a[i] * b[i]
	}
	fmt.Println("Option 1:", time.Since(t).Nanoseconds(), "ns")

	// Option 2
	num := runtime.NumCPU() - 1
	load := count / num
	if count%num != 0 {
		load++
	}

	t = time.Now()
	var wg sync.WaitGroup
	total := count
	pos := 0
	for total > 0 {
		if total >= load {
			total -= load
		} else {
			load = total
			total = 0
		}

		wg.Add(1)
		go func(pos, load int) {
			defer func() {
				wg.Done()
			}()
			for i := pos; load > 0; load-- {
				c[i] = a[i] * b[i]
			}
		}(pos, load)
	}
	wg.Wait()
	fmt.Println("Option 2:", time.Since(t).Nanoseconds(), "ns")

}
