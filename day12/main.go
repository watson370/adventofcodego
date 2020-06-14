package main

import (
	"fmt"
)

type Position struct {
	X int
	Y int
	Z int
}
type Velocity struct {
	X int
	Y int
	Z int
}

// type Moon struct{
// 	Pos Position,
// 	Vel Velocity,
// }
func main() {
	// moonPositions := []Position{
	// 	Position{
	// 		X: -1,
	// 		Y: 0,
	// 		Z: 2,
	// 	},
	// 	Position{
	// 		X: 2,
	// 		Y: -10,
	// 		Z: -7,
	// 	},
	// 	Position{
	// 		X: 4,
	// 		Y: -8,
	// 		Z: 8,
	// 	},
	// 	Position{
	// 		X: 3,
	// 		Y: 5,
	// 		Z: -1,
	// 	},
	// }
	// moonVelocities := []Velocity{
	// 	Velocity{
	// 		X: 0,
	// 		Y: 0,
	// 		Z: 0,
	// 	},
	// 	Velocity{
	// 		X: 0,
	// 		Y: 0,
	// 		Z: 0,
	// 	},
	// 	Velocity{
	// 		X: 0,
	// 		Y: 0,
	// 		Z: 0,
	// 	},
	// 	Velocity{
	// 		X: 0,
	// 		Y: 0,
	// 		Z: 0,
	// 	},
	// }
	// calcAndPrint(moonPositions, moonVelocities)

	// moonPositions2 := []Position{
	// 	Position{
	// 		X: -8,
	// 		Y: -10,
	// 		Z: 0,
	// 	},
	// 	Position{
	// 		X: 5,
	// 		Y: 5,
	// 		Z: 10,
	// 	},
	// 	Position{
	// 		X: 2,
	// 		Y: -7,
	// 		Z: 3,
	// 	},
	// 	Position{
	// 		X: 9,
	// 		Y: -8,
	// 		Z: -3,
	// 	},
	// }
	// moonVelocities2 := []Velocity{
	// 	Velocity{
	// 		X: 0,
	// 		Y: 0,
	// 		Z: 0,
	// 	},
	// 	Velocity{
	// 		X: 0,
	// 		Y: 0,
	// 		Z: 0,
	// 	},
	// 	Velocity{
	// 		X: 0,
	// 		Y: 0,
	// 		Z: 0,
	// 	},
	// 	Velocity{
	// 		X: 0,
	// 		Y: 0,
	// 		Z: 0,
	// 	},
	// }
	// calcAndPrint(moonPositions2, moonVelocities2)

	//myinput
	// <x=6, y=10, z=10>
	// <x=-9, y=3, z=17>
	// <x=9, y=-4, z=14>
	// <x=4, y=14, z=4>
	moonPositions3 := []Position{
		Position{
			X: 6,
			Y: 10,
			Z: 10,
		},
		Position{
			X: -9,
			Y: 3,
			Z: 17,
		},
		Position{
			X: 9,
			Y: -4,
			Z: 14,
		},
		Position{
			X: 4,
			Y: 14,
			Z: 4,
		},
	}
	moonVelocities3 := []Velocity{
		Velocity{
			X: 0,
			Y: 0,
			Z: 0,
		},
		Velocity{
			X: 0,
			Y: 0,
			Z: 0,
		},
		Velocity{
			X: 0,
			Y: 0,
			Z: 0,
		},
		Velocity{
			X: 0,
			Y: 0,
			Z: 0,
		},
	}
	calcAndPrint(moonPositions3, moonVelocities3)
}

func calcAndPrint(moonPositionsIn []Position, moonVelocitiesIn []Velocity) {
	moonPositions := make([]Position, len(moonPositionsIn))
	moonVelocities := make([]Velocity, len(moonVelocitiesIn))
	copy(moonPositions, moonPositionsIn)
	copy(moonVelocities, moonVelocitiesIn)

	xFirstReturn, yFirstReturn, zFirstReturn := 0, 0, 0

	printTotalEnergies(moonPositions, moonVelocities, 0)
	for timestep := 1; xFirstReturn == 0 || yFirstReturn == 0 || zFirstReturn == 0; timestep++ {
		//apply gravity to each velocity
		for i := 0; i < len(moonPositions); i++ {
			for j := i + 1; j < len(moonPositions); j++ {
				if moonPositions[i].X < moonPositions[j].X {
					moonVelocities[i].X = moonVelocities[i].X + 1
					moonVelocities[j].X = moonVelocities[j].X - 1
				} else if moonPositions[i].X > moonPositions[j].X {
					moonVelocities[i].X = moonVelocities[i].X - 1
					moonVelocities[j].X = moonVelocities[j].X + 1
				}
				if moonPositions[i].Y < moonPositions[j].Y {
					moonVelocities[i].Y = moonVelocities[i].Y + 1
					moonVelocities[j].Y = moonVelocities[j].Y - 1
				} else if moonPositions[i].Y > moonPositions[j].Y {
					moonVelocities[i].Y = moonVelocities[i].Y - 1
					moonVelocities[j].Y = moonVelocities[j].Y + 1
				}
				if moonPositions[i].Z < moonPositions[j].Z {
					moonVelocities[i].Z = moonVelocities[i].Z + 1
					moonVelocities[j].Z = moonVelocities[j].Z - 1
				} else if moonPositions[i].Z > moonPositions[j].Z {
					moonVelocities[i].Z = moonVelocities[i].Z - 1
					moonVelocities[j].Z = moonVelocities[j].Z + 1
				}
			}
		}
		//apply velocity to position
		for i := range moonVelocities {
			moonPositions[i].X = moonPositions[i].X + moonVelocities[i].X
			moonPositions[i].Y = moonPositions[i].Y + moonVelocities[i].Y
			moonPositions[i].Z = moonPositions[i].Z + moonVelocities[i].Z
		}

		x, y, z := true, true, true
		for i := range moonVelocities {
			if xFirstReturn != 0 || moonVelocities[i].X != moonVelocitiesIn[i].X || moonPositions[i].X != moonPositionsIn[i].X {
				x = false
			}
			if yFirstReturn != 0 || moonVelocities[i].Y != moonVelocitiesIn[i].Y || moonPositions[i].Y != moonPositionsIn[i].Y {
				y = false
			}
			if zFirstReturn != 0 || moonVelocities[i].Z != moonVelocitiesIn[i].Z || moonPositions[i].Z != moonPositionsIn[i].Z {
				z = false
			}
		}
		if x {
			xFirstReturn = timestep
		}
		if y {
			yFirstReturn = timestep
		}
		if z {
			zFirstReturn = timestep
		}

		//for part two I had to look at the answer, they keep track of the number iterations until x returns to initial, y to initial, and z to initial independently, then calc the smallest number they are all factors of

		if timestep%10000000 == 0 {
			printTotalEnergies(moonPositions, moonVelocities, timestep)

		}
	}

	fmt.Printf("x %d  y %d  z %d\n", xFirstReturn, yFirstReturn, zFirstReturn)
	fmt.Printf("with a least common multiple of %d\n", leastCommonMultiple([]int{xFirstReturn, yFirstReturn, zFirstReturn}))

	//find the

}

