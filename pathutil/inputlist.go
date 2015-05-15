package pathutil

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type StepJSONInputStruct struct {
	MappedTo 		string 		`json:"mapped_to"`
	Title 			string 		`json:"title"`
	Description 	string 		`json:"description"`
	ValueOptions 	[]string 	`json:"value_options"`
	Value 			string		`json:"value"`
	IsExpand 		string 		`json:"is_expand"`
	IsRequired 		string 		`yaml:"is_required"`
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