package main

import (
	"fmt"
	"log"

	"github.com/ribesthlm/AdventOfCode2021/internal/utils"
)

type plugin string

var puzzle []int

func load() {

	fmt.Println("Hello from day one - init")

	resp, err := utils.LoadInput("one")
	if err != nil {
		log.Fatalln("day one: error loading input ", err)
	}

	puzzle, err = utils.ReadInput(resp)
	if err != nil {
		log.Fatalln("day one: error reading input ", err)
	}
}

func (p plugin) Run() {

	load()

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
	log.Println("Answer-1:", increase)

	// Sliding threes
	e := len(puzzle)
	log.Println("Length of input", e)
	var m2 []int

	for n, m := range puzzle {
		switch {
		case n < 2:
			fmt.Println(n, m)
		default:
			m2 = append(m2, (puzzle[n] + puzzle[n-1] + puzzle[n-2]))
			fmt.Println(puzzle[n], puzzle[n-1], puzzle[n-2], (puzzle[n] + puzzle[n-1] + puzzle[n-2]))
		}
	}

	depth = 0
	increase = 0
	e = len(m2)
	log.Println("Length of input", e)

	for _, m := range m2 {
		switch {
		case depth == 0:
			depth = m
		case m > depth:
			increase = increase + 1
		}
		depth = m
	}
	log.Println("Answer-2:", increase)

}

// export symbol Plugin
var Plugin plugin