package game

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Character_stat struct {
	Name       string
	MaxValue   int
	CurValue   int
	ClassBonus int
}

func (cs *Character_stat) UpdateStat(value int) int {
	cs.CurValue += value
	if cs.CurValue > cs.MaxValue {
		cs.CurValue = cs.MaxValue
	} else if cs.CurValue <= 0 {
		cs.CurValue = 0
		return 0
	}
	return 1
}

func (c *Character) UpdateExp(value int) {
	if value == 0 {
		return
	}
	if c.Xp.CurValue >= c.Xp.MaxValue {
		c.AddLevel()
		c.Xp.CurValue -= c.Xp.MaxValue
		rest := c.Xp.CurValue - c.Xp.MaxValue
		c.UpdateExp(rest)
	}
}

func new_char_stat(Name string, MaxValue, CurValue int) *Character_stat {
	char_stat := new(Character_stat)

	char_stat.Name = Name
	char_stat.MaxValue = MaxValue
	char_stat.CurValue = CurValue

	return char_stat
}

type Character struct {
	Name       string
	Class      *Class
	Currency   *Character_stat
	level      *Character_stat
	lives      *Character_stat
	Hp         *Character_stat
	Mana       *Character_stat
	Xp         *Character_stat
	Inventory  *Inventory
	SkillTree  *SkillTree
	Initiative *Character_stat
}

func NewCharacter() *Character {
	var c *Character = new(Character)

	c.Currency = new_char_stat("Golds", 999, 150)
	c.level = new_char_stat("Level", 20, 1)
	c.Initiative = new_char_stat("Initiative", 100, 100)
	c.lives = new_char_stat("Lives", 3, 2)
	c.Hp = new_char_stat("Health", 100, 20)
	c.Mana = new_char_stat("Mana", 500, 500)
	c.Xp = new_char_stat("Experience", 50, 0)
	c.Inventory = InitBaseInventory()
	c.SkillTree = InitSkillTree()

	return c
}

// Set the class of the player and update
// its stats accordingly
func (c *Character) SetClass(class *Class) {
	c.Class = class
	c.Hp.ClassBonus = class.Stats[0].Value
	c.Hp.CurValue += c.Hp.ClassBonus
	c.Hp.MaxValue += c.Hp.ClassBonus
}

// Function to call to see if the player needs
// to level up or has its hp set to 0
func (c *Character) CheckCharacter() {
	if c.is_dead() {
		c.dead()
	}
	if c.Xp.CurValue >= c.Xp.MaxValue {
		c.AddLevel()
	}
}

func (c *Character) AddLevel() {
	if c.level.CurValue < c.level.MaxValue {
		fmt.Printf("\n%s gained a level.\n\n", c.Name)
		c.level.UpdateStat(1)
	}
}

func (c Character) is_dead() bool {
	return c.Hp.CurValue == 0
}

func (c *Character) dead() {
	c.lives.UpdateStat(-1)

	if c.lives.CurValue <= 0 {
		fmt.Println("You have no live left. You are dead.")
		os.Exit(1)
	}

	c.Hp.UpdateStat((c.Hp.MaxValue + c.Hp.ClassBonus) / 2)
	tmp := "lives"
	if c.lives.CurValue <= 1 {
		tmp = "live"
	}
	fmt.Printf("You respawned, you have %d %s left.\n", c.lives.CurValue, tmp)
}

// Get the loot of the mob
// Update the success progress depending
// on mob type.
// Randomly generated xp and golds rewardds.
func (c *Character) GetLoot(mob *Mob) {
	switch mob.Name {
	case "Goblin":
		Successes["Sock seeker"].UpdateSucces(c, 1)
	case "Goose":
		Successes["Sock seeker"].UpdateSucces(c, 10)
	}
	exp := rand.Intn(mob.Loot.Exp)
	golds := rand.Intn(mob.Loot.Golds)
	c.Xp.CurValue += exp
	c.UpdateExp(exp)
	c.Currency.UpdateStat(golds)
	fmt.Printf("\n%s gained %d exp.\n", c.Name, exp)
	fmt.Printf("\n%s gained %d golds.\n\n", c.Name, golds)
}

func (c *Character) GetInitiative() int {
	return c.Initiative.CurValue
}

func (c *Character) GetName() string {
	return c.Name
}

// Deal damage to the player or Heal it if the value is negative
func (c *Character) TakeDamage(opponent Entity, damage int) {
	c.Hp.UpdateStat(-damage)
	if damage >= 0 {
		fmt.Printf("%s suffered %d damage from %s.\n", c.Name, damage, opponent.GetName())
		return
	}
	fmt.Printf("%s healed %d health from %s.\n", c.Name, -damage, opponent.GetName())
}

