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
	return fmt.Sprintf("%.3f", mat.Formatted(m, mat.Prefix("")))
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
		return 0, errorsutils.BuildError(nil, "dimension mismatch (%v,%v) (%v,%v)", arows, acols, brows, bcols)
	}

	if acols != 1 || bcols != 1 {
		return 0, errorsutils.BuildError(nil, "not a vector")
	}

	s := mat.Dot(a, b)
	return s, nil
}

// Product between two matrices m and n,
//
// Parameters:
//   - m, n: matrices (mat.Matrix)
//
// Return:
//   - a matrix (mat.MAtrix) of dimentions: r.Dims() == {m.Dims()[0], n.Dims()[1]}
//   - an error (error) if m.Dims()[0] != n.Dims()[1]
func ProductMatrix(m, n mat.Matrix) (mat.Matrix, error) {
	mrows, mcols := m.Dims()
	nrows, ncols := n.Dims()

	if mcols != nrows {
		return nil, errorsutils.BuildError(nil, "dimension mismatch (%v,%v) (%v,%v)", mrows, mcols, nrows, ncols)
	}

	o := mat.NewDense(mrows, ncols, nil)
	o.Product(m, n)
	return o, nil
}

// Product between a matrice m and vector v,
//
// Parameters:
//   - m: matrice (mat.Matrix)
//   - v: vector (mat.Vector)
//
// Return:
//   - a vector (mat.Vector) of dimentions: v.Dims() == {v.Dims()[0], 1}
//   - an error (error) if m.Dims()[1] != n.Dims()[0]
func ProductMatVec(m mat.Matrix, v mat.Vector) (mat.Vector, error) {
	mr, mc := m.Dims()
	nr, nc := v.Dims()

	fmt.Printf("(%v,%v) (%v,%v)\n", mr, mc, nr, nc)

	if mc != nr {
		return nil, errorsutils.BuildError(nil, "dimension mismatch (%v,%v) (%v,%v)", mr, mc, nr, nc)
	}

	r := mat.NewVecDense(mr, nil)
	r.MulVec(m, v)
	return r, nil
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
	mr, mc := m.Dims()
	vr, vc := v.Dims()

	if vr != mr {
		return nil, errorsutils.BuildError(nil, "dimension mismatch (%v,%v) (%v,%v)", mr, mc, vr, vc)
	}

	result := mat.NewDense(mr, mc, nil)

	for r := 0; r < vr; r++ {
		for c := 0; c < mc; c++ {
			result.Set(r, c, m.At(r, c)+v.At(r, 0))
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
			return nil, errorsutils.BuildError(nil, "dimension mismatch (%v,%v) (%v,%v), for the %vth matrice", rows, cols, mr, mc, i)
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
			return nil, errorsutils.BuildError(nil, "dimension mismatch (%v,%v) (%v,%v), for the %vth matrice", rows, cols, mr, mc, i)
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
