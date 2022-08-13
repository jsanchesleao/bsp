package main

import (
	"bsp/bsp"
	"fmt"
)

const fps = 60

func main() {

	sdlRenderer, err := bsp.NewSDLRenderer("BSP Test", 300, 200, 4)
	if err != nil {
		panic(err)
	}
	defer sdlRenderer.Destroy()

	bsp.Init()

	game := bsp.NewGame(Update, Render, 60, &sdlRenderer)

	game.Loop()
}

var x, y = 0, 0

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
			x--
			fmt.Println("left")
		}
		if game.Inputs.D {
			x++
			fmt.Println("right")
		}
		if game.Inputs.W {
			y++
			fmt.Println("up")
		}
		if game.Inputs.S {
			y--
			fmt.Println("down")
		}
	}
	if game.Inputs.COMMA {
		fmt.Println("strafe left")
	}
	if game.Inputs.DOT {
		fmt.Println("strafe right")
	}

	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}

}

func Render(game *bsp.Game) {
	game.Renderer.DrawPixel(&bsp.Color{R: 200, G: 0, B: 0}, &bsp.Position{X: int32(x), Y: int32(y)})
}
