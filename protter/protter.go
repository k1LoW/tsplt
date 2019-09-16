package protter

import (
	"image/color"

	"github.com/k1LoW/tsplt/timeseries"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

var colors = []color.Color{
	color.RGBA{0, 115, 146, 255},   // #007392
	color.RGBA{30, 152, 185, 255},  // #1e98b9
	color.RGBA{86, 191, 225, 255},  // #56bfe1
	color.RGBA{131, 231, 255, 255}, // #83e7ff
	color.RGBA{224, 106, 59, 255},  // #e06a3b
	color.RGBA{255, 154, 104, 255}, // #ff9a68
	color.RGBA{255, 204, 150, 255}, // #ffcc96
}

// Plot ...
func Plot(data *timeseries.Data, outPath string) error {
	st := data.Points[0][0].X
	et := data.Points[0][len(data.Points[0])-1].X
	pCols := len(data.Points)
	pRows := len(data.Points[0])

	layout := "2006-01-02\n15:04"
	switch {
	case st.Year() == et.Year() && st.Month() != et.Month():
		layout = "01-02\n15:04"
	case st.Year() == et.Year() && st.Month() == et.Month() && st.Day() != et.Day():
		layout = "02 15:04"
	case st.Format("2006-01-02") == et.Format("2006-01-02") && st.Format("2006-01-02 15:04") != et.Format("2006-01-02 15:04"):
		layout = "15:04:05"
	case st.Format("2006-01-02 15:04") == et.Format("2006-01-02 15:04"):
		layout = "04:05"
	}

	xticks := plot.TimeTicks{Format: layout}

	p, err := plot.New()
	if err != nil {
		return err
	}
	p.X.Tick.Marker = xticks

	for i := 0; i < pCols; i++ {
		pts := make(plotter.XYs, pRows)
		for row := 0; row < pRows; row++ {
			t := data.Points[i][row].X
			v := data.Points[i][row].Y
			pts[row].X = float64(t.Unix())
			pts[row].Y = v
		}
		line, points, err := plotter.NewLinePoints(pts)
		if err != nil {
			return err
		}
		line.Color = colors[i]
		points.Color = colors[i]
		points.Shape = NoneGlyph{}
		p.Add(line, points)
		thum := [2]plot.Thumbnailer{line, points}
		p.Legend.Add(data.YLabels[i], thum[0], thum[1])
	}

	width := 8 * pRows
	if width < 256 {
		width = 256
	}

	return p.Save(vg.Length(width), 256, outPath)
}

type NoneGlyph struct{}

func (NoneGlyph) DrawGlyph(*draw.Canvas, draw.GlyphStyle, vg.Point) {}
