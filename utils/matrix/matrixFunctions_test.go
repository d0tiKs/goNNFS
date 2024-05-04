package matfunc

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestDotVectorSuccess(test *testing.T) {
	v1 := mat.NewVecDense(3, []float64{
		1, 2, 3,
	})

	v2 := mat.NewVecDense(3, []float64{
		2, 3, 4,
	})

	got, err := DotVector(v1, v2)
	want := mat.NewVecDense(1, []float64{
		20,
	})

	if err != nil {

		test.Fatalf("got an error: %s", err.Error())
		return
	}

	if !mat.EqualApprox(got, want, 0.0001) {
		test.Errorf("got:\n%.4f\n want:\n%.4f\n", mat.Formatted(got, mat.Prefix("  ")), mat.Formatted(want, mat.Prefix("  ")))
	}
}

func TestDotVectorError(test *testing.T) {
	v1 := mat.NewVecDense(2, []float64{
		1, 2,
	})

	v2 := mat.NewVecDense(3, []float64{
		2, 3, 4,
	})

	_, err := DotVector(v1, v2)

	if err == nil {
		test.Errorf("Should return an error")
	}
}

func TestDotMatrixSuccess(test *testing.T) {
	m1 := mat.NewDense(1, 3, []float64{
		1, 2, 3,
	})

	m2 := mat.NewDense(3, 1, []float64{
		2, 3, 4,
	})

	got, err := DotMatrix(m1, m2)
	want := mat.NewDense(1, 1, []float64{
		20,
	})

	if err != nil {

		test.Fatalf("got an error: %s", err.Error())
		return
	}

	if !mat.EqualApprox(got, want, 0.0001) {
		test.Errorf("got:\n%.4f\n want:\n%.4f\n", mat.Formatted(got, mat.Prefix("  ")), mat.Formatted(want, mat.Prefix("  ")))
	}
}

func TestDotMatrixError(test *testing.T) {
	m1 := mat.NewDense(1, 3, []float64{
		1, 2, 3,
	})

	m2 := mat.NewDense(3, 2, []float64{
		2, 3, 4, 5, 6, 7,
	})

	_, err := DotMatrix(m1, m2)

	if err == nil {
		test.Errorf("Should return an error")
	}
}

func TestAddScalar(test *testing.T) {
	m := mat.NewDense(3, 3, []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	s := 42.0

	got := AddScalar(s, m)

	want := mat.NewDense(3, 3, []float64{
		43, 44, 45, 46, 47, 48, 49, 50, 51,
	})

	if !mat.EqualApprox(got, want, 0.0001) {
		test.Errorf("got:\n%.4f\n want:\n%.4f\n", mat.Formatted(got, mat.Prefix("  ")), mat.Formatted(want, mat.Prefix("  ")))
	}
}

func TestAddVector(test *testing.T) {
	m := mat.NewDense(3, 3, []float64{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
	})

	v := mat.NewVecDense(3, []float64{
		1, -1, 0,
	})

	got, err := AddVector(v, m)
	if err != nil {
		test.Fatalf("got an error: %s", err.Error())
		return
	}

	want := mat.NewDense(3, 3, []float64{
		2, 1, 3, 5, 4, 6, 8, 7, 9,
	})

	if !mat.EqualApprox(got, want, 0.0001) {
		test.Errorf("got:\n%.4f\n want:\n%.4f\n", mat.Formatted(got, mat.Prefix("  ")), mat.Formatted(want, mat.Prefix("  ")))
	}
}

func TestMergeColumnsSuccess(test *testing.T) {
	m1 := mat.NewDense(3, 1, []float64{
		1, 2, 3,
	})

	m2 := mat.NewDense(3, 2, []float64{
		2, 5, 3, 6, 4, 7,
	})

	got, err := MergeColumns(m1, m2)
	if err != nil {
		test.Fatalf("got an error: %s", err.Error())
		return
	}

	want := mat.NewDense(3, 3, []float64{
		1, 2, 5, 2, 3, 6, 3, 4, 7,
	})

	if !mat.EqualApprox(got, want, 0.0001) {
		test.Errorf("got:\n  %.4f\n want:\n  %.4f\n", mat.Formatted(got, mat.Prefix("  ")), mat.Formatted(want, mat.Prefix("  ")))
	}
}

func TestMergeColumnsError(test *testing.T) {
	m1 := mat.NewDense(3, 1, []float64{
		1, 2, 3,
	})

	m2 := mat.NewDense(2, 2, []float64{
		2, 5, 3, 6,
	})

	_, err := MergeColumns(m1, m2)
	if err == nil {
		test.Errorf("Should return an error")
	}
}

func TestMergeRowsSuccess(test *testing.T) {
	m1 := mat.NewDense(1, 3, []float64{
		1, 2, 3,
	})

	m2 := mat.NewDense(2, 3, []float64{
		2, 3, 4, 5, 6, 7,
	})

	got, err := MergeRows(m1, m2)
	if err != nil {
		test.Fatalf("got an error: %s", err.Error())
		return
	}

	want := mat.NewDense(3, 3, []float64{
		1, 2, 3, 2, 3, 4, 5, 6, 7,
	})

	if !mat.EqualApprox(got, want, 0.0001) {
		test.Errorf("got:\n  %.4f\n want:\n  %.4f\n", mat.Formatted(got, mat.Prefix("  ")), mat.Formatted(want, mat.Prefix("  ")))
	}
}

func TestMergeRowsError(test *testing.T) {
	m1 := mat.NewDense(3, 1, []float64{
		1, 2, 3,
	})

	m2 := mat.NewDense(2, 2, []float64{
		2, 5, 3, 6,
	})

	_, err := MergeRows(m1, m2)
	if err == nil {
		test.Errorf("Should return an error")
	}
}
