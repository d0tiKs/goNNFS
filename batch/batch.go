package main

import (
	matfunc "NNFS/utils/matrix"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	input := mat.NewDense(3, 4, []float64{
		1, 2, 3, 2.5, 2, 5, -1, 2, -1.5, 2.7, 3.3, -.8,
	})
	fmt.Printf("input :\n  %.4f\n", mat.Formatted(input, mat.Prefix("  ")))

	weights := mat.NewDense(3, 4, []float64{
		0.2, 0.8, -0.5, 1.0, 0.5, -0.91, 0.26, -0.5, -0.26, -0.27, 0.17, 0.87,
	})
	fmt.Printf("weights :\n  %.4f\n", mat.Formatted(weights, mat.Prefix("  ")))

	bias := mat.NewVecDense(3, []float64{
		2, 3, .5,
	})
	fmt.Printf("bias :\n  %.4f\n", mat.Formatted(bias, mat.Prefix("  ")))

	prod, _ := matfunc.DotMatrix(input, weights.T())
	output, _ := matfunc.AddVector(bias, prod)

	fmt.Printf("output :\n  %.4f\n", mat.Formatted(output, mat.Prefix("  ")))
}
