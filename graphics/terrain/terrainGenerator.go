package terrain

import (
	"github.com/ojrac/opensimplex-go"
	"math/rand"
	"fmt"
)

// Generator contains the seeded noise instance
type Generator struct{
	noise opensimplex.Noise
}

// NewGenerator initializes the generator with a random seed
func NewGenerator() *Generator {
	return &Generator{noise: opensimplex.NewNormalized(rand.Int63())}
}

// GetData provides terrain data for a given point
func (g Generator) GetData(x,y float64) float64{
	return g.noise.Eval2(x,y)
}

// Chunk contains terrain data for a 32x32 section
type Chunk struct{
	X, Y float64
	Data [32][32]float64
}

// NewChunk generates a new chunk starting in the given point
func (g Generator) NewChunk(x, y float64) *Chunk{
	chunk := Chunk{X: x, Y: y}

	for xi, yarr := range chunk.Data {
		for yi := range yarr {
			chunk.Data[xi][yi] = g.GetData(x + float64(xi), y + float64(yi))
		}
	}

	fmt.Println(chunk.Data)

	return &chunk
}