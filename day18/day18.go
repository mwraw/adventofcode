package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	mapStart := readFile("data.txt")

	for t := 1; t <= 10000; t++ {
		mapEnd := make([][]string, len(mapStart))

		for y := range mapStart {
			mapEnd[y] = make([]string, len(mapStart[y]))

			for x := range mapStart[y] {
				if mapStart[y][x] == "." && countAdjacent(mapStart, x, y, "|") >= 3 {
					mapEnd[y][x] = "|"
					continue
				}
				if mapStart[y][x] == "|" && countAdjacent(mapStart, x, y, "#") >= 3 {
					mapEnd[y][x] = "#"
					continue
				}
				if mapStart[y][x] == "#" && (countAdjacent(mapStart, x, y, "#") == 0 || countAdjacent(mapStart, x, y, "|") == 0) {
					mapEnd[y][x] = "."
					continue
				}
				mapEnd[y][x] = mapStart[y][x]
			}
		}

		mapStart = mapEnd
		fmt.Println(t, score(mapStart))
	}
}

func readFile(fileName string) [][]string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	mapData := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mapData = append(mapData, strings.Split(scanner.Text(), ""))
	}

	return mapData
}

func printMap(mapData [][]string) {
	for y := range mapData {
		fmt.Println(strings.Join(mapData[y], ""))
	}
}
func countAdjacent(mapData [][]string, baseX int, baseY int, lookFor string) int {
	count := 0

	for y := max(baseY-1, 0); y <= min(baseY+1, len(mapData)-1); y++ {
		for x := max(baseX-1, 0); x <= min(baseX+1, len(mapData[y])-1); x++ {
			if y == baseY && x == baseX {
				continue
			}
			if mapData[y][x] == lookFor {
				count++
			}
		}
	}

	return count
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

func score(mapData [][]string) int {
	wood := 0
	lumberyard := 0

	for y := range mapData {
		for x := range mapData[y] {
			if mapData[y][x] == "|" {
				wood++
				continue
			}
			if mapData[y][x] == "#" {
				lumberyard++
				continue
			}
		}
	}

	return wood * lumberyard
}
