package gildedrose

type ConjuredPassesBuilder struct {
	SellIn  int
	Quality int
}

func newConjuredBuilder() *ConjuredPassesBuilder {
	return &ConjuredPassesBuilder{}
}

func (b *ConjuredPassesBuilder) calculSellIn(item *Item) {
	b.SellIn = item.SellIn - 1
}

func (b *ConjuredPassesBuilder) calculQuality(item *Item) {

	b.Quality = item.Quality

	if b.Quality > 0 {
		b.Quality = b.Quality - 2
	}
	if b.SellIn < 0 {
		if b.Quality > 0 {
			b.Quality = b.Quality - 2
		}
	}
}

func (b *ConjuredPassesBuilder) getCalculItem() CulcalItem {
	return CulcalItem{
		SellIn:  b.SellIn,
		Quality: b.Quality,
	}
}
