package handlers

import (
	. "api/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

var productStore = make(map[string]Product)
var id int = 0

// HTTP POST- /api/products
func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println(err)
	}

	product.CreatedOn = time.Now()
	id++
	product.ID = id
	key := strconv.Itoa(id)
	productStore[key] = product

	data, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// HTTP GET- /api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product
	for _, product := range productStore {
		products = append(products, product)
	}

	data, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)

}

// HTTP GET- /api/product/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	for _, prd := range productStore {
		if prd.ID == key {
			product = prd
		}
	}

	data, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func PutProductHandler(w http.ResponseWriter, r *http.Request) {

}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

}
