package main

import (
	matfunc "NNFS/utils/math/matrix"
	neuralnetwork "NNFS/utils/neuralNetwork"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	inputSize := 2
	neuronCount := 3

	input := mat.NewVecDense(inputSize, []float64{
		1.0, 2.0,
	})

	layer := neuralnetwork.DenseLayer{}

	layer.Init(uint(inputSize), uint(neuronCount))
	layer.Input = input

	fmt.Println("")
	fmt.Println(matfunc.Format(layer.Weights))

	layer.Forward()
	fmt.Println(matfunc.Format(layer.Output))
}
