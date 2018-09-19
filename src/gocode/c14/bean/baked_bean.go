package bean

type BakedBean struct {
	Level    string `schema:"level"`
	GoesWith string `schema:"goesWith"`
}

func NewBakedBean() *BakedBean {
	return &BakedBean{
		Level:    "half-baked",
		GoesWith: "hot dogs",
	}
}
