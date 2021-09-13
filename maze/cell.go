package maze

type CellType string

const (
	Goal     CellType = "G"
	Way               = "W"
	Obstacle          = "O"
	Bomb              = "B"
	Player            = "~"
)

type Cell struct {
	index      int
	cellType   CellType
	playerFlag bool
}

func (c *Cell) String() int {
	return c.index
}

func (c *Cell) GetCurrentState() string {
	if c.playerFlag {
		return string(Player)
	}
	return string(c.cellType)
}

func (c *Cell) GetPosition() (int, int) {
	// Bad code!! Wont work unless the
	// total cell count is 25
	if c.index < 5 {
		return 0, c.index
	}

	if c.index < 10 {
		return 1, c.index - 5
	}

	if c.index < 15 {
		return 2, c.index - 10
	}

	if c.index < 20 {
		return 3, c.index - 15
	}

	return 4, c.index - 20
}

func (c *Cell) GetIndexFrom(row, col int) int {

	// Bad code!! Wont work unless the
	// total cell count is 25
	return row + col + (4 * row)
}
