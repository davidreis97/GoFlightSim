package airplane

import (
	"github.com/g3n/engine/loader/obj"
	"github.com/g3n/engine/graphic"
)

// Init loads the spaceship .obj file
func Init() (*graphic.Mesh, error) {
	decoder, err := obj.Decode("assets/spaceModels/spaceCraft1.obj", "assets/spaceModels/spaceCraft1.mtl")

	if err != nil {
		return nil, err
	}
	
	return decoder.NewMesh(&decoder.Objects[0])
}