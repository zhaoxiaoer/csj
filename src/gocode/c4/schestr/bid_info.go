package schestr

type BidInfo struct {
	ItemID        string  `schema:"itemID"`
	ItemName      string  `schema:"itemName"`
	BidderName    string  `schema:"bidderName"`
	EmailAddress  string  `schema:"emailAddress"`
	BidPrice      float64 `schema:"bidPrice"`
	AutoIncrement bool    `schema:"autoIncrement"`
}

func (bi *BidInfo) IsComplete() bool {
	return (bi.hasValue(bi.ItemID) &&
		bi.hasValue(bi.ItemName) &&
		bi.hasValue(bi.BidderName) &&
		bi.hasValue(bi.EmailAddress) &&
		(bi.BidPrice > 0))
}

func (bi *BidInfo) IsPartlyComplete() bool {
	return (bi.hasValue(bi.ItemID) ||
		bi.hasValue(bi.ItemName) ||
		bi.hasValue(bi.BidderName) ||
		bi.hasValue(bi.EmailAddress) ||
		(bi.BidPrice > 0) ||
		bi.AutoIncrement)
}

func (bi *BidInfo) hasValue(v string) bool {
	return v != ""
}
