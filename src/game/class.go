package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Class struct {
	Name  string
	Stats []*Stat
}

// Parse the json file and put classes in a global map
func InitClasses() {
	Classes = make(map[string]*Class, 0)

	filename := "../config/classes.json"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error parsing file", filename, err)
	}

	var payload map[string][]*Class
	err = json.Unmarshal(content, &payload)
	if err != nil {
		fmt.Println("Error Unmarshal()", filename, err)
	}

	for _, class := range payload["Classes"] {
		Classes[class.Name] = class
	}
}
