package game

import (
	"fmt"
	"strconv"
)

type SkillMenu struct {
	Id				int
	PrevMenu 		Menu
	PlayerSkillTree *SkillTree
}

func InitSkillMenu(g *Game) (*SkillMenu) {
	var menu *SkillMenu = new(SkillMenu)
	menu.Id = SKILL_TREE
	menu.PlayerSkillTree = g.Char.SkillTree
	return menu
}

func (m SkillMenu) DisplayChoices(*Game) {
	m.PlayerSkillTree.DisplaySkills()
	fmt.Printf("%d\\Quit\n\n", len(m.PlayerSkillTree.Skills))
}

func (m SkillMenu) Choice(g *Game, index int) any {
	if index >= len(m.PlayerSkillTree.Skills) || index < 0 {
		m.Quit(g)
		return nil
	}
	skill := g.Char.SkillTree.Skills[index]
	target := m.ChoseTarget(g)
	g.Char.UseSkill(skill, target)
	if g.CurrentFight != nil {
		g.PrevMenu()
	}
	return skill
}

// Choose the target of the spell and apply effects to it
func (m SkillMenu) ChoseTarget(g *Game) Entity {
	if g.CurrentFight == nil {
		return g.Char
	}

	fmt.Printf("--- Choose Target ---\n")
	fmt.Printf("0\\ %s (You)\n", g.Char.GetName())
	fmt.Printf("1\\ %s (Enemy)\n", g.CurrentFight.Enemy.GetName())

	var choice string
	fmt.Scanf("%s\n", &choice)
	fmt.Println()
	_choice, err := strconv.Atoi(choice)

	if err != nil {
		fmt.Println("Please, enter a valid choice")
		return m.ChoseTarget(g)
	}

	if _choice == 0 {
		return g.Char
	}

	return g.CurrentFight.Enemy
}

func (m *SkillMenu) GetPrevMenu() Menu {
	return m.PrevMenu
}

func (m *SkillMenu) SetPrevMenu(menu Menu) {
	m.PrevMenu = menu
}

func (m *SkillMenu) GetId() int {
	return m.Id
}

func (m *SkillMenu) Quit(g *Game) {
	g.PrevMenu()
}
