// Author: Shayan Salehe <shay.sale86@gmail.com>
// Licence: MIT
package main

import (
	"fmt"

	"github.com/gen2brain/raylib-go/raylib"
)

const WIN_HEIGHT int32 = 1000

const WIN_WIDTH int32 = 700

type RunningMode int

const (
	MainMenu RunningMode = iota
	Editing
	Simulating
)

func main() {
	// Initialize window
	rl.InitWindow(WIN_HEIGHT, WIN_WIDTH, "Gravity Simulation")
	defer rl.CloseWindow()

	// Set target FPS
	rl.SetTargetFPS(60)
	planets := []planet{}
	// camera := rl.NewCamera2D(rl.NewVector2(float32(WIN_HEIGHT/2), float32(WIN_WIDTH/2)), rl.Vector2Zero(), 0, 0.1)
	currentMode := Editing
	for !rl.WindowShouldClose() {
		// Update your planetects here (for example, physics updates)
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
		}
		// Begin drawing
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		// rl.BeginMode2D(camera)
		// cameraControl(&camera)

		// End drawing
		rl.EndMode2D()
		rl.EndDrawing()
	}
}
