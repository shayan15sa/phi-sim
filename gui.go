package main

import (
	"fmt"
	"strconv"

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
var massNumBox = NewNumBox(rl.NewRectangle(20, 10, 90, 40), 6, rl.Black, rl.SkyBlue)

func editPlanetUI() {
	DrawPlanets(editPlanets)
	massNumBox.showNumBox()
	cursorPos := rl.GetMousePosition()
	scrollMove := rl.GetMouseWheelMove()
	if scrollMove != 0 {
		cradius += scrollMove
	}
	if !massNumBox.mouseOnBox {
		rl.DrawCircle(int32(cursorPos.X), int32(cursorPos.Y), cradius, colors[ccolor])
	}
	fmt.Println("edit and draw")
	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		editPlanets = append(editPlanets, newPlanet(cursorPos, cradius, rl.Vector2Zero(), rl.Vector2Zero(), float32(massNumBox.getMass()), colors[ccolor]))
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

type NumBox struct {
	rect            rl.Rectangle
	numLen          int
	backgroundColor rl.Color
	foregroundColor rl.Color
	mouseOnBox      bool // set it to false; TODO: I couldn't find a better to handle this thing
	input           string
}

func NewNumBox(rect rl.Rectangle, numLen int, backgroundColor rl.Color, foregroundColor rl.Color) NumBox {
	return NumBox{rect, numLen, backgroundColor, foregroundColor, false, ""}
}

func (nb *NumBox) showNumBox() {
	// text handeling
	if rl.CheckCollisionPointRec(rl.GetMousePosition(), nb.rect) {
		nb.mouseOnBox = true
	} else {
		nb.mouseOnBox = false
	}
	if nb.mouseOnBox {
		rl.SetMouseCursor(rl.MouseCursorIBeam)
		key := rl.GetCharPressed()
		// Check if more characters have been pressed on the same frame
		for key > 0 {
			// NOTE: Only allow keys in range [32..125]
			if (key >= 48) && (key <= 57) && (len(nb.input) < nb.numLen) {
				nb.input += string(key)
			}
			key = rl.GetCharPressed() // Check next character in the queue
		}
		if rl.IsKeyPressed(rl.KeyBackspace) {
			if len(nb.input) > 0 {
				nb.input = nb.input[:len(nb.input)-1]
			}
		}
	} else {
		rl.SetMouseCursor(rl.MouseCursorDefault)
	}
	// drawing the box
	rl.DrawRectangleRec(nb.rect, nb.backgroundColor)
	if nb.mouseOnBox {
		rl.DrawRectangleLines(int32(nb.rect.X), int32(nb.rect.Y), int32(nb.rect.Width), int32(nb.rect.Height), rl.Red)
	} else {
		rl.DrawRectangleLines(int32(nb.rect.X), int32(nb.rect.Y), int32(nb.rect.Width), int32(nb.rect.Height), rl.LightGray)
	}
	if nb.input != "" {
		rl.DrawText(nb.input, int32(nb.rect.X+5), int32(nb.rect.Y+8), 26, nb.foregroundColor)
	}else{
		rl.DrawText("mass", int32(nb.rect.X+5), int32(nb.rect.Y+8), 26, rl.LightGray)
}
}

func (nb *NumBox) getMass() int {
	i, err := strconv.Atoi(nb.input)
	if err != nil {
		fmt.Println("AHHHHHHHHHHHHHHHHHH line 127 gui.go")
		// TODO : crach the program
	}
	return i
}
