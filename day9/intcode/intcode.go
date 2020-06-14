package intcode

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

//Process   1, 2, 99, or error
//ints must be large enough to hold itself while processing, maybe just increase size to largest number?will it always be sparse? create new data type?Expand on demand here?
func Process(ints []int64) ([]int64, error) {
	set := make(map[int]string)
	set[1] = "addition"
	set[2] = "multiplication"
	set[3] = "input"
	set[4] = "output"
	set[5] = "jump-if-true"
	set[6] = "jump-if-false"
	set[7] = "less-than"
	set[8] = "equal"
	set[9] = "relativeBaseOffset"
	set[99] = "terminate"
	instructionpointer := int64(0) //pointer to current memory address
	instruction := ""              //instruction
	ok := true
	firstoperand := int64(0)
	secondoperand := int64(0)
	result := int64(0)
	resultpos := int64(0)
	var opcodeandparammodes []int64
	var relativeBase int64
	pass := 0

	for {
		opcodeandparammodes = SplitOpCode(ints[instructionpointer], make([]int64, 4))

		instruction, ok = set[int(opcodeandparammodes[0])]
		if !ok {
			return ints, errors.New("unexpected value " + strconv.FormatInt(ints[instructionpointer], 10) + " at position " + strconv.FormatInt(instructionpointer, 10))
		}

		//perform instruction
		switch instruction {
		case "addition":
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1, relativeBase)  //ints[params[0]]
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2, relativeBase) //ints[params[1]]
			resultpos = GetResultPos(ints, opcodeandparammodes, instructionpointer, 3, relativeBase)
			// resultpos = ints[instructionpointer+3]                                                 // params[2]
			result = firstoperand + secondoperand
			// fmt.Printf("I'm in addition, first op %d, second op %d, res pos %d result %d\n", firstoperand, secondoperand, resultpos, result)
			ints = checkSize(ints, resultpos)
			ints[resultpos] = result
			instructionpointer += 4
		case "multiplication":
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1, relativeBase)    //ints[params[0]]
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2, relativeBase)   //ints[params[1]]
			resultpos = GetResultPos(ints, opcodeandparammodes, instructionpointer, 3, relativeBase) // params[2]
			result = firstoperand * secondoperand
			ints = checkSize(ints, resultpos)
			ints[resultpos] = result
			instructionpointer += 4
		case "input":
			resultpos = GetResultPos(ints, opcodeandparammodes, instructionpointer, 1, relativeBase)
			result = GetInput()
			ints = checkSize(ints, resultpos)
			ints[resultpos] = result
			instructionpointer += 2
		case "output":
			result = GetVal(ints, opcodeandparammodes, instructionpointer, 1, relativeBase)
			OutputResult(result)
			instructionpointer += 2
		case "jump-if-true":
			// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1, relativeBase)
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2, relativeBase)
			if firstoperand != 0 {
				// ints = checkSize(ints, secondoperand)
				instructionpointer = secondoperand
			} else {
				// ints = checkSize(ints, instructionpointer + 3)
				instructionpointer += 3
			}
		case "jump-if-false":
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1, relativeBase)
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2, relativeBase)
			if firstoperand == 0 {
				// ints = checkSize(ints, secondoperand)
				instructionpointer = secondoperand
			} else {
				// ints = checkSize(ints, instructionpointer + 3)
				instructionpointer += 3
			}
		case "less-than":
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1, relativeBase)
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2, relativeBase)
			resultpos = GetResultPos(ints, opcodeandparammodes, instructionpointer, 3, relativeBase)
			ints = checkSize(ints, resultpos)
			if firstoperand < secondoperand {
				ints[resultpos] = 1
			} else {
				ints[resultpos] = 0
			}
			instructionpointer += 4
		case "equal":
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1, relativeBase)
			secondoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 2, relativeBase)
			resultpos = GetResultPos(ints, opcodeandparammodes, instructionpointer, 3, relativeBase)
			ints = checkSize(ints, resultpos)
			if firstoperand == secondoperand {
				ints[resultpos] = 1
			} else {
				ints[resultpos] = 0
			}
			instructionpointer += 4
		case "relativeBaseOffset":
			//takes a single parameter and adjusts the relative base by it
			firstoperand = GetVal(ints, opcodeandparammodes, instructionpointer, 1, relativeBase)
			relativeBase = relativeBase + firstoperand
			instructionpointer += 2
		case "terminate":
			return ints, nil
		default:
			return ints, errors.New("unexpected instruction value " + instruction)
		}
		// fmt.Printf("pass %d offset %d intcode %v\n", pass, relativeBase, ints)
		pass++
	}
}