// //pass in valueHolder with enough capacity to avoid allocation
// func calcState(pos []Position, vel []Velocity, valueHolder *[24]int) {
// 	// var smStorage [24]int16
// 	var currentPos int
// 	for _, v := range pos {

// 		valueHolder[currentPos] = v.X
// 		currentPos++
// 		valueHolder[currentPos] = v.Y
// 		currentPos++
// 		valueHolder[currentPos] = v.Z
// 		currentPos++
// 	}
// 	for _, v := range vel {
// 		valueHolder[currentPos] = v.X
// 		currentPos++
// 		valueHolder[currentPos] = v.Y
// 		currentPos++
// 		valueHolder[currentPos] = v.Z
// 		currentPos++
// 	}
// 	//fill into valueholder
// 	// for i := range smStorage { //0-23
// 	// 	arrPos := i / 4
// 	// 	inIntPos := i % 4 //0 will be first, 3 will be last
// 	// 	switch inIntPos {
// 	// 	case 0:
// 	// 		valueHolder[arrPos] = valueHolder[arrPos] | (int64(smStorage[i]) << 48)
// 	// 	case 1:
// 	// 		valueHolder[arrPos] = valueHolder[arrPos] | (int64(smStorage[i]) << 32)
// 	// 	case 2:
// 	// 		valueHolder[arrPos] = valueHolder[arrPos] | (int64(smStorage[i]) << 16)
// 	// 	case 3:
// 	// 		valueHolder[arrPos] = valueHolder[arrPos] | (int64(smStorage[i]))
// 	// 	}
// 	// }
// }

//compare two "states"
// func compare(left []int, right []int) int { //can I put a []int into a hash table?
// 	if len(left) != len(right) {
// 		panic("can only compare two states of equal len!")
// 	}
// 	for i := range left {
// 		if left[i] != right[i] {
// 			if left[i] < right[i] {
// 				return -1
// 			}
// 			return 1
// 		}
// 	}
// 	return 0
// }

func printTotalEnergies(pos []Position, vel []Velocity, timestep int) {
	sumOfTotalEnergies := 0
	for i := range pos {
		//potential
		potential := absOfInt(pos[i].X) + absOfInt(pos[i].Y) + absOfInt(pos[i].Z)
		//kinetic
		kinetic := absOfInt(vel[i].X) + absOfInt(vel[i].Y) + absOfInt(vel[i].Z)
		//total
		total := potential * kinetic
		sumOfTotalEnergies += total
		fmt.Printf("timestep %d planet index %d position <%d><%d><%d> velocity <%d><%d><%d> total energy %d\n", timestep, i, pos[i].X, pos[i].Y, pos[i].Z, vel[i].X, vel[i].Y, vel[i].Z, total)

	}
	fmt.Printf("timestep %d sum of energies for all planets %d\n", timestep, sumOfTotalEnergies)

}
func absOfInt(val int) int {
	if val < 0 {
		return -val
	}
	return val
}
func greatestCommonDenominator(left int, right int) int {
	// fmt.Printf("%d\t%d\n", left, right)
	for {
		left, right = right, left%right
		// fmt.Printf("%d\t%d\n", left, right)
		if right == 0 {
			// fmt.Printf("Returning %d\n", left)
			return left
		}
	}

}
func leastCommonMultiple(nums []int) int {
	running := 1
	for _, val := range nums {
		running = (running * val) / greatestCommonDenominator(running, val)
	}
	return running
}

// func leastCommonMultiple(left int, right int) int {
// 	return (left * right) / greatestCommonDenominator(left, right)
// }
