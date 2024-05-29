package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"https://github.com/MohsenNasertabrizi/Spring2024/tree/main/Web/Web_midrerm/model"
	"https://github.com/MohsenNasertabrizi/Spring2024/tree/main/Web/Web_midrerm/request"
	"github.com/gorilla/mux"
)

func GetBaskets(w http.ResponseWriter, r *http.Request) {
	baskets, err := model.GetAllBaskets()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(baskets)
}

func GetBasket(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid basket ID", http.StatusBadRequest)
		return
	}

	basket, err := model.GetBasketByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if basket == nil {
		http.Error(w, "Basket not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(basket)
}

func CreateBasket(w http.ResponseWriter, r *http.Request) {
	var req request.BasketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	basket := &model.Basket{Status: req.Status}
	if err := model.CreateBasket(basket); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(basket)
}

func UpdateBasket(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid basket ID", http.StatusBadRequest)
		return
	}

	var req request.BasketRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	basket, err := model.GetBasketByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if basket == nil {
		http.Error(w, "Basket not found", http.StatusNotFound)
		return
	}

	if basket.Status == "COMPLETED" {
		http.Error(w, "Cannot update a completed basket", http.StatusBadRequest)
		return
	}

	basket.Status = req.Status
	if err := model.UpdateBasket(basket); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(basket)
}

func DeleteBasket(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid basket ID", http.StatusBadRequest)
		return
	}

	if err := model.DeleteBasket(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
