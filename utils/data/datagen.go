package data

import (
	mathutils "NNFS/utils/math"
	"fmt"
	"math"

	"gonum.org/v1/plot/plotter"
)

type Label struct {
	text string
	id   int
}

type Dataset struct {
	Points  plotter.XYs
	Labels  []Label
	Classes uint
}

// Generate a spiral dataset for classification.
//
// Parameters:
//   - count (uint): Number of points per class
//   - classes (uint): Number of classes
//   - noise (float64): The standard deviation of Gaussian noise added to the angles
//
// Return:
//   - a array of points ([count*classes][2]float64)
//   - label vector ([]Label) of length (count * classes)
func GenerateSpiral(count, classes uint, noise float64) Dataset {
	dataset := make([]plotter.XY, count*classes)
	labels := make([]Label, count*classes)

	radius := LinearSpace(0, 1, count, 4)
	for class := 0; class < int(classes); class++ {
		theta := LinearSpace(float64(class*4), float64((class+1)*4), count, 4)
		for i := 0; i < int(count); i++ {
			idx := (class * int(count)) + i

			dataset[idx] = plotter.XY{
				X: radius[i] * math.Sin(theta[i]),
				Y: radius[i] * math.Cos(theta[i]),
			}
			labels[idx] = Label{
				text: fmt.Sprintf("%v", class),
				id:   class,
			}
		}
	}

	return Dataset{Points: dataset, Labels: labels, Classes: classes}
}

func GenerateSpiral2(count, classes uint) Dataset {
	dataset := make([]plotter.XY, count*classes)
	labels := make([]Label, count*classes)

	for i := 0; i < int(count); i++ {
		theta := float64(i) * (4 * math.Pi / float64(count))
		radius := float64(i) / float64(count) * 10
		x := radius * math.Cos(theta)
		y := radius * math.Sin(theta)

		// Determine the class by angle
		class := int((theta/(2*math.Pi))*numSlices) % numSlices
		dataset[i] = plotter.XY{
			X: x,
			Y: y,
		}
		labels[i] = Label{
			text: fmt.Sprintf("%v", class),
			id:   class,
		}
	}
	return Dataset{Points: dataset, Labels: labels, Classes: classes}
}

func LinearSpace(start float64, stop float64, count uint, precision uint) []float64 {
	if count == 0 {
		return []float64{}
	}

	samples := make([]float64, count)
	step := (stop - start) / float64(count-1)

	for i := 0; i < int(count-1); i++ {
		samples[i] = mathutils.Round(start+float64(i)*step, int(precision))
	}
	samples[count-1] = mathutils.Round(stop, int(precision))

	return samples
}

func LinearSpace2(min, max float64, count int, step float64, precision uint) []float64 {
	// step = math.Pow(10, -step)
	result := make([]float64, count)
	for i := range result {
		result[i] = mathutils.Round(min+(max-min)*float64(i+1)/float64(count+1), int(precision))
	}
	return result
}
