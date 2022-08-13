package bsp

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type SDLRenderer struct {
	width    int
	height   int
	window   *sdl.Window
	renderer *sdl.Renderer
	scale    int
}

func NewSDLRenderer(title string, width int, height int, scale int) (SDLRenderer, error) {

	sdlRenderer := SDLRenderer{
		scale:  scale,
		width:  width,
		height: height,
	}

	window, err := sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		int32(width*scale),
		int32(height*scale),
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		return sdlRenderer, err
	}
	sdlRenderer.window = window

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return sdlRenderer, err
	}
	sdlRenderer.renderer = renderer

	return sdlRenderer, nil
}

func (s *SDLRenderer) DrawPixel(c *Color, p *Position) {
	s.renderer.SetDrawColor(c.R, c.G, c.B, 255)
	s.renderer.FillRect(&sdl.Rect{
		X: p.X * int32(s.scale),
		Y: (int32(s.height) - p.Y) * int32(s.scale),
		W: int32(s.scale),
		H: int32(s.scale),
	})
}

func (s *SDLRenderer) Loop(game *Game) {
	running := true
	skips := 0
	clearColor := Color{R: 0, G: 0, B: 0}
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.KeyboardEvent:
				if t.Repeat > 0 {
					break
				}
				switch t.Keysym.Sym {
				case 119:
					game.Inputs.W = (t.State == sdl.PRESSED)
				case 97:
					game.Inputs.A = (t.State == sdl.PRESSED)
				case 115:
					game.Inputs.S = (t.State == sdl.PRESSED)
				case 100:
					game.Inputs.D = (t.State == sdl.PRESSED)
				case 109:
					game.Inputs.M = (t.State == sdl.PRESSED)
				case 44:
					game.Inputs.COMMA = (t.State == sdl.PRESSED)
				case 46:
					game.Inputs.DOT = (t.State == sdl.PRESSED)
				}
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		game.Renderer.Clear(&clearColor)

		now := time.Now().UnixMilli()
		game.Update(game)
		if skips == 0 {
			game.Render(game)
			s.renderer.Present()
		} else {
			skips--
		}

		elapsedTime := time.Now().UnixMilli() - now
		now = elapsedTime
		delay := int64(1000/game.FPS) - elapsedTime

		if delay < 0 {
			game.Update(game)
		}
		for delay < 0 {
			delay += int64(1000 / game.FPS)
			skips++
		}
		sdl.Delay(uint32(delay))
	}

}

func (s *SDLRenderer) Clear(c *Color) {
	s.renderer.SetDrawColor(c.R, c.G, c.B, 255)
	s.renderer.Clear()
}

func (s *SDLRenderer) Destroy() {
	if s.renderer != nil {
		s.renderer.Destroy()
	}
	if s.window != nil {
		s.window.Destroy()
	}
}
