package optimizer

import (
	"math"
	"strings"
)

var size int

//Solve start solving algo
func (d Data) Solve() string {
	// initial size := sqrt(# of tetrominoes * 4 characters per tetromino)
	size = int(math.Sqrt(float64(4 * (d.tetrominoes.Tail.depth))))
	for ; size < size*size; size++ {
		square := strings.Repeat(".", size*size)
		// fmt.Println(size)
		if solution := Next(square, d.tetrominoes.Head); solution != "" {
			return solution
		}
	}
	panic("No solution")
}

//Next move our solver to next step
func Next(square string, tetro *Tetromino) string {
	if tetro == nil {
		return square
	}

	length := int(math.Sqrt(float64(len(square))))
	var blacklist = []string{}

	for j := 0; j+tetro.jMax < length; j++ {

	iLoop:
		for i := 0; i+tetro.iMax < length; i++ {
			// Check if teetromino repeating
			for _, item := range blacklist {
				if item == tetro.id {
					return ""
				}
			}
			// Place a Tetromino
			tmpSquare := []byte(square)
			for k := range tetro.iCoords {
				cell := &tmpSquare[length*(j+tetro.jCoords[k])+(i+tetro.iCoords[k])]
				if *cell != '.' {
					continue iLoop
				}
				*cell = tetro.depth
			}
			tmp := string(tmpSquare)
			if sol := Next(tmp, tetro.Next); sol != "" {
				return sol
			}
			blacklist = append(blacklist, tetro.Next.id)
		}
	}
	return ""
}
