package matfunc

import (
	errorsutils "NNFS/utils/errors"

	"gonum.org/v1/gonum/mat"
)

func Dot(m, n mat.Matrix) (mat.Matrix, error) {
	mrows, _ := m.Dims()
	_, ncols := n.Dims()

	if mrows != ncols {
		return nil, errorsutils.BuildError(nil, "dimension mismatch")
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

func MergeColumns(matrices ...mat.Matrix) (mat.Matrix, error) {
	rows, cols := matrices[0].Dims()
	result := mat.NewDense(rows*len(matrices), cols, nil)

	for i, m := range matrices {
		mr, mc := m.Dims()

		if mr != rows && mc != cols {
			return nil, errorsutils.BuildError(nil, "dimension mismatch, for the %vth matrice", i)
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
			return nil, errorsutils.BuildError(nil, "dimension mismatch, for the %vth matrice", i)
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
