package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var SockArt string = `
                                      ░░▒▒▒▒            
                        ▒▒▒▒▒▒▒▒▒▒▒▒▒▒░░░░░░▒▒░░        
                        ░░░░░░░░░░░░░░░░░░░░▒▒▒▒        
                        ░░░░░░░░░░░░░░░░░░░░▒▒▒▒░░      
                      ░░░░░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▓▓░░      
                      ░░░░░░░░░░▒▒▒▒▒▒▒▒▒▒▓▓▓▓▓▓        
                      ░░░░░░░░▒▒▒▒▓▓▓▓▓▓████▒▒          
                    ░░░░░░░░▒▒▒▒▓▓▓▓▓▓██▒▒              
                    ░░░░░░▒▒▒▒▓▓████▒▒                  
                  ░░░░░░▒▒▒▒▓▓▓▓██                      
                ░░░░░░░░▒▒▓▓▓▓▓▓                        
            ░░░░░░░░░░▒▒▓▓▓▓██                          
        ░░░░░░░░░░░░▒▒▓▓▓▓▓▓                            
      ▒▒▒▒░░░░▒▒▒▒▒▒▓▓▓▓▓▓                              
      ▒▒▒▒▒▒▒▒▒▒▓▓▓▓██▓▓                                
        ░░  ▒▒▒▒▒▒▒▒
`

type Progress struct {
	CurVal    int
	NeededVal int
}

type Success struct {
	Name        string
	Description string
	Discovered  bool
	Progress    *Progress
	Reward      *Item
}

// Parse the json file and put successes in a global map
func InitSuccesses() {
	Successes = make(map[string]*Success, 0)

	filename := "../config/successes.json"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error parsing file", filename, err)
	}
	var payload map[string][]*Success
	err = json.Unmarshal(content, &payload)
	if err != nil {
		fmt.Println("Error Unmarshal()", filename, err)
	}
	for _, success := range payload["Successes"] {
		Successes[success.Name] = success
	}
}

func (s *Success) UpdateSucces(c *Character, val int) {
	s.Progress.CurVal += val
	if s.Progress.CurVal >= s.Progress.NeededVal && !s.Discovered {
		s.Discovered = true
		fmt.Printf("\nYou achieved the success \"%s\"\n\n", s.Name)
		c.Inventory.Slots += 1
		c.Inventory.AddItem(s.Reward, 1)
	}
}

func DisplaySuccesses() {
	fmt.Printf("--- Successes ---\n\n")
	for _, success := range Successes {
		success_desc := "?"
		if success.Discovered {
			success_desc = success.Description
		}
		fmt.Printf("Success: %s\nDescription: %s\nDiscovered: %t\n", success.Name, success_desc, success.Discovered)
		fmt.Printf("Progress: %d/%d\n\n", success.Progress.CurVal, success.Progress.NeededVal)
	}
}
