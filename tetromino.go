package optimizer

import (
	utils "./utils"
)

type Tetromino struct {
	id      string
	depth   uint8
	jMax    int
	iMax    int
	iCoords []int
	jCoords []int
	Next    *Tetromino
}

type Sequence struct {
	Head *Tetromino
	Tail *Tetromino
}

var length = 0

//PushBack push node back to the linked list
func (s *Sequence) PushBack(data [][]byte) {
	node := &Tetromino{depth: 1}
	node.Validate(data)
	if s.Head == nil {
		s.Head = node
		s.Tail = node
		return
	}
	node.setDepth(s.Tail)
	s.Tail.Next = node
	s.Tail = node
}

//Validate validating tetrominoes
func (t *Tetromino) Validate(shape [][]byte) {
	// 1. Count touching sides
	// 2. Get Active Coords
	// 3. Validate
	length = len(shape)
	touchingSides := 0
	iActive, jActive := make([]int, length), make([]int, length) // Contain coordinates of "#" symbol

	ind := 0
	for j := 0; j < length; j++ {
		for i := 0; i < length; i++ {
			if shape[j][i] == '#' {
				if ind == 4 {
					panic("ERROR")
				}
				touchingSides += utils.CountTouchingSides(shape, i, j, length)
				iActive[ind] = i
				jActive[ind] = j
				ind++
			}
		}
	}
	utils.ValidateTouchingSides(touchingSides)
	//Shift tetrominoes to the left-up corner
	t.setICoords(iActive)
	t.setJCoords(jActive)
	//Find the right-down corner coords of shifted tetrominoes
	t.setIMax()
	t.setJMax()
	//Set unique id for checking repeats
	t.setId()
}

//setICoords set shifted i coords
func (t *Tetromino) setICoords(iActive []int) {
	if len(iActive) == 0 {
		return
	}
	_min := utils.Min(iActive)
	for k := 0; k < len(iActive); k++ {
		iActive[k] -= _min
	}
	t.iCoords = iActive
}

//setJCoords set shifted j coords
func (t *Tetromino) setJCoords(jActive []int) {
	if len(jActive) == 0 {
		return
	}
	_min := utils.Min(jActive)
	for k := 0; k < len(jActive); k++ {
		jActive[k] -= _min
	}
	t.jCoords = jActive
}

//setDepth set depth which used for printing letters
func (t *Tetromino) setDepth(prev *Tetromino) {
	t.depth = prev.depth + 1
}

//setIMax set right down corner j coord of tetromineos
func (t *Tetromino) setIMax() {
	t.iMax = utils.Max(t.iCoords)
}

//setJMax set right down corner j coord of tetromineos
func (t *Tetromino) setJMax() {
	t.jMax = utils.Max(t.jCoords)
}

//setId set id for blacklisting(repeat)
func (t *Tetromino) setId() {
	iStr, jStr := "", ""
	for i := 0; i < len(t.jCoords); i++ {
		iStr += string(t.iCoords[i])
		jStr += string(t.jCoords[i])
	}
	t.id = iStr + jStr
}
