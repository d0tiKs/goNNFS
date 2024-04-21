package main

import (
	"errors"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func dot(m, n mat.Matrix) (mat.Matrix, error) {
	mrows, _ := m.Dims()
	_, ncols := n.Dims()

	if mrows != ncols {
		return nil, errors.New("dimension mismatch")
	}

	o := mat.NewDense(mrows, ncols, nil)
	o.Product(m, n)
	return o, nil
}

func addScalar(s float64, m mat.Matrix) mat.Matrix {
	rows, cols := m.Dims()

	values := make([]float64, rows*cols)

	for x := 0; x < rows*cols; x++ {
		values[x] = s
	}

	result := mat.NewDense(rows, cols, values)
	result.Add(result, m)

	return result
}

func main() {
	input := mat.NewVecDense(3, []float64{
		1, 2, 3,
	})

	weights := mat.NewDense(1, 3, []float64{
		0.2, 0.8, -0.5,
	})

	bias := 2.0

	prod, _ := dot(weights, input)
	fmt.Printf("prod :\n  %.4f\n", mat.Formatted(prod, mat.Prefix("  ")))

	output := addScalar(bias, prod)
	fmt.Printf("output :\n  %.4f\n", mat.Formatted(output, mat.Prefix("  ")))
}
