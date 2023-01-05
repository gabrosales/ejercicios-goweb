package services

import (
	"ejercicios-goweb/services/models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Variable que contendra la BD en memoria.
var products []models.Product

// Funcion de carga de datos
func LoadProducts() {
	jsonFile, err := os.ReadFile("products.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(jsonFile, &products)

}

func GetAllProducts() []models.Product {
	return products
}

func GetProductById(id int) (models.Product, error) {

	for _, pro := range products {
		if pro.ID == id {
			return pro, nil
		}
	}
	return models.Product{}, errors.New("error: product not found")
}

func SearchProductsByPrice(price int) []models.Product {

	var productsFiltered []models.Product

	for _, pro := range products {
		if pro.Price > float64(price) {
			productsFiltered = append(productsFiltered, pro)
		}
	}

	return productsFiltered

}

func Ping() string {
	return "pong"
}
