// Author: Shayan Salehe <shay.sale86@gmail.com>
// Licence: MIT
package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var WIN_HEIGHT int32 = 1000

var WIN_WIDTH int32 = 700

type RunningMode int

const (
	MainMenu RunningMode = iota
	Editing
	Simulating
)

func main() {
	// Initialize window
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(WIN_HEIGHT, WIN_WIDTH, "Gravity Simulation")
	defer rl.CloseWindow()

	// Set target FPS
	rl.SetTargetFPS(60)
	planets := []planet{}
	camera := rl.NewCamera2D(rl.NewVector2(float32(WIN_HEIGHT/2), float32(WIN_WIDTH/2)), rl.Vector2Zero(), 0, 1)
	currentMode := Editing
	for !rl.WindowShouldClose() {
		WIN_HEIGHT = int32(rl.GetScreenHeight())
		WIN_WIDTH = int32(rl.GetScreenWidth())
		if rl.IsKeyPressed(rl.KeyQ) {
			rl.CloseWindow()
		}
		if rl.IsKeyPressed(rl.KeyC) {
			emptyAddedPlanets()
			currentMode = Editing
		}
		// Update your planetects here (for example, physics updates)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		switch currentMode {
		case Editing:
			editPlanetUI()
			if rl.IsKeyPressed(rl.KeyEnter) {
				planets = getAddedPlanets()
				currentMode = Simulating
				fmt.Println("mode changing")
			}
		case Simulating:
			updatePlanets(planets)
			DrawPlanets(planets)
			cameraControl(&camera)
			if rl.IsKeyPressed(rl.KeyP) {
				currentMode = Editing
			}
		}
		// Begin drawing

		// End drawing
		rl.EndDrawing()
	}
}
