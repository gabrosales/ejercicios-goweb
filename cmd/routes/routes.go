package routes

import (
	"ejercicios-goweb/cmd/handlers"
	"ejercicios-goweb/internal/domain"
	"ejercicios-goweb/internal/product"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	//Load .env
	_ = godotenv.Load()

	// instances
	rp := product.NewRepository(r.db, 500)
	sv := product.NewService(rp)
	h := handlers.NewProduct(sv)

	pro := r.en.Group("/products")
	pro.GET("", h.GetAll())
	pro.GET(":id", h.GetByID())
	pro.GET("/search", h.SearchProductsByPrice())
	pro.POST("", h.Create())
	pro.PUT(":id", h.Update())
	pro.PATCH(":id", h.PartialUpdate())
	pro.DELETE(":id", h.Delete())

}
