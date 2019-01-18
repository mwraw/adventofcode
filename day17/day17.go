package main

import (
	"fmt"

	"github.com/mwraw/adventofcode/day17/renderer"
)

func main() {
	space, waterX := renderer.RenderData("data.txt", 500)
	fmt.Println(waterX, space)
}
