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
