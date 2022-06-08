package gildedrose

type AbstractItem interface {
	Update()
}

type ItemSulf struct {
	item *Item
}
type ItemBrie struct {
	item *Item
}
type ItemPass struct {
	item *Item
}
type ItemConj struct {
	item *Item
}
type ItemNormal struct {
	item *Item
}

func (a *ItemSulf) Update() {
}

func (a *ItemBrie) Update() {
	a.item.SellIn -= 1
	a.item.Quality += 1
	if a.item.SellIn < 0 {
		a.item.Quality += 1
	}
	PostCheck(a.item)
}
func (a *ItemPass) Update() {
	a.item.SellIn -= 1
	if a.item.SellIn < 0 {
		a.item.Quality = 0
		return
	}

	a.item.Quality += 1
	if a.item.SellIn < 5 {
		a.item.Quality += 2
	} else if a.item.SellIn < 10 {
		a.item.Quality += 1
	}
	PostCheck(a.item)
}
func (a *ItemConj) Update() {
	a.item.SellIn -= 1
	a.item.Quality -= 2

	if a.item.SellIn < 0 {
		a.item.Quality -= 2
	}
	PostCheck(a.item)
}
func (a *ItemNormal) Update() {
	a.item.SellIn -= 1
	a.item.Quality -= 1
	if a.item.SellIn < 0 {
		a.item.Quality -= 1
	}
	PostCheck(a.item)
}

func PostCheck(item *Item) {
	if item.Quality < 0 {
		item.Quality = 0
	} else if item.Quality > 50 {
		item.Quality = 50
	}
}

func UpdateItemQuality(item AbstractItem) {
	item.Update()
}

func convert(item *Item) AbstractItem {
	var cust AbstractItem

	if item.Name == "Sulfuras, Hand of Ragnaros" {
		cust = &ItemSulf{item}
	} else if item.Name == "Aged Brie" {
		cust = &ItemBrie{item}
	} else if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		cust = &ItemPass{item}
	} else if item.Name == "Conjured Mana Cake" {
		cust = &ItemConj{item}
	} else {
		cust = &ItemNormal{item}
	}

	return cust
}

func UpdateQualityInterface(items []*Item) {
	for _, item := range items {
		abi := convert(item)
		UpdateItemQuality(abi)
	}
}
