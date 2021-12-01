package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"
)

type Plugin interface {
	Run()
}

func main() {
	fmt.Println("AdventOfCode2021 in golang this year :D")

	var days []string

	err := filepath.Walk("./bin", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		} 

		if !info.IsDir() && filepath.Ext(path) == ".so" {
			days = append(days, path)
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