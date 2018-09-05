package c9

type KidsBooksPage struct {
	CatalogPage
}

func NewKidsBooksPage() *KidsBooksPage {
	kbp := &KidsBooksPage{}
	kbp.SetItem([]string{"lewis001", "alexander001",
		"rowling001"})
	kbp.SetTitle("All-Time Best Children's Fantasy Books")
	return kbp
}
