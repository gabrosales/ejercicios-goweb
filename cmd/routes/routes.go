package routes

import (
	"ejercicios-goweb/cmd/handlers"
	"ejercicios-goweb/internal/domain"
	"ejercicios-goweb/internal/product"

	"github.com/gin-gonic/gin"
)

type Router struct {
	db *[]domain.Product
	en *gin.Engine
}

func NewRouter(en *gin.Engine, db *[]domain.Product) *Router {
	return &Router{en: en, db: db}
}

func (r *Router) SetRoutes() {
	r.SetProduct()
}

// product
func (r *Router) SetProduct() {
	// instances
	rp := product.NewRepository(r.db, 500)
	sv := product.NewService(rp)
	h := handlers.NewProduct(sv)

	pro := r.en.Group("/products")
	pro.GET("", h.GetAll())
	pro.GET("/:id", h.GetByID())
	pro.GET("/search", h.SearchProductsByPrice())
	pro.POST("", h.Create())
}
