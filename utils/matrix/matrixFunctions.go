package matfunc

import (
	errorsutils "NNFS/utils/errors"

	"gonum.org/v1/gonum/mat"
)

func DotVector(a, b mat.Vector) (mat.Vector, error) {
	arows, acols := a.Dims()
	brows, bcols := b.Dims()

	if arows != brows {
		return nil, errorsutils.BuildError(nil, "dimension mismatch")
	}

	if acols != 1 || bcols != 1 {
		return nil, errorsutils.BuildError(nil, "not a vector")
	}

	s := mat.Dot(a, b)
	o := mat.NewVecDense(1, []float64{s})

	return o, nil
}

// Dot product between the two mat.Matrix m and n,
// returns an error if m.Dims()[0] != n.Dims()[1],
// return a matrix of size r.Dims() == {m.Dims()[0], n.Dims()[1]}.
func DotMatrix(m, n mat.Matrix) (mat.Matrix, error) {
	mrows, _ := m.Dims()
	_, ncols := n.Dims()

	if mrows != ncols {
		return nil, errorsutils.BuildError(nil, "dimension mismatch")
	}

	o := mat.NewDense(mrows, ncols, nil)
	o.Product(m, n)
	return o, nil
}

// Add the value of a float64 s to each element of the mat.Matrix m,
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

// Add the value of the transposed mat.Vector v to the mat.Matrix m,
// returns an error if m.Dims()[0] != v.Dims()[0].
func AddVector(v mat.Vector, m mat.Matrix) (mat.Matrix, error) {
	rows, cols := m.Dims()
	vr, _ := v.Dims()

	vt := v.T()

	if vr != rows {
		return nil, errorsutils.BuildError(nil, "dimension mismatch")
	}

	result := mat.NewDense(rows, cols, nil)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			result.Set(r, c, m.At(r, c)+vt.At(0, c))
		}
	}

	return result, nil
}

func MergeColumns(matrices ...mat.Matrix) (mat.Matrix, error) {
	rows, cols := matrices[0].Dims()

	rcols := 0
	for _, m := range matrices {
		_, mc := m.Dims()
		rcols += mc
	}
	result := mat.NewDense(rows, rcols, nil)

	merged := 0
	for i, m := range matrices {
		mr, mc := m.Dims()
		if mr != rows && mc != cols {
			return nil, errorsutils.BuildError(nil, "dimension mismatch, for the %vth matrice", i)
		}

		for c := 0; c < mc; c++ {
			for r := 0; r < mr; r++ {
				el := m.At(r, c)
				result.Set(r, c+merged, el)
			}
		}

		merged += mc
	}

	return result, nil
}

func MergeRows(matrices ...mat.Matrix) (mat.Matrix, error) {
	rows, cols := matrices[0].Dims()

	rrows := 0
	for _, m := range matrices {
		mr, _ := m.Dims()
		rrows += mr
	}
	result := mat.NewDense(rrows, cols, nil)
	merged := 0
	for i, m := range matrices {
		mr, mc := m.Dims()

		if mr != rows && mc != cols {
			return nil, errorsutils.BuildError(nil, "dimension mismatch, for the %vth matrice", i)
		}

		for r := 0; r < mr; r++ {
			for c := 0; c < mc; c++ {
				el := m.At(r, c)
				result.Set(r+merged, c, el)
			}
		}
		merged += mr
	}

	return result, nil
}
