package inputlist

import (
	"github.com/gkiki90/steplib-cli/pathutil"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"errors"
	"os"
)

type StepInputYMLStruct struct {
	MappedTo 		string 		`yaml:"mapped_to"`
	Title 			string 		`yaml:"title"`
	Description 	string 		`yaml:"description"`
	ValueOptions 	[]string 	`yaml:"value_options"`
	Value 			string		`yaml:"value"`
	IsExpand 		string 		`yaml:"is_expand"`
	IsRequired 		string 		`yaml:"is_required"`
}

type StepInputsYMLStruct struct {
	Inputs []StepInputYMLStruct `yaml:"inputs"`
}

func ReadSetpInputListYMLFromFile(fpath string) (StepInputsYMLStruct, error) {
	source, err := ioutil.ReadFile(fpath)
	if err != nil {
		return StepInputsYMLStruct{}, err;
	}

	var stepYMLInputStruct StepInputsYMLStruct

	err = yaml.Unmarshal(source, &stepYMLInputStruct)

    return stepYMLInputStruct, err
}

func generateFormattedJSONForInputList(inputList StepInputsYMLStruct) ([]byte, error) {
	jsonContBytes, err := json.MarshalIndent(inputList, "", "\t")
	if err != nil {
		return []byte{}, err
	}
	return jsonContBytes, nil
}

func WriteInputListToFile(fpath string, inputList StepInputsYMLStruct) error {
	if fpath == "" {
		return errors.New("No path provided")
	}

	isExists, err := pathutil.IsPathExists(fpath)
	if err != nil {
		return err
	}
	if isExists {
		// return errors.New("Inputlist file already exists!")
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



