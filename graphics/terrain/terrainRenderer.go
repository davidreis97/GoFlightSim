package terrain

import (
	"fmt"

	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/core"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/renderer"
	"github.com/g3n/engine/texture"
)

var terrainMap = make(map[string]*graphic.Mesh)
var centerChunkKey = ""
var generator = RandomGenerator()
var grassTexture *texture.Texture2D

// Init loads resources required for the terrain rendering
func Init(appRenderer *renderer.Renderer){
	var err error
	grassTexture, err = texture.NewTexture2DFromImage("assets/textures/grass.jpg")
	if err != nil {
        fmt.Println(err)
    }
}

// RenderTerrain renders the terrain around a given coordinate
func RenderTerrain(x, y float32, scene *core.Node){
	currentcx := GetChunkIndex(x)
	currentcy := GetChunkIndex(y)

	currentChunkKey := GenChunkKey(currentcx,currentcy)

	if currentChunkKey == centerChunkKey{
		return
	}
	centerChunkKey = currentChunkKey

	setForUnrender := make(map[string]bool, len(terrainMap))
	for k := range terrainMap {
		setForUnrender[k] = true
	}

	for cx := currentcx - 3; cx <= currentcx + 2; cx++ {
		for cy := currentcy - 3; cy <= currentcy + 2; cy++{
			currentChunkKey := GenChunkKey(cx,cy)
			delete(setForUnrender, currentChunkKey)
			if chunk, exists := terrainMap[currentChunkKey]; !exists {
				go RenderChunk(cx,cy,scene)
			}else{
				chunk.SetVisible(true)
			}
		}
	}

	for key := range setForUnrender {
		terrainMap[key].SetVisible(false)
	}
}

// RenderChunk renders the chunk in the given coordinate
func RenderChunk(cx, cy int, scene *core.Node){
	currentChunkKey := GenChunkKey(cx,cy)

	chunkData := generator.NewChunk(float64(-cx * chunkSize), float64(cy * chunkSize))

	plane := geometry.NewSegmentedPlane(float32(chunkSize),float32(chunkSize),chunkSize,chunkSize)

	plane.OperateOnVertices(func (vertex *math32.Vector3) bool {
		vertex.Z = float32(chunkData[int(vertex.X)+chunkSize/2][int(vertex.Y)+chunkSize/2])
		return false
	})
	mat := material.NewStandard(&math32.Color{1,1,1})
	mat.AddTexture(grassTexture)
	newChunk := graphic.NewMesh(plane, mat)
	newChunk.RotateX(-math32.Pi/2)
	newChunk.RotateZ(math32.Pi)
	newChunk.SetPosition(float32(cx * chunkSize), 0, float32(cy * chunkSize))
	scene.Add(newChunk)

	terrainMap[currentChunkKey] = newChunk
}

// GenChunkKey generates the chunk key for the given chunk coordinates
func GenChunkKey(cx, cy int) string{
	return fmt.Sprintf("%d:%d", cx, cy)
}

// GetChunkIndex generates the chunk coordinate for the given coordinate
func GetChunkIndex(coord float32) int{
	return int(coord) / chunkSize
}