package gildedrose

type ItemOp interface {
	Update()
}
type ItemBase struct {
	op   ItemOp
	item *Item
}

type ItemSulf2 struct {
	item *Item
}
type ItemBrie2 struct {
	item *Item
}
type ItemPass2 struct {
	item *Item
}
type ItemConj2 struct {
	item *Item
}
type ItemNormal2 struct {
	item *Item
}

func (a *ItemSulf2) Update() {
}
func (a *ItemBrie2) Update() {
	a.item.SellIn -= 1
	a.item.Quality += 1
	if a.item.SellIn < 0 {
		a.item.Quality += 1
	}
}
func (a *ItemPass2) Update() {
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
func (a *ItemConj2) Update() {
	a.item.SellIn -= 1
	a.item.Quality -= 2

	if a.item.SellIn < 0 {
		a.item.Quality -= 2
	}
}
func (a *ItemNormal2) Update() {
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
	base := &ItemBase{item: item, op: cust}
	return base
}

func convertIB(item *Item) *ItemBase {
	var cust ItemOp

	if item.Name == "Sulfuras, Hand of Ragnaros" {
		cust = &ItemSulf2{item}
	} else if item.Name == "Aged Brie" {
		cust = &ItemBrie2{item}
	} else if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		cust = &ItemPass2{item}
	} else if item.Name == "Conjured Mana Cake" {
		cust = &ItemConj2{item}
	} else {
		cust = &ItemNormal2{item}
	}

	return ItemBaseInit(item, cust)
}

func UpdateQualityTemplate(items []*Item) {
	for _, item := range items {
		ib := convertIB(item)
		UpdateQualityIB(ib)
	}
}
