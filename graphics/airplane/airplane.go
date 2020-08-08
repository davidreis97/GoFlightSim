package airplane

import (
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/graphic"
)

// Init loads the spaceship .obj file
func Init() (*graphic.Mesh, error) {
	decoder, err := obj.Decode("assets/spaceModels/spaceCraft1.obj", "")

	if err != nil {
		return nil, err
	}

	for _, mat := range decoder.Materials {
		mat.Ambient = mat.Diffuse
	}

	mesh, err := decoder.NewMesh(&decoder.Objects[0])

	if err != nil {
		return nil, err
	}

	return mesh, err
}