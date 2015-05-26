package steputil

import (

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

type StepInputJSONStruct struct {
	MappedTo 		string 		`json:"mapped_to"`
	Title 			string 		`json:"title"`
	Description 	string 		`json:"description"`
	ValueOptions 	[]string 	`json:"value_options"`
	Value 			string		`json:"value"`
	IsExpand 		string 		`json:"is_expand"`
	IsRequired 		string 		`json:"is_required"`
}

type StepInputsJSONStruct struct {
	Inputs []StepInputYMLStruct `json:"inputs"`
}