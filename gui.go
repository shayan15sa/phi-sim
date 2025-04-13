package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func cameraControl(camera *rl.Camera2D) {
	scroll := rl.GetMouseWheelMove()
	if scroll > 0 {
		camera.Zoom += 0.1 // Zoom in
	} else if scroll < 0 {
		camera.Zoom -= 0.1 // Zoom out
	}
	if rl.IsKeyPressed(rl.KeyK) {
		camera.Zoom += 0.05 // Zoom in
	}
	if rl.IsKeyPressed(rl.KeyJ) {
		camera.Zoom -= 0.05 // Zoom in
	}
	if rl.IsKeyDown(rl.KeyW) {
		camera.Offset.Y += 10
	}
	if rl.IsKeyDown(rl.KeyS) {
		camera.Offset.Y -= 10
	}
	if rl.IsKeyDown(rl.KeyA) {
		camera.Offset.X += 10
	}
	if rl.IsKeyDown(rl.KeyD) {
		camera.Offset.X -= 10
	}
}

var cradius float32 = 10
var colors = []rl.Color{rl.Red, rl.Blue, rl.Black, rl.Orange, rl.SkyBlue, rl.Purple, rl.Pink, rl.Green}
var ccolor = 0
var editPlanets = []planet{}

func editPlanetUI() {
	DrawPlanets(editPlanets)
	cursorPos := rl.GetMousePosition()
	scrollMove := rl.GetMouseWheelMove()
	if scrollMove != 0 {
		cradius += scrollMove
	}
	rl.DrawCircle(int32(cursorPos.X), int32(cursorPos.Y), cradius, colors[ccolor])
	fmt.Println("edit and draw")
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		editPlanets = append(editPlanets, newPlanet(cursorPos, cradius, rl.Vector2Zero(), rl.Vector2Zero(), 100, colors[ccolor]))
		// TODO: Add a better way to indicate that the planet has been added
		rl.DrawCircle(int32(cursorPos.X), int32(cursorPos.Y), cradius*2, rl.Yellow)
	}
	if rl.IsMouseButtonPressed(rl.MouseButtonRight) {
		ccolor = (ccolor + 1) % len(colors)
	}
}
func getAddedPlanets() []planet {
	return editPlanets
}
func emptyAddedPlanets() {
	editPlanets = []planet{}
}

// TODO: add a textbox for inputing the math https://www.raylib.com/examples/text/loader.html?name=text_input_box
