// main
package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	cnt = 100
)

func main() {
	var x time.Time
	fmt.Println(x, x.IsZero())

	var t2 time.Time
	t := time.Now()

	var wg sync.WaitGroup
	wg.Add(1)
	time.AfterFunc(time.Millisecond*110, func() {
		defer func() {
			wg.Done()
		}()
		t2 = time.Now()
	})
	wg.Wait()

	fmt.Println(t)
	fmt.Println(t2)
	fmt.Println(t2.Sub(t).Seconds()*1000, "ms")
}
