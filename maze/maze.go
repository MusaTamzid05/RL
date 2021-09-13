package maze

import (
	"fmt"
)

type Maze struct {
	cells       []Cell
	totalCells  int
	playerIndex int
}

func (m *Maze) initCells() {
	m.totalCells = 25
	m.playerIndex = 0

	for i := 0; i < m.totalCells; i += 1 {
		m.cells = append(m.cells, Cell{index: i, cellType: Way})
	}

	m.cells[m.playerIndex].playerFlag = true

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

func (m *Maze) Update(action Action) bool {
	doneFlag := false
	playerRow, playerCol := m.cells[m.playerIndex].GetPosition()
	nextRow, nextCol := m.getNextPositionFrom(action, playerRow, playerCol)

	if !m.isValid(nextRow, nextCol) {
		return doneFlag
	}

	nextIndex := m.getIndexFrom(nextRow, nextCol)

	if m.shouldUpdate(nextIndex) {
		m.update(nextIndex)

	}

	if m.goalReached() {
		doneFlag = true
	}

	return doneFlag

}

func (m *Maze) getNextPositionFrom(action Action, row, col int) (int, int) {

	if action == Up {
		row -= 1
	}

	if action == Down {
		row += 1
	}

	if action == Left {
		col -= 1
	}

	if action == Right {
		col += 1
	}

	return row, col
}

func (m *Maze) isValid(row, col int) bool {
	// Wont work if cells count is
	// not 25

	if row < 0 || row > 4 || col < 0 || col > 4 {
		return false
	}

	return true
}

func (m *Maze) getIndexFrom(row, col int) int {

	// Bad code!! Wont work unless the
	// total cell count is 25
	return row + col + (4 * row)
}
func (m *Maze) shouldUpdate(index int) bool {
	if m.cells[index].cellType == Obstacle {
		return false
	}
	return true
}

func (m *Maze) update(index int) {
	m.cells[m.playerIndex].playerFlag = false
	m.playerIndex = index
	m.cells[m.playerIndex].playerFlag = true
}

func (m *Maze) goalReached() bool {
	if m.cells[m.playerIndex].cellType == Goal || m.cells[m.playerIndex].cellType == Bomb {
		return true
	}

	return false
}

func New() *Maze {
	maze := Maze{}
	maze.initCells()
	return &maze
}
