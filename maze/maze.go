package maze

import (
	"log"
)

type Maze struct {
	cells []Cell
}

func (m *Maze) initCells() {

	for i := 0; i < 12; i += 1 {
		m.cells = append(m.cells, Cell{index: i})
	}
}

func (m *Maze) Test() {
	log.Println("This is a test")
}

func New() *Maze {
	maze := Maze{}
	maze.initCells()
	return &maze
}
