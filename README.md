# ebiten_path

![docs/screenshot.png](docs/screenshot.png)

```go
p := path.NewPath()
p.Arc(200, 200, 80, 0, 270, 40)
p.LineTo(400, 100)
p.CurveTo(500, 300, 20)
p.BezierTo(300, 400, 500, 400, 600, 200, 30)
p.Close()

percent := float32(0.4)

// get point
pt := g.p.GetPointAtPercent(percent)
// get degrees
deg := g.p.GetDegreesAtPercent(percent)

/// ...

func (g *Game) Draw(screen *ebiten.Image) {
    // draw path
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

Except `Draw()`, almost dependent from Ebitengine.