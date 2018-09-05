package shoppingcart

import (
	"encoding/gob"
)

type ItemOrder struct {
	item     CatalogItem
	numItems int
}

func init() {
	gob.Register(&ItemOrder{})
}

func NewItemOrder(item CatalogItem) *ItemOrder {
	itemOrder := &ItemOrder{}
	itemOrder.item = item
	itemOrder.numItems = 1
	return itemOrder
}

func (itemO *ItemOrder) GetItem() CatalogItem {
	return itemO.item
}

func (itemO *ItemOrder) SetItem(item CatalogItem) {
	itemO.item = item
}

func (itemO *ItemOrder) GetItemID() string {
	return itemO.item.ItemID
}

func (itemO *ItemOrder) GetShortDescription() string {
	return itemO.item.ShortDescription
}

func (itemO *ItemOrder) GetLongDescription() string {
	return itemO.item.LongDescription
}

func (itemO *ItemOrder) GetUnitCost() float64 {
	return itemO.item.Cost
}

func (itemO *ItemOrder) GetNumItems() int {
	return itemO.numItems
}

func (itemO *ItemOrder) SetNumItems(n int) {
	itemO.numItems = n
}

func (itemO *ItemOrder) IncrementNumItems() {
	itemO.numItems++
}

func (itemO *ItemOrder) CancelOrder() {
	itemO.numItems = 0
}

func (itemO *ItemOrder) GetTotalCost() float64 {
	return float64(itemO.numItems) * itemO.item.Cost
}
