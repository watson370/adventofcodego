package main

import (
	"fmt"

	"github.com/watson370/adventofcodego/day10/slope"
	"github.com/watson370/adventofcodego/day11/intcode"
)

func main() {
	var myInput []int64 = []int64{3, 8, 1005, 8, 311, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 102, 1, 8, 28, 1, 1104, 0, 10, 1006, 0, 71, 2, 1002, 5, 10, 2, 1008, 5, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 102, 1, 8, 66, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 1, 8, 10, 4, 10, 102, 1, 8, 87, 1006, 0, 97, 2, 1002, 6, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 102, 1, 8, 116, 1006, 0, 95, 1, 1009, 10, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 1, 8, 10, 4, 10, 102, 1, 8, 145, 1, 1002, 19, 10, 2, 1109, 7, 10, 1006, 0, 18, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1001, 8, 0, 179, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 102, 1, 8, 200, 1, 1105, 14, 10, 1, 1109, 14, 10, 2, 1109, 11, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 102, 1, 8, 235, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1002, 8, 1, 257, 2, 101, 9, 10, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 108, 0, 8, 10, 4, 10, 101, 0, 8, 282, 2, 1109, 19, 10, 1, 105, 0, 10, 101, 1, 9, 9, 1007, 9, 1033, 10, 1005, 10, 15, 99, 109, 633, 104, 0, 104, 1, 21102, 937268368140, 1, 1, 21102, 328, 1, 0, 1106, 0, 432, 21102, 1, 932700599052, 1, 21101, 0, 339, 0, 1105, 1, 432, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21101, 0, 209421601831, 1, 21102, 1, 386, 0, 1106, 0, 432, 21102, 235173604443, 1, 1, 21102, 1, 397, 0, 1106, 0, 432, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21101, 825439855372, 0, 1, 21102, 1, 420, 0, 1106, 0, 432, 21101, 0, 988220907880, 1, 21102, 431, 1, 0, 1106, 0, 432, 99, 109, 2, 22101, 0, -1, 1, 21101, 40, 0, 2, 21102, 1, 463, 3, 21102, 453, 1, 0, 1106, 0, 496, 109, -2, 2105, 1, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 458, 459, 474, 4, 0, 1001, 458, 1, 458, 108, 4, 458, 10, 1006, 10, 490, 1102, 1, 0, 458, 109, -2, 2106, 0, 0, 0, 109, 4, 2102, 1, -1, 495, 1207, -3, 0, 10, 1006, 10, 513, 21102, 0, 1, -3, 22102, 1, -3, 1, 21202, -2, 1, 2, 21102, 1, 1, 3, 21101, 532, 0, 0, 1105, 1, 537, 109, -4, 2105, 1, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 560, 2207, -4, -2, 10, 1006, 10, 560, 21201, -4, 0, -4, 1106, 0, 628, 22102, 1, -4, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21102, 1, 579, 0, 1106, 0, 537, 21202, 1, 1, -4, 21102, 1, 1, -1, 2207, -4, -2, 10, 1006, 10, 598, 21101, 0, 0, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 620, 21201, -1, 0, 1, 21102, 1, 620, 0, 105, 1, 495, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2105, 1, 0}

	panelColors := make(map[slope.Coord]int)
	currentPos := slope.Coord{
		X: 0,
		Y: 0,
	}
	//paint first panel white
	panelColors[currentPos] = 1
	currentDirection := 0 //0 up,1 right, 2 down, 3 left

	hullPainter := intcode.NewHullPainter(myInput)
	for !hullPainter.HasHalted() {
		hullPainter.Process()
		// if hullPainter.RequestsInput() { this program should always be input output input output
		if color, ok := panelColors[currentPos]; ok {
			hullPainter.GiveInput(int64(color))
		} else {
			hullPainter.GiveInput(0) //all panels start off black
		}
		hullPainter.Process()
		// }
		output := hullPainter.GetOutput() //0 is color and 1 is direction
		panelColors[currentPos] = int(output[0])
		//0 is left, 1 is right
		if output[1] == 0 { //+3
			currentDirection = (currentDirection + 3) % 4
		} else if output[1] == 1 { //+1
			currentDirection = (currentDirection + 1) % 4
		}
		//move in that direction
		switch currentDirection {
		case 0:
			currentPos = slope.Coord{
				X: currentPos.X,
				Y: currentPos.Y + 1,
			}
		case 1:
			currentPos = slope.Coord{
				X: currentPos.X + 1,
				Y: currentPos.Y,
			}
		case 2:
			currentPos = slope.Coord{
				X: currentPos.X,
				Y: currentPos.Y - 1,
			}
		case 3:
			currentPos = slope.Coord{
				X: currentPos.X - 1,
				Y: currentPos.Y,
			}
		}
	}
	// for coord, value := range panelColors {
	// 	color := "black"
	// 	if value == 1 {
	// 		color = "white"
	// 	}
	// 	fmt.Printf("panel %d, %d is the color %s\n", coord.X, coord.Y, color)
	// }
	fmt.Printf("the number of panels painted is %d \n", len(panelColors))
	minx, maxx, miny, maxy := 0, 0, 0, 0
	for c := range panelColors {
		if c.X > maxx {
			maxx = c.X
		}
		if c.X < minx {
			minx = c.X
		}
		if c.Y < miny {
			miny = c.Y
		}
		if c.Y > maxy {
			maxy = c.Y
		}
	}

	reparr := make([][]int, maxy-miny+1, maxy-miny+1)
	for i := range reparr {
		reparr[i] = make([]int, maxx-minx+1, maxx-minx+1)
	}
	for c := range panelColors {
		reparr[c.Y-miny][c.X-minx] = panelColors[c]
	}
	lines := make([]string, 0, 100)
	for i := range reparr {
		line := ""
		for j := range reparr[i] {
			if reparr[i][j] == 1 {
				line += "#"
			} else {
				line += " "
			}
		}
		lines = append(lines, line)
	}
	//print highest y values on top
	for i := len(lines) - 1; i >= 0; i-- {
		fmt.Println(lines[i])
	}

}
