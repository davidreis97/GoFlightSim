package terrain

import (
	"github.com/ojrac/opensimplex-go"
	"math/rand"
)

const chunkSize int = 180

// Generator contains the seeded noise instance
type Generator struct{
	noise opensimplex.Noise
	Seed int64
}

// RandomGenerator initializes the generator with a random seed
func RandomGenerator() *Generator {
	seed := rand.Int63()
	return NewGenerator(seed)
}

// NewGenerator initializes the generator with a given seed
func NewGenerator(seed int64) *Generator {
	return &Generator{noise: opensimplex.New(seed), Seed: seed}
}

// GetData provides terrain data for a given point
func (g Generator) GetData(x,y float64) float64{
	return g.noise.Eval2(x,y)
}

// NewChunk generates a new chunk starting in the given point
func (g Generator) NewChunk(x, y float64) *[chunkSize+1][chunkSize+1]float64{
	var chunk [chunkSize+1][chunkSize+1]float64

	frequency := 90.0
	intensity := 60.0

	for xi, yarr := range chunk {
		for yi := range yarr {
			chunk[xi][yi] = g.GetData((x + float64(xi))/frequency, (y + float64(yi))/frequency) * intensity
		}
	}

	return &chunk
}