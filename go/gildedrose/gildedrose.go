package gildedrose

import "strings"

const (
	defaultMaxQuality = 50

	agedBrieName        = "Aged Brie"
	sulfurasName        = "Sulfuras, Hand of Ragnaros"
	backstagePassesName = "Backstage passes to a TAFKAL80ETC concert"
	conjuredPrefix      = "Conjured"
)

type DepreciatingItem interface {
	UpdateQuality()
	GetName() string
	GetSellIn() int
	GetQuality() int
}

type Item struct {
	Name            string
	SellIn, Quality int
}

func NewItem(name string, sellIn, quality int) DepreciatingItem {
	i := Item{Name: name, SellIn: sellIn, Quality: quality}
	switch {
	case i.Name == agedBrieName:
		return &AgedBrieItem{i}
	case i.Name == sulfurasName:
		return &SulfurasItem{i}
	case i.Name == backstagePassesName:
		return &BackstagePassesItem{i}
	case strings.HasPrefix(i.Name, conjuredPrefix):
		return &ConjuredItem{i}
	default:
		return &i
	}
}

func (i *Item) GetName() string {
	return i.Name
}

func (i *Item) GetSellIn() int {
	return i.SellIn
}

func (i *Item) GetQuality() int {
	return i.Quality
}

func (i *Item) UpdateQuality() {
	i.SellIn -= 1

	decrement := 1
	if i.SellIn < 0 {
		decrement = 2
	}
	i.Quality = max(i.Quality-decrement, 0)
}

type AgedBrieItem struct {
	Item
}

func (i *AgedBrieItem) UpdateQuality() {
	i.SellIn -= 1

	increment := 1
	if i.SellIn < 0 {
		increment = 2
	}
	i.Quality = min(i.Quality+increment, defaultMaxQuality)
}

type SulfurasItem struct {
	Item
}

func (i *SulfurasItem) UpdateQuality() {
	// do nothing, Quality nor SellIn changes
}

type BackstagePassesItem struct {
	Item
}

func (i *BackstagePassesItem) UpdateQuality() {
	i.SellIn -= 1
	if i.SellIn < 0 {
		i.Quality = 0
		return
	}

	increment := 1
	if i.SellIn < 5 {
		increment = 3
	} else if i.SellIn < 10 {
		increment = 2
	}
	i.Quality = min(i.Quality+increment, defaultMaxQuality)
}

type ConjuredItem struct {
	Item
}

func (i *ConjuredItem) UpdateQuality() {
	prevQuality := i.GetQuality()
	i.Item.UpdateQuality()
	// degrade twice as fast as regular item
	decrement := prevQuality - i.Quality
	i.Quality = max(i.Quality-decrement, 0)
}

func UpdateQuality(items []DepreciatingItem) {
	for _, item := range items {
		item.UpdateQuality()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
