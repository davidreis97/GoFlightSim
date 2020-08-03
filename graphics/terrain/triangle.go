package terrain

import (
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/gls"
)

// NewTriangle creates a new triangle geometry with the given vertices
func NewTriangle(vertices math32.ArrayF32) *geometry.Geometry {
	geom := geometry.NewGeometry()
	triVBO := gls.NewVBO(vertices).AddAttrib(gls.VertexPosition)
	geom.AddVBO(triVBO)
	return geom
}