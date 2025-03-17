package cobic

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameSettings struct {
	ScreenWidth  int32
	ScreenHeight int32
	Background   rl.Color
	TargetFPS    int32
}

func NewSettings(ScreenWidth int32, ScreenHeight int32, Background rl.Color, TargetFPS int32) GameSettings {
	return GameSettings{ScreenWidth, ScreenWidth, Background, TargetFPS}
}

func (settings GameSettings) StartGame(funcs ...func()) {
	rl.InitWindow(settings.ScreenWidth, settings.ScreenHeight, "Raylib-go Example")
	defer rl.CloseWindow()

	rl.SetTargetFPS(settings.TargetFPS)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(settings.Background)
		for _, fn := range funcs {
			fn()
		}
		rl.EndDrawing()
	}
}
