package bean

import (
	"fmt"
	"strconv"
)

type SaleEntry struct {
	ItemID       string  `schema:"itemID"`
	NumItems     int     `schema:"numItems"`
	DiscountCode float64 `schema:"discountCode"`
}

func (se *SaleEntry) GetItemCost() float64 {
	var cost float64
	if se.ItemID == "a1234" {
		cost = 12.99 * se.DiscountCode
	} else {
		cost = -9999
	}
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", cost), 64)
	return v
}

func (se *SaleEntry) GetTotalCost() float64 {
	return se.GetItemCost() * float64(se.NumItems)
}
