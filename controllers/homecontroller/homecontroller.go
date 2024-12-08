package homecontroller

import (
	"net/http"
	"text/template"
	"go-web-native/models/productmodel"
	"strconv"
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

func Detail(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	product := productmodel.Detail(id)
	data := map[string]any{
		"product": product,
	}

	temp, err := template.ParseFiles("views/home/detail.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}