package main
import(
	"math"
	"github.com/gen2brain/raylib-go/raylib"
)
const G float32 = 50
type planet struct {
	pos      rl.Vector2
	radius   float32
	velocity rl.Vector2
	acc      rl.Vector2
	mass     float32
	color    rl.Color
}

func newPlanet(pos rl.Vector2, radius float32, velocity rl.Vector2, acc rl.Vector2, mass float32, color rl.Color) planet {
	return planet{pos, radius, velocity, acc, mass, color}
}
func (p *planet) DrawPlanet() {
	rl.DrawCircle(int32(p.pos.X), int32(p.pos.Y), p.radius, p.color)
}
func (p *planet) calcAcc(op *planet) rl.Vector2 {
	r := rl.Vector2Subtract(op.pos, p.pos)
	if rl.Vector2Length(r) <= 300 {
		return rl.Vector2Zero()
	}
	g := (G * op.mass) / float32(math.Pow(float64(rl.Vector2Length(r)), 2))
	return rl.Vector2Scale(rl.Vector2Normalize(r), g)
}
func (p *planet) updateVelocity() {
	p.velocity = rl.Vector2Add(p.velocity, p.acc)
}
func (p *planet) updatePos() {
	p.pos = rl.Vector2Add(p.pos, p.velocity)
}
func (p *planet) updateAcc(planets []planet) {
	addAcc := rl.Vector2Zero()
	for i := range len(planets) {
		// Skip the planet itself
		if p == &(planets)[i] {
			continue
		}
		addAcc = rl.Vector2Add(addAcc, p.calcAcc(&(planets)[i]))
	}
	p.acc = addAcc
}
