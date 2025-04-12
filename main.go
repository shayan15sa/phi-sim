// Author: Shayan Salehe <shay.sale86@gmail.com>
// Licence: MIT
package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

func cameraControl(camera *rl.Camera2D){
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
func main() {
	// Initialize window
	rl.InitWindow(1000, 700, "Gravity Simulation")
	defer rl.CloseWindow()

	// Set target FPS
	rl.SetTargetFPS(60)
	planets := []planet{}

	camera := rl.NewCamera2D(rl.NewVector2(500, 350), rl.Vector2Zero(), 0, 0.1)

	for !rl.WindowShouldClose() {
		// Update your planetects here (for example, physics updates)
		for i := range planets {
			planet := &planets[i]
			planet.updateAcc(planets)
			planet.updateVelocity()
			planet.updatePos()
		}
		// Begin drawing
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode2D(camera)
		cameraControl(&camera)
		// Draw planets
		for _, p := range planets {
			p.DrawPlanet()
		}
		// End drawing
		rl.EndMode2D()
		rl.EndDrawing()
	}
}
