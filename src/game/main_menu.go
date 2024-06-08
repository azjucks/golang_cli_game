package game

import (
	"fmt"
	"os"
	"math/rand"
	"strings"
)

const (
	CHARACTER_INFOS int = iota
	CHARACTER_INVENTORY
	CHARACTER_SKILL_TREE
	SHOPKEEPER_MENU
	BLACKSMITH_MENU
	FIGHT_MENU
	SUCCESSES
	WHO_ARE_THEY
	QUIT
)

type MainMenu struct {
	Id		int
	Choices map[int]any
}

func InitMainMenu(g *Game) *MainMenu {
	var menu *MainMenu = new(MainMenu)

	menu.Id = MAIN
	menu.Choices = make(map[int]any, 0)
	menu.Choices[CHARACTER_INFOS] = g.Char.DisplayInfos
	menu.Choices[CHARACTER_INVENTORY] = g.ChangeMenu
	menu.Choices[CHARACTER_SKILL_TREE] = g.ChangeMenu
	menu.Choices[SHOPKEEPER_MENU] = g.ChangeMenu
	menu.Choices[BLACKSMITH_MENU] = g.ChangeMenu
	menu.Choices[FIGHT_MENU] = g.ChangeMenu
	menu.Choices[SUCCESSES] = DisplaySuccesses
	menu.Choices[WHO_ARE_THEY] = WhoAreThey
	menu.Choices[QUIT] = menu.Quit

	return menu
}

func WhoAreThey() {
	fmt.Println("ABBA, Steven Spielberg")
}

func (m MainMenu) DisplayChoices(*Game) {
	fmt.Printf("\n")
	fmt.Printf("%d\\ Character infos\n", CHARACTER_INFOS)
	fmt.Printf("%d\\ Character inventory\n", CHARACTER_INVENTORY)
	fmt.Printf("%d\\ Skill tree\n", CHARACTER_SKILL_TREE)
	fmt.Printf("%d\\ Shopkeeper\n", SHOPKEEPER)
	fmt.Printf("%d\\ Blacksmith\n", BLACKSMITH)
	fmt.Printf("%d\\ Fight\n", FIGHT_MENU)
	fmt.Printf("%d\\ Successes\n", SUCCESSES)
	fmt.Printf("%d\\ Who are they?\n", WHO_ARE_THEY)
	fmt.Printf("%d\\ Exit Game\n\n", QUIT)
}

func (m MainMenu) Choice(g *Game, index int) any {
	switch index {
		case CHARACTER_INFOS:
			val, _ := m.Choices[CHARACTER_INFOS]
			val.(func())()
		case CHARACTER_INVENTORY:
			val, _ := m.Choices[CHARACTER_INVENTORY]
			val.(func(Menu))(Menus[INVENTORY])
		case SHOPKEEPER_MENU:
			val, _ := m.Choices[SHOPKEEPER_MENU]
			fmt.Println(ShopkeeperArt)
			fmt.Println(strings.Replace(Npcs["Shopkeeper"].Intro + "\n", "%s", g.Char.Name, -1))
			val.(func(Menu))(Menus[SHOPKEEPER])
		case BLACKSMITH_MENU:
			val, _ := m.Choices[BLACKSMITH_MENU]
			fmt.Println(BlacksmithArt)
			fmt.Println(strings.Replace(Npcs["Blacksmith"].Intro + "\n", "%s", g.Char.Name, -1))
			val.(func(Menu))(Menus[BLACKSMITH])
		case CHARACTER_SKILL_TREE:
			val, _ := m.Choices[CHARACTER_SKILL_TREE]
			val.(func(Menu))(Menus[SKILL_TREE])
		case FIGHT_MENU:
			val, _ := m.Choices[FIGHT_MENU]
			var mob *Mob = Mobs["Goblin"]
			n := rand.Intn(6)
			if n == 5 {
				mob = Mobs["Goose"]
			}
			val.(func(Menu))(InitFightMenu(g, g.Char, mob))
		case SUCCESSES:
			val, _ := m.Choices[SUCCESSES]
			val.(func())()
		case WHO_ARE_THEY:
			val, _ := m.Choices[WHO_ARE_THEY]
			val.(func())()
		case QUIT:
			val, _ := m.Choices[QUIT]
			val.(func(*Game))(g)
		default:
			fmt.Println("Please, enter a valid choice")
	}
	return nil
}

func (m *MainMenu) GetPrevMenu() Menu {
	return nil
}

func (m *MainMenu) SetPrevMenu(Menu) {
}

func (m *MainMenu) GetId() int {
	return m.Id
}

func (m *MainMenu) Quit(g *Game) {
	os.Exit(1)
}
