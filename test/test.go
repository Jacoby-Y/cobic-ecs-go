package testing

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
		&Text{Text: "Hello from my ECS"},
	)

	settings.StartGame(
		func() {
			cobic.RunSystems(&ctx)
		},
	)
}

func testSystems() {
	// ctx := ent.NewContext()

	// ctx.Add(
	// 	&comp.Position{X: 0, Y: 0},
	// 	&comp.Velocity{X: 1, Y: 1},
	// )
	// ctx.Add(
	// 	&comp.Position{X: 0, Y: 0},
	// 	&comp.Velocity{X: 5, Y: 5},
	// )
	// ctx.Add(
	// 	&comp.Position{X: 0, Y: 0},
	// 	&comp.Velocity{X: 7, Y: 7},
	// )

	// system.AddSystems(movePosition, printPosition)

	// for i := 0; i < 10; i++ {
	// 	system.RunSystems(&ctx)
	// }

	// testQuery()
}

func testQuery() {
	// ctx := ent.NewContext()

	// ctx.Add(&comp.Position{X: 0, Y: 0})
	// ctx.Add(
	// 	&comp.Position{X: 0, Y: 0},
	// 	&comp.Velocity{X: 0, Y: 0},
	// )
	// ctx.Add(
	// 	&comp.Position{X: 0, Y: 0},
	// 	&comp.Velocity{X: 0, Y: 0},
	// )
	// ctx.Add(&comp.Position{X: 0, Y: 0})

	// qry := ctx.QueryList(
	// 	&comp.Position{},
	// 	&comp.Velocity{},
	// )

	// fmt.Printf("Query: %v\n", qry)

	// for _, comps := range qry {
	// 	fmt.Printf("[\n")
	// 	for _, c := range comps {
	// 		fmt.Printf("  %v,\n", c)
	// 	}
	// 	fmt.Printf("]\n")
	// }
}

func doWin() {
	// y := 300

	// settings := setup.NewSettings(800, 600, rl.RayWhite, 60)

	// settings.StartGame(
	// 	func() {
	// 		y++
	// 		rl.DrawText("Testing", 300, int32(y), 20, rl.Black)
	// 	},
	// )
}
