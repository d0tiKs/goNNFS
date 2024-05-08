package matfunc

import (
	errorsutils "NNFS/utils/errors"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// Format the data of a matrix in a humanreadble way
//
// Parameters:
//   - m: a matrix (mat.Matrix)
//
// Return:
//   - a string (string)
func Format(m mat.Matrix) string {
	return fmt.Sprintf("%.4f", mat.Formatted(m, mat.Prefix("")))
}

// Dot product between the two vectors a and b,
//
// Parameters:
//   - a, b: vectors (mat.Vector)
//
// Return:
//   - a scalar (float64)
//   - an error (error) if a.Dims()[0] != b.Dims()[0],
func DotVector(a, b mat.Vector) (float64, error) {
	arows, acols := a.Dims()
	brows, bcols := b.Dims()

	if arows != brows {
		return 0, errorsutils.BuildError(nil, "dimension mismatch")
	}

	if acols != 1 || bcols != 1 {
		return 0, errorsutils.BuildError(nil, "not a vector")
	}

	s := mat.Dot(a, b)
	return s, nil
}

// Product between the two matrices m and n,
//
// Parameters:
//   - m, n: matrices (mat.Matrix)
//
// Return:
//   - a matrix (mat.MAtrix) of dimentions: r.Dims() == {m.Dims()[0], n.Dims()[1]}
//   - an error (error) if m.Dims()[0] != n.Dims()[1]
func ProductMatrix(m, n mat.Matrix) (mat.Matrix, error) {
	mrows, _ := m.Dims()
	_, ncols := n.Dims()

	if mrows != ncols {
		return nil, errorsutils.BuildError(nil, "dimension mismatch")
	}

	o := mat.NewDense(mrows, ncols, nil)
	o.Product(m, n)
	return o, nil
}

// Add the value of a sclacar s to each element of the matrix m,
//
// Parameters:
//   - s: a scalar (float64)
//   - m: a matrix (mat.Matrix)
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

// Add the value of the transposed vector v to the matrix m,
//
// Parameters:
//   - v: a vector (mat.Vector)
//   - m: a matrix (mat.Matrix)
//
// Return:
//   - a matrix (mat.Matrix)
//   - an error error if m.Dims()[0] != v.Dims()[0].
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

// Merge a list of matrices collums wise
//
// Parameters:
//   - matrices: a slice of matrix ([]mat.Matrix)
//
// Return:
//   - a matrice (mat.Matrix)
//   - an error (error) if at least once of the matrices doesn't match the rows dimention of the first matrix in the list:
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

// Merge a list of matrices rows wise
//
// Parameters:
//   - matrices: a slice of matrix ([]mat.Matrix)
//
// Return:
//   - a matrice (mat.Matrix)
//   - an error (error) if at least once of the matrices doesn't match the collumn dimention of the first matrix in the list:
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
