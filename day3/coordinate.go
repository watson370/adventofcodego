package day3

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

//Intersection a coordinate and the combined steps to get to that coordinate
type Intersection struct {
	CombinedSteps int
	Coordinate
}
