package main

import "fmt"

const players = 470
const highestMarble = 7217000

type marble struct {
	value    int
	next     *marble
	previous *marble
}

func main() {
	var scores [players]int
	player := 0

	currentMarble := &marble{0, nil, nil}
	currentMarble.next = currentMarble
	currentMarble.previous = currentMarble

	for marbleValue := 1; marbleValue <= highestMarble; marbleValue++ {
		if marbleValue%23 == 0 {
			currentMarble = moveMarble(currentMarble, 7, -1)
			scores[player] += marbleValue + currentMarble.value
			currentMarble = deleteMarble(currentMarble)
		} else {
			currentMarble = moveMarble(currentMarble, 1, 1)
			currentMarble = addMarble(marbleValue, currentMarble)
		}
		player = (player + 1) % players
	}

	fmt.Println(arrayMax(scores[:]))
}

func arrayMax(arr []int) int {
	maxValue := arr[0]

	for _, val := range arr {
		if val > maxValue {
			maxValue = val
		}
	}

	return maxValue
}

func addMarble(marbleValue int, after *marble) *marble {
	newMarble := &marble{marbleValue, after.next, after}

	after.next.previous = newMarble
	after.next = newMarble

	return newMarble
}

func deleteMarble(currentMarble *marble) *marble {
	currentMarble.previous.next = currentMarble.next
	currentMarble.next.previous = currentMarble.previous

	return currentMarble.next
}

func moveMarble(currentMarble *marble, distance int, direction int) *marble {
	for i := 0; i < distance; i++ {
		if direction == 1 {
			currentMarble = currentMarble.next
		} else {
			currentMarble = currentMarble.previous
		}
	}
	return currentMarble
}
