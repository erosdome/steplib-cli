package main

import (
	"github.com/erosdome/steplib-cli/pathutil"
	"github.com/erosdome/steplib-cli/steputil"
	"github.com/erosdome/steplib-cli/inputlist"
	"code.google.com/p/go.crypto/ssh/terminal"
	"github.com/codegangsta/cli"
	"fmt"
	"os"
	"os/exec"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

var stdinValue string

func initCommand(c *cli.Context) {
	ymlPath := "step.yml"

	stepYMLInputStruct, error := inputlist.ReadSetpInputListYMLFromFile(ymlPath)
	if error != nil {
		fmt.Println("Error: %s", error)
		return
	}

	err := inputlist.WriteInputListToFile("./inputlist.json", stepYMLInputStruct)
	if err != nil {
		fmt.Println("Error: %s", err)
		return
	}
	fmt.Println("inputlist.json file saved")

	return
}

func runCommand(c *cli.Context) {
	doesExist, err := pathutil.IsPathExists("./step.sh")
	if err != nil {
		fmt.Println("Error: %s", err)
		return
	}

	if doesExist == false {
		fmt.Println("No step.sh script file found in current directory")
		return
	}

	out, err := exec.Command("./step.sh").Output()
	fmt.Printf("%s", out)
}

func convertCommand(c *cli.Context) {
	stepPath := c.String("path")
	stepFormat := c.String("format")
	sourceName := "step.yml"
	targetName := "step.json"

	if stepPath == "" {
		fmt.Println("Error: path not defined")
		return
	}

	if stepFormat == "" {
		fmt.Println("Error: format not defined")
		return
	}

	if stepFormat != "json" {
		fmt.Println("Invalid target format")
		return
	}

	doesExist, err := pathutil.IsPathExists(stepPath+sourceName)
	if err != nil {
		fmt.Println("Error: %s", err)
		return
	}
	if doesExist == false {
		fmt.Println("Error: %s not found", sourceName)
		return
	}

	sourceFile, err := ioutil.ReadFile("step.yml")
	if err != nil {
		fmt.Println("Error: %s", err)
		return
	}

	var stepYMLStruct steputil.StepYMLStruct

	err = yaml.Unmarshal(sourceFile, &stepYMLStruct)

	if err != nil {
		fmt.Println("Error: %s", err)
		return
	}

	jsonContBytes, err := json.MarshalIndent(stepYMLStruct, "", "\t")

	if err != nil {
		fmt.Println("Error: %s", err)
		return
	}

	file, err := os.Create(stepPath+targetName)
	if err != nil {
		fmt.Println("Error: %s", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonContBytes)
	if err != nil {
		fmt.Println("Error: %s", err)
		return
	}

	return
}

func main() {

	// Read piped data
	stdinValue = ""
	if ! terminal.IsTerminal(0) {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Print("Failed to read stdin, err: %s", err)
		}
		stdinValue = string(bytes)
	} 

	// Parse cl 
	app := cli.NewApp()
	app.Name = "steplib-cli"
	app.Usage = "Steplib step client."

	app.Commands = []cli.Command {
		{
			Name:  "init",
			SkipFlagParsing: true,
			Action: initCommand,
		},
		{
			Name:  "run",
			SkipFlagParsing: true,
			Action: runCommand,
		},
		{
			Name:  "convert",
			Flags: []cli.Flag {
				cli.StringFlag {
					Name: "path",
					Value: "",
				},
				cli.StringFlag {
					Name: "format",
					Value: "",
				},
			},
			Action: convertCommand,
		},
	}

	app.Run(os.Args)
}