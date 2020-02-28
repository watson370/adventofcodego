package day3

import (
	"log"
	"strconv"
	"testing"
)

func TestIntcode(t *testing.T) {
	// 	R75,D30,R83,U83,L12,D49,R71,U7,L72
	// U62,R66,U55,R34,D71,R55,D58,R83 = distance 159

	// R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51
	// U98,R91,D20,R16,D67,R40,U7,R15,U6,R7 = distance 135
	inputs := []WirePairInput{
		WirePairInput{
			Wire1:               "R75,D30,R83,U83,L12,D49,R71,U7,L72",
			Wire2:               "U62,R66,U55,R34,D71,R55,D58,R83",
			ExpectedResult:      159,
			ExpectedStepsResult: 610,
		},
		WirePairInput{
			Wire1:               "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51",
			Wire2:               "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			ExpectedResult:      135,
			ExpectedStepsResult: 410,
		},
	}

	for _, wp := range inputs {
		wp1coords := BuildCoordList(wp.Wire1)
		wp2coords := BuildCoordList(wp.Wire2)
		// log.Println(wp1coords)
		// log.Println("\n\n ---------------")
		// log.Println(wp2coords)
		intersections := FindIntersections(wp1coords, wp2coords)
		min, minSteps := MinManhattenDist(intersections)
		log.Println(intersections)
		log.Printf("The minimum manhatten dist is %d and the min steps is %d", min, minSteps)
		if wp.ExpectedResult != min {
			t.Error("was expecting " + strconv.Itoa(wp.ExpectedResult) + " but got " + strconv.Itoa(min))
		}
		if wp.ExpectedStepsResult != minSteps {
			t.Error("was expecting " + strconv.Itoa(wp.ExpectedStepsResult) + " but got " + strconv.Itoa(minSteps))
		}
	}
}

type WirePairInput struct {
	Wire1               string
	Wire2               string
	ExpectedResult      int
	ExpectedStepsResult int
}
