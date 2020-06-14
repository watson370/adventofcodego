package intcode

import (
	"errors"
	"strconv"
)

//HullPainter a hull painter use pointer semantics
type HullPainter struct {
	intcode            []int64
	instructionpointer int64
	relativeBase       int64
	terminated         bool
	output             []int64
	input              int64
	validInput         bool
	needsInput         bool
	err                error
}

//HasHalted if the intcode has finished running
func (state *HullPainter) HasHalted() bool {
	return state.terminated
}

//RequestsInput if the program needs input, provide by calling GiveInput
func (state *HullPainter) RequestsInput() bool {
	return state.needsInput
}

//GiveInput give the program an input
func (state *HullPainter) GiveInput(input int64) {
	state.input = input
	state.validInput = true
}

//GetOutput gets and CLEARS the output,
func (state *HullPainter) GetOutput() []int64 {
	toReturn := state.output
	state.output = make([]int64, 0, 0)
	return toReturn
}

//NewHullPainter a factory function for a new hull painter
func NewHullPainter(intcodeVal []int64) HullPainter {
	return HullPainter{
		intcode: intcodeVal,
		output:  make([]int64, 0, 0),
	}
}

//Process   1, 2, 99, or error
//ints must be large enough to hold itself while processing, maybe just increase size to largest number?will it always be sparse? create new data type?Expand on demand here?
func (state *HullPainter) Process() error {
	if state.terminated {
		return nil
	}
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
	// instructionpointer := int64(0) //pointer to current memory address
	instruction := "" //instruction
	ok := true
	var firstoperand int64
	var secondoperand int64
	var result int64
	var resultpos int64
	var opcodeandparammodes []int64

	for {
		opcodeandparammodes = SplitOpCode(state.intcode[state.instructionpointer], make([]int64, 4))

		instruction, ok = set[int(opcodeandparammodes[0])]
		if !ok {
			return errors.New("unexpected value " + strconv.FormatInt(state.intcode[state.instructionpointer], 10) + " at position " + strconv.FormatInt(state.instructionpointer, 10))
		}

		//perform instruction
		switch instruction {
		case "addition":
			firstoperand = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 1, state.relativeBase)  //ints[params[0]]
			secondoperand = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 2, state.relativeBase) //ints[params[1]]
			resultpos = GetResultPos(state.intcode, opcodeandparammodes, state.instructionpointer, 3, state.relativeBase)
			result = firstoperand + secondoperand
			state.intcode = checkSize(state.intcode, resultpos)
			state.intcode[resultpos] = result
			state.instructionpointer += 4
		case "multiplication":
			firstoperand = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 1, state.relativeBase)    //ints[params[0]]
			secondoperand = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 2, state.relativeBase)   //ints[params[1]]
			resultpos = GetResultPos(state.intcode, opcodeandparammodes, state.instructionpointer, 3, state.relativeBase) // params[2]
			result = firstoperand * secondoperand
			state.intcode = checkSize(state.intcode, resultpos)
			state.intcode[resultpos] = result
			state.instructionpointer += 4
		case "input":
			if !state.validInput {
				state.needsInput = true
				return nil //next call to procees shows up here
			}
			resultpos = GetResultPos(state.intcode, opcodeandparammodes, state.instructionpointer, 1, state.relativeBase)
			result = state.input
			state.validInput = false
			state.intcode = checkSize(state.intcode, resultpos)
			state.intcode[resultpos] = result
			state.instructionpointer += 2
		case "output":
			result = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 1, state.relativeBase)
			state.output = append(state.output, result)
			state.instructionpointer += 2
		case "jump-if-true":
			// Opcode 5 is jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			firstoperand = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 1, state.relativeBase)
			secondoperand = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 2, state.relativeBase)
			if firstoperand != 0 {
				state.instructionpointer = secondoperand
			} else {
				state.instructionpointer += 3
			}
		case "jump-if-false":
			firstoperand = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 1, state.relativeBase)
			secondoperand = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 2, state.relativeBase)
			if firstoperand == 0 {
				state.instructionpointer = secondoperand
			} else {
				state.instructionpointer += 3
			}
		case "less-than":
			firstoperand = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 1, state.relativeBase)
			secondoperand = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 2, state.relativeBase)
			resultpos = GetResultPos(state.intcode, opcodeandparammodes, state.instructionpointer, 3, state.relativeBase)
			state.intcode = checkSize(state.intcode, resultpos)
			if firstoperand < secondoperand {
				state.intcode[resultpos] = 1
			} else {
				state.intcode[resultpos] = 0
			}
			state.instructionpointer += 4
		case "equal":
			firstoperand = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 1, state.relativeBase)
			secondoperand = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 2, state.relativeBase)
			resultpos = GetResultPos(state.intcode, opcodeandparammodes, state.instructionpointer, 3, state.relativeBase)
			state.intcode = checkSize(state.intcode, resultpos)
			if firstoperand == secondoperand {
				state.intcode[resultpos] = 1
			} else {
				state.intcode[resultpos] = 0
			}
			state.instructionpointer += 4
		case "relativeBaseOffset":
			//takes a single parameter and adjusts the relative base by it
			firstoperand = GetVal(state.intcode, opcodeandparammodes, state.instructionpointer, 1, state.relativeBase)
			state.relativeBase = state.relativeBase + firstoperand
			state.instructionpointer += 2
		case "terminate":
			state.terminated = true
			return nil
		default:
			return errors.New("unexpected instruction value " + instruction)
		}
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

//resize will cause allocation, but we can't expect the user to know the max size needed, so allocaiton will be necessary
func checkSize(ints []int64, size int64) []int64 {
	if int64(len(ints)-1) < size {
		toReturn := make([]int64, size+1, size+1)
		copy(toReturn, ints)
		return toReturn
	}
	return ints
}
