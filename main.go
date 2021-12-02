package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

type Plugin interface {
	Run()
}

func main() {
	fmt.Println("AdventOfCode2021 in golang this year :D")

	var days []string

	dayPtr := flag.String("day", "all", "day to run")

	flag.Parse()

	err := filepath.Walk("./bin", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		if *dayPtr != "all" {
			if !info.IsDir() && filepath.Ext(path) == ".so" && strings.TrimSuffix(filepath.Base(path),".so") == *dayPtr {
				days = append(days, path)
			}
		} else {
			if !info.IsDir() && filepath.Ext(path) == ".so" {
				days = append(days, path)
			}
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, day := range days {
		// load module
		plug, err := plugin.Open(day)
		if err != nil {
			log.Fatalln("Could not load plugin", err)
		}

		// Lookup symbol
		symPlugin, err := plug.Lookup("Plugin")
		if err != nil {
			log.Fatalln(err)
		}

		// Assert symbols
		var plugin Plugin
		plugin, ok := symPlugin.(Plugin)
		if !ok {
			log.Fatalln("unexpected type from module")
		}

		plugin.Run()
	}
}