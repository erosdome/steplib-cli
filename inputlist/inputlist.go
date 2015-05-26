package inputlist

import (
	"github.com/gkiki90/stepman/steputil"
	"github.com/gkiki90/stepman/pathutil"
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"errors"
	"os"
)

func ReadSetpInputListYMLFromFile(fpath string) (steputil.StepInputsYMLStruct, error) {
	source, err := ioutil.ReadFile(fpath)
	if err != nil {
		return steputil.StepInputsYMLStruct{}, err;
	}

	var stepYMLInputStruct steputil.StepInputsYMLStruct

	err = yaml.Unmarshal(source, &stepYMLInputStruct)

    return stepYMLInputStruct, err
}

func generateFormattedJSONForInputList(inputList steputil.StepInputsYMLStruct) ([]byte, error) {
	jsonContBytes, err := json.MarshalIndent(inputList, "", "\t")
	if err != nil {
		return []byte{}, err
	}
	return jsonContBytes, nil
}

func WriteInputListToFile(fpath string, inputList steputil.StepInputsYMLStruct) error {
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



