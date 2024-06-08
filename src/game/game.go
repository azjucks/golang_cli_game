package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Create a menu for the player
const (
	MAIN int = iota
	CHAR_CREATION
	INVENTORY
	SHOPKEEPER
	BLACKSMITH
	SKILL_TREE
	FIGHT
)

// Create maps for the menu
var Menus map[int]Menu
var Npcs map[string]*Npc
var Items map[string]*Item
var Skills map[string]*Skill
var Classes map[string]*Class
var Mobs map[string]*Mob
var Successes map[string]*Success

type Game struct {
	CurrentMenu  Menu
	Char         *Character
	ValidIndex   bool
	CurrentFight *FightMenu
}

// Init all values in the game
func init_game() *Game {
	game := new(Game)
	InitClasses()
	InitItems()
	InitSkills()
	InitSuccesses()
	c := NewCharacter()
	game.Char = c
	InitNpcs()
	InitMobs()
	InitMenus(game)

	char_crea_menu := InitCharCreaMenu()

	game.ChangeMenu(char_crea_menu)
	return game
}

// Main loop function
func StartGame() {
	var game *Game = init_game()

	var index string

	rand.Seed(time.Now().UnixNano())

	for {
		game.CheckGame()
		game.CurrentMenu.DisplayChoices(game)

		fmt.Scanf("%s", &index)
		fmt.Println()
		_index, err := strconv.Atoi(index)
		if err != nil {
			fmt.Println("Please, enter a valid choice")
			continue
		}

		game.CurrentMenu.Choice(game, _index)
	}
}

// Check if the player is in a fight
// If so check enemy health
func (g *Game) CheckGame() {
	g.Char.CheckCharacter()
	if g.CurrentFight != nil {
		g.CurrentFight.CheckEnemyHealth(g)
	}
}

// Set the current menu to the main menu
func (g *Game) HardChangeMenu(menu Menu) {
	g.CurrentMenu = menu
}

// Change the menu and set the current as its previous
func (g *Game) ChangeMenu(menu Menu) {
	if menu.GetId() == FIGHT {
		g.CurrentFight = menu.(*FightMenu)
	}
	menu.SetPrevMenu(g.CurrentMenu)
	g.CurrentMenu = menu
}

// Go back to the previous menu
func (g *Game) PrevMenu() {
	var tmp Menu = g.CurrentMenu
	if tmp.GetId() == FIGHT {
		g.CurrentFight = nil
	}
	g.CurrentMenu = g.CurrentMenu.GetPrevMenu()
	tmp.SetPrevMenu(nil)
}
