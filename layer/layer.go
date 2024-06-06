package main

import (
	"NNFS/utils/data"
	matfunc "NNFS/utils/math/matrix"
	neuralnetwork "NNFS/utils/neuralNetwork"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	points := 100
	classes := 3
	inputSize := 2
	neuronCount := 3

	X := data.GenerateSpiral(uint(points), uint(classes), 0.2)

	data := mat.NewDense(points*classes, 2, nil)

	for i, point := range X.Points {
		data.SetRow(i, []float64{point.X, point.Y})
	}

	layer := neuralnetwork.DenseLayer{}

	layer.Init(uint(inputSize), uint(neuronCount), neuralnetwork.ReLU)

	layer.Forward(data)
	fmt.Println(matfunc.Format(layer.Output))

	layer.Activate()
	fmt.Println(matfunc.Format(layer.Output))
}
