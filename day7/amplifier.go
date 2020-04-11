package day7

import (
	"log"
)

func RunPartTwo(intcodeprogram []int, phases []int) int {
	chin := make(chan int, 1)
	chout := make(chan int, 1)
	chab := make(chan int, 1)
	chbc := make(chan int, 1)
	chcd := make(chan int, 1)
	chde := make(chan int, 1)
	chlast := make(chan int)
	//create a sixth routine that just pases the value from chout back into chin, and upon closing writes the last value to an answer channel
	go ProcessPartTwo(chin, chab, intcodeprogram, 0)
	go ProcessPartTwo(chab, chbc, intcodeprogram, 1)
	go ProcessPartTwo(chbc, chcd, intcodeprogram, 2)
	go ProcessPartTwo(chcd, chde, intcodeprogram, 3)
	go ProcessPartTwo(chde, chout, intcodeprogram, 4)
	//setup the loop back
	go func(chout chan int, chin chan int, chlast chan int) {
		last := 0
		for {
			tmp, ok := <-chout
			if !ok {
				log.Printf("stoping loopback, tmp %d ok %v\n", tmp, ok)
				close(chin)
				chlast <- last
				close(chlast)
				return
			}
			log.Printf("looping back value %d\n", tmp)
			last = tmp
			chin <- last
			log.Printf("value %d was looped\n", tmp)
		}
	}(chout, chin, chlast)

	//use the channels to input the phases
	// log.Println("SSSSSSSSSSSSSSStarting initial vals")
	go func(chin chan int, chab chan int, chbc chan int, chcd chan int, chde chan int, phases []int) {
		// log.Printf("putting %d on chin", phases[0])
		chin <- phases[0]
		// log.Printf("putting %d on chab\n", phases[1])
		chab <- phases[1]
		// log.Printf("putting %d on chbc\n", phases[2])
		chbc <- phases[2]
		// log.Printf("putting %d on chcd\n", phases[3])
		chcd <- phases[3]
		// log.Printf("putting %d on chde\n", phases[4])
		chde <- phases[4]
		//they should all be blocking waiting for input
		// log.Println("adding initial 0 to start program")
		chin <- 0
	}(chin, chab, chbc, chcd, chde, phases)

	result := <-chlast
	log.Printf("I got a final answer of %d for phases %v", result, phases)
	return result

}

func Setup(intcodeprogram []int) (chan func() (int, []int), chan func() (int, []int)) {

	//create nodes a through e
	//create input channel, output channel, and intermediate channels
	chin := make(chan func() (int, []int))
	chout := make(chan func() (int, []int))
	chab := make(chan func() (int, []int))
	chbc := make(chan func() (int, []int))
	chcd := make(chan func() (int, []int))
	chde := make(chan func() (int, []int))

	go Amplify(chin, chab, 0, intcodeprogram)
	go Amplify(chab, chbc, 1, intcodeprogram)
	go Amplify(chbc, chcd, 2, intcodeprogram)
	go Amplify(chcd, chde, 3, intcodeprogram)
	go Amplify(chde, chout, 4, intcodeprogram)
	return chin, chout

}
func Run(chin chan func() (int, []int), chout chan func() (int, []int), phases []int) int {
	//only input is 0
	chin <- func() (int, []int) { return 0, phases }
	resultfunc := <-chout
	result, _ := resultfunc()
	return result
}
func Amplify(in chan func() (int, []int), out chan func() (int, []int), amplifier int, intcodeprogram []int) {
	//use a select with the in chan and a signal to kill chan, maybe a waitgrop?
	for {
		inputfunc, ok := <-in
		if !ok {
			log.Printf("GGGGGGGG closing goroutine %d\n", amplifier)
			close(out)
			return
		}
		input, phases := inputfunc()
		// run calc
		//make a copy
		tmp := make([]int, len(intcodeprogram))
		copy(tmp, intcodeprogram)

		result, e := Process(tmp, []int{phases[amplifier], input})
		//here I need to get the output, but then be able to provide more input
		if e != nil {

		}
		// log.Printf("amplifier %d has a phase of %d  had an input of %d is now putting %d on the out\n", amplifier, phases[amplifier], input, result)
		// log.Printf("program %v\n", intcodeprogram)
		out <- func() (int, []int) { return result, phases }
	}

}
