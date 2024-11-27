package controllers

import (
	"encoding/json"
	"go-pharmacy-api/models"
	"go-pharmacy-api/views"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetAllProducts()
	if err != nil {
		views.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	views.RespondJSON(w, http.StatusOK, products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		views.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := models.CreateProduct(&product); err != nil {
		views.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	views.RespondJSON(w, http.StatusCreated, product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		views.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := models.UpdateProduct(&product, id); err != nil {
		views.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	views.RespondJSON(w, http.StatusOK, map[string]string{"message": "Product updated successfully"})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := models.DeleteProduct(id); err != nil {
		views.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	views.RespondJSON(w, http.StatusOK, map[string]string{"message": "Product deleted successfully"})
}

func ResetProductsHandler(w http.ResponseWriter, r *http.Request) {
	if err := models.ResetProducts(); err != nil {
		views.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	views.RespondJSON(w, http.StatusOK, map[string]string{"message": "Products reset successfully"})
}
