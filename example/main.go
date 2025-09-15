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
	g.p.Draw(screen, color.Black, 3)
	for i := 0; i <= 9; i++ {
		percent := float32(i) / 10
		pt := g.p.GetPointAtPercent(percent)
		vector.DrawFilledRect(screen, pt.X-3, pt.Y-3, 6, 6, color.RGBA{255, 0, 0, 255}, true)
		ebitenutil.DebugPrintAt(screen,
			fmt.Sprintf("%.1f", percent),
			int(pt.X)+8, int(pt.Y)-8)

		deg := g.p.GetDegreesAtPercent(percent)
		ebitenutil.DebugPrintAt(screen,
			fmt.Sprintf("%.1f°", deg),
			int(pt.X)+8, int(pt.Y)+8)

		rad := g.p.GetRadiansAtPercent(percent)
		dx := float32(30) * math.Cos(rad)
		dy := float32(30) * math.Sin(rad)
		vector.StrokeLine(screen, pt.X, pt.Y, pt.X+dx, pt.Y+dy, 2, color.RGBA{0, 0, 255, 255}, true)
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
