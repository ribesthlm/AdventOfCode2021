package main

import (
	"log"

	"github.com/ribesthlm/AdventOfCode2021/internal/utils"
)

const (
	day = "1"
)

type plugin string
var puzzle []int

func init() {
	resp, err := utils.LoadInput(day)
	if err != nil {
		log.Fatalln("day one: error loading input ", err)
	}

	puzzle, err = utils.ByteToInt(resp)
	if err != nil {
		log.Fatalln("day one: error reading input ", err)
	}
}

func solve(puzzle []int) int {
	
	depth := 0
	increase := 0

	for _, m := range puzzle {
		switch {
		case depth == 0:
			depth = m
		case m > depth:
			increase = increase + 1
		}
		depth = m
	}
	return increase
}

func (p plugin) Run() {
	log.Println("day-one-1:", solve(puzzle))

	// Sliding threes
	var m2 []int

	for n := range puzzle {
		switch {
		case n < 2:
			// do nothing
		default:
			m2 = append(m2, (puzzle[n] + puzzle[n-1] + puzzle[n-2]))
		}
	}

	log.Println("day-one-2:", solve(m2))

}

// export symbol Plugin
var Plugin plugin