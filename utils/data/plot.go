package data

import (
	errorsutils "NNFS/utils/errors"
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
	Title   string
	Xaxis   string
	Yaxis   string
	Colors  []color.RGBA
	plot    *plot.Plot
	Dataset Dataset
	Width   int
	Height  int
}

func SetupPlot(data Dataset) PlotSettings {
	settings := PlotSettings{
		Colors: []color.RGBA{
			{R: 255, G: 0, B: 0, A: 255},     // Red
			{R: 0, G: 255, B: 0, A: 255},     // Green
			{R: 0, G: 0, B: 255, A: 255},     // Blue
			{R: 255, G: 165, B: 0, A: 255},   // Orange
			{R: 128, G: 0, B: 128, A: 255},   // Dark Purple
			{R: 255, G: 215, B: 0, A: 255},   // Bright Yellow
			{R: 0, G: 128, B: 0, A: 255},     // Forest Green
			{R: 255, G: 0, B: 128, A: 255},   // Pink
			{R: 128, G: 128, B: 128, A: 255}, // Gray
			{R: 0, G: 0, B: 0, A: 255},       // Black
			{R: 255, G: 255, B: 0, A: 255},   // Yellow
			{R: 128, G: 0, B: 0, A: 255},     // Dark Brown
			{R: 0, G: 0, B: 128, A: 255},     // Dark Blue
			{R: 255, G: 0, B: 255, A: 255},   // Magenta
			{R: 0, G: 128, B: 128, A: 255},   // Sea Green
			{R: 255, G: 165, B: 225, A: 255}, // Pastel Pink
			{R: 128, G: 128, B: 192, A: 255}, // Soft Gray
			{R: 0, G: 64, B: 0, A: 255},      // Dark Brown
			{R: 255, G: 240, B: 160, A: 255}, // Pastel Orange
			{R: 128, G: 192, B: 128, A: 255}, // Mint Green
			{R: 0, G: 96, B: 0, A: 255},      // Dark Brown
			{R: 255, G: 235, B: 204, A: 255}, // Light Peach
		},

		Title:   "Categorized Spiral Plot",
		Xaxis:   "X-axis",
		Yaxis:   "Y-axis",
		Width:   600,
		Height:  600,
		plot:    plot.New(),
		Dataset: data,
	}

	settings.plot.Add(plotter.NewGrid())
	return settings
}

func GeneratePlot(settings PlotSettings) {
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
}

func SavePlotImage(settings PlotSettings, filename string) {
	err := settings.plot.Save(10*vg.Inch, 10*vg.Inch, filename)
	if err != nil {
		panic(errorsutils.BuildError(err, "error while saving plot at '%s'", filename))
	}
}
