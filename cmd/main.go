package main

import (
	"ejercicios-goweb/cmd/routes"
	"ejercicios-goweb/internal/domain"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// func main() {

// 	services.LoadProducts()

// 	// Server gin
// 	r := gin.Default()

// 	//router
// 	products := r.Group("/products")
// 	products.GET("/ping", handlers.Ping)
// 	products.GET("", handlers.GetAllProducts)
// 	products.GET("/:id", handlers.GetProductById)
// 	products.GET("/search", handlers.SearchProductsByPrice)
// 	products.POST("", handlers.CreateProduct)
// 	r.Run() // escucha en 0.0.0.0:8080 por defecto

// }

var db []domain.Product

func LoadProducts() {
	jsonFile, err := os.ReadFile("products.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(jsonFile, &db)

}

func main() {

	// instances
	LoadProducts()

	en := gin.Default()
	rt := routes.NewRouter(en, &db)
	rt.SetRoutes()

	if err := en.Run(); err != nil {
		log.Fatal(err)
	}
}
