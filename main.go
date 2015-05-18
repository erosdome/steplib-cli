package main

import (
	"fmt"
	"flag"
	"os"
	"log"
    "github.com/gkiki90/steplib-cli/inputlist"
)

type Command struct {
	Name  string
	Usage string
	Run   func() error
}

var (
	availableCommands = []Command{
		Command{
			Name:  "init",
			Usage: "init - creates an inputlist.json file in the current folder",
			Run:   doInitCommand,
		},
	}
)


func usage() {
	fmt.Println("Called usage")
	fmt.Fprintf(os.Stderr, "Usage: %s command [FLAGS]", os.Args[0])
	flag.PrintDefaults()
	fmt.Println("Available commands:")
	for _, cmd := range availableCommands {
		fmt.Println(" *", cmd.Name)
		fmt.Println("    ", cmd.Usage)
	}
}

func doInitCommand() error {
	ymlPath := "step.yml"

	stepYMLInputStruct, error := inputlist.ReadYMLInputListFromFile(ymlPath)
	if error != nil {
		return error
	}

    err := inputlist.WriteInputListToFile("./inputlist.json", stepYMLInputStruct)
	if err != nil {
		return err
	}
	fmt.Println("inputlist.json file saved")

	return nil
}

func main() {
	fmt.Println("steplib-cli")

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	theCommandName := args[0]

	for _, cmd := range availableCommands {
		if cmd.Name == theCommandName {
			err := cmd.Run()
			if err != nil {
				log.Fatalln(err)
			}
			return
		}
	}
}