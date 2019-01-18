package main

import (
	"adventofcode/day15/pathfinding"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const initialHp = 200

type point = pathfinding.Point

type creature struct {
	pos         point
	hp          int
	team        string
	turnsTaken  int
	attackPower int
}

func main() {
	initMapData := parseFile("data.txt")

	winner := "G"
	creatures := []creature{}
	mapData := [][]string{}

	for attackPower := 3; lowestHp(creatures, "E") <= 0; attackPower++ {
		mapData = make([][]string, len(initMapData))
		for y := range initMapData {
			mapData[y] = make([]string, len(initMapData[y]))
			copy(mapData[y], initMapData[y])
		}

		creatures = findCreatures(mapData, 3, attackPower)

		changedLastTurn := true

	TurnLoop:
		for {
			changedInTurn := false
			sortReadingOrder(creatures)

			for i := range creatures {
				creatures[i].turnsTaken++

				if creatures[i].hp <= 0 {
					continue
				}

				target := targetCreature(creatures, i)

				if (changedLastTurn || changedInTurn) && target == -1 {
					changedInTurn = moveCreature(creatures, i, mapData) || changedInTurn
					target = targetCreature(creatures, i)
				}

				if target != -1 {
					creatures[target].hp -= creatures[i].attackPower
					if creatures[target].hp <= 0 {
						changedInTurn = true
						mapData[creatures[target].pos.Y][creatures[target].pos.X] = "."
						if gameOver(creatures, creatures[target].team) {
							winner = creatures[i].team
							fmt.Printf("Attack power=%v, winner=%v, turn=%v, score=%v\n", attackPower, winner, creatures[i].turnsTaken, score(creatures))
							break TurnLoop
						}
					}
				}
			}
			changedLastTurn = changedInTurn
		}
	}
}

func lowestHp(creatures []creature, team string) int {
	minHp := -999

	for i := range creatures {
		if creatures[i].team == team && (minHp == -999 || creatures[i].hp < minHp) {
			minHp = creatures[i].hp
		}
	}

	return minHp
}

func score(creatures []creature) int {
	totalHp := 0
	endTurn := 0
	for i := range creatures {
		if creatures[i].hp > 0 {
			totalHp += creatures[i].hp
			if endTurn == 0 || creatures[i].turnsTaken < endTurn {
				endTurn = creatures[i].turnsTaken
			}
		}
	}
	return totalHp * endTurn
}

func printMapData(mapData [][]string) {
	for y := range mapData {
		for x := range mapData[y] {
			print(mapData[y][x])
		}
		print("\n")
	}
}

func manhattan(p1 point, p2 point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func targetCreature(creatures []creature, j int) int {
	target := -1

	for i := range creatures {
		if i != j && manhattan(creatures[i].pos, creatures[j].pos) == 1 && creatures[i].team == oppositeTeam(creatures[j].team) && creatures[i].hp > 0 {
			if target == -1 || creatures[i].hp < creatures[target].hp || creatures[i].hp == creatures[target].hp && creatureLowerReadingOrder(creatures, i, target) {
				target = i
			}
		}
	}

	return target
}

func sortReadingOrder(creatures []creature) {
	sort.Slice(creatures, func(i, j int) bool {
		return creatureLowerReadingOrder(creatures, i, j)
	})
}

func lowerReadingOrder(p1 point, p2 point) bool {
	return p1.Y < p2.Y || p1.Y == p2.Y && p1.X < p2.X
}

func creatureLowerReadingOrder(creatures []creature, i int, j int) bool {
	return lowerReadingOrder(creatures[i].pos, creatures[j].pos)
}

func parseFile(fileName string) [][]string {
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

func findCreatures(mapData [][]string, goblinAP int, elfAP int) []creature {
	creatures := []creature{}

	for y := range mapData {
		for x := range mapData[y] {
			if mapData[y][x] == "G" {
				creatures = append(creatures, creature{point{X: x, Y: y}, initialHp, mapData[y][x], 0, goblinAP})
			}
			if mapData[y][x] == "E" {
				creatures = append(creatures, creature{point{X: x, Y: y}, initialHp, mapData[y][x], 0, elfAP})
			}
		}
	}

	return creatures
}

func oppositeTeam(team string) string {
	if team == "E" {
		return "G"
	}
	return "E"
}

func gameOver(creatures []creature, teamToCheck string) bool {
	for i := range creatures {
		if creatures[i].hp > 0 && creatures[i].team == teamToCheck {
			return false
		}
	}
	return true
}

func moveCreature(creatures []creature, currentCreature int, mapData [][]string) bool {
	ends := targetSquares(creatures, currentCreature, mapData)
	route, foundPath := pathfinding.BestPath(creatures[currentCreature].pos, ends, mapData)

	if foundPath {
		mapData[route[0].Y][route[0].X] = "."
		mapData[route[1].Y][route[1].X] = creatures[currentCreature].team
		creatures[currentCreature].pos = route[1]
	}

	return foundPath
}

func targetSquares(creatures []creature, currentCreature int, mapData [][]string) []point {
	ends := []point{}

	for i := range creatures {
		if i == currentCreature || creatures[i].team == creatures[currentCreature].team || creatures[i].hp <= 0 {
			continue
		}

		adj := pathfinding.AdjacentSquares(creatures[i].pos, mapData)

		for j := range adj {
			if !inListOfPoints(ends, adj[j]) {
				ends = append(ends, adj[j])
			}
		}
	}

	return ends
}

func inListOfPoints(listToSearch []point, p point) bool {
	for i := range listToSearch {
		if listToSearch[i].X == p.X && listToSearch[i].Y == p.Y {
			return true
		}
	}
	return false
}
