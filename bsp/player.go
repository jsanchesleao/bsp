package bsp

func (game *Game) UpdatePlayerPosition() {
	if game.Inputs.M {
		if game.Inputs.A {
			game.Player.Look -= 1
		}
		if game.Inputs.D {
			game.Player.Look += 1
		}
		if game.Inputs.W {
			game.Player.Z += 4
		}
		if game.Inputs.S {
			game.Player.Z -= 4
		}
	} else {
		if game.Inputs.A {
			game.Player.Angle -= 4
			if game.Player.Angle < 0 {
				game.Player.Angle += 360
			}
		}
		if game.Inputs.D {
			game.Player.Angle += 4
			if game.Player.Angle >= 360 {
				game.Player.Angle -= 360
			}
		}
		dx := Sin[game.Player.Angle] * 10
		dy := Cos[game.Player.Angle] * 10
		if game.Inputs.W {
			game.Player.X += int32(dx)
			game.Player.Y += int32(dy)
		}
		if game.Inputs.S {
			game.Player.X -= int32(dx)
			game.Player.Y -= int32(dy)
		}
	}
	dx := Sin[game.Player.Angle] * 10
	dy := Cos[game.Player.Angle] * 10
	if game.Inputs.COMMA {
		game.Player.X -= int32(dy)
		game.Player.Y += int32(dx)
	}
	if game.Inputs.DOT {
		game.Player.X += int32(dy)
		game.Player.Y -= int32(dx)
	}
}
