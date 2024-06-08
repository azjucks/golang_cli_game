package game

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CharCreaMenu struct {
	Id        int
	CharName  string
	CharClass *Class
	Choices   []*Class
}

// Init the character
func InitCharCreaMenu() *CharCreaMenu {
	var menu *CharCreaMenu = new(CharCreaMenu)
	menu.Id = CHAR_CREATION
	menu.Choices = make([]*Class, 0)
	return menu
}

// Check if the name contains only letters
// Title the name if so
func check_name(name *string) bool {
	for _, _rune := range *name {
		if (_rune < 65 || _rune > 90) && (_rune < 97 || _rune > 122) {
			return false
		}
	}
	caser := cases.Title(language.BrazilianPortuguese)
	*name = caser.String(*name)
	return true
}

func (m *CharCreaMenu) DisplayChoices(*Game) {
	for m.CharName == "" {
		tmp := ""
		fmt.Printf("Choose your name: ")
		fmt.Scanf("%s\n", &tmp)
		fmt.Println("\n")
		if check_name(&tmp) {
			m.CharName = tmp
		} else {
			fmt.Println("Invalid name. It must contain only letters.")
		}
	}
	var j int
	for key, val := range Classes {
		m.Choices = append(m.Choices, val)
		fmt.Printf("%d\\ %s\n", j, key)
		j++
	}
	fmt.Println()
}

func (m *CharCreaMenu) Choice(g *Game, index int) any {
	if index >= len(m.Choices) || index < 0 {
		return nil
	}

	m.CharClass = m.Choices[index]
	fmt.Printf("You chose: %s\n", m.CharClass.Name)

	g.Char.Name = m.CharName
	g.Char.SetClass(m.CharClass)
	g.ChangeMenu(Menus[MAIN])
	return nil
}

func (m CharCreaMenu) GetPrevMenu() Menu {
	return nil
}

func (m *CharCreaMenu) SetPrevMenu(menu Menu) {
}

func (m CharCreaMenu) GetId() int {
	return m.Id
}

func (m CharCreaMenu) Quit(g *Game) {
}
