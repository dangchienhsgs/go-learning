package main

import (
	"errors"
	"fmt"
	"math"
)

type kmeans struct {
	array     [][]float64
	num_elems int
	dimension int
}

func (model *kmeans) fit(input [][]float64) {
	model.array = input
	model.num_elems = len(input)
	model.dimension = len(input[0])

	fmt.Println("Input matrix have ", model.num_elems, " number of elements")
	fmt.Println("Dimension of vectors: ", model.dimension)
}

func (model *kmeans) cluster() {

}

func distance(x, y []float64) float64 {
	if len(x) != len(y) {
		errors.New("Two vector don't have same dimension")
		return 0.0
	} else {
		sum := 0.0
		for i := 0; i < len(x); i++ {
			sum = sum + math.Pow((x[i]-y[i]), 2.0)
		}

		return math.Sqrt(sum)
	}
}

func main() {
	x := [][]float64{
		[]float64{1, 2, 3},
		[]float64{2, 3, 4},
		[]float64{3, 5, 6},
	}

	model := kmeans{}
	model.fit(x)

	fmt.Println(distance(x[1], x[0]))
}
