package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/watson370/adventofcodego/day10/slope"
)

//calc slope between source and target, calc slope between source and obstruction,
//if the slope is the same and x and y move in the same direction as the slope, it blocks

//make a list of the slope between source and all possible obstructions, check is slope of source to target is in the list
func main() {
	//why does the y coordinate value go up in a positive direction when you move down???
	processFile("firstSample.txt")
	processFile("sampleInput2.txt")
	processFile("sampleInput3.txt")
	processFile("sampleInput4.txt")
	processFile("myInput.txt")
	processFilePartTwo("sampleInput4.txt")
	processFilePartTwo("myInput.txt")
}

func processFile(file string) ([]slope.Coord, int) {

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	coords := make([]slope.Coord, 0, 100)
	for yVal := 0; s.Scan(); yVal++ {
		line := s.Text()
		if len(line) > 0 {
			for xVal, v := range s.Text() {
				if v == rune('#') {
					coords = append(coords, slope.Coord{X: xVal, Y: yVal})
				}
			}
		}
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	visibleCounts := make([]int, len(coords), len(coords))
	calcVis(coords, visibleCounts)
	//get coord with highest visibility
	highest := 0
	highestIndex := 0
	for i, v := range visibleCounts {
		if v > highest {
			highest = v
			highestIndex = i
		}
	}
	fmt.Printf("The highest count for file %s was %d for coord %v\n", file, visibleCounts[highestIndex], coords[highestIndex])
	return coords, highestIndex

}

func processFilePartTwo(file string) {

	coords, highestIndex := processFile(file)
	asteroidBlasterCounter := 0
	coordsToProcess := coords
	blasterStationIndex := highestIndex
	blasterStationCoord := coords[highestIndex]
	for {
		visibleCoords, blockedCoords := calcCoordVisFrom(coordsToProcess, blasterStationIndex)
		for c := range visibleCoords {
			asteroidBlasterCounter++
			fmt.Printf("Blasting asteroid number %d  coord %v\n", asteroidBlasterCounter, visibleCoords[c])
		}
		if len(blockedCoords) == 0 {
			break
		}
		coordsToProcess = append(blockedCoords, blasterStationCoord)
		blasterStationIndex = len(coordsToProcess) - 1
	}
}

//visible counts should be the same length as coords, the index into the passed in coords will be the index to the visibleCounts, passed in to avoid allocation,
func calcVis(coords []slope.Coord, visibleCounts []int) {
	for i := range coords {
		//make a list of slopes from coord i
		slopes := make(map[slope.Coord]slope.Fraction)
		for j := i + 1; j < len(coords); j++ {
			slopes[coords[j]] = slope.Slope(coords[i], coords[j])
		}
		//I want to know if there is a coord with the same slope as current calculated, and if there is then I want to know if that coord is between start and end
		for k := i + 1; k < len(coords); k++ {
			slopeToMatch := slope.Slope(coords[i], coords[k])
			blocker := make([]slope.Coord, 0, 5)
			for s := range slopes {
				if slopes[s].Equals(slopeToMatch) && !s.Equals(coords[k]) { //the slope matches and it is not the same coord as the one we are trying to test
					blocker = append(blocker, s)
				}
			}
			blocked := false
			for blockerIndex := range blocker {
				//has same slope, now test that it is in between start and end
				if (coords[i].X <= blocker[blockerIndex].X && blocker[blockerIndex].X <= coords[k].X ||
					coords[i].X >= blocker[blockerIndex].X && blocker[blockerIndex].X >= coords[k].X) &&
					(coords[i].Y <= blocker[blockerIndex].Y && blocker[blockerIndex].Y <= coords[k].Y ||
						coords[i].Y >= blocker[blockerIndex].Y && blocker[blockerIndex].Y >= coords[k].Y) {
					blocked = true
				}
			}
			if !blocked {
				visibleCounts[i] = visibleCounts[i] + 1
				visibleCounts[k] = visibleCounts[k] + 1
			}
		}
	}
}

//coordIndex int is the index of the coord in coords, returns []coord that are visible and [] coord blocked
func calcCoordVisFrom(coords []slope.Coord, coordIndex int) ([]slope.Coord, []slope.Coord) {
	visibleCoords := make([]slope.Coord, 0, len(coords)) //can't pass in because we append to it
	blockedCoords := make([]slope.Coord, 0, len(coords))

	//make a list of slopes from coord i
	slopes := make(map[slope.Coord]slope.Fraction)
	for j := range coords {
		if j != coordIndex {
			slopes[coords[j]] = slope.Slope(coords[coordIndex], coords[j])
		}
	}
	//I want to know if there is a coord with the same slope as current calculated, and if there is then I want to know if that coord is between start and end
	for k := range coords {
		if k == coordIndex { //for each coord other than start
			continue
		}
		slopeToMatch := slope.Slope(coords[coordIndex], coords[k])
		blocker := make([]slope.Coord, 0, 5)

		for s := range slopes {
			if slopes[s].Equals(slopeToMatch) && !s.Equals(coords[k]) { //the slope matches and it is not the same coord as the one we are trying to test
				blocker = append(blocker, s)
			}
		}
		blocked := false
		for blockerIndex := range blocker {
			//has same slope, now test that it is in between start and end
			if (coords[coordIndex].X <= blocker[blockerIndex].X && blocker[blockerIndex].X <= coords[k].X ||
				coords[coordIndex].X >= blocker[blockerIndex].X && blocker[blockerIndex].X >= coords[k].X) &&
				(coords[coordIndex].Y <= blocker[blockerIndex].Y && blocker[blockerIndex].Y <= coords[k].Y ||
					coords[coordIndex].Y >= blocker[blockerIndex].Y && blocker[blockerIndex].Y >= coords[k].Y) {
				blocked = true
			}
		}
		if blocked {
			blockedCoords = append(blockedCoords, coords[k])
		} else {
			visibleCoords = append(visibleCoords, coords[k])
		}
	}
	//sort by degrees from vertical using slope,
	sort.Slice(visibleCoords, func(i, j int) bool {
		//compare slope from coords[coordIndex] to visibleCoords[i] and visibleCoords[j]
		//return true is i is less than j
		//based on angle from coords[coordIndex]
		slopeI := slope.Slope(coords[coordIndex], visibleCoords[i])
		slopeJ := slope.Slope(coords[coordIndex], visibleCoords[j])
		if slopeI.Equals(slopeJ) {
			return false
		}
		c1, s1 := specialCase(slopeI)
		c2, s2 := specialCase(slopeJ)
		if s1 || s2 || c1 != c2 {
			return c1 < c2
		}
		//same quadrant, compare fractions,
		inum := slopeI.Numerator * slopeJ.Denominator
		if inum < 0 {
			inum = -inum
		}
		jnum := slopeJ.Numerator * slopeI.Denominator
		if jnum < 0 {
			jnum = -jnum
		}
		switch c1 {
		case 45, 225:
			return inum > jnum
		case 135, 315:
			return inum < jnum
		}
		panic("How did we get here?")
	})
	return visibleCoords, blockedCoords
}

//is it worth a func call?
func less(start slope.Coord, left slope.Coord, right slope.Coord) bool {
	slopeI := slope.Slope(start, left)
	slopeJ := slope.Slope(start, right)
	if slopeI.Equals(slopeJ) {
		fmt.Printf("starting coord %v, I coord %v calculated slope %v j coord %v calculated slope %v LESS %v  got EQUAL\n", start, left, slopeI, right, slopeJ, false)
		return false
	}
	c1, s1 := specialCase(slopeI)
	c2, s2 := specialCase(slopeJ)
	if s1 || s2 || c1 != c2 {
		fmt.Printf("starting coord %v, I coord %v calculated slope %v j coord %v calculated slope %v LESS %v  special case or quadrant\n", start, left, slopeI, right, slopeJ, c1 < c2)

		return c1 < c2
	}
	//same quadrant, compare fractions,
	inum := slopeI.Numerator * slopeJ.Denominator
	if inum < 0 {
		inum = -inum
	}
	jnum := slopeJ.Numerator * slopeI.Denominator
	if jnum < 0 {
		jnum = -jnum
	}
	switch c1 {
	case 45, 225:
		fmt.Printf("starting coord %v, I coord %v calculated slope %v j coord %v calculated slope %v LESS %v  quadrant 1 or 3\n", start, left, slopeI, right, slopeJ, inum > jnum)

		return inum > jnum
	case 135, 315:
		fmt.Printf("starting coord %v, I coord %v calculated slope %v j coord %v calculated slope %v LESS %v  quadrant 2 or 4\n", start, left, slopeI, right, slopeJ, inum < jnum)

		return inum < jnum
	}
	panic("should not get here")
}

//returns the quadrant(degrees to the middle of it), or if a special case the degree and true
//-1 is none,
//0 straight up,
//90 straight right
//180 straight down
//270 straight left
//bool is true if a special case, assumes not comparing to itself, slope 0/0
func specialCase(f slope.Fraction) (int, bool) {
	switch {
	case f.Denominator == 0 && f.Numerator > 0:
		return 0, true
	case f.Numerator == 0 && f.Denominator > 0:
		return 90, true
	case f.Denominator == 0 && f.Numerator < 0:
		return 180, true
	case f.Numerator == 0 && f.Denominator < 0:
		return 270, true
		//q1
	case f.Numerator > 0 && f.Denominator > 0:
		return 45, false
		//q2
	case f.Numerator < 0 && f.Denominator > 0:
		return 135, false
		//q3
	case f.Numerator < 0 && f.Denominator < 0:
		return 225, false
		//q4
	case f.Numerator > 0 && f.Denominator < 0:
		return 315, false
	default:
		panic("this should never happen, all cases covered")

	}
}
