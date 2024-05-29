package main

import (
	"log"
	"net/http"
	"https://github.com/MohsenNasertabrizi/Spring2024/tree/main/Web/Web_midrerm/handler"
	"https://github.com/MohsenNasertabrizi/Spring2024/tree/main/Web/Web_midrerm/model"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	model.InitDB("shopping_cart.db")
	defer model.CloseDB()

	r := mux.NewRouter()
	r.HandleFunc("/basket/", handler.GetBaskets).Methods("GET")
	r.HandleFunc("/basket/", handler.CreateBasket).Methods("POST")
	r.HandleFunc("/basket/{id}", handler.UpdateBasket).Methods("PATCH")
	r.HandleFunc("/basket/{id}", handler.GetBasket).Methods("GET")
	r.HandleFunc("/basket/{id}", handler.DeleteBasket).Methods("DELETE")

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
