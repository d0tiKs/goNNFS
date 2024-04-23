package main

import (
	"errors"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func BuildError(err error, format string, vars ...interface{}) error {
	errorMessage := fmt.Sprintf(format, vars...)

	if err == nil {
		return errors.New(errorMessage)
	}

	embeddedError := fmt.Sprintf("\nSee error bellow:\n%s", err.Error())
	return errors.New(errorMessage + embeddedError)
}

func Dot(m, n mat.Matrix) (mat.Matrix, error) {
	mrows, _ := m.Dims()
	_, ncols := n.Dims()

	if mrows != ncols {
		return nil, errors.New("dimension mismatch")
	}

	o := mat.NewDense(mrows, ncols, nil)
	o.Product(m, n)
	return o, nil
}

func AddScalar(s float64, m mat.Matrix) mat.Matrix {
	rows, cols := m.Dims()

	values := make([]float64, rows*cols)

	for x := 0; x < rows*cols; x++ {
		values[x] = s
	}

	result := mat.NewDense(rows, cols, values)
	result.Add(result, m)

	return result
}

func MergeCollumns(matrices ...mat.Matrix) (mat.Matrix, error) {
	rows, cols := matrices[0].Dims()
	result := mat.NewDense(rows*len(matrices), cols, nil)

	for i, m := range matrices {
		mr, mc := m.Dims()

		if mr != rows && mc != cols {
			return nil, BuildError(nil, "dimension mismatch, for the %vth matrice", i)
		}
		for x := 0; x < rows; x++ {
			for y := 0; y < cols; y++ {
				el := m.At(x, y)
				result.Set(i*rows+x, y, el)
			}
		}
	}

	return result, nil
}

func MergeRows(matrices ...mat.Matrix) (mat.Matrix, error) {
	rows, cols := matrices[0].Dims()

	result := mat.NewDense(rows, cols*len(matrices), nil)
	for i, m := range matrices {

		mr, mc := m.Dims()

		if mr != rows && mc != cols {
			return nil, BuildError(nil, "dimension mismatch, for the %vth matrice", i)
		}

		for x := 0; x < rows; x++ {
			for y := 0; y < cols; y++ {
				el := m.At(x, y)
				result.Set(x, i*rows+y, el)
			}
		}
	}

	return result, nil
}

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

	prod1, _ := Dot(weights1, input)
	neuron1 := AddScalar(bias1, prod1)
	prod2, _ := Dot(weights2, input)
	neuron2 := AddScalar(bias2, prod2)
	prod3, _ := Dot(weights3, input)
	neuron3 := AddScalar(bias3, prod3)

	output, _ := MergeRows(neuron1, neuron2, neuron3)
	fmt.Printf("output :\n  %.4f\n", mat.Formatted(output, mat.Prefix("  ")))
}
