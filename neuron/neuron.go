package main

import (
	matfunc "NNFS/utils/matrix"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	input := mat.NewVecDense(3, []float64{
		1, 2, 3,
	})

	weights := mat.NewDense(1, 3, []float64{
		0.2, 0.8, -0.5,
	})

	bias := 2.0

	prod, _ := matfunc.DotMatrix(weights, input)
	fmt.Printf("prod :\n  %.4f\n", mat.Formatted(prod, mat.Prefix("  ")))

	output := matfunc.AddScalar(bias, prod)
	fmt.Printf("output :\n  %.4f\n", mat.Formatted(output, mat.Prefix("  ")))
}
