package main

import (
	"github.com/gkiki90/steplib-cli/pathutil"
	"github.com/gkiki90/steplib-cli/inputlist"
	"fmt"
	"flag"
	"os"
	"log"
	"errors"
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
		Command{
			Name:  "run",
			Usage: "run - perform step with given inputlist.json",
			Run:   doRunCommand,
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

	stepYMLInputStruct, error := inputlist.ReadSetpInputListYMLFromFile(ymlPath)
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

func doRunCommand() error {
	fmt.Println("run")

	isExists, err := pathutil.IsPathExists("./inputlist.json")
	if err != nil {
		return err
	}
	if isExists == false {
		return errors.New("Inputlist file dos not exists!")
	}

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