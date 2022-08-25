package bsp

var movSpeed = 3.0
var angleSpeed = 2
var lookSpeed = 1
var zSpeed int32 = 4

func (game *Game) UpdatePlayerPosition() {
	if game.Inputs.M {
		if game.Inputs.A {
			game.Player.Look -= lookSpeed
		}
		if game.Inputs.D {
			game.Player.Look += lookSpeed
		}
		if game.Inputs.W {
			game.Player.Z += zSpeed
		}
		if game.Inputs.S {
			game.Player.Z -= zSpeed
		}
	} else {
		if game.Inputs.A {
			game.Player.Angle -= angleSpeed
			if game.Player.Angle < 0 {
				game.Player.Angle += 360
			}
		}
		if game.Inputs.D {
			game.Player.Angle += angleSpeed
			if game.Player.Angle >= 360 {
				game.Player.Angle -= 360
			}
		}
		dx := Sin[game.Player.Angle] * movSpeed
		dy := Cos[game.Player.Angle] * movSpeed
		if game.Inputs.W {
			game.Player.X += int32(dx)
			game.Player.Y += int32(dy)
		}
		if game.Inputs.S {
			game.Player.X -= int32(dx)
			game.Player.Y -= int32(dy)
		}
	}
	dx := Sin[game.Player.Angle] * movSpeed
	dy := Cos[game.Player.Angle] * movSpeed
	if game.Inputs.COMMA {
		game.Player.X -= int32(dy)
		game.Player.Y += int32(dx)
	}
	if game.Inputs.DOT {
		game.Player.X += int32(dy)
		game.Player.Y -= int32(dx)
	}
}
