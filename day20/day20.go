package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	regex := readFile("data.txt")
	allRoutes := routes(regex)
	fmt.Println(len(allRoutes))
	destinations := destinations(allRoutes)
	fmt.Println(allRoutes, destinations)
	fmt.Println(maxDoors(allRoutes))
}

func readFile(fileName string) string {
	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return string(contents[1 : len(contents)-1])
}

func routes(regex string) []string {
	if len(regex) == 0 {
		return []string{""}
	}

	allRoutes := []string{}
	if isDirection(string(regex[0])) {
		for _, route := range routes(regex[1:]) {
			allRoutes = append(allRoutes, string(regex[0])+route)
		}
		return allRoutes
	}

	if string(regex[0]) == "(" {
		i := 1
		branchPoint := 1
		for openBrackets := 1; openBrackets > 0; i++ {
			if string(regex[i]) == "(" {
				openBrackets++
			} else if string(regex[i]) == ")" {
				openBrackets--
			} else if string(regex[i]) == "|" && openBrackets == 1 {
				branchPoint = i
			}
		}
		routesAfterRejoin := routes(regex[i:])
		for _, route := range routes(regex[1:branchPoint]) {
			for _, routeAfterRejoin := range routesAfterRejoin {
				allRoutes = append(allRoutes, route+routeAfterRejoin)
			}
		}
		for _, route := range routes(regex[branchPoint+1 : i-1]) {
			for _, routeAfterRejoin := range routesAfterRejoin {
				allRoutes = append(allRoutes, route+routeAfterRejoin)
			}
		}
		return allRoutes
	}
	return nil
}

func isDirection(s string) bool {
	return s == "N" || s == "E" || s == "S" || s == "W"
}

func destinations(allRoutes []string) [][2]int {
	dest := make([][2]int, len(allRoutes))

	for i := range allRoutes {
		for j := range allRoutes[i] {
			switch string(allRoutes[i][j]) {
			case "N":
				dest[i][1]++
			case "E":
				dest[i][0]++
			case "S":
				dest[i][1]--
			case "W":
				dest[i][0]--
			}
		}
	}

	return dest
}

func maxDoors(allRoutes []string) int {
	max := 0
	for _, route := range allRoutes {
		if len(route) > max {
			max = len(route)
		}
	}
	return max
}
