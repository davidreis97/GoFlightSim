package physics

import (
	"github.com/g3n/engine/math32"
	"github.com/davidreis97/GoFlightSim/controller"

	"time"
	//"fmt"
)

// Plane represents the physics status of a plane
type Plane struct {
	Transform *math32.Matrix4
	Velocity *math32.Matrix4
}

// NewPlane creates a new plane physics object with a given position
func NewPlane(pos math32.Vector3) *Plane {
	return &Plane{
		Transform: math32.NewMatrix4().SetPosition(&pos),
		Velocity: math32.NewMatrix4(),
	}
}

// Step processes a physics step for the given plane
func (p *Plane) Step(timeElapsed time.Duration, input controller.Input){
	tmpMatrix := math32.NewMatrix4()

	p.Transform.Multiply(tmpMatrix.MakeTranslation(0,0,input.Thrust * 0.01))
	p.Transform.Multiply(tmpMatrix.MakeRotationX(input.Pitch * 0.01))
	p.Transform.Multiply(tmpMatrix.MakeRotationY(input.Yaw * 0.01))
	p.Transform.Multiply(tmpMatrix.MakeRotationZ(input.Roll * 0.01))
}