package game

import "fmt"

type Stuff struct {
	Helmet     *Item
	Chestplate *Item
	Boots      *Item
}

func InitStuff() *Stuff {
	var e *Stuff = new(Stuff)

	return e
}

func (s Stuff) DisplayStuff() {
	fmt.Println("--- Stuff ---")
	helmet_name := ""
	if s.Helmet != nil {
		helmet_name = s.Helmet.Name
	}
	chestplate_name := ""
	if s.Chestplate != nil {
		chestplate_name = s.Chestplate.Name
	}
	boots_name := ""
	if s.Boots != nil {
		boots_name = s.Boots.Name
	}

	fmt.Printf("Helmet: %s\n", helmet_name)
	fmt.Printf("Chestplate: %s\n", chestplate_name)
	fmt.Printf("Boots: %s\n", boots_name)
}

func (s *Stuff) Equip(item *Item) *Item {
	var tmp *Item
	switch item.Type {
	case HELMET:
		tmp = s.Helmet
		s.Helmet = item
	case CHESTPLATE:
		tmp = s.Chestplate
		s.Chestplate = item
	case BOOTS:
		tmp = s.Boots
		s.Boots = item
	}
	return tmp
}
