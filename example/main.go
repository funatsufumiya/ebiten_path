package main

import (
	"fmt"
	"image/color"
	"log"

	math "github.com/chewxy/math32"

	"github.com/funatsufumiya/ebiten_path/path"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	width  = 640
	height = 480
)

type Game struct {
	p *path.Path
}

func NewGame() *Game {
	p := path.NewPath()
	p.Arc(200, 200, 80, 0, 270, 40)
	p.LineTo(400, 100)
	p.CurveTo(500, 300, 20)
	p.BezierTo(300, 400, 500, 400, 600, 200, 30)
	p.Close()
	return &Game{p: p}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	bg := color.RGBA{160, 160, 160, 255}
	screen.Fill(bg)
	opts := &ebiten.DrawImageOptions{}

	// g.p.DrawFilled(screen, color.RGBA{200, 220, 255, 128}, opts)
	g.p.DrawStroke(screen, color.Black, 3, opts)

	for i := 0; i <= 9; i++ {
		percent := float32(i) / 10

		pt := g.p.GetPointAtPercent(percent)
		x, y := opts.GeoM.Apply(float64(pt.X), float64(pt.Y))

		vector.DrawFilledRect(screen, float32(x)-3, float32(y)-3, 6, 6, color.RGBA{255, 0, 0, 255}, true)
		ebitenutil.DebugPrintAt(screen,
			fmt.Sprintf("%.1f", percent),
			int(x)+8, int(y)-8)

		deg := g.p.GetDegreesAtPercent(percent)
		ebitenutil.DebugPrintAt(screen,
			fmt.Sprintf("%.1f°", deg),
			int(x)+8, int(y)+8)

		rad := g.p.GetRadiansAtPercent(percent)
		dx := float32(30) * math.Cos(rad)
		dy := float32(30) * math.Sin(rad)
		vector.StrokeLine(screen, float32(x), float32(y), float32(x)+dx, float32(y)+dy, 2, color.RGBA{0, 0, 255, 255}, true)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return width, height
}

func main() {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Path Example")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
