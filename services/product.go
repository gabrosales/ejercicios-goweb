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

func Ping() string {
	return "pong"
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

func GetLastId() int {
	product := products[len(products)-1]
	return product.ID
}

func ExistProduct(code_value string) bool {
	for _, p := range products {
		if p.Code_value == code_value {
			return true
		}
	}
	return false
}

func Create(name string, quantity int, code_value string, is_published bool, expiration string, price float64) (models.Product, error) {
	// validations
	if ExistProduct(code_value) {
		return models.Product{}, errors.New("error: product already exist")
	}

	lastId := GetLastId()
	lastId++
	product := models.Product{
		ID:           lastId,
		Name:         name,
		Quantity:     quantity,
		Code_value:   code_value,
		Is_published: is_published,
		Expiration:   expiration,
		Price:        price,
	}

	products = append(products, product)
	return product, nil
}
