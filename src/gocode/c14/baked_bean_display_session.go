package c14

import (
	"fmt"
	"html/template"
	"net/http"

	"../util"
	"./bean"
)

type BakedBeanDisplaySession struct{}

func (bbds *BakedBeanDisplaySession) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = bbds.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (bbds *BakedBeanDisplaySession) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	session := util.SM.GetSession(w, r)
	bakedBean := session.Get("bakedBean")
	if bakedBean == nil {
		bakedBean = bean.NewBakedBean()
		session.Set("bakedBean", bakedBean)
	}

	err = util.PopulateScheStr(bakedBean, r)
	if err != nil {
		return
	}

	tpl := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Baked Bean Values: page-based Sharing</title>
</head>
<body>
<h1>Baked Bean Values: page-based Sharing</h1>
<h2>Bean level: {{.Level}}</h2>
<h2>Dish bean goes with: {{.GoesWith}}</h2>
</body>
</html>`
	t, err := template.New("csj").Parse(tpl)
	if err != nil {
		return
	}
	return t.Execute(w, bakedBean)
}
