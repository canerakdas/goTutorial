package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"waste/controllers"
	"waste/models"
	"waste/util"
)

func main() {
	var productController controllers.Products
	var result = make([]models.Product,1,5)
	err := util.Conn.Find(nil).Sort("-timestamp").All(&result)

	fmt.Printf("%v",result)

	// Route
	r := mux.NewRouter()

	r.HandleFunc("/product", productController.Get).Methods("GET")
	r.HandleFunc("/product/{id}", productController.GetWithId).Methods("GET")

	r.HandleFunc("/product", productController.Post).Methods("POST")
	r.HandleFunc("/product/{id}", productController.PostWithId).Methods("POST")

	r.HandleFunc("/product/{id}", productController.Put).Methods("PUT")

	r.HandleFunc("/product/{id}", productController.Patch).Methods("PATCH")

	r.HandleFunc("/product/{id}", productController.Delete).Methods("DELETE")

	err = http.ListenAndServe(":8080",r)

	if err != nil{
		log.Fatal("Went wrong")
	}

}