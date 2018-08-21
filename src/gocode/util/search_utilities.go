package util

var commonSpecs []*SearchSpec = []*SearchSpec{
	&SearchSpec{"Google", "http://www.google.com/search?q="},
	&SearchSpec{"Yahoo", "http://search.yahoo.com/bin/search?p="},
	&SearchSpec{"Baidu", "http://www.baidu.com/baidu?word="},
}

func GetCommonSpecs() []*SearchSpec {
	return commonSpecs
}

func MakeURL(searchEngineName, searchString string) string {
	searchURL := ""
	for i := 0; i < len(commonSpecs); i++ {
		spec := commonSpecs[i]
		if spec.Name == searchEngineName {
			searchURL = spec.MakeURL(searchString)
			break
		}
	}
	return searchURL
}
