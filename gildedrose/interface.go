package gildedrose

type ItemOp interface {
	Update()
}
type ItemBase struct {
	op   ItemOp
	item *Item
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
}
func (a *ItemConj) Update() {
	a.item.SellIn -= 1
	a.item.Quality -= 2

	if a.item.SellIn < 0 {
		a.item.Quality -= 2
	}
}
func (a *ItemNormal) Update() {
	a.item.SellIn -= 1
	a.item.Quality -= 1
	if a.item.SellIn < 0 {
		a.item.Quality -= 1
	}
}

func UpdateQualityIB(ib *ItemBase) {
	ib.op.Update()
	ib.postCheck()
}

func (a *ItemBase) postCheck() {
	if a.item.Name == "Sulfuras, Hand of Ragnaros" {
		return
	}

	if a.item.Quality < 0 {
		a.item.Quality = 0
	} else if a.item.Quality > 50 {
		a.item.Quality = 50
	}
}

func ItemBaseInit(item *Item, cust ItemOp) *ItemBase {
	base := &ItemBase{item: item}
	base.op = cust
	return base
}

func convert(item *Item) *ItemBase {
	var cust ItemOp

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

	return ItemBaseInit(item, cust)
}

func UpdateQualityTemplate(items []*Item) {
	for _, item := range items {
		ib := convert(item)
		UpdateQualityIB(ib)
	}
}
