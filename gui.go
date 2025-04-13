package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
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

var radius float32 = 10
var editPlanets = []planet{}

func editPlanetUI() {
	cursorPos := rl.GetMousePosition()
	scrollMove := rl.GetMouseWheelMove()
	if scrollMove != 0 {
		radius += scrollMove
	}
	rl.DrawCircle(int32(cursorPos.X), int32(cursorPos.Y), radius, rl.Red)
	fmt.Println("edit and draw")
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		editPlanets = append(editPlanets, newPlanet(cursorPos, radius, rl.Vector2Zero(), rl.Vector2Zero(), 100, rl.Red))
		// TODO: Add a better way to indicate that the planet has been added
		rl.DrawCircle(int32(cursorPos.X), int32(cursorPos.Y), radius*2, rl.Yellow)
	}
	DrawPlanets(editPlanets)
}
func getAddedPlanets() []planet {
	return editPlanets
}
