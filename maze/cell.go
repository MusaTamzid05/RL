package maze

type CellType string

const (
	Goal     CellType = "G"
	Way               = "W"
	Obstacle          = "O"
	Bomb              = "B"
)

type Cell struct {
	index    int
	cellType CellType
}

func (c *Cell) String() int {
	return c.index
}

func (c *Cell) GetCurrentState() string {
	return string(c.cellType)
}
