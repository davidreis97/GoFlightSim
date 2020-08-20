package skybox

import (
	"fmt"
	"github.com/g3n/engine/graphic"
)

// Generate imports textures and generates a graphics.Skybox
func Generate() *graphic.Skybox{
	skyboxData := graphic.SkyboxData{
		DirAndPrefix: "assets/skybox/",
		Extension: "png",
		Suffixes: [6]string{"Front", "Back", "Top", "Bottom", "Left", "Right"},
	}

	skybox, err := graphic.NewSkybox(skyboxData)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return skybox
}