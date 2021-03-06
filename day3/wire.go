package day3

import (
	"strconv"
	"strings"
)

// R75,D30,R83,U83,L12,D49,R71,U7,L72
// U62,R66,U55,R34,D71,R55,D58,R83

//BuildCoordList stuff
func BuildCoordList(s string) []Coordinate {
	instruction := strings.Split(s, ",")
	coords := make([]Coordinate, 0, 100)
	coords = append(coords, Coordinate{
		X: 0,
		Y: 0,
	})
	lastCoord := Coordinate{
		X: 0,
		Y: 0,
	}
	for _, currentinst := range instruction {
		// fmt.Println(currentinst)
		direction := currentinst[0]
		quantity, _ := strconv.Atoi(currentinst[1:])
		// fmt.Println(quantity)
		switch direction {
		case 'L':
			for l := 1; l <= quantity; l++ {
				coords = append(coords, Coordinate{
					X: lastCoord.X - l,
					Y: lastCoord.Y,
				})
			}
			lastCoord.X = lastCoord.X - quantity

		case 'R':
			for l := 1; l <= quantity; l++ {
				coords = append(coords, Coordinate{
					X: lastCoord.X + l,
					Y: lastCoord.Y,
				})
			}
			lastCoord.X = lastCoord.X + quantity

		case 'U':
			for l := 1; l <= quantity; l++ {
				coords = append(coords, Coordinate{
					X: lastCoord.X,
					Y: lastCoord.Y + l,
				})
			}
			lastCoord.Y = lastCoord.Y + quantity

		case 'D':
			for l := 1; l <= quantity; l++ {
				coords = append(coords, Coordinate{
					X: lastCoord.X,
					Y: lastCoord.Y - l,
				})
			}
			lastCoord.Y = lastCoord.Y - quantity
		}

	}

	return coords
}

//BuildCoordMap   builds a map with the string representation of a coord as key and the value is the number of steps to get there
func BuildCoordMap(s string) map[string]int {
	instruction := strings.Split(s, ",")

	coords := make(map[string]int)
	lastCoord := Coordinate{
		X: 0,
		Y: 0,
	}
	coords[lastCoord.String()] = 0
	steps := 0

	for _, currentinst := range instruction {
		// fmt.Println(currentinst)
		direction := currentinst[0]
		quantity, _ := strconv.Atoi(currentinst[1:])
		// fmt.Println(quantity)
		switch direction {
		case 'L':
			for l := 1; l <= quantity; l++ {
				tmp := Coordinate{
					X: lastCoord.X - l,
					Y: lastCoord.Y,
				}
				steps++
				coords[tmp.String()] = steps
			}
			lastCoord.X = lastCoord.X - quantity

		case 'R':
			for l := 1; l <= quantity; l++ {
				tmp := Coordinate{
					X: lastCoord.X + l,
					Y: lastCoord.Y,
				}
				steps++
				coords[tmp.String()] = steps
			}
			lastCoord.X = lastCoord.X + quantity

		case 'U':
			for l := 1; l <= quantity; l++ {
				tmp := Coordinate{
					X: lastCoord.X,
					Y: lastCoord.Y + l,
				}
				steps++
				coords[tmp.String()] = steps
			}
			lastCoord.Y = lastCoord.Y + quantity

		case 'D':
			for l := 1; l <= quantity; l++ {
				tmp := Coordinate{
					X: lastCoord.X,
					Y: lastCoord.Y - l,
				}
				steps++
				coords[tmp.String()] = steps
			}
			lastCoord.Y = lastCoord.Y - quantity
		}

	}

	return coords
}

// func addStep(coords map[string]int, x int, y int){
// 	tmp := Coordinate{
// 		X: x,
// 		Y: y,
// 	}
// 	steps++
// 	coords[tmp.String()] = steps
// }

//FindIntersections finds the intersections in left and right, not deduped
func FindIntersections(left []Coordinate, right []Coordinate) []Intersection {
	intersections := make([]Intersection, 0, 100)
	rightMap := make(map[string]int)
	for i, c := range right {
		rightMap[c.String()] = i
	}
	for li, l := range left {
		if rightsteps, ok := rightMap[l.String()]; ok {
			intersections = append(intersections, Intersection{
				CombinedSteps: li + rightsteps,
				Coordinate:    l,
			})
		}
	}
	return intersections
}

//FindIntersectionsQuadratic finds the intersections in left and right, not deduped
func FindIntersectionsQuadratic(left []Coordinate, right []Coordinate) []Intersection {
	intersections := make([]Intersection, 0, 100)
	rightMap := make(map[string]int)
	for i, c := range right {
		rightMap[c.String()] = i
	}
	for li, l := range left {
		for ri, r := range right {
			if l.Equals(r) {
				intersections = append(intersections, Intersection{
					CombinedSteps: li + ri,
					Coordinate:    l,
				})
			}
		}
	}
	return intersections
}

//FindIntersectionsMapBased the map based version
func FindIntersectionsMapBased(left map[string]int, right map[string]int) []Intersection {
	intersections := make([]Intersection, 0, 100)
	for key, leftsteps := range left {
		if rightsteps, ok := right[key]; ok {
			intersections = append(intersections, Intersection{
				CombinedSteps: leftsteps + rightsteps,
				Coordinate:    MakeCoordinate(key),
			})
		}
	}
	return intersections
}

//FindIntersectionsHybrid use a slice for the one you have to traverse, and a map for the one you look up, pass it in instead of calcing it in func
func FindIntersectionsHybrid(left []Coordinate, right *map[string]int) []Intersection {
	intersections := make([]Intersection, 0, 100)
	rightMap := *right
	for li, l := range left {
		if rightsteps, ok := rightMap[l.String()]; ok {
			intersections = append(intersections, Intersection{
				CombinedSteps: li + rightsteps,
				Coordinate:    l,
			})
		}
	}
	return intersections

}

//MinManhattenDist returns the lowest manhatten dist, not including 0, -1 in no nonzero values were found
func MinManhattenDist(locations []Intersection) (int, int) {
	min := -1
	minSteps := -1
	for _, v := range locations {
		localmd := v.ManhattenDist()
		if localmd == 0 {
			continue
		} else if min == -1 || localmd < min {
			min = localmd
		}
	}
	for _, v := range locations {
		localms := v.CombinedSteps
		if localms == 0 {
			continue
		} else if minSteps == -1 || localms < minSteps {
			minSteps = localms
		}
	}
	return min, minSteps
}
