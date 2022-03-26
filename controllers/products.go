package controllers

import (
	"html/template"
	"log"
	"net/http"
	"store/models"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.GetAllProducts()

	temp.ExecuteTemplate(w, "index", allProducts)

}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("price convertion error : ", err)
		}
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			log.Println("quantity conversion error: ", err)
		}

		models.NewProduct(name, description, price, quantity)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	product := models.EditProduct(id)
	temp.ExecuteTemplate(w, "edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("id conversion error: ", err)
		}
		name := r.FormValue("name")
		description := r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("price convertion error : ", err)
		}
		quantity, err := strconv.Atoi(r.FormValue("quantity"))
		if err != nil {
			log.Println("quantity conversion error: ", err)
		}

		models.UpdatedProduct(name, description, price, id, quantity)
	}

	http.Redirect(w, r, "/", 301)
}
