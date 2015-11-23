package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

type kmeans struct {
	elems        [][]float64
	num_elems    int
	dimension    int
	num_clusters int

	labels  []int
	centers [][]float64
}

func (model *kmeans) fit(input [][]float64, num_clusters int) {
	model.elems = input
	model.num_elems = len(input)
	model.dimension = len(input[0])
	model.num_clusters = num_clusters

	fmt.Println("Input matrix have ", model.num_elems, " number of elements")
	fmt.Println("Dimension of vectors: ", model.dimension)

	model.cluster()

	fmt.Println("Result: ", model.labels)
}

func (model *kmeans) cluster() {
	model.init_model()

	for i := 0; i < 10; i++ {
		model.update_centers()
		model.update_labels()
	}
}

func (model *kmeans) update_centers() {
	for i := 0; i < model.num_clusters; i++ {
		model.centers[i] = create_zeros_vector(model.dimension)
	}

	count_label := create_zeros_vector(model.num_clusters)
	for i := 0; i < model.num_elems; i++ {
		label := model.labels[i]

		model.centers[label] = sum_two_vector(
			model.centers[label],
			model.elems[i])

		count_label[label] = count_label[label] + 1
	}

	for i := 0; i < model.num_clusters; i++ {
		model.centers[i] = divide_vector(model.centers[i], count_label[i])

		fmt.Println("Update center ", i, " : ", model.centers[i])
	}
}

func divide_vector(vector []float64, divisor float64) []float64 {
	for i := 0; i < len(vector); i++ {
		vector[i] = vector[i] / divisor
	}
	return vector
}

func (model *kmeans) update_labels() {
	for i := 0; i < model.num_elems; i++ {
		index_min := 0
		value_min := distance(model.elems[i], model.centers[0])

		for j := 1; j < model.num_clusters; j++ {
			temp := distance(model.elems[i], model.centers[j])
			if temp < value_min {
				value_min = temp
				index_min = j
			}
		}

		model.labels[i] = index_min
	}

	fmt.Println("Update labels: ", model.labels)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func create_zeros_vector(dimension int) []float64 {
	vector := make([]float64, dimension)
	return vector
}

func sum_two_vector(x, y []float64) []float64 {
	z := create_zeros_vector(len(x))
	if len(x) != len(y) {
		errors.New("Two vector do not have same dimension")
	} else {
		for i := 0; i < len(x); i++ {
			z[i] = x[i] + y[i]
		}
	}

	return z
}

func (model *kmeans) init_model() {
	model.labels = make([]int, model.num_elems)

	// to sure that every label must have at least 1 element
	for i := 0; i < model.num_clusters; i++ {
		model.labels[i] = i
	}

	for i := model.num_clusters; i < model.num_elems; i++ {
		model.labels[i] = random(0, model.num_clusters-1)
	}

	fmt.Println("Init model labels: ", model.labels)
	// init centers
	model.centers = make([][]float64, model.num_clusters)
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
	rand.Seed(time.Now().Unix())

	x := [][]float64{
		[]float64{1, 2, 3},
		[]float64{2, 3, 4},
		[]float64{3, 5, 6},
		[]float64{0, 0, 1},
		[]float64{1, 2, 3},
		[]float64{1, 1, 3},
		[]float64{1, 1, 5},
		[]float64{5, 2, 3},
		[]float64{1, 1, 7},
		[]float64{0, 1, 7},
	}

	model := kmeans{}
	model.fit(x, 3)
}
