package main

import (
	"NNFS/utils/data"
)

func main() {
	count := uint(100)
	classes := uint(4)
	noise := 1.0
	data1 := data.GenerateSpiral(count, classes, noise)
	// data2 := data.GenerateSpiral2(count, classes)

	set1 := data.SetupPlot(data1)
	// set2 := data.SetupPlot(data2)

	data.DrawPlot(set1)
	// data.DrawPlot(set2)
}
