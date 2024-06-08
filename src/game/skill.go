package game

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Skill struct {
	Name 		string
	Description string
	ManaCost	int
	Level 		int
	Stats 		[]*Stat
}

// Parse the json file and put skills in a global map
func InitSkills() {
	Skills = make(map[string]*Skill, 0)

	filename := "../config/skills.json"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error parsing file", filename, err)
	}
	var payload map[string][]*Skill
	err = json.Unmarshal(content, &payload)
	if err != nil {
		fmt.Println("Error Unmarshal()", filename, err)
	}
	for _, skill := range payload["Skills"] {
		Skills[skill.Name] = skill
	}
}

type SkillTree struct {
	Skills []*Skill
}

func InitSkillTree() (*SkillTree) {
	var skill_tree *SkillTree = new(SkillTree)

	skill_tree.Skills = make([]*Skill, 0)
	skill_tree.Skills = append(skill_tree.Skills, Skills["Punch"])

	return skill_tree
}

func (st SkillTree) HasSkill(skill *Skill) (bool) {
	for _, _skill := range st.Skills {
		if skill.Name == _skill.Name {
			return true
		}
	}
	return false
}
func (st *SkillTree) AddSkill(skill *Skill) {
	if skill == nil {
		return
	}
	if st.HasSkill(skill) {
		fmt.Printf("You already know the skill \"%s\".\n\n", skill.Name)
		return
	}

	st.Skills = append(st.Skills, skill)
	fmt.Printf("You learned the skill \"%s\".\n\n", skill.Name)
}

func (st SkillTree) DisplaySkills() {
	fmt.Println("--- Skills ---")
	for j, skill := range st.Skills {
		fmt.Printf("%d\\ \"%s\"\n\tLevel: %d\n", j, skill.Name, skill.Level)
		fmt.Printf("\tDescription: %s\n\tEffects: ", skill.Description)
		for i, stat := range skill.Stats {
			fmt.Printf("\"%s\": %d", stat.Name, stat.Value)
			if i != len(skill.Stats) - 1 {
				fmt.Printf(",")
			}
		}
		fmt.Println("\n")
	}
}
