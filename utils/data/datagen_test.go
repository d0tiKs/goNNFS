package data

import (
	"reflect"
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

	if !reflect.DeepEqual(got, want) {
		test.Errorf("got: %v,\nwant: %v", got, want)
	}
}
