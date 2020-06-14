package main

import (
	"fmt"
	"io"
	"os"
)

func count(s []byte, match byte) int {
	count := int(0)
	for i := 0; i < len(s); i++ {
		if s[i] == match {
			count++
		}
	}
	return count
}

func main() {
	const NumPixels = 25 * 6
	const BufferSize = NumPixels

	file, err := os.Open("myinput.txt")
	if err != nil {
		panic(err)
	}

	var image [][]byte = make([][]byte, 0, 100)
	layer := 0

	for {
		image = append(image, make([]byte, NumPixels, NumPixels))

		n, err := file.Read(image[layer])

		if n == 0 {
			image = image[:len(image)-1]
		}
		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}
		layer++
	}
	fmt.Printf("len of image %d, len of layer %d\n", len(image), len(image[0]))
	//find the layer with the LEAST zeroes
	leastZerosLayerIndex := 0
	leastZeros := NumPixels // number of zero pixel values can't exceed number of pixel values
	for i, layer := range image {
		currentZeroes := count(layer, '0')
		if currentZeroes < leastZeros {
			leastZeros = currentZeroes
			leastZerosLayerIndex = i
		}
	}
	numberOfOnes := count(image[leastZerosLayerIndex], '1')
	numberOfTwos := count(image[leastZerosLayerIndex], '2')
	fmt.Printf("\nThe layer with the LEAST zeroes was layer %d, it had %d ones and %d twos, %d times %d equals %d\n",
		leastZerosLayerIndex, numberOfOnes, numberOfTwos, numberOfOnes, numberOfTwos, numberOfOnes*numberOfTwos)

	//0 is black, 1 is white, and 2 is transparent
	finalImage := make([]byte, 150, 150)
	for i := range finalImage {
		finalImage[i] = '2'
	}
	for _, l := range image {
		update(finalImage, l)
	}
	//reshape to [6][25]
	reshaped := make([][]byte, 6, 6)
	for i := 0; i < 6; i++ {
		reshaped[i] = make([]byte, 25, 25)
		copy(reshaped[i], finalImage[(i*25):((i+1)*25)])
	}

	for _, rowvals := range reshaped {
		for _, v := range rowvals {
			switch v {
			case '0':
				fmt.Print(" ")
			case '1':
				fmt.Print("W")
			case '2':
				fmt.Print("T")
			}
		}
		fmt.Println()
	}

}

//updates the current with vals from update if the current val is transparent, i.e. '2'
func update(current []byte, update []byte) {
	for i, v := range current {
		if v == '2' {
			current[i] = update[i]
		}
	}
}
