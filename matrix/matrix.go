// matrix
package main

import (
	"fmt"

	"github.com/gonum/matrix/mat64"
)

func main() {
	t3()
}

func t3() {
	m := mat64.NewDense(4, 3, []float64{
		1, 0, 0,
		1, 0, 1,
		1, 1, 0,
		1, 1, 1,
	})
	m.Apply(func(i, j int, v float64) float64 {
		return v * 0.01
	}, m)
	fmt.Println(mat64.Formatted(m))
}

func t2() {
	inp := mat64.NewDense(4, 3, []float64{
		1, 0, 0,
		1, 0, 1,
		1, 1, 0,
		1, 1, 1,
	})
	wgh := mat64.NewDense(3, 2, []float64{
		1, 1,
		1, 1,
		1, 1,
	})
	out := mat64.NewDense(4, 2, nil)

	out.Mul(inp, wgh)
	fmt.Println(mat64.Formatted(out, mat64.Prefix("")))

	out.Apply(func(i, j int, v float64) float64 {
		return v * 1.5
	}, out)
	fmt.Println(mat64.Formatted(out, mat64.Prefix("")))

}

func t1() {

	z1 := mat64.NewDense(2, 2, nil)
	z2 := mat64.NewDense(2, 2, nil)
	yi := mat64.NewDense(2, 2, nil)

	//
	// z1 = x . y
	//

	x := mat64.NewDense(2, 2, []float64{1, 2, 3, 4})
	y := mat64.NewDense(2, 2, []float64{-1, 7, 4, 2})
	z1.Mul(x, y)
	fmt.Printf("%f %f %f %f\n", z1.At(0, 0), z1.At(0, 1), z1.At(1, 0), z1.At(1, 1))
	fmt.Println(mat64.Formatted(z1))

	//
	// z2 = z1 . yinverse = x.y.yinverse = x.I = x
	//

	Panic(yi.Inverse(y))
	z2.Mul(z1, yi)
	fmt.Printf("%f %f %f %f\n", z2.At(0, 0), z2.At(0, 1), z2.At(1, 0), z2.At(1, 1))
	fmt.Println(mat64.Formatted(z2))

}

func Panic(e error) {
	if e != nil {
		panic(e)
	}
}
