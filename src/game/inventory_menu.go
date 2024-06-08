package game

import (
	"fmt"
	"strconv"
)

type InventoryMenu struct {
	Id			int
	PrevMenu 	Menu
	PlayerInv 	*Inventory
	Choices 	[]*Item
}

func InitInvMenu(g *Game) (*InventoryMenu) {
	var menu *InventoryMenu = new(InventoryMenu)
	menu.Id = INVENTORY
	menu.PlayerInv = g.Char.Inventory
	menu.Choices = make([]*Item, 0)
	return menu
}

func (m *InventoryMenu) DisplayChoices(*Game) {
	m.Choices = m.PlayerInv.DisplayItems()
	fmt.Printf("%d\\Quit\n\n", len(m.PlayerInv.Items))
}

func (m InventoryMenu) Choice(g *Game, index int) any {
	if (index >= len(m.Choices) || index < 0) {
		m.Quit(g)
		return nil
	}
	item := m.Choices[index]
	target := m.ChoseTarget(g)
	g.Char.ConsumeItem(target, item)
	if g.CurrentFight != nil {
		g.PrevMenu()
	}
	return item
}


// Choose the target of the item and apply effects to it
func (m InventoryMenu) ChoseTarget(g *Game) Entity {
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

func (m *InventoryMenu) GetPrevMenu() Menu {
	return m.PrevMenu
}

func (m *InventoryMenu) SetPrevMenu(menu Menu) {
	m.PrevMenu = menu
}

func (m *InventoryMenu) GetId() int {
	return m.Id
}

func (m *InventoryMenu) Quit(g *Game) {
	g.PrevMenu()
}
