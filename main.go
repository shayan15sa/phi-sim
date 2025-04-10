// Author: Shayan Salehe <shay.sale86@gmail.com>
// Licence: MIT
package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"math"
	"math/rand"
)

type planet struct {
	pos      rl.Vector2
	radius   float32
	velocity rl.Vector2
	acc      rl.Vector2
	mass     float32
	color    rl.Color
}

const G float32 = 50

func newPlanet(pos rl.Vector2, radius float32, velocity rl.Vector2, acc rl.Vector2, mass float32, color rl.Color) planet {
	return planet{pos, radius, velocity, acc, mass, color}
}

func DrawPlanet(o *planet) {
	rl.DrawCircle(int32(o.pos.X), int32(o.pos.Y), o.radius, o.color)
}
func calcAcc(p *planet, op *planet) rl.Vector2 {
	r := rl.Vector2Subtract(op.pos, p.pos)
	if rl.Vector2Length(r) <= 300 {
		return rl.Vector2Zero()
	}
	g := (G * op.mass) / float32(math.Pow(float64(rl.Vector2Length(r)), 2))
	return rl.Vector2Scale(rl.Vector2Normalize(r), g)
}
func updateAcc(planets *[]planet) {
	for i := 0; i < len(*planets); i++ {
		p := &(*planets)[i] // Get a pointer to the planet in the slice
		addAcc := rl.Vector2Zero()

		// Inner loop should use 'o' as the index for other planets
		for o := 0; o < len(*planets); o++ {
			// Skip the planet itself
			if p == &(*planets)[o] {
				continue
			}
			addAcc = rl.Vector2Add(addAcc, calcAcc(p, &(*planets)[o]))
		}

		// Update the planet's acceleration
		p.acc = addAcc
	}
}
func updateVelocity(planets *[]planet) {
	for i := 0; i < len(*planets); i++ {
		(*planets)[i].velocity = rl.Vector2Add((*planets)[i].velocity, (*planets)[i].acc)
	}
}
func updatePos(planets *[]planet) {
	for i := 0; i < len(*planets); i++ {
		(*planets)[i].pos = rl.Vector2Add((*planets)[i].pos, (*planets)[i].velocity)
	}
}
func main() {
	// Initialize window
	rl.InitWindow(1000, 700, "Gravity Simulation")

	// Set target FPS
	rl.SetTargetFPS(60)
	planets := []planet{}
	sun := newPlanet(rl.NewVector2(500, 350), 50, rl.Vector2Zero(), rl.Vector2Zero(), 1000, rl.Yellow)
	planets = append(planets, sun)

	// Planet Data: [radius, mass, distance from Sun, orbital velocity, color]
	planetData := [][]float32{
		// name: radius, mass, distance from Sun, orbital velocity, color
		{10, 0.33, 58, 0.07, 255},  // Mercury
		{12, 4.87, 108, 0.05, 255}, // Venus
		{13, 5.97, 150, 0.03, 0},   // Earth
		{11, 0.642, 228, 0.02, 255},// Mars
		{35, 1898, 778, 0.01, 255}, // Jupiter
		{30, 568, 1427, 0.009, 255}, // Saturn
		{25, 86.8, 2871, 0.006, 255}, // Uranus
		{25, 102, 4495, 0.004, 255},  // Neptune
	}

	// Create planets for the solar system
	for _, p := range planetData {
		// Simplified orbital velocity using the formula (G * Sun's mass / distance)^(1/2)
		orbitalVelocity := float32(math.Sqrt(float64(G * sun.mass / p[2]))) // Orbital velocity based on distance

		// Calculate planet's initial position based on distance from Sun
		planetPos := rl.NewVector2(sun.pos.X+p[2], sun.pos.Y)

		// Color randomization
		color := rl.Color{R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255}

		// Create the planet and add to the planets array
		planet := newPlanet(planetPos, p[0], rl.NewVector2(0, orbitalVelocity), rl.Vector2Zero(), p[1], color)
		planets = append(planets, planet)
	}
	camera := rl.NewCamera2D(rl.NewVector2(500, 350), rl.Vector2Zero(), 0, 0.1)
	// planets = append(planets, newPlanet(rl.NewVector2(300, 300), 35, rl.NewVector2(0.2, 0), rl.NewVector2(0, 0), 3, rl.Red))
	// planets = append(planets, newPlanet(rl.NewVector2(900, 600), 35, rl.NewVector2(0, -0.3), rl.NewVector2(0, 0), 6, rl.Blue))
	for !rl.WindowShouldClose() {
		// Update your planetects here (for example, physics updates)
		updateAcc(&planets)
		updateVelocity(&planets)
		updatePos(&planets)
		// Begin drawing
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.BeginMode2D(camera)
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
		// Draw a circle
		for _, o := range planets {
			DrawPlanet(&o)
			fmt.Println(o)
		}
		// End drawing
		rl.EndMode2D()
		rl.EndDrawing()
	}

	// Close window
	rl.CloseWindow()
}
