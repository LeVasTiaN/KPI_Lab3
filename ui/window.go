package ui

import (
	"image"
	"image/color"
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/imageutil"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/image/draw"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
)

type Visualizer struct {
	Title         string
	Debug         bool
	OnScreenReady func(s screen.Screen)

	w    screen.Window
	tx   chan screen.Texture
	done chan struct{}

	sz  size.Event
	pos image.Rectangle

	FigurePos image.Rectangle
}

func (pw *Visualizer) Main() {
	pw.tx = make(chan screen.Texture)
	pw.done = make(chan struct{})
	pw.pos.Max.X = 200
	pw.pos.Max.Y = 200

	pw.FigurePos = image.Rect(300, 300, 550, 500)

	driver.Main(pw.run)
}

func (pw *Visualizer) drawDefaultUI() {
	pw.w.Fill(pw.sz.Bounds(), color.White, draw.Src)

	if pw.FigurePos.Empty() {
		centerX := 800 / 2
		centerY := 800 / 2
		pw.FigurePos = image.Rect(
			centerX-125, centerY-100,
			centerX+125, centerY+100,
		)
	}

	figureColor := color.RGBA{B: 255, A: 255}

	verticalBar := image.Rect(
		pw.FigurePos.Min.X, pw.FigurePos.Min.Y,
		pw.FigurePos.Min.X+50, pw.FigurePos.Max.Y,
	)

	horizontalBar := image.Rect(
		pw.FigurePos.Min.X, pw.FigurePos.Min.Y+75,
		pw.FigurePos.Max.X, pw.FigurePos.Min.Y+125,
	)

	pw.w.Fill(verticalBar, figureColor, draw.Src)
	pw.w.Fill(horizontalBar, figureColor, draw.Src)

	for _, br := range imageutil.Border(pw.sz.Bounds(), 10) {
		pw.w.Fill(br, color.White, draw.Src)
	}
}

func (pw *Visualizer) handleEvent(e any, t screen.Texture) {
	switch e := e.(type) {

	case size.Event:
		pw.sz = e

	case error:
		log.Printf("ERROR: %s", e)

	case mouse.Event:
		if t == nil {
			// TODO: Реалізувати реакцію на натискання кнопки миші.
		}

	case paint.Event:
		if t == nil {
			pw.drawDefaultUI()
		} else {
			pw.w.Scale(pw.sz.Bounds(), t, t.Bounds(), draw.Src, nil)
		}
		pw.w.Publish()
	}
}

func (pw *Visualizer) drawDefaultUI() {
	pw.w.Fill(pw.sz.Bounds(), color.Black, draw.Src) // Фон.

	// TODO: Змінити колір фону та додати відображення фігури у вашому варіанті.

	// Малювання білої рамки.
	for _, br := range imageutil.Border(pw.sz.Bounds(), 10) {
		pw.w.Fill(br, color.White, draw.Src)
	}
}
