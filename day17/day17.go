package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func main() {
	lines := readData("data.txt")
	minX, minY, maxX, maxY := boundingBox(lines)
	space := emptySpace(minX, minY, maxX, maxY)
	addLines(space, lines, minX, minY)
	space[0][500-minX] = "+"
	printSpace(space)
}

func readData(fileName string) []line {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := []line{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ln := scanner.Text()

		axis, _ := strconv.Atoi(ln[2:strings.Index(ln, ",")])
		start, _ := strconv.Atoi(ln[strings.LastIndex(ln, "=")+1 : strings.Index(ln, "..")])
		end, _ := strconv.Atoi(ln[strings.Index(ln, "..")+2:])

		if string(ln[0]) == "x" {
			data = append(data, line{axis, start, axis, end})
		} else {
			data = append(data, line{start, axis, end, axis})
		}
	}

	return data
}

func boundingBox(lines []line) (int, int, int, int) {
	minX, minY, maxX, maxY := lines[0].x1, lines[0].y1, lines[0].x2, lines[0].y2

	for i := range lines {
		minX, minY = min(lines[i].x1, minX), min(lines[i].y1, minY)
		maxX, maxY = max(lines[i].x2, maxX), max(lines[i].y2, maxY)
	}

	return minX, minY, maxX, maxY
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	return -min(-a, -b)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func emptySpace(minX int, minY int, maxX int, maxY int) [][]string {
	space := make([][]string, maxY-minY+1)
	for y := range space {
		space[y] = make([]string, maxX-minX+1)
		for x := range space[y] {
			space[y][x] = "."
		}
	}
	return space
}

func printSpace(space [][]string) {
	for y := range space {
		fmt.Println(strings.Join(space[y], ""))
	}
}

func addLines(space [][]string, lines []line, xOffset int, yOffset int) {
	for i := range lines {
		xDir := 0
		if lines[i].x2 != lines[i].x1 {
			xDir = 1
		}

		for j := 0; j <= lineLength(lines[i]); j++ {
			space[lines[i].y1+j*(1-xDir)-yOffset][lines[i].x1+j*xDir-xOffset] = "#"
		}
	}
}

func lineLength(l line) int {
	return abs(l.x2-l.x1) + abs(l.y2-l.y1)
}
