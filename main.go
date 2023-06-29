package main

import (
	"log"
	"net/http"

	"github.com/autumnleaf-ra/go-crud-web/config"
	"github.com/autumnleaf-ra/go-crud-web/controllers/categorycontroller"
	"github.com/autumnleaf-ra/go-crud-web/controllers/homecontroller"
	"github.com/autumnleaf-ra/go-crud-web/controllers/productcontroller"
)

func main() {
	config.ConnectDB()

	// Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// Categories
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", categorycontroller.Add)
	http.HandleFunc("/categories/edit", categorycontroller.Edit)
	http.HandleFunc("/categories/delete", categorycontroller.Delete)

	// Products
	http.HandleFunc("/products", productcontroller.Index)
	http.HandleFunc("/products/add", productcontroller.Add)
	http.HandleFunc("/products/detail", productcontroller.Detail)
	http.HandleFunc("/products/edit", productcontroller.Edit)
	http.HandleFunc("/products/delete", productcontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)

}
