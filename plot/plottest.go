package main

import "NNFS/utils/data"

func main() {
	count := uint(200)
	classes := uint(11)
	noise := 0.0
	data1 := data.GenerateSpiral(count, classes, noise)

	set1 := data.SetupPlot(data1)

	data.GeneratePlot(set1)
	data.SavePlotImage(set1, "spiral1.png")
}
