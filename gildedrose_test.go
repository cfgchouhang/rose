package main

import (
	//"fmt"
	"gildedrose/gildedrose"
	"testing"
)

func Test_Origin(t *testing.T) {
	/*var itemsa = []*gildedrose.Item{
		{"apple", 5, 20},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Aged Brie", 10, 2},
		{"Aged Brie", 1, 0},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 10},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 10},
		{"+5 Dexterity Vest", 10, 20},
		//{"Conjured Mana Cake", 20, 40},
	}*/
	itemsa := gildedrose.GetItems()
	itemsb := gildedrose.GetItems()
	itemsc := gildedrose.GetItems()
	itemsd := gildedrose.GetItems()
	itemse := gildedrose.GetItems()

	for day := 0; day < 15; day++ {
		gildedrose.UpdateQuality(itemsa)
		gildedrose.UpdateQualityRev(itemsb)
		gildedrose.UpdateQualityMap(itemsc)
		gildedrose.UpdateQualityInterface(itemsd)
		gildedrose.UpdateQualityTemplate(itemse)

		for i := 0; i < len(itemsa); i++ {
			//fmt.Printf("day%d, %s left %d %d\t %d %d\n", day+1, itemsa[i].Name, itemsa[i].SellIn, itemsb[i].SellIn, itemsa[i].Quality, itemsb[i].Quality)
			if itemsa[i].Quality != itemsb[i].Quality {
				t.Errorf("After %d days, rev %d:%s quality a:%d not match with b:%d", day, i, itemsa[i].Name, itemsa[i].Quality, itemsb[i].Quality)
			}
			if itemsa[i].SellIn != itemsb[i].SellIn {
				t.Errorf("After %d days, rev %d:%s sellin a:%d not match with b:%d", day, i, itemsa[i].Name, itemsa[i].SellIn, itemsb[i].SellIn)
			}

			if itemsa[i].Quality != itemsc[i].Quality {
				t.Errorf("After %d days, map %d:%s quality a:%d not match with b:%d", day, i, itemsa[i].Name, itemsa[i].Quality, itemsc[i].Quality)
			}
			if itemsa[i].SellIn != itemsc[i].SellIn {
				t.Errorf("After %d days, map %d:%s sellin a:%d not match with b:%d", day, i, itemsa[i].Name, itemsa[i].SellIn, itemsc[i].SellIn)
			}

			if itemsa[i].Quality != itemsd[i].Quality {
				t.Errorf("After %d days, interface %d:%s quality a:%d not match with b:%d", day, i, itemsa[i].Name, itemsa[i].Quality, itemsd[i].Quality)
			}
			if itemsa[i].SellIn != itemsd[i].SellIn {
				t.Errorf("After %d days, interface %d:%s sellin a:%d not match with b:%d", day, i, itemsa[i].Name, itemsa[i].SellIn, itemsd[i].SellIn)
			}

			if itemsa[i].Quality != itemse[i].Quality {
				t.Errorf("After %d days, templ %d:%s quality a:%d not match with b:%d", day, i, itemsa[i].Name, itemsa[i].Quality, itemse[i].Quality)
			}
			if itemsa[i].SellIn != itemse[i].SellIn {
				t.Errorf("After %d days, templ %d:%s sellin a:%d not match with b:%d", day, i, itemsa[i].Name, itemsa[i].SellIn, itemse[i].SellIn)
			}
		}
	}
}

func Test_ConjuredRev(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", 5, 10},
		{"Conjured Mana Cake", 5, 0},
		{"Conjured Mana Cake", 0, 10},
	}

	expected := [][]int{{4, 8}, {4, 0}, {-1, 6}}

	gildedrose.UpdateQualityRev(items)
	for i := 0; i < len(items); i++ {
		if items[i].Quality != expected[i][1] {
			t.Errorf("items %d quality expect %d but %d", i, expected[i][1], items[i].Quality)
		}
		if items[i].SellIn != expected[i][0] {
			t.Errorf("items %d sellin expect %d but %d", i, expected[i][0], items[i].Quality)
		}
	}
}

func Test_ConjuredMap(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", 5, 10},
		{"Conjured Mana Cake", 5, 0},
		{"Conjured Mana Cake", 0, 10},
	}

	expected := [][]int{{4, 8}, {4, 0}, {-1, 6}}

	gildedrose.UpdateQualityMap(items)
	for i := 0; i < len(items); i++ {
		if items[i].Quality != expected[i][1] {
			t.Errorf("items %d quality expect %d but %d", i, expected[i][1], items[i].Quality)
		}
		if items[i].SellIn != expected[i][0] {
			t.Errorf("items %d sellin expect %d but %d", i, expected[i][0], items[i].Quality)
		}
	}
}

func Test_ConjuredInterface(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", 5, 10},
		{"Conjured Mana Cake", 5, 0},
		{"Conjured Mana Cake", 0, 10},
	}

	expected := [][]int{{4, 8}, {4, 0}, {-1, 6}}

	gildedrose.UpdateQualityInterface(items)
	for i := 0; i < len(items); i++ {
		if items[i].Quality != expected[i][1] {
			t.Errorf("items %d quality expect %d but %d", i, expected[i][1], items[i].Quality)
		}
		if items[i].SellIn != expected[i][0] {
			t.Errorf("items %d sellin expect %d but %d", i, expected[i][0], items[i].Quality)
		}
	}
}

func Test_ConjuredTemplate(t *testing.T) {
	var items = []*gildedrose.Item{
		{"Conjured Mana Cake", 5, 10},
		{"Conjured Mana Cake", 5, 0},
		{"Conjured Mana Cake", 0, 10},
	}

	expected := [][]int{{4, 8}, {4, 0}, {-1, 6}}

	gildedrose.UpdateQualityTemplate(items)
	for i := 0; i < len(items); i++ {
		if items[i].Quality != expected[i][1] {
			t.Errorf("items %d quality expect %d but %d", i, expected[i][1], items[i].Quality)
		}
		if items[i].SellIn != expected[i][0] {
			t.Errorf("items %d sellin expect %d but %d", i, expected[i][0], items[i].Quality)
		}
	}
}
