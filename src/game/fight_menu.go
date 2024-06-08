package game

import "fmt"

const (
	ATTACK_FIGHT int = iota
	INVENTORY_FIGHT
	QUIT_FIGHT
)

type FightMenu struct {
	Id				int
	PrevMenu		Menu
	Choices 		map[int]any
	TurnCount		int
	PlayerTurn		bool
	PlayerStarted	bool
	Player 			Entity
	Enemy 			Entity
}

func InitFightMenu(g *Game, player *Character, mob *Mob) (*FightMenu) {
	var menu *FightMenu = new(FightMenu)

	menu.Id = FIGHT
	menu.TurnCount = 1
	menu.Choices = make(map[int]any, 0)
	menu.Choices[ATTACK_FIGHT] = g.ChangeMenu
	menu.Choices[INVENTORY_FIGHT] = g.ChangeMenu
	menu.Choices[QUIT_FIGHT] = menu.Quit
	menu.Player = player
	menu.Enemy = mob

	if player.GetInitiative() > mob.GetInitiative() {
		menu.PlayerTurn = true
		menu.PlayerStarted = true
	}
	switch mob.Name {
	case "Goblin":
		fmt.Println(GoblinArt)
	case "Goose":
		fmt.Println(GooseArt)
	}
	return menu
}

func (m *FightMenu) CheckEnemyHealth(g *Game) {
	switch m.Enemy.(type) {
	case *Mob:
		var enemy *Mob = m.Enemy.(*Mob)
		if enemy.Hp.CurValue <= 0 {
			m.Quit(g)
		}
	}
}

func (m *FightMenu) CheckTurn(g *Game) {
	if m.PlayerTurn {
		return
	}

	m.Enemy.Turn(m.Player, m.TurnCount)

	m.PlayerTurn = true

	m.Player.(*Character).CheckCharacter()

	if m.PlayerStarted {
		m.TurnCount++
	}
}

func (m *FightMenu) DisplayChoices(game *Game) {
	m.CheckTurn(game)

	var player *Character = m.Player.(*Character)
	var enemy *Mob = m.Enemy.(*Mob)

	fmt.Printf("\n--- Turn %d ---\n", m.TurnCount)
	fmt.Printf("-- %s -- HP: %d/%d", player.GetName(), player.Hp.CurValue, player.Hp.MaxValue)
	fmt.Printf(" | Mana: %d/%d\n", player.Mana.CurValue, player.Mana.MaxValue)
	fmt.Printf("-- %s -- HP: %d/%d\n", enemy.GetName(), enemy.Hp.CurValue, enemy.Hp.MaxValue)

	fmt.Printf("\n")
	fmt.Printf("%d\\ Attack\n", ATTACK_FIGHT)
	fmt.Printf("%d\\ Inventory\n", INVENTORY_FIGHT)
	fmt.Printf("%d\\ Run\n\n", QUIT_FIGHT)

	if !m.PlayerStarted {
		m.TurnCount++
	}
}

func (m *FightMenu) Choice(g *Game, index int) any {
	m.PlayerTurn = false
	switch index {
	case ATTACK_FIGHT:
		val, _ := m.Choices[ATTACK_FIGHT]
		val.(func(Menu))(Menus[SKILL_TREE])
	case INVENTORY_FIGHT:
		val, _ := m.Choices[INVENTORY_FIGHT]
		val.(func(Menu))(Menus[INVENTORY])
	case QUIT_FIGHT:
		val, _ := m.Choices[QUIT_FIGHT]
		val.(func(*Game))(g)
	}
	return nil
}

func (m *FightMenu) GetPrevMenu() Menu {
	return m.PrevMenu
}

func (m *FightMenu) SetPrevMenu(menu Menu) {
	m.PrevMenu = menu
}

func (m *FightMenu) GetId() int {
	return m.Id
}

func (m *FightMenu) Quit(g *Game) {
	switch m.Enemy.(type) {
	case *Mob:
		var enemy *Mob = m.Enemy.(*Mob)
		if enemy.Hp.CurValue <= 0 {
			m.Player.(*Character).GetLoot(enemy)
		}
		enemy.Reset()
	}
	g.PrevMenu()
}
