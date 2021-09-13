package maze

import "fmt"

type Maze struct {
	cells      []Cell
	totalCells int
}

func (m *Maze) initCells() {
	m.totalCells = 12

	for i := 0; i < m.totalCells; i += 1 {
		m.cells = append(m.cells, Cell{index: i, cellType: Wall})
	}
}

func (m *Maze) Show() {

	for i := 0; i < m.totalCells; i += 1 {
		if i != 0 && i%4 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%s", m.cells[i].GetCurrentState())
	}
	fmt.Println("")

}

func New() *Maze {
	maze := Maze{}
	maze.initCells()
	return &maze
}
