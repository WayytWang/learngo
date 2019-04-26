package main

import (
	"fmt"
	"os"
)

func readMaze() [][]int {
	file, err := os.Open("maze.in")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var row, col int
	fmt.Fscanf(file, "%d %d\n", &row, &col)
	fmt.Printf("row:%d col%d \n", row, col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

type point struct {
	//i : 行  j : 列
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grip [][]int) (int, bool) {
	//上下越界
	if p.i < 0 || p.i >= len(grip) {
		return 0, false
	}

	//左右越界
	if p.j < 0 || p.j >= len(grip[0]) {
		return 0, false
	}

	return grip[p.i][p.j], true
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		for _, dir := range dirs {
			next := cur.add(dir)

			//maze at next is 0
			//and steps at next is 0
			//and next != start
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze()

	steps := walk(maze, point{i: 0, j: 0}, point{i: len(maze) - 1, j: len(maze[0]) - 1})
	for i := range steps {
		for j := range steps[i] {
			fmt.Print(steps[i][j])
			fmt.Print("  ")
		}
		fmt.Println("")
	}
}
