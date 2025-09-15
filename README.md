# ebiten_path

![docs/screenshot.png](docs/screenshot.png)

```go
import "github.com/funatsufumiya/ebiten_path/path"

p := path.NewPath()
p.Arc(200, 200, 80, 0, 270, 40)
p.LineTo(400, 100)
p.CurveTo(500, 300, 20)
p.BezierTo(300, 400, 500, 400, 600, 200, 30)
p.Close()

percent := float32(0.4)

// get point
pt := p.GetPointAtPercent(percent)

// get degrees
deg := p.GetDegreesAtPercent(percent)

func (g *Game) Draw(screen *ebiten.Image) {
    opts := &ebiten.DrawImageOptions{}

	// draw filled path
	// p.DrawFilled(screen, color.RGBA{200,220,255,128}, opts)

	// draw stroke path
	p.DrawStroke(screen, color.Black, 3, opts)
	g.p.Draw(screen, color.Black, 3)
}
```

Ebitengine port from [ofPolyline](https://openframeworks.cc/documentation/graphics/ofPolyline/) (openFrameworks)

Useful for path animation using `GetPointAtPercent()` / `GetDegreesAtPercent()` / `GetRadiansAtPercent()`

> [!WARNING]
> Go port was almost done by GitHub Copilot. Use with care.

## Example

```bash
$ go run ./example/main.go
```

## Note

Except `DrawFilled()` / `DrawStroke()`, almost dependent from Ebitengine.