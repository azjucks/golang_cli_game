package game

import (
	"fmt"
	"time"
)

type NpcMenu struct {
	Id			int
	PrevMenu 	Menu
	Npc 		*Npc
}

func InitNpcMenu(npc *Npc, id int) (*NpcMenu) {
	var menu *NpcMenu = new(NpcMenu)
	menu.Id = id
	menu.Npc = npc
	return menu
}

func (m NpcMenu) DisplayChoices(*Game) {
	m.Npc.NpcShop.DisplaySells()
	fmt.Printf("%d\\Quit\n\n", len(m.Npc.NpcShop.Sells))
}

func (m NpcMenu) Choice(g *Game, index int) any {
	if index >= len(m.Npc.NpcShop.Sells) || index < 0 {
		m.Quit(g)
		return nil
	}
	sell := m.Npc.NpcShop.Sells[index]
	if sell.Count <= 0 {
		return nil
	}
	var can_buy bool
	defer func() {
		if can_buy {
			m.Npc.NpcShop.UseSell(sell)
			g.Char.BuyItem(sell)
			fmt.Printf("You bought: %s.\n", sell.ShopItem.Name)
		}
		time.Sleep(600 * time.Millisecond)
	}()
	if sell.Price.Value > g.Char.Currency.CurValue {
		fmt.Println("You do not have enough golds.")
		return nil
	}
	switch sell.ShopItem.Name {
		case "Book: Fireball":
			fireball_skill := SkillFromName("Fireball")
			g.Char.SkillTree.AddSkill(fireball_skill)
			can_buy = true
			break
		case "Book: Thunder Strike":
			thunder_skill := SkillFromName("Thunder Strike")
			g.Char.SkillTree.AddSkill(thunder_skill)
			can_buy = true
			break
		default:
			can_buy = g.Char.CanBuy(sell)
	}
	return nil
}

func (m *NpcMenu) GetPrevMenu() Menu {
	return m.PrevMenu
}

func (m *NpcMenu) SetPrevMenu(menu Menu) {
	m.PrevMenu = menu
}

func (m *NpcMenu) GetId() int {
	return m.Id
}

func (m *NpcMenu) Quit(g *Game) {
	g.PrevMenu()
}
