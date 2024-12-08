package homecontroller

import (
	"net/http"
	"text/template"
	"go-web-native/models/productmodel"
)

func Welcome(w http.ResponseWriter, r *http.Request) {

	products := productmodel.GetAll()
	data := map[string]any{
		"products": products,
	}

	temp, err := template.ParseFiles("views/home/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}