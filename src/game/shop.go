package game

import "fmt"

type Price struct {
	Value int
	Items map[string]int
}

type Sell struct {
	ShopItem *Item
	Price    *Price
	Count    int
}

func InitSell(item *Item, price *Price, count int) *Sell {
	var sell *Sell = new(Sell)
	sell.ShopItem = item
	sell.Price = price
	sell.Count = count
	return sell
}

// Sell the item
func (s *Sell) UseSell() {
	if s.Count-1 >= 0 {
		s.Count -= 1
	}
}

type Shop struct {
	Sells []*Sell
}

func InitShop() *Shop {
	var shop *Shop = new(Shop)
	shop.Sells = make([]*Sell, 0)
	return shop
}

func (s *Shop) DisplaySells() {
	for i, sell := range s.Sells {
		fmt.Printf("%d\\ %s\n", i, sell.ShopItem.Name)
		fmt.Printf("\tPrice: %d\n", sell.Price.Value)
		fmt.Printf("\tCount: %d\n", sell.Count)
		if len(sell.Price.Items) > 0 {
			fmt.Printf("\tItems:\n")
		}
		for name, count := range sell.Price.Items {
			fmt.Printf("\t\t%s : %d\n", name, count)
		}
	}
	fmt.Println()
}

func (s *Shop) AddSells(sells []*Sell) {
	s.Sells = append(s.Sells, sells...)
}

func (s *Shop) AddSell(sell *Sell) {
	s.Sells = append(s.Sells, sell)
}

func (s *Shop) UseSell(sell *Sell) {
	if sell.Count > 0 {
		sell.UseSell()
	}
}
