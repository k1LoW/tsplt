package timeseries

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/araddon/dateparse"
)

type Point struct {
	X time.Time
	Y float64
}

type Data struct {
	XLabel  string
	YLabels []string
	Points  [][]Point
}

// Build plot data
func Build(in io.Reader, delimiter rune) (*Data, error) {
	header := false
	r := csv.NewReader(in)
	r.Comma = delimiter
	rs, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	_, err = strconv.ParseFloat(rs[0][1], 64)
	if err != nil {
		header = true
	}

	dCols := len(rs[0]) - 1
	dRows := len(rs) - 1
	sIdx := 0
	xLabel := "time"
	yLabels := []string{}
	for i := 0; i < dCols; i++ {
		yLabels = append(yLabels, fmt.Sprintf("d%d", i))
	}

	if header {
		sIdx = 1
		dRows = len(rs) - 2
		if rs[0][0] != "" {
			xLabel = rs[0][0]
		}
		yLabels = rs[0][1:]
	}

	points := [][]Point{}

	for i := 0; i < dCols; i++ {
		pts := []Point{}
		for row := 0; row < dRows; row++ {
			pt := Point{}
			t, err := dateparse.ParseAny(rs[sIdx+row][0])
			if err != nil {
				return nil, err
			}
			v, err := strconv.ParseFloat(rs[sIdx+row][i+1], 64)
			if err != nil {
				return nil, err
			}
			pt.X = t
			pt.Y = v
			pts = append(pts, pt)
		}
		points = append(points, pts)
	}

	return &Data{
		XLabel:  xLabel,
		YLabels: yLabels,
		Points:  points,
	}, nil
}
