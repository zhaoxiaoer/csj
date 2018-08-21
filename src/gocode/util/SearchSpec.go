package util

type SearchSpec struct {
	Name    string
	baseURL string
}

func (ss *SearchSpec) MakeURL(searchString string) string {
	return ss.baseURL + searchString
}
