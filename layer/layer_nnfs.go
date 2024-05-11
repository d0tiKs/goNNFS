package main

import (
	matfunc "NNFS/utils/math/matrix"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	inputs := mat.NewVecDense(4, []float64{
		1.0, 2.0, 3.0, 2.5,
	})
	fmt.Println(matfunc.Format(inputs))

	weights := mat.NewDense(3, 4, []float64{
		.2, .8, -.5, 1.0, .5, -0.91, 0.26, -0.5, -0.26, -0.27, 0.17, 0.87,
	})
	fmt.Println(matfunc.Format(weights))

	biases := mat.NewVecDense(3, []float64{
		2, 3, .5,
	})
	fmt.Println(matfunc.Format(biases))

	product, err := matfunc.ProductMatVec(weights, inputs)
	if err != nil {
		panic(err)
	}
	fmt.Println(matfunc.Format(product))

	outputs, err := matfunc.AddVector(biases, product)
	if err != nil {
		panic(err)
	}
	fmt.Println(matfunc.Format(outputs))
}
