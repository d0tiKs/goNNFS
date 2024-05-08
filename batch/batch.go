package main

import (
	matfunc "NNFS/utils/math/matrix"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	input := mat.NewDense(3, 4, []float64{
		1, 2, 3, 2.5, 2, 5, -1, 2, -1.5, 2.7, 3.3, -.8,
	})
	fmt.Printf("input: \n%s\n", matfunc.Format(input))

	weights := mat.NewDense(3, 4, []float64{
		0.2, 0.8, -0.5, 1.0, 0.5, -0.91, 0.26, -0.5, -0.26, -0.27, 0.17, 0.87,
	})
	fmt.Printf("weights: \n%s\n", matfunc.Format(weights))

	biases := mat.NewVecDense(3, []float64{
		2, 3, .5,
	})
	fmt.Printf("bias: \n%s\n", matfunc.Format(biases))

	prod, _ := matfunc.ProductMatrix(input, weights.T())
	output, _ := matfunc.AddVector(biases, prod)

	fmt.Printf("output: \n%s\n", matfunc.Format(output))
}
