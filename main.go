package main

import (
	"ejercicios-goweb/handlers"
	"ejercicios-goweb/services"

	"github.com/gin-gonic/gin"
)

func main() {

	services.LoadProducts()

	// Server gin
	r := gin.Default()

	//router
	products := r.Group("/products")
	products.GET("/ping", handlers.Ping)
	products.GET("", handlers.GetAllProducts)
	products.GET("/:id", handlers.GetProductById)
	products.GET("/search", handlers.SearchProductsByPrice)
	// products.POST("", handlers.CreateProduct)
	r.Run() // escucha en 0.0.0.0:8080 por defecto

}