func (c *Character) Attack(target Entity, damage int) {
	target.TakeDamage(target, damage)
}

func (c *Character) Turn(opponent Entity, turn int) {
	c.Attack(opponent, turn)
}

func (c *Character) UseSkill(skill *Skill, target Entity) {
	if skill.ManaCost > c.Mana.CurValue {
		fmt.Printf("You cannot use %s. You have not enough mana.\n", skill.Name)
		return
	}
	c.Mana.UpdateStat(-skill.ManaCost)
	fmt.Printf("You used %s.\n", skill.Name)
	ApplyStatsValuesFromNames(c, target, skill.Stats)
}

// Check if the player hass enough golds
// or enough items to build/craft something
func (c *Character) CanBuy(sell *Sell) bool {
	if sell.Price.Value > c.Currency.CurValue {
		fmt.Println("You do not have enough golds.")
		return false
	}
	if (len(c.Inventory.Items) - len(sell.Price.Items)) >= c.Inventory.Slots {
		fmt.Println("You cannot buy this Item. Some resources are missing.")
		return false
	}
	var item_count int
	for item_name, s_count := range sell.Price.Items {
		for c_item, c_count := range c.Inventory.Items {
			if c_item.Name != item_name {
				continue
			}
			item_count++
			if s_count > c_count {
				fmt.Println("You cannot buy this Item. Some resources are missing.")
				return false
			}
		}
	}
	can_buy := item_count == len(sell.Price.Items)
	if !can_buy {
		fmt.Println("You cannot buy this Item. Some resources are missing.")
	}
	return can_buy
}

func (c *Character) BuyItem(sell *Sell) {
	c.Currency.UpdateStat(-sell.Price.Value)
	for item_name, count := range sell.Price.Items {
		c.Inventory.UseItem(Items[item_name], count)
	}
	// If the item is a book, learn the spell directly and do not add it
	// to the player's inventory
	if strings.Contains(sell.ShopItem.Name, "Book:") {
		return
	}
	c.Inventory.AddItem(sell.ShopItem, 1)
}

func (c *Character) DisplayInfos() {
	fmt.Println("--- Player infos ---")

	fmt.Println("Name:", c.Name)
	fmt.Println("Class:", c.Class.Name)
	fmt.Printf("Lives: %d\n", c.lives.CurValue)
	fmt.Printf("Golds: %d/%d\n", c.Currency.CurValue, c.Currency.MaxValue)
	fmt.Printf("Level: %d/%d\n", c.level.CurValue, c.level.MaxValue)
	fmt.Printf("Experience: %d/%d\n", c.Xp.CurValue, c.Xp.MaxValue)
	fmt.Printf("HP: %d/%d\n", c.Hp.CurValue, c.Hp.MaxValue)
	fmt.Printf("Mana: %d/%d\n", c.Mana.CurValue, c.Mana.MaxValue)
	fmt.Printf("Initiative: %d\n", c.Initiative.CurValue)
}

// Equip an Item and get back the other one equipped
// Update stats accordingly
func (c *Character) EquipItem(item *Item) {
	tmp := c.Inventory.Stuff.Equip(item)
	c.Hp.MaxValue += item.Stats[0].Value
	c.Hp.UpdateStat(item.Stats[0].Value)
	if tmp != nil {
		c.Hp.MaxValue -= tmp.Stats[0].Value
		c.Hp.UpdateStat(-tmp.Stats[0].Value)
	}
	c.Inventory.AddItem(tmp, 1)
}

// Consume item and do something accordingly to the item
func (c *Character) ConsumeItem(target Entity, item *Item) {
	name := item.Name
	stats := item.Stats
	if item.Type == 0 {
		switch name {
		case "Health Potion":
			target.TakeDamage(c, -stats[0].Value)
			break
		case "Mana Potion":
			c.Mana.UpdateStat(stats[0].Value)
			break
		case "Poison Potion":
			var dots int = 3
			for i := 0; i < dots; i++ {
				target.TakeDamage(c, stats[0].Value/dots)
				time.Sleep(1 * time.Second)
			}
		case "Upgrade inventory slot":
			c.Inventory.UpgradeInventorySlot()
		case "Bebou-chan's sock":
			fmt.Println(SockArt)
			fmt.Println("\nAn unpleaseant odor escapes... It's as hard as rock.\n")
			return
		}
		fmt.Printf("You used %s on %s.\n\n", name, target.GetName())
		c.Inventory.UseItem(item, 1)
	} else {
		c.Inventory.UseItem(item, 1)
		c.EquipItem(item)
	}
}
