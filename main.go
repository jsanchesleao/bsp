package main

import (
	"bsp/bsp"
)

var colors []bsp.Color = []bsp.Color{
	{R: 0, G: 0, B: 0},
	{R: 255, G: 0, B: 0},
	{R: 0, G: 255, B: 0},
	{R: 0, G: 0, B: 255},
	{R: 255, G: 255, B: 255},
}

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

	game := bsp.NewGame(update, render, 60, &sdlRenderer)

	game.Player.X = 70
	game.Player.Y = -110
	game.Player.Z = 20
	game.Player.Angle = 0
	game.Player.Look = 0

	game.Loop()
}

func clipBehindPlayer(x1 *int32, y1 *int32, z1 *int32, x2 int32, y2 int32, z2 int32) {
	da := float64(*y1)
	db := float64(y2)
	d := da - db
	if d == 0 {
		d = 1
	}
	s := da / d

	*x1 = int32(float64(*x1) + s*float64(x2-(*x1)))
	*y1 = int32(float64(*y1) + s*float64(y2-(*y1)))
	*z1 = int32(float64(*z1) + s*float64(z2-(*z1)))

	if *y1 == 0 {
		*y1 = 1
	}

}

func drawWall(game *bsp.Game, x1 int32, x2 int32, b1 int32, b2 int32, t1 int32, t2 int32) {
	var x, y int32
	yBottomDistance := b2 - b1
	yTopDistance := t2 - t1
	xDistance := x2 - x1
	if xDistance == 0 {
		xDistance = 1
	}
	startX := x1

	var padding int32 = 1

	// Clip X
	if x1 < padding {
		x1 = padding
	}
	if x2 < padding {
		x2 = padding
	}
	if x1 > game.Renderer.GetWidth() {
		x1 = game.Renderer.GetWidth() - padding
	}
	if x2 > game.Renderer.GetWidth() {
		x2 = game.Renderer.GetWidth() - padding
	}

	for x = x1; x < x2; x++ {
		y1 := int32(float64(yBottomDistance)*(float64(x)-float64(startX)+0.5)/float64(xDistance)) + b1
		y2 := int32(float64(yTopDistance)*(float64(x)-float64(startX)+0.5)/float64(xDistance)) + t1
		//Clip Y
		if y1 < padding {
			y1 = padding
		}
		if y2 < padding {
			y2 = padding
		}
		if y1 > game.Renderer.GetHeight() {
			y1 = game.Renderer.GetHeight() - padding
		}
		if y2 > game.Renderer.GetHeight() {
			y2 = game.Renderer.GetHeight() - padding
		}
		for y = y1; y < y2; y++ {
			game.Renderer.DrawPixel(&colors[2], &bsp.Position{X: x, Y: y})
		}
		game.Renderer.DrawPixel(&colors[3], &bsp.Position{X: x, Y: y1})
		game.Renderer.DrawPixel(&colors[3], &bsp.Position{X: x, Y: y2})
	}
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
	wx[0] = int32(float64(x1)*CS - float64(y1)*SN)
	wx[1] = int32(float64(x2)*CS - float64(y2)*SN)
	wx[2] = wx[0]
	wx[3] = wx[1]

	//world y position
	wy[0] = int32(float64(y1)*CS + float64(x1)*SN)
	wy[1] = int32(float64(y2)*CS + float64(x2)*SN)
	wy[2] = wy[0]
	wy[3] = wy[1]

	//world z height
	wz[0] = 0 - game.Player.Z + ((int32(game.Player.Look) * wy[0]) / 32.0)
	wz[1] = 0 - game.Player.Z + ((int32(game.Player.Look) * wy[1]) / 32.0)
	wz[2] = wz[0] + 40
	wz[3] = wz[1] + 40

	// wont draw if behind player
	if wy[0] < 1 && wy[1] < 1 {
		return
	}

	// point 1 is behind the player, so we clip
	if wy[0] < 1 {
		clipBehindPlayer(&wx[0], &wy[0], &wz[0], wx[1], wy[1], wz[1]) // clipping bottom line
		clipBehindPlayer(&wx[2], &wy[2], &wz[2], wx[3], wy[3], wz[3]) // clipping top line
	}
	// point 2 is behind the player, we also clip
	if wy[1] < 1 {
		clipBehindPlayer(&wx[1], &wy[1], &wz[1], wx[0], wy[0], wz[0]) // clipping bottom line
		clipBehindPlayer(&wx[3], &wy[3], &wz[3], wx[2], wy[2], wz[2]) // clipping top line
	}

	//screen positions
	wx[0] = wx[0]*200/wy[0] + HalfWidth
	wy[0] = wz[0]*200/wy[0] + HalfHeight
	wx[1] = wx[1]*200/wy[1] + HalfWidth
	wy[1] = wz[1]*200/wy[1] + HalfHeight
	wx[2] = wx[2]*200/wy[2] + HalfWidth
	wy[2] = wz[2]*200/wy[2] + HalfHeight
	wx[3] = wx[3]*200/wy[3] + HalfWidth
	wy[3] = wz[3]*200/wy[3] + HalfHeight

	//draw points
	drawWall(game, wx[0], wx[1], wy[0], wy[1], wy[2], wy[3])

}
