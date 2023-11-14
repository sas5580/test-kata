package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_RegularItem(t *testing.T) {
	items := []gildedrose.DepreciatingItem{
		gildedrose.NewItem("foo", 5, 20),
	}

	updateQualityN(items, 1)
	checkItem(t, items[0], 4, 19)

	updateQualityN(items, 5)
	checkItem(t, items[0], -1, 13)

	updateQualityN(items, 6)
	checkItem(t, items[0], -7, 1)

	updateQualityN(items, 1)
	checkItem(t, items[0], -8, 0)

	updateQualityN(items, 10)
	checkItem(t, items[0], -18, 0)
}

func Test_AgedBrie(t *testing.T) {
	items := []gildedrose.DepreciatingItem{
		gildedrose.NewItem("Aged Brie", 5, 30),
	}

	updateQualityN(items, 5)
	checkItem(t, items[0], 0, 35)

	updateQualityN(items, 1)
	checkItem(t, items[0], -1, 37)

	updateQualityN(items, 6)
	checkItem(t, items[0], -7, 49)

	updateQualityN(items, 13)
	checkItem(t, items[0], -20, 50)
}

func Test_Sulfuras(t *testing.T) {
	items := []gildedrose.DepreciatingItem{
		gildedrose.NewItem("Sulfuras, Hand of Ragnaros", 0, 80),
	}

	updateQualityN(items, 2)
	checkItem(t, items[0], 0, 80)

	updateQualityN(items, 4)
	checkItem(t, items[0], 0, 80)

	updateQualityN(items, 8)
	checkItem(t, items[0], 0, 80)
}

func Test_BackstagePasses(t *testing.T) {
	items := []gildedrose.DepreciatingItem{
		gildedrose.NewItem("Backstage passes to a TAFKAL80ETC concert", 15, 20),
	}

	updateQualityN(items, 5)
	checkItem(t, items[0], 10, 25)

	updateQualityN(items, 1)
	checkItem(t, items[0], 9, 27)
	updateQualityN(items, 4)
	checkItem(t, items[0], 5, 35)

	updateQualityN(items, 1)
	checkItem(t, items[0], 4, 38)
	updateQualityN(items, 4)
	checkItem(t, items[0], 0, 50)

	updateQualityN(items, 1)
	checkItem(t, items[0], -1, 0)
	updateQualityN(items, 19)
	checkItem(t, items[0], -20, 0)
}

func Test_Conjured(t *testing.T) {
	items := []gildedrose.DepreciatingItem{
		gildedrose.NewItem("Conjured Laptop", 5, 20),
	}

	updateQualityN(items, 1)
	checkItem(t, items[0], 4, 18)

	updateQualityN(items, 4)
	checkItem(t, items[0], 0, 10)

	updateQualityN(items, 1)
	checkItem(t, items[0], -1, 6)

	updateQualityN(items, 1)
	checkItem(t, items[0], -2, 2)

	updateQualityN(items, 1)
	checkItem(t, items[0], -3, 0)

	updateQualityN(items, 17)
	checkItem(t, items[0], -20, 0)
}

func checkItem(t *testing.T, item gildedrose.DepreciatingItem, expectedSellIn, expectedQuality int) {
	if item.GetSellIn() != expectedSellIn {
		t.Errorf("SellIn: expected %v, got %v", expectedSellIn, item.GetSellIn())
	}

	if item.GetQuality() != expectedQuality {
		t.Errorf("Quality: expected %v, got %v", expectedQuality, item.GetQuality())
	}
}

func updateQualityN(items []gildedrose.DepreciatingItem, n int) {
	for i := 0; i < n; i++ {
		gildedrose.UpdateQuality(items)
	}
}
