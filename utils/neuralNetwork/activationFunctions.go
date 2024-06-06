package neuralnetwork

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

// Apply the rectified linear function to the value of a dense matrix
//
// Notes:
//   - x   > 0, R(x) -> x
//   - x <= 0, R(x) -> 0.0
//
// Parameters:
//   - m(*mat.Dense): the matrix
func ReLU(m *mat.Dense) {
	mr, mc := m.Dims()

	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			if m.At(i, j) < 0 {
				m.Set(i, j, 0)
			}
		}
	}
}

// Apply the Softmax function to the value of a dense matrix
//
// Notes:
//   - exp(x) -> e^x
//   - sum(f, v, k) -> f(v[0]) + (v[1]) + ... + f(v[k])
//   - l = m.Dims().Collums
//   - S(i, j) -> exp(m[i][j]) / sum(exp, m[i], l)
//
// Parameters:
//   - m(*mat.Dense): the matrix
func Softmax(m *mat.Dense) {
	mr, mc := m.Dims()
	exp := mat.NewDense(mr, mc, nil)
	sumexp := make([]float64, mc)
	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			value := math.Exp(m.At(i, j))
			sumexp[i] += value
			exp.Set(i, j, value)
		}
	}

	for i := 0; i < mr; i++ {
		for j := 0; j < mc; j++ {
			m.Set(i, j, exp.At(i, j)/sumexp[i])
		}
	}
}
