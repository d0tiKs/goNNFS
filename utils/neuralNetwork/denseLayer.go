package neuralnetwork

import (
	errorsutils "NNFS/utils/errors"
	matfunc "NNFS/utils/math/matrix"
	"math/rand"

	"gonum.org/v1/gonum/mat"
)

type DenseLayer struct {
	Weights    *mat.Dense
	Biases     *mat.VecDense
	Output     *mat.Dense
	Activation func(*mat.Dense)
}

// Initialize a Dense layer with random weights
//
// Parameters:
//   - inputSize(uint): the numbers of features of the input
//   - neuronsCount(uint): the number of neuron in the layer
func (layer *DenseLayer) Init(inputSize, neuronsCount uint, activation func(*mat.Dense)) {
	layer.Biases = mat.NewVecDense(int(neuronsCount), nil)
	layer.Weights = mat.NewDense(int(inputSize), int(neuronsCount), nil)

	for i := 0; i < int(inputSize); i++ {
		for j := 0; j < int(neuronsCount); j++ {
			layer.Weights.Set(i, j, .1*rand.Float64())
		}
	}

	layer.Activation = activation
}

// Forward the input to the layer and compute the output
//
// Parameters:
//   - inputs(*mat.Dense): the vector of inputs
func (layer *DenseLayer) Forward(inputs *mat.Dense) {
	product, err := matfunc.ProductMatrix(inputs, layer.Weights)
	if err != nil {
		// FIX: return error instead of panic
		panic(errorsutils.BuildError(err, "error while forwarding : inputs * weights"))
	}

	layer.Output, err = matfunc.AddVector(layer.Biases, product)
	if err != nil {
		// FIX: return error instead of panic
		panic(errorsutils.BuildError(err, "error while forwarding : product + bias"))
	}
}

func (layer *DenseLayer) Activate() {
	layer.Activation(layer.Output)
}
