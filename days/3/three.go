package main

import (
	"log"
	"strconv"

	"github.com/ribesthlm/AdventOfCode2021/internal/utils"
)

const (
	day = "3"
)

type plugin string
var puzzle interface{}

func init() {
	resp, err := utils.LoadInput(day)
	if err != nil {
		log.Fatalln(day, "error loading input")
	}

	puzzle, err = utils.ByteToString(resp)
	if err != nil {
		log.Fatalln(day, "erro converting input")
	}
	
}

// solve
func solve(a int, b int) int {
	return a * b
}

// Run is called from main.go
func (p plugin) Run() {
	// Array size of 12 with uint16 is within input size
	var counters [12]int
	
	loopCount := 0
	for n, l := range puzzle.([]string) {
		num, _ := strconv.ParseInt(l, 2, 32)

		// LSB for every position added together
		for n := 0; n < 12; n++ {
			if ((num >> (11-n)) & 1) > 0 {
				counters[n] += 1
			}
		}
		loopCount = n
	}

	var gamma int = 0
	
	// As loopcount is odd there is an answer
	for n, counter := range counters {
		//fmt.Println(n, counter)
		// Set bits accourding to counters
		if (int(loopCount) - int(counter) > 499) {
			// Create a 1, left shifting to correct position
			// OR sets the bit
			gamma |= 1 << (11-n)
		} 
		//fmt.Printf("%12b %d\n", gamma, gamma)
	}

	log.Println("day-three-gamma:", gamma)

	// 0xFFF as we have 12 bits
	epsilon := gamma ^ 0xFFF
	
	log.Println("day-three-epsilon:", epsilon)

	log.Println("day-three-1:", solve(int(gamma), int(epsilon)))
}

// export symbol Plugin
var Plugin plugin