package game

import "fmt"

type Inventory struct {
	Slots int
	Items map[*Item]int
	Stuff *Stuff
}

// Add item to the inventory of the player
// Prevent if the inventory is full
func (i *Inventory) AddItem(item *Item, count int) bool {
	if item == nil {
		return true
	}
	value, ok := i.Items[item]
	if ok {
		i.Items[item] = value + count
		return true
	}
	if len(i.Items)+1 <= i.Slots {
		i.Items[item] = count
		return true
	}
	fmt.Println("You have no place left.")
	return false
}

// Use item in the inventory
func (i *Inventory) UseItem(item *Item, count int) {
	value, ok := i.Items[item]
	if ok && value >= count {
		i.Items[item] = value - count
		if value-count == 0 {
			delete(i.Items, item)
		}
	}
}

// Print the element in the inventory
func (i Inventory) DisplayItems() []*Item {
	i.Stuff.DisplayStuff()

	var item_arr []*Item = make([]*Item, 0)

	fmt.Println("--- Inventory ---")
	var j int = 0
	for key, val := range i.Items {
		item_arr = append(item_arr, key)
		fmt.Printf("%d\\ \"%s\": %d\n", j, key.Name, val)
		j++
	}

	return item_arr
}

// Init the inventory at the beginning
func InitBaseInventory() *Inventory {
	inv := new(Inventory)

	inv.Slots = 10
	inv.Items = make(map[*Item]int, 10)
	inv.Stuff = InitStuff()

	inv.AddItem(Items["Health Potion"], 3)

	return inv
}

// Upgrade how many slots the inventory can contain
func (i *Inventory) UpgradeInventorySlot() {
	i.Slots += 10
}
