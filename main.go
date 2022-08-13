package main

import (
	"bsp/bsp"
	"fmt"
)

func main() {

	sdlRenderer, err := bsp.NewSDLRenderer("BSP Test", 300, 200, 4)
	if err != nil {
		panic(err)
	}
	defer sdlRenderer.Destroy()

	bsp.Init()

	update := func(game *bsp.Game) {
		game.UpdatePlayerPosition()
	}

	game := bsp.NewGame(update, render, 30, &sdlRenderer)

	game.Player.X = 70
	game.Player.Y = 110
	game.Player.Z = 20
	game.Player.Angle = 0
	game.Player.Look = 0

	game.Loop()
}

func render(game *bsp.Game) {
	var wx, wy, wz [4]int32
	HalfWidth := game.Renderer.GetWidth() / 2
	HalfHeight := game.Renderer.GetHeight() / 2
	CS := bsp.Cos[game.Player.Angle]
	SN := bsp.Sin[game.Player.Angle]

	x1 := 40 - game.Player.X
	y1 := 10 - game.Player.Y
	x2 := 40 - game.Player.X
	y2 := 290 - game.Player.Y

	//world x position
	wx[0] = int32(float64(x1)*CS) - int32(float64(y1)*SN)
	wx[1] = int32(float64(x2)*CS) - int32(float64(y2)*SN)

	//world y position
	wy[0] = int32(float64(y1)*CS) - int32(float64(x1)*SN)
	wy[1] = int32(float64(y2)*CS) - int32(float64(x2)*SN)

	//world height
	wz[0] = 0 - game.Player.Z + ((int32(game.Player.Look) * wy[0]) / 32)
	wz[1] = 0 - game.Player.Z + ((int32(game.Player.Look) * wy[1]) / 32)

	//screen positions
	wx[0] = wx[0]*200/wy[0] + HalfWidth
	wy[0] = wz[0]*200/wy[0] + HalfHeight
	wx[1] = wx[1]*200/wy[1] + HalfWidth
	wy[1] = wz[1]*200/wy[1] + HalfHeight

	//draw points
	color := bsp.Color{R: 255, G: 0, B: 0}
	if wx[0] > 0 && wx[0] < game.Renderer.GetWidth() && wy[0] > 0 && wy[0] < game.Renderer.GetHeight() {
		game.Renderer.DrawPixel(&color, &bsp.Position{X: wx[0], Y: wy[0]})
	}
	if wx[1] > 0 && wx[1] < game.Renderer.GetWidth() && wy[1] > 0 && wy[1] < game.Renderer.GetHeight() {
		game.Renderer.DrawPixel(&color, &bsp.Position{X: wx[1], Y: wy[1]})
	}

}
