package pathfinding

import (
	"sort"
)

// Point (x,y)
type Point struct {
	X int
	Y int
}

type distance struct {
	p    Point
	dist int
}

// BestPath returns the shortest route to the closest point with the lowest reading order. If no route exists, false is returned as the second value.
func BestPath(start Point, ends []Point, mapData [][]string) ([]Point, bool) {
	route := []Point{}
	minLength := -1

	distances := allDistances(start, mapData)

	for _, end := range ends {
		candidateRoute, foundRoute := shortestPath(start, end, distances)

		if !foundRoute || len(candidateRoute) > minLength && minLength != -1 {
			continue
		}

		if minLength == -1 || len(candidateRoute) < minLength || len(candidateRoute) == minLength && lowerReadingOrder(candidateRoute[len(candidateRoute)-1], route[len(route)-1]) {
			minLength = len(candidateRoute)
			route = candidateRoute
		}
	}

	return route, len(route) > 0
}

func shortestPath(start Point, end Point, dist []distance) ([]Point, bool) {
	if routeExists, _ := isInDistances(dist, end); !routeExists {
		return []Point{}, false
	}

	candidates := candidatePoints(start, end, dist)
	route := []Point{}
	curDist := 0

	for i := range candidates {
		if candidates[len(candidates)-i-1].dist == curDist {
			route = append(route, candidates[len(candidates)-i-1].p)
			curDist++
		}
	}

	return route, true
}

func manhattan(p1 Point, p2 Point) int {
	return abs(p1.X-p2.X) + abs(p1.Y-p2.Y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func allDistances(start Point, mapData [][]string) []distance {
	dist := []distance{}
	open := []distance{distance{p: start, dist: 0}}

	for len(open) > 0 {
		sort.Slice(open, func(i, j int) bool { return open[i].dist < open[j].dist })
		base := open[0]
		dist = append(dist, base)
		open = open[1:]
		adj := AdjacentSquares(base.p, mapData)

		for i := range adj {
			if inClosed, _ := isInDistances(dist, adj[i]); inClosed {
				continue
			}

			if inOpen, j := isInDistances(open, adj[i]); inOpen {
				if base.dist+1 < open[j].dist {
					open[j].dist = base.dist + 1
				}
				continue
			}
			open = append(open, distance{p: adj[i], dist: base.dist + 1})
		}
	}

	return dist
}

// AdjacentSquares returns a slice containing all the points adjacent to the base, excluding any points that are either occupied or contain walls
func AdjacentSquares(base Point, mapData [][]string) []Point {
	adj := []Point{}

	for y := -1; y <= 1; y++ {
		for x := -1 + abs(y); x <= 1-abs(y); x++ {
			if !(x == 0 && y == 0) && mapData[base.Y+y][base.X+x] == "." {
				adj = append(adj, Point{X: base.X + x, Y: base.Y + y})
			}
		}
	}

	return adj
}

func isInDistances(open []distance, p Point) (bool, int) {
	for i := range open {
		if manhattan(open[i].p, p) == 0 {
			return true, i
		}
	}
	return false, -1
}

func candidatePoints(start Point, end Point, dist []distance) []distance {
	sortDistances(dist)
	_, endIdx := isInDistances(dist, end)
	candidates := []distance{dist[endIdx]}

	for i := endIdx - 1; i >= 0; i-- {
		for j := range candidates {
			if dist[i].dist == candidates[j].dist-1 && manhattan(dist[i].p, candidates[j].p) == 1 {
				candidates = append(candidates, dist[i])
				break
			}
		}
	}

	return candidates
}

func sortDistances(dist []distance) {
	sort.Slice(dist, func(i, j int) bool {
		return dist[i].dist < dist[j].dist || dist[i].dist == dist[j].dist && lowerReadingOrder(dist[i].p, dist[j].p)
	})
}

func lowerReadingOrder(p1 Point, p2 Point) bool {
	return p1.Y < p2.Y || p1.Y == p2.Y && p1.X < p2.X
}
