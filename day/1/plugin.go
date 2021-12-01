package main

import (
	"fmt"
	"log"

	"github.com/ribesthlm/internal/utils"
)

type plugin string

func (p plugin) Run() {
	fmt.Println("Hello from day one")

	resp, err := utils.LoadInput("one")
	if err != nil {
		log.Println("day one: error loading input")
	}
	fmt.Print(resp)
}

// export symbol Plugin
var Plugin plugin