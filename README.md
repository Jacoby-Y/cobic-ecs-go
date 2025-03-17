## Cobic (Coby's Epic) ECS in Go and Raylib

To get started, run these commands (or similar commands for your OS)
```bash
cd path/to/my-projects
mkdir my-game
cd my-game
go mod init my-game
go get github.com/Jacoby-Y/cobic-ecs-go
touch main.go
# Write code in main.go
go run main.go
```

```go
package main

import (
	cobic "github.com/Jacoby-Y/cobic-ecs-go"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	cobic.BaseComponent
	Text string
}

func SysDrawText(pos *cobic.Position, text *Text) {
	rl.DrawText(text.Text, int32(pos.X), int32(pos.Y), 20, rl.Black)
	pos.X++
}

func main() {
	rl.SetTraceLogLevel(rl.LogError)

	ctx := cobic.NewContext()
	settings := cobic.NewSettings(800, 600, rl.RayWhite, 60)

	cobic.AddSystems(SysDrawText)

	ctx.AddEntity(
		&cobic.Position{X: 20, Y: 20},
		&Text{Text: "Hello, World!"},
	)

	settings.StartGame(
		func() {
			cobic.RunSystems(&ctx)
		},
	)
}
```