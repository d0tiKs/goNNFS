package main

import (
	matfunc "NNFS/utils/math/matrix"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	input := mat.NewVecDense(4, []float64{
		1, 2, 3, 2.5,
	})

	weights1 := mat.NewDense(1, 4, []float64{
		0.2, 0.8, -0.5, 1,
	})

	weights2 := mat.NewDense(1, 4, []float64{
		0.5, -0.91, 0.26, -0.5,
	})

	weights3 := mat.NewDense(1, 4, []float64{
		-0.26, -0.27, 0.17, 0.87,
	})
	bias1 := 2.0
	bias2 := 3.0
	bias3 := 0.5

	prod1, _ := matfunc.ProductMatrix(weights1, input)
	neuron1 := matfunc.AddScalar(bias1, prod1)
	prod2, _ := matfunc.ProductMatrix(weights2, input)
	neuron2 := matfunc.AddScalar(bias2, prod2)
	prod3, _ := matfunc.ProductMatrix(weights3, input)
	neuron3 := matfunc.AddScalar(bias3, prod3)

	output, _ := matfunc.MergeRows(neuron1, neuron2, neuron3)
	fmt.Printf("output :\n  %.4f\n", mat.Formatted(output, mat.Prefix("  ")))
}
