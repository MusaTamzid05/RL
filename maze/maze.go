package maze

import "fmt"

type Maze struct {
	cells      []Cell
	totalCells int
}

func (m *Maze) initCells() {
	m.totalCells = 25

	for i := 0; i < m.totalCells; i += 1 {
		m.cells = append(m.cells, Cell{index: i, cellType: Way})
	}

	m.cells[4].cellType = Goal
	m.cells[1].cellType = Obstacle
	m.cells[2].cellType = Obstacle
	m.cells[7].cellType = Obstacle
	m.cells[9].cellType = Obstacle
	m.cells[10].cellType = Obstacle
	m.cells[14].cellType = Obstacle
	m.cells[18].cellType = Obstacle
	m.cells[19].cellType = Obstacle
	m.cells[20].cellType = Obstacle
	m.cells[24].cellType = Bomb
}

func (m *Maze) Show() {

	for i := 0; i < m.totalCells; i += 1 {
		if i != 0 && i%5 == 0 {
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