//SplitOpCode takes in the integer opcode and return an int slice with the opcode in position 0, and the parameter modes in the following positions
//result must be atleast length 4
func SplitOpCode(opcode int64, result []int64) []int64 {
	strval := strconv.FormatInt(opcode, 10)
	// toreturn := make([]int64, 4)
	if len(strval) == 1 {
		opcode, _ = strconv.ParseInt(strval, 10, 64)
	} else {
		opcode, _ = strconv.ParseInt(strval[len(strval)-2:], 10, 64)
	}
	result[0] = opcode
	//i index to write to r index to read from
	//i starts at 1, 0 is the opcode, r -1 for zero based, -2 for the opcode
	for i, r := 1, len(strval)-2-1; r >= 0; i, r = i+1, r-1 {
		intval, _ := strconv.ParseInt(string(strval[r]), 10, 64)
		result[i] = intval
	}
	return result
}

//GetVal takes in the []int of the program, the []int that holds the parameter input modes, the index of the current instruction, and the position requested, then returns the correct value based on the mode
//the position requested is 1, 2, or 3, not index based, but the number of the param, but turns out to be the index in the opcodeandparammodes
func GetVal(ints []int64, opcodeandparammodes []int64, instructionpointer int64, position int64, relativeBase int64) int64 {
	mode := opcodeandparammodes[position]
	position = getPosCommon(ints, mode, instructionpointer, position, relativeBase)
	if int64(len(ints)-1) < position {
		return 0 //pretend it is there, but don't allocate unless writing
	}
	return ints[position]
}

func getPosCommon(ints []int64, mode int64, instructionpointer int64, paramPosition int64, relativeBase int64) int64 {
	//todo verify valid mode

	position := instructionpointer + paramPosition
	if mode == 0 { //position mode
		position = ints[position]
	} else if mode == 2 { //relative mode
		// fmt.Println("Using relative base mode")
		position = ints[position] + relativeBase
	}

	return position

}

//GetResultPos basically GetVal, but result position will never be in immediate mode
func GetResultPos(ints []int64, opcodeandparammodes []int64, instructionpointer int64, position int64, relativeBase int64) int64 {
	mode := opcodeandparammodes[position]
	if mode == 1 { //default immediate mode to position mode
		mode = 0
	}
	return getPosCommon(ints, mode, instructionpointer, position, relativeBase)
}

//GetInput asks the user for the ID of the system to test by running an input instruction
func GetInput() int64 {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter the ID of the system to test or run: ")
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		res, e := strconv.ParseInt(text, 10, 64)
		// res, e := strconv.Atoi(text)
		if e != nil {
			fmt.Printf("The system id is an int, for example 1 for the air conditioning system. %s is not valid input", text)
		} else {
			//close scanner ?
			return res
		}

	}
}

// OutputResult oututs the int to standard out
func OutputResult(result int64) {
	outputstring := strconv.FormatInt(result, 10)
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, outputstring)
	fmt.Fprint(w, "\n")
	w.Flush()
}

//resize will cause allocation, but we can't expect the user to know the max size needed, so allocaiton will be necessary
func checkSize(ints []int64, size int64) []int64 {
	if int64(len(ints)-1) < size {
		toReturn := make([]int64, size+1, size+1)
		copy(toReturn, ints)
		return toReturn
	}
	return ints
}
