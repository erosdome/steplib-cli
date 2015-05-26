package main

import (
	"github.com/gkiki90/stepman/pathutil"
	"github.com/gkiki90/stepman/inputlist"
	"fmt"
	"flag"
	"os"
	"log"
	"errors"
)

type Command struct {
	Name  string
	Usage string
	Arguments []Argument
	Run   func(params[] string) error
}

type Argument struct {
	Name string
	Usage string
}

var (
	availableCommands = []Command{
		Command{
			Name:  "init",
			Usage: "init - creates an inputlist.json file in the current folder",
			Arguments: nil,
			Run:   doInitCommand,
		},
		Command{
			Name:  "run",
			Usage: "run - perform step with given inputlist.json",
			Arguments: nil,
			Run:   doRunCommand,
		},
		Command{
			Name:  "convert",
			Usage: "convert - convert specified json/yml to yml/json",
			Arguments: []Argument{
				Argument{
					Name: "-step",
					Usage: "-step - methods on step",
				},
			},
			Run:   doConvertCommand,
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

func doInitCommand(params[] string) error {
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

func doRunCommand(params[] string) error {
	isExists, err := pathutil.IsPathExists("./inputlist.json")
	if err != nil {
		return err
	}
	if isExists == false {
		return errors.New("Inputlist file dos not exists!")
	}

	return nil
}

func doConvertCommand(params[] string) error {
	fmt.Println(params)
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

	var theParams []string

	if len(args) > 1 {
		theParams = args[1:len(args)]
	}

	for _, cmd := range availableCommands {
		if cmd.Name == theCommandName {
			err := cmd.Run(theParams)
			if err != nil {
				log.Fatalln(err)
			}
			return
		}
	}
}