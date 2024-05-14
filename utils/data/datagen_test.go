package data

import (
	"testing"
)

func TestLinearSpace(test *testing.T) {
	start := 0.0
	stop := 1.0
	count := 11

	got := LinearSpace(start, stop, uint(count), 2)
	want := [11]float64{
		0, 0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1,
	}

	for i := 0; i < count; i++ {
		if got[i] != want[i] {
			test.Errorf("got: %v,\nwant: %v", got, want)
			break
		}
	}
}
