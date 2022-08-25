package bsp

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var Sin [360]float64
var Cos [360]float64

type Game struct {
	Inputs   Inputs
	Update   UpdateFunc
	Render   RenderFunc
	FPS      int
	Renderer Renderer
	Player   Player
	Colors   []*Color
	Walls    []*Wall
	Sectors  []*Sector
}

type Player struct {
	X, Y, Z int32
	Angle   int
	Look    int
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
	GetWidth() int32
	GetHeight() int32
}

type Color struct {
	R, G, B uint8
}

type Position struct {
	X, Y int32
}

type Wall struct {
	X1, Y1, X2, Y2 int32
	Color          int
}

type Sector struct {
	WallStart, WallEnd int
	Z1, Z2             int32
	CenterX, CenterY   int32
	Distance           float64
}

func NewGame(update UpdateFunc, render RenderFunc, fps int, renderer Renderer) *Game {
	return &Game{
		Inputs:   Inputs{},
		Update:   update,
		Render:   render,
		FPS:      fps,
		Renderer: renderer,
		Player:   Player{},
		Colors:   []*Color{},
		Walls:    []*Wall{},
		Sectors:  []*Sector{},
	}
}

func (game *Game) LoadColors(colors []Color) {
	game.Colors = []*Color{}
	for _, color := range colors {
		c := color
		game.Colors = append(game.Colors, &c)
	}
}

func (game *Game) LoadWalls(walls []Wall) {
	game.Walls = []*Wall{}
	for _, wall := range walls {
		w := wall
		game.Walls = append(game.Walls, &w)
	}
}

func (game *Game) LoadSectors(sectors []Sector) {
	game.Sectors = []*Sector{}
	for _, sector := range sectors {
		s := sector
		game.Sectors = append(game.Sectors, &s)
	}
}

func (game *Game) Loop() {
	game.Renderer.Loop(game)
}

func Init() {
	for x := 0; x < 360; x++ {
		rads := float64(x) * math.Pi / 180
		Sin[x] = math.Sin(rads)
		Cos[x] = math.Cos(rads)
	}
}
