package maze

import (
	"fmt"
	"math/rand"
	"time"
)

type Action int

const (
	Left  Action = 0
	Right        = 1
	Down         = 2
	Up           = 3
)

func generateNumber(maxNum int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	num := (r1.Intn(10) % maxNum) + 1
	return num

}

type Agent struct {
	maze *Maze
}

func (a *Agent) Run() {

	done := false

	for !done {
		nextAction := Action(generateNumber(4))
		done = a.maze.Update(nextAction)
		fmt.Println("=====================")
		a.maze.Show()
		fmt.Println("=====================")

	}

}

func NewAgent() *Agent {
	return &Agent{maze: New()}
}
