package bsp

import (
	"github.com/jsanchesleao/pixel"
)

type PixelRenderer struct {
	engine *pixel.Engine
}

func NewPixelRenderer(title string, width int, height int, scale int, fps int) (PixelRenderer, error) {
	renderer := PixelRenderer{}
	engine, err := pixel.NewEngine(title, int32(width), int32(height), scale, fps)
	if err != nil {
		return renderer, err
	}

	renderer.engine = engine

	return renderer, nil
}

func (r *PixelRenderer) Destroy() {
	r.engine.Destroy()
}

func (r *PixelRenderer) Clear(color *Color) {
	for x := 0; x < int(r.engine.Width); x++ {
		for y := 0; y < int(r.engine.Height); y++ {
			r.engine.Draw(x, y, color.R, color.G, color.B)
		}
	}
}

func (r *PixelRenderer) DrawPixel(color *Color, position *Position) {
	r.engine.Draw(int(position.X), int(position.Y), color.R, color.G, color.B)
}

func (r *PixelRenderer) Loop(game *Game) {
	update := func(e *pixel.Engine) {
		game.Inputs.W = r.engine.Inputs.W
		game.Inputs.A = r.engine.Inputs.A
		game.Inputs.S = r.engine.Inputs.S
		game.Inputs.D = r.engine.Inputs.D
		game.Inputs.M = r.engine.Inputs.M
		game.Inputs.COMMA = r.engine.Inputs.Comma
		game.Inputs.DOT = r.engine.Inputs.Dot

		game.Update(game)
	}
	render := func(e *pixel.Engine) {
		r.Clear(&Color{0, 0, 0})
		game.Render(game)
	}
	r.engine.Loop(update, render)
}

func (r *PixelRenderer) GetWidth() int32 {
	return r.engine.Width
}

func (r *PixelRenderer) GetHeight() int32 {
	return r.engine.Height
}
