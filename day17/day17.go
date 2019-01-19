package main

import (
	"fmt"

	"github.com/mwraw/adventofcode/day17/renderer"
)

const initX = 500

func main() {
	space, waterX := renderer.RenderData("data.txt", initX)
	//renderer.PrintSpace(space)
	waterPath(space, waterX, 0, 1)
	//renderer.PrintSpace(space)
	fmt.Printf("Part 1: %v. Part 2: %v.\n", countSpaces(space, []string{"~", "|"}), countSpaces(space, []string{"~"}))
}

func waterPath(space [][]string, x int, y int, dirY int) {
	if y+1 == len(space) {
		return
	}

	if space[y+1][x] == "|" {
		return
	}

	if space[y+1][x] == "." {
		space[y+1][x] = "|"
		waterPath(space, x, y+1, 1)
		return
	}

	if space[y+1][x] == "~" && dirY == 1 {
		waterPath(space, x, y+1, 1)
		return
	}

	if space[y+1][x] == "#" || (space[y+1][x] == "~" || space[y+1][x] == "|") && dirY == -1 {
		escaped := false
		for dirX := 1; dirX >= -1; dirX -= 2 {
			for i := 0; i == 0 || space[y][x+dirX*i] == "." || space[y][x+dirX*i] == "~" || space[y][x+dirX*i] == "|"; i++ {
				if space[y+1][x+dirX*i] != "#" && space[y+1][x+dirX*i] != "~" {
					space[y][x+dirX*i] = "|"
					escaped = true
					waterPath(space, x+dirX*i, y, 1)
					break
				}
				space[y][x+dirX*i] = "~"
			}
		}

		if escaped && escapable(space, x, y) {
			for dirX := 1; dirX >= -1; dirX -= 2 {
				for i := 0; space[y][x+dirX*i] == "~" || i == 0; i++ {
					space[y][x+dirX*i] = "|"
				}
			}
		}
		if !escaped {
			waterPath(space, x, y-1, -1)
		}
	}
}

func escapable(space [][]string, x int, y int) bool {
	for dirX := 1; dirX >= -1; dirX -= 2 {
		for i := 1; space[y][x+dirX*i] == "~" || space[y][x+dirX*i] == "|"; i++ {
			if space[y][x+dirX*i] == "|" {
				return true
			}
		}
	}
	return false
}

func countSpaces(space [][]string, targets []string) int {
	totalWater := 0

	for y := range space {
		for x := range space[y] {
			for i := range targets {
				if space[y][x] == targets[i] {
					totalWater++
					break
				}
			}
		}
	}

	return totalWater
}
