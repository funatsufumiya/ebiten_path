package path

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Point struct {
	X, Y float32
}

type Path struct {
	points []Point
	closed bool
}

func NewPath() *Path {
	return &Path{points: []Point{}}
}

func (p *Path) AddVertex(x, y float32) {
	p.points = append(p.points, Point{X: x, Y: y})
}

func (p *Path) LineTo(x, y float32) {
	p.AddVertex(x, y)
}

func (p *Path) Close() {
	p.closed = true
}

func (p *Path) segmentLengths() []float32 {
	var lengths []float32
	for i := 0; i < len(p.points)-1; i++ {
		l := float32(math.Hypot(
			float64(p.points[i+1].X-p.points[i].X),
			float64(p.points[i+1].Y-p.points[i].Y),
		))
		lengths = append(lengths, l)
	}
	if p.closed && len(p.points) > 1 {
		l := float32(math.Hypot(
			float64(p.points[0].X-p.points[len(p.points)-1].X),
			float64(p.points[0].Y-p.points[len(p.points)-1].Y),
		))
		lengths = append(lengths, l)
	}
	return lengths
}

func (p *Path) TotalLength() float32 {
	lengths := p.segmentLengths()
	total := float32(0)
	for _, l := range lengths {
		total += l
	}
	return total
}

func (p *Path) GetPointAtPercent(percent float32) Point {
	if len(p.points) == 0 {
		return Point{0, 0}
	}
	lengths := p.segmentLengths()
	total := float32(0)
	for _, l := range lengths {
		total += l
	}
	target := percent * total
	accum := float32(0)
	for i := 0; i < len(lengths); i++ {
		if accum+lengths[i] >= target {
			t := (target - accum) / lengths[i]
			p0 := p.points[i]
			var p1 Point
			if i+1 < len(p.points) {
				p1 = p.points[i+1]
			} else {
				p1 = p.points[0] // if closed
			}
			return Point{
				X: p0.X + (p1.X-p0.X)*t,
				Y: p0.Y + (p1.Y-p0.Y)*t,
			}
		}
		accum += lengths[i]
	}
	return p.points[len(p.points)-1]
}

func (p *Path) GetPointAtLength(length float32) Point {
	if len(p.points) == 0 {
		return Point{0, 0}
	}
	lengths := p.segmentLengths()
	accum := float32(0)
	for i := 0; i < len(lengths); i++ {
		if accum+lengths[i] >= length {
			t := (length - accum) / lengths[i]
			p0 := p.points[i]
			var p1 Point
			if i+1 < len(p.points) {
				p1 = p.points[i+1]
			} else {
				p1 = p.points[0]
			}
			return Point{
				X: p0.X + (p1.X-p0.X)*t,
				Y: p0.Y + (p1.Y-p0.Y)*t,
			}
		}
		accum += lengths[i]
	}
	return p.points[len(p.points)-1]
}

func (p *Path) Arc(cx, cy, radius, startAngle, endAngle float32, segments int) {
	if segments < 2 {
		segments = 20
	}
	angleStep := (endAngle - startAngle) / float32(segments)
	for i := 0; i <= segments; i++ {
		angle := startAngle + angleStep*float32(i)
		x := cx + radius*float32(math.Cos(float64(angle)*math.Pi/180))
		y := cy + radius*float32(math.Sin(float64(angle)*math.Pi/180))
		p.AddVertex(x, y)
	}
}

func (p *Path) CurveTo(x, y float32, segments int) {
	if len(p.points) < 2 {
		p.AddVertex(x, y)
		return
	}
	if segments < 2 {
		segments = 20
	}
	p0 := p.points[len(p.points)-2]
	p1 := p.points[len(p.points)-1]
	p2 := Point{X: x, Y: y}
	for i := 1; i <= segments; i++ {
		t := float32(i) / float32(segments)
		pt := catmullRom(p0, p1, p2, t)
		p.AddVertex(pt.X, pt.Y)
	}
}

func (p *Path) BezierTo(cx1, cy1, cx2, cy2, x, y float32, segments int) {
	if len(p.points) == 0 {
		p.AddVertex(x, y)
		return
	}
	if segments < 2 {
		segments = 20
	}
	p0 := p.points[len(p.points)-1]
	p1 := Point{X: cx1, Y: cy1}
	p2 := Point{X: cx2, Y: cy2}
	p3 := Point{X: x, Y: y}
	for i := 1; i <= segments; i++ {
		t := float32(i) / float32(segments)
		pt := cubicBezier(p0, p1, p2, p3, t)
		p.AddVertex(pt.X, pt.Y)
	}
}

func catmullRom(p0, p1, p2 Point, t float32) Point {
	return Point{
		X: (1-t)*p1.X + t*p2.X,
		Y: (1-t)*p1.Y + t*p2.Y,
	}
}

func cubicBezier(p0, p1, p2, p3 Point, t float32) Point {
	u := 1 - t
	return Point{
		X: u*u*u*p0.X + 3*u*u*t*p1.X + 3*u*t*t*p2.X + t*t*t*p3.X,
		Y: u*u*u*p0.Y + 3*u*u*t*p1.Y + 3*u*t*t*p2.Y + t*t*t*p3.Y,
	}
}

func (p *Path) Draw(dst *ebiten.Image, clr color.Color, strokeWidth float32) {
	for i := 0; i < len(p.points)-1; i++ {
		vector.StrokeLine(dst, p.points[i].X, p.points[i].Y, p.points[i+1].X, p.points[i+1].Y, strokeWidth, clr, true)
	}
	if p.closed && len(p.points) > 1 {
		vector.StrokeLine(dst, p.points[len(p.points)-1].X, p.points[len(p.points)-1].Y, p.points[0].X, p.points[0].Y, strokeWidth, clr, true)
	}
}
