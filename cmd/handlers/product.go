package handlers

import (
	"ejercicios-goweb/internal/product"
	"ejercicios-goweb/pkg/response"
	"net/http"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	layout = "02/01/2006"
)

func DateValidation(fl validator.FieldLevel) bool {
	_, err := time.Parse(layout, fl.Field().String())
	if err != nil {
		return false
	}
	return true
}

type Product struct {
	sv product.Service
}

func NewProduct(sv product.Service) *Product {
	return &Product{sv: sv}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request

		// process
		products, err := p.sv.GetAll()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.Err(err))
			return
		}

		// response
		ctx.JSON(http.StatusOK, response.Ok("succeed to get products", products))
	}
}

func (p *Product) GetByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Err(err))
			return
		}

		// process
		product, err := p.sv.GetByID(id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.Err(err))
			return
		}

		// response
		ctx.JSON(http.StatusOK, response.Ok("succeed to get product", product))
	}
}

func (p *Product) SearchProductsByPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request
		priceGt, err := strconv.Atoi(ctx.Query("priceGt"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Err(err))
			return
		}

		// process
		products, err := p.sv.SearchProductsByPrice(priceGt)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.Err(err))
			return
		}

		// response
		ctx.JSON(http.StatusOK, response.Ok("succeed to get products", products))
	}
}

func (p *Product) Create() gin.HandlerFunc {
	type Request struct {
		Name         string  `json:"name" validate:"required"`
		Quantity     int     `json:"quantity" validate:"required"`
		Code_value   string  `json:"code_value" validate:"required"`
		Is_published bool    `json:"is_published"`
		Expiration   string  `json:"expiration" validate:"required,date"`
		Price        float64 `json:"price" validate:"required"`
	}

	return func(ctx *gin.Context) {
		// request
		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.Err(err))
			return
		}

		//validation
		validate := validator.New()
		validate.RegisterValidation("date", DateValidation)
		if err := validate.Struct(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
			return
		}

		// process
		pro, err := p.sv.Create(req.Name, req.Quantity, req.Code_value, req.Is_published, req.Expiration, req.Price)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.Err(err))
			return
		}

		// response
		ctx.JSON(http.StatusOK, response.Ok("succeed to create product", pro))

	}
}

func (p *Product) Update() gin.HandlerFunc {
	type RequestPut struct {
		Name         string  `json:"name" validate:"required"`
		Quantity     int     `json:"quantity" validate:"required"`
		Code_value   string  `json:"code_value" validate:"required"`
		Is_published bool    `json:"is_published"`
		Expiration   string  `json:"expiration" validate:"required,date"`
		Price        float64 `json:"price" validate:"required"`
	}

	return func(ctx *gin.Context) {
		// request
		var req RequestPut
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.Err(err))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Err(err))
			return
		}

		//validation
		validate := validator.New()
		validate.RegisterValidation("date", DateValidation)
		if err := validate.Struct(&req); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
			return
		}

		// process
		pro, err := p.sv.Update(id, req.Name, req.Quantity, req.Code_value, req.Is_published, req.Expiration, req.Price)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.Err(err))
			return
		}

		// response
		ctx.JSON(http.StatusOK, response.Ok("succeed to update product", pro))

	}
}

func (p *Product) PartialUpdate() gin.HandlerFunc {
	type Request struct {
		Name         string  `json:"name,omitempty"`
		Quantity     int     `json:"quantity,omitempty"`
		Code_value   string  `json:"code_value,omitempty"`
		Is_published bool    `json:"is_published,omitempty"`
		Expiration   string  `json:"expiration,omitempty"`
		Price        float64 `json:"price,omitempty"`
	}
	return func(ctx *gin.Context) {
		// request
		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, response.Err(err))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, response.Err(err))
			return
		}

		// process
		pro, err := p.sv.Update(id, req.Name, req.Quantity, req.Code_value, req.Is_published, req.Expiration, req.Price)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, response.Err(err))
			return
		}

		// response
		ctx.JSON(http.StatusOK, response.Ok("succeed to partial update product", pro))
	}
}
