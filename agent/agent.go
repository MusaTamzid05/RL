package agent

import (
	"math/rand"
	"rl/maze"
	"time"
)

func generateNumber(maxNum int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	num := (r1.Intn(10) % maxNum) + 1
	return num

}

type Agent struct {
	maze *maze.Maze
}

func (a *Agent) Run(maze_ *maze.Maze) {
	nextAction := generateNumber(4)
	a.maze.Update(nextAction)

}
