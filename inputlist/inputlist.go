package inputlist

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"github.com/gkiki90/steplib-cli/pathutil"
	"encoding/json"
	"errors"
	"os"
)

type StepJSONInputStruct struct {
	MappedTo 		string 		`json:"mapped_to"`
	Title 			string 		`json:"title"`
	Description 	string 		`json:"description"`
	ValueOptions 	[]string 	`json:"value_options"`
	Value 			string		`json:"value"`
	IsExpand 		string 		`json:"is_expand"`
	IsRequired 		string 		`json:"is_required"`
}

type StepJSONInputsStruct struct {
	Inputs []StepJSONInputStruct `json:"inputs"`
}

type StepYMLInputStruct struct {
	MappedTo 		string 		`yaml:"mapped_to"`
	Title 			string 		`yaml:"title"`
	Description 	string 		`yaml:"description"`
	ValueOptions 	[]string 	`yaml:"value_options"`
	Value 			string		`yaml:"value"`
	IsExpand 		string 		`yaml:"is_expand"`
	IsRequired 		string 		`yaml:"is_required"`
}

type StepYMLInputsStruct struct {
	Inputs []StepYMLInputStruct `yaml:"inputs"`
}


func ReadYMLInputListFromFile(fpath string) (StepYMLInputsStruct, error) {
	source, err := ioutil.ReadFile(fpath)
	if err != nil {
		return StepYMLInputsStruct{}, err;
	}

	var stepYMLInputStruct StepYMLInputsStruct

	err = yaml.Unmarshal(source, &stepYMLInputStruct)

    return stepYMLInputStruct, err
}

func generateFormattedJSONForInputList(inputList StepYMLInputsStruct) ([]byte, error) {
	jsonContBytes, err := json.MarshalIndent(inputList, "", "\t")
	if err != nil {
		return []byte{}, err
	}
	return jsonContBytes, nil
}

func WriteInputListToFile(fpath string, inputList StepYMLInputsStruct) error {
	if fpath == "" {
		return errors.New("No path provided")
	}

	isExists, err := pathutil.IsPathExists(fpath)
	if err != nil {
		return err
	}
	if isExists {
		//return errors.New("Inputlist file already exists!")
	}

	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonContBytes, err := generateFormattedJSONForInputList(inputList)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonContBytes)
	if err != nil {
		return err
	}

	return nil
}