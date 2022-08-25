package main

import (
	"bsp/bsp"
	"fmt"
	"math"
)

func main() {

	renderer, err := bsp.NewPixelRenderer("BSP Test", 320, 240, 5, 60)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	bsp.Init()

	update := func(game *bsp.Game) {
		game.UpdatePlayerPosition()
	}

	game := bsp.NewGame(update, render, 60, &renderer)

	game.Player.X = 70
	game.Player.Y = -110
	game.Player.Z = 20
	game.Player.Angle = 0
	game.Player.Look = 0

	game.LoadColors([]bsp.Color{
		{R: 200, G: 0, B: 0},
		{R: 180, G: 0, B: 0},
		{R: 0, G: 200, B: 0},
		{R: 0, G: 180, B: 0},
		{R: 0, G: 0, B: 200},
		{R: 0, G: 0, B: 180},
		{R: 150, G: 0, B: 150},
		{R: 120, G: 0, B: 120},
	})

	game.LoadWalls([]bsp.Wall{
		{X1: 0, Y1: 0, X2: 32, Y2: 0, Color: 0},
		{X1: 32, Y1: 0, X2: 32, Y2: 32, Color: 1},
		{X1: 32, Y1: 32, X2: 0, Y2: 32, Color: 0},
		{X1: 0, Y1: 32, X2: 0, Y2: 0, Color: 1},

		{X1: 64, Y1: 0, X2: 96, Y2: 0, Color: 2},
		{X1: 96, Y1: 0, X2: 96, Y2: 32, Color: 3},
		{X1: 96, Y1: 32, X2: 64, Y2: 32, Color: 2},
		{X1: 64, Y1: 32, X2: 64, Y2: 0, Color: 3},

		{X1: 64, Y1: 64, X2: 96, Y2: 64, Color: 4},
		{X1: 96, Y1: 64, X2: 96, Y2: 96, Color: 5},
		{X1: 96, Y1: 96, X2: 64, Y2: 96, Color: 4},
		{X1: 64, Y1: 96, X2: 64, Y2: 64, Color: 5},

		{X1: 0, Y1: 64, X2: 32, Y2: 64, Color: 6},
		{X1: 32, Y1: 64, X2: 32, Y2: 96, Color: 7},
		{X1: 32, Y1: 96, X2: 0, Y2: 96, Color: 6},
		{X1: 0, Y1: 96, X2: 0, Y2: 64, Color: 7},
	})

	game.LoadSectors([]bsp.Sector{
		{WallStart: 0, WallEnd: 4, Z1: 0, Z2: 40},
		{WallStart: 4, WallEnd: 8, Z1: 0, Z2: 40},
		{WallStart: 8, WallEnd: 12, Z1: 0, Z2: 40},
		{WallStart: 12, WallEnd: 16, Z1: 0, Z2: 40},
	})

	fmt.Printf("%+v\n", game)
	game.Loop()
}

func dist(x1 int32, y1 int32, x2 int32, y2 int32) float64 {
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
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

func drawWall(game *bsp.Game, x1 int32, x2 int32, b1 int32, b2 int32, t1 int32, t2 int32, color *bsp.Color) {
	var x, y int32
	yBottomDistance := b2 - b1
	yTopDistance := t2 - t1
	xDistance := x2 - x1
	if xDistance == 0 {
		xDistance = 1
	}
	startX := x1

	var padding int32 = 5

	// Clip X
	if x1 < padding {
		x1 = padding
	}
	if x2 < padding {
		x2 = padding
	}
	if x1+padding > game.Renderer.GetWidth() {
		x1 = game.Renderer.GetWidth() - padding
	}
	if x2+padding > game.Renderer.GetWidth() {
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
		if y1+padding > game.Renderer.GetHeight() {
			y1 = game.Renderer.GetHeight() - padding
		}
		if y2+padding > game.Renderer.GetHeight() {
			y2 = game.Renderer.GetHeight() - padding
		}
		for y = y1; y < y2; y++ {
			game.Renderer.DrawPixel(color, &bsp.Position{X: x, Y: y})
		}
	}
}

func render(game *bsp.Game) {
	HalfWidth := game.Renderer.GetWidth() / 2
	HalfHeight := game.Renderer.GetHeight() / 2
	CS := bsp.Cos[game.Player.Angle]
	SN := bsp.Sin[game.Player.Angle]

	//order sectors
	for i := 0; i < len(game.Sectors)-1; i++ {
		for j := 0; j < len(game.Sectors)-i-1; j++ {
			if game.Sectors[j].Distance < game.Sectors[j+1].Distance {
				st := game.Sectors[j]
				game.Sectors[j] = game.Sectors[j+1]
				game.Sectors[j+1] = st
			}
		}
	}

	//draw sectors
	for _, sector := range game.Sectors {
		sector.Distance = 0
		for i := sector.WallStart; i < sector.WallEnd; i++ {
			var wx, wy, wz [4]int32
			x1 := game.Walls[i].X1 - game.Player.X
			y1 := game.Walls[i].Y1 - game.Player.Y
			x2 := game.Walls[i].X2 - game.Player.X
			y2 := game.Walls[i].Y2 - game.Player.Y

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

			sector.Distance += dist(0, 0, (wx[0]+wx[1])/2, (wy[0] + wy[1]/2))

			//world z height
			wz[0] = sector.Z1 - game.Player.Z + ((int32(game.Player.Look) * wy[0]) / 32.0)
			wz[1] = sector.Z1 - game.Player.Z + ((int32(game.Player.Look) * wy[1]) / 32.0)
			wz[2] = wz[0] + (sector.Z2 - sector.Z1)
			wz[3] = wz[1] + (sector.Z2 - sector.Z1)

			// wont draw if behind player
			if wy[0] < 1 && wy[1] < 1 {
				continue
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
			drawWall(game, wx[0], wx[1], wy[0], wy[1], wy[2], wy[3], game.Colors[game.Walls[i].Color])
		}
		sector.Distance = sector.Distance / (float64(sector.WallEnd - sector.WallStart))
	}

}
