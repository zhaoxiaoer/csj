package shoppingcart

import (
	"sync"
)

import (
	"encoding/gob"
)

type ShoppingCart struct {
	sync.Mutex
	itemsOrdered []*ItemOrder
}

func init() {
	gob.Register(&ShoppingCart{})
}

func NewShoppingCart() *ShoppingCart {
	sc := &ShoppingCart{}
	sc.itemsOrdered = make([]*ItemOrder, 0)
	return sc
}

// 获取商品订单
func (sc *ShoppingCart) GetItemsOrdered() []*ItemOrder {
	return sc.itemsOrdered
}

func (sc *ShoppingCart) AddItem(itemID string) {
	sc.Lock()
	defer sc.Unlock()

	for i := 0; i < len(sc.itemsOrdered); i++ {
		order := sc.itemsOrdered[i]
		if order.GetItemID() == itemID {
			order.IncrementNumItems()
			return
		}
	}
	newOrder := NewItemOrder(*GetItem(itemID))
	sc.itemsOrdered = append(sc.itemsOrdered, newOrder)
}

func (sc *ShoppingCart) SetNumOrdered(itemID string, numOrdered int) {
	sc.Lock()
	defer sc.Unlock()

	//	fmt.Printf("len(sc.itemsOrdered): %v\n", len(sc.itemsOrdered))

	for i := 0; i < len(sc.itemsOrdered); i++ {
		order := sc.itemsOrdered[i]
		if order.GetItemID() == itemID {
			if numOrdered <= 0 {
				sc.itemsOrdered = append(sc.itemsOrdered[:i], sc.itemsOrdered[i+1:]...)
			} else {
				order.SetNumItems(numOrdered)
			}
			return
		}
	}
	newOrder := NewItemOrder(*GetItem(itemID))
	newOrder.SetNumItems(numOrdered)
	sc.itemsOrdered = append(sc.itemsOrdered, newOrder)
}
