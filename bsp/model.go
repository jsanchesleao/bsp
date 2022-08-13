package bsp

import "github.com/veandco/go-sdl2/sdl"

type Game struct {
	Inputs      Inputs
	Update      UpdateFunc
	Render      RenderFunc
	HandleEvent HandleEventFunc
}

type UpdateFunc func(*Game)
type RenderFunc func(*Game, *sdl.Renderer)
type HandleEventFunc func(*Game, sdl.Event)

type Inputs struct {
	W, A, S, D, M, COMMA, DOT bool
}
