package bsp

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func GameLoop(renderer *sdl.Renderer, game *Game, fps int) {
	running := true
	skips := 0
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			game.HandleEvent(game, event)
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		now := time.Now().UnixMilli()
		game.Update(game)
		if skips == 0 {
			game.Render(game, renderer)
			renderer.Present()
		} else {
			skips--
		}

		elapsedTime := time.Now().UnixMilli() - now
		now = elapsedTime
		delay := int64(1000/fps) - elapsedTime

		if delay < 0 {
			game.Update(game)
		}
		for delay < 0 {
			delay += int64(1000 / fps)
			skips++
		}
		sdl.Delay(uint32(delay))
	}
}
