package main

import (
	"fmt"
	"math/rand"
	"time"
	"x/matrix"
)

type (
	arrayOfIndex struct {
		I []int
	}
)

func main() {

	inputData := matrix.Floats(4, 2, []float64{
		0, 0,
		0, 1,
		1, 0,
		1, 1,
	})
	targetData := matrix.Floats(4, 1, []float64{
		0.2,
		0.8,
		0.8,
		0.2,
	})

	randomizeMiniBatch(inputData, targetData, 2,
		func(i, t *matrix.Matrix) {
			fmt.Println("Halo:")
			i.Print()
			t.Print()
		})

}

func randomizeMiniBatch(
	inputData,
	targetData *matrix.Matrix,
	miniBatchSize int,
	mini func(*matrix.Matrix, *matrix.Matrix)) {

	rand.Seed(time.Now().UnixNano())

	if inputData.Rows() != targetData.Rows() {
		panic("Error")
	}

	nog1, size1, nog2, size2 := calcMiniBatchSize(inputData.Rows(), miniBatchSize)
	arrOfIdx := randomizeIndex(inputData.Rows())

	var arrIdx int
	x := func(nog, size int) {
		for i := 0; i < nog; i++ {
			inpMini := matrix.Empty(size, inputData.Cols())
			tgtMini := matrix.Empty(size, targetData.Cols())
			for j := 0; j < size; j++ {
				idx := arrOfIdx.I[arrIdx]
				arrIdx++

				inpMini.SetRow(j, inputData.Row(idx))
				tgtMini.SetRow(j, targetData.Row(idx))
			}
			mini(inpMini, tgtMini)
		}
	}
	x(nog1, size1)
	x(nog2, size2)
}

func calcMiniBatchSize(
	totalSize, expectedMiniBatchSize int) (
	numOfGroups1, gSize1, numOfGroups2, gSize2 int) {

	var numOfGroups int
	nog := totalSize / expectedMiniBatchSize
	if totalSize%expectedMiniBatchSize == 0 {
		numOfGroups = nog
	} else {
		numOfGroups = nog + 1
	}

	gSize := totalSize / numOfGroups
	rem := totalSize % numOfGroups
	if rem == 0 {
		numOfGroups1 = numOfGroups
		gSize1 = gSize
	} else {
		numOfGroups1 = rem
		gSize1 = gSize + 1
		numOfGroups2 = numOfGroups - rem
		gSize2 = gSize
	}
	return numOfGroups1, gSize1, numOfGroups2, gSize2
}

func randomizeIndex(arrLen int) *arrayOfIndex {
	idx := &arrayOfIndex{make([]int, arrLen)}
	for i := 0; i < arrLen; i++ {
		idx.I[i] = i
	}
	for i := 0; i < arrLen; i++ {
		j := rand.Intn(arrLen)

		// Swap
		a := idx.I[i]
		b := idx.I[j]
		idx.I[i] = b
		idx.I[j] = a
	}
	return idx
}
