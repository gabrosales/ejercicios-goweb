package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Estructura
type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Quantity     int     `json:"quantity"`
	Code_value   string  `json:"code_value"`
	Is_published bool    `json:"is_published"`
	Expiration   string  `json:"expiration"`
	Price        float64 `json:"price"`
}

// Variable que contendra la BD en memoria.
var products []Product

// Funcion de carga de datos
func LoadProducts() {
	jsonFile, err := os.ReadFile("products.json")
	if err != nil {
		fmt.Println(err)
	}
	// defer jsonFile.Close()

	// byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(jsonFile, &products)

}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GetAllProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Fail to parse id."})
		return
	}

	for _, pro := range products {
		if pro.ID == id {
			c.IndentedJSON(http.StatusOK, pro)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Id not found"})
}

func SearchProductsByPrice(c *gin.Context) {
	priceGt, err := strconv.Atoi(c.Query("priceGt"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Fail to parse price."})
		return
	}

	var productsFiltered []Product

	for _, pro := range products {
		if pro.Price > float64(priceGt) {
			productsFiltered = append(productsFiltered, pro)
		}
	}

	if productsFiltered != nil {
		c.IndentedJSON(http.StatusOK, productsFiltered)
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "There is no products with that price."})

}

func main() {
	LoadProducts()

	// Server gin
	r := gin.Default()

	//PING
	r.GET("/ping", Ping)
	//Produtcs
	r.GET("/products", GetAllProducts)
	//Product:id
	r.GET("/products/:id", GetProductById)

	//Product search x price
	r.GET("/products/search", SearchProductsByPrice)

	r.Run() // escucha en 0.0.0.0:8080 por defecto

}
