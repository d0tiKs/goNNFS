package data

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

const (
	numPoints = 400
	numSlices = 4
)

type PlotSettings struct {
	Colors  []color.RGBA
	plot    *plot.Plot
	Dataset Dataset

	Title  string
	Xaxis  string
	Yaxis  string
	Width  int
	Height int
}

var i int

func SetupPlot(data Dataset) PlotSettings {
	settings := PlotSettings{
		Colors: []color.RGBA{
			{R: 255, G: 0, B: 0, A: 255},   // Red
			{R: 0, G: 255, B: 0, A: 255},   // Green
			{R: 0, G: 0, B: 255, A: 255},   // Blue
			{R: 255, G: 165, B: 0, A: 255}, // Orange
		},

		Title:   "Categorized Spiral Plot",
		Xaxis:   "X-axis",
		Yaxis:   "Y-axis",
		Width:   600,
		Height:  400,
		plot:    plot.New(),
		Dataset: data,
	}

	settings.plot.Add(plotter.NewGrid())
	i++

	return settings
}

func DrawPlot(settings PlotSettings) {
	scatters := make([]*plotter.Scatter, settings.Dataset.Classes)
	for i := 0; i < int(settings.Dataset.Classes); i++ {
		scatters[i], _ = plotter.NewScatter(plotter.XYZs{})
		scatters[i].Shape = draw.CircleGlyph{}
		scatters[i].Color = settings.Colors[i]
		scatters[i].Radius = vg.Points(3)
		settings.plot.Add(scatters[i])
	}

	for idx, point := range settings.Dataset.Points {
		classID := settings.Dataset.Labels[idx].id
		scatters[classID].XYs = append(scatters[classID].XYs, point)
	}

	err := settings.plot.Save(10*vg.Inch, 10*vg.Inch, fmt.Sprintf("spiral_plot%v.png", i))
	if err != nil {
		panic(err)
	}
}
