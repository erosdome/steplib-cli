package main

import (
	"fmt"
	"flag"
	"os"
	"log"
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type Command struct {
	Name  string
	Usage string
	Run   func() error
}

// Parsing step.yml (https://github.com/steplib/steplib/blob/master/docs/step_format.md)
type StepInput struct {
	MappedTo string `yaml:"mapped_to"`
	Title string 
	Description string 
	ValueOptions []string 
	Value string
	IsExpand string `yaml:"is_expand"`
	IsRequired string `yaml:"is_required"`
}

type StepConfig struct {
	Inputs []StepInput
}

var (
	availableCommands = []Command{
		Command{
			Name:  "init",
			Usage: "init - creates a deplist.json file in the current folder",
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
	source, err := ioutil.ReadFile("step.yml")
	if err != nil {
		return err;
	}

	var stepConfig StepConfig

	err = yaml.Unmarshal(source, &stepConfig)
	if err != nil {
        panic(err)
    }

    fmt.Printf("StepConfig: %s\n", stepConfig)

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