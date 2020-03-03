package day3

import (
	"strconv"
	"strings"
)

//Coordinate represents an x y coordinate
type Coordinate struct {
	X int
	Y int
}

//ManhattenDist returns the manhatten distance of this coordinate from 0, 0, does not account for values large enough to overflow int
func (c Coordinate) ManhattenDist() int {
	var horizontal int
	var vertical int
	if c.X < 0 {
		horizontal = -c.X
	} else {
		horizontal = c.X
	}
	if c.Y < 0 {
		vertical = -c.Y
	} else {
		vertical = c.Y
	}
	return horizontal + vertical
}

//Equals returns true in in equals this cooord
func (c Coordinate) Equals(in Coordinate) bool {
	if c.X == in.X && c.Y == in.Y {
		return true
	}
	return false
}

//String returns a string representation
func (c Coordinate) String() string {
	return strconv.Itoa(c.X) + "," + strconv.Itoa(c.Y)
}

//MakeCoordinate takes in the strig rep of a coord, and returns a coord
func MakeCoordinate(s string) Coordinate {
	vals := strings.Split(s, ",")
	x, _ := strconv.Atoi(vals[0])
	y, _ := strconv.Atoi(vals[1])
	return Coordinate{
		X: x,
		Y: y,
	}
}

//Intersection a coordinate and the combined steps to get to that coordinate
type Intersection struct {
	CombinedSteps int
	Coordinate
}

//Intersection a coordinate and the combined steps to get to that coordinate
type IntersectionAttributes struct {
	CombinedSteps     int
	ManhattenDistance int
}
