package c9

type TechBooksPage struct {
	CatalogPage
}

func NewTechBooksPage() *TechBooksPage {
	tbp := &TechBooksPage{}
	tbp.SetItem([]string{"hall001", "hall002"})
	tbp.SetTitle("All-Time Best Computer Books")
	return tbp
}
