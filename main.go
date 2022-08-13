package main

import (
	"bsp/bsp"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const fps = 60

func main() {

	window, err := sdl.CreateWindow("bsp", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	game := bsp.Game{
		Inputs:      bsp.Inputs{},
		Update:      Update,
		Render:      Render,
		HandleEvent: HandleEvent,
	}

	bsp.GameLoop(renderer, &game, 60)

}

func HandleEvent(game *bsp.Game, event sdl.Event) {
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
	}
}

func Update(game *bsp.Game) {
	if game.Inputs.M {
		if game.Inputs.A {
			fmt.Println("look down")
		}
		if game.Inputs.D {
			fmt.Println("loop up")
		}
		if game.Inputs.W {
			fmt.Println("move up")
		}
		if game.Inputs.S {
			fmt.Println("move down")
		}
	} else {
		if game.Inputs.A {
			fmt.Println("left")
		}
		if game.Inputs.D {
			fmt.Println("right")
		}
		if game.Inputs.W {
			fmt.Println("up")
		}
		if game.Inputs.S {
			fmt.Println("down")
		}
	}
	if game.Inputs.COMMA {
		fmt.Println("strafe left")
	}
	if game.Inputs.DOT {
		fmt.Println("strafe right")
	}
}

func Render(game *bsp.Game, r *sdl.Renderer) {
	r.SetDrawColor(255, 0, 0, 255)
	r.FillRect(&sdl.Rect{X: 10, Y: 10, W: 50, H: 50})
}
