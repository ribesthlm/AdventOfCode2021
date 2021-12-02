package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/ribesthlm/AdventOfCode2021/internal/utils"
)

const (
	day = "2"
)

type plugin string
var puzzle interface{}

func init() {
	resp, err := utils.LoadInput(day)
	if err != nil {
		log.Fatalln("day two: error loading input ", err)
	}

	puzzle, err = utils.ByteToString(resp)
	if err != nil {
		log.Fatalln("day two: error reading input ", err)
	}
	
}

type position struct {
	depth int
	horizontal int
	aim int
}

// solve returns the first answer
func solve(p *position) int {
	return p.depth * p.horizontal
}

func (p plugin) Run() {

	pos := position {
		depth: 0,
		horizontal: 0,
	}

	for _, l := range puzzle.([]string) {
		vals := strings.Split(l," ")

		action := vals[0]
		amount, err := strconv.Atoi(vals[1])
		if err != nil {
			log.Println("two.go: error", err)
		}

		switch action {
		case "forward":
			pos.horizontal += amount
		case "down":
			pos.depth += amount
		case "up":
			pos.depth -= amount
		}
	}
	log.Println("day-two-1:", solve(&pos))

	pos.depth = 0
	pos.horizontal = 0
	pos.aim = 0

	for _, l := range puzzle.([]string) {
		vals := strings.Split(l," ")

		action := vals[0]
		amount, err := strconv.Atoi(vals[1])
		if err != nil {
			log.Println("two.go: error", err)
		}

		switch action {
		case "forward":
			pos.horizontal += amount
			pos.depth += (pos.aim * amount)
		case "down":
			pos.aim += amount
		case "up":
			pos.aim -= amount
		}
	}
	log.Println("day-two-2:", solve(&pos))
}

// export symbol Plugin
var Plugin plugin