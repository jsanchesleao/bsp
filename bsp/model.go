package bsp

import "github.com/veandco/go-sdl2/sdl"

type Game struct {
	Scale       int
	Inputs      Inputs
	Update      UpdateFunc
	Render      RenderFunc
	HandleEvent HandleEventFunc
	FPS         int
	Renderer    Renderer
}

type UpdateFunc func(*Game)
type RenderFunc func(*Game)
type HandleEventFunc func(*Game, sdl.Event)

type Inputs struct {
	W, A, S, D, M, COMMA, DOT bool
}

type Renderer interface {
	Clear(*Color)
	DrawPixel(*Color, *Position)
	Loop(*Game)
}

type Color struct {
	R, G, B uint8
}

type Position struct {
	X, Y int32
}

func (game *Game) Loop() {
	game.Renderer.Loop(game)
}
