package slope

//Coord represents an x, y coordinate
type Coord struct {
	X int
	Y int
}

//Equals returns true if this coord equals the other coord
func (c Coord) Equals(other Coord) bool {
	return c.X == other.X && c.Y == other.Y
}

//Fraction a struct to represent an exact fraction
type Fraction struct {
	Numerator   int
	Denominator int
}

//Equals true if the fractions are equal, false otherwise
func (f Fraction) Equals(compareTo Fraction) bool {
	//todo handle 0/0 and anumber/0
	myNum := f.Numerator
	myDen := f.Denominator

	yourNum := compareTo.Numerator
	yourDen := compareTo.Denominator

	if myNum == 0 && yourNum == 0 { //line y = 3 for example
		return true
	}
	if myDen == 0 && yourDen == 0 { //line x=3 for example
		return true
	}

	myNum = myNum * compareTo.Denominator
	myDen = myDen * compareTo.Denominator

	yourNum = yourNum * f.Denominator
	yourDen = yourDen * f.Denominator
	// fmt.Printf("mynum %d myDen %d yourNum %d yourDen %d", myNum, myDen, yourNum, yourDen)

	if myNum == yourNum && myDen == yourDen {
		return true
	}

	return false
}

//Slope calculates the slope from begin to end
func Slope(start Coord, end Coord) Fraction {
	// rise := end.Y - start.Y  they flipped the y axis, so a y value of 10 is above a y value of 12
	rise := start.Y - end.Y
	run := end.X - start.X
	return Fraction{Numerator: rise, Denominator: run}
}
