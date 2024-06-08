package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	NONE int = iota
	HELMET
	CHESTPLATE
	BOOTS
)

type Item struct {
	Name  string
	Type  int
	Stats []*Stat
}

func NewItem(name string, stats []*Stat) *Item {
	item := new(Item)

	item.Name = name
	item.Stats = stats

	return item
}

// Duplicate items
func (i *Item) CopyItem(item Item) {
	i.Name = item.Name
	i.Stats = make([]*Stat, 0)
	i.Stats = append(i.Stats, item.Stats...)
}

// Parse the json file and put items in a global map
func InitItems() {
	Items = make(map[string]*Item, 0)

	filename := "../config/items.json"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error parsing file", filename, err)
	}
	var payload map[string][]*Item
	err = json.Unmarshal(content, &payload)
	if err != nil {
		fmt.Println("Error Unmarshal()", filename, err)
	}
	for _, item := range payload["Items"] {
		Items[item.Name] = item
	}
}
