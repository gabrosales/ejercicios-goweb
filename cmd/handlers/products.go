package handlers

// import (
// 	"errors"
// 	"net/http"
// 	"strconv"
// 	"time"

// 	"ejercicios-goweb/pkg/response"
// 	"ejercicios-goweb/services"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// )

// const SECRET_KEY = "ABC"

// var (
// 	ErrUnauthorized = errors.New("error: invalid token")
// 	ErrParameter    = errors.New("error: wrong parameter")
// 	ErrNotFound     = errors.New("error: products not found")
// 	ErrAlreadyExist = errors.New("error: product already exist")
// )

// func Ping(c *gin.Context) {
// 	result := services.Ping()
// 	c.JSON(http.StatusOK, response.Ok("suceed", result))
// }

// func GetAllProducts(c *gin.Context) {
// 	token := c.GetHeader("token")
// 	if token != SECRET_KEY {
// 		c.JSON(http.StatusUnauthorized, response.Err(ErrUnauthorized))
// 		return
// 	}

// 	products := services.GetAllProducts()

// 	c.JSON(http.StatusOK, response.Ok("succeed to get websites", products))
// }

// func GetProductById(c *gin.Context) {
// 	token := c.GetHeader("token")

// 	if token != SECRET_KEY {
// 		c.JSON(http.StatusUnauthorized, response.Err(ErrUnauthorized))
// 		return
// 	}

// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, response.Err(ErrParameter))
// 		return
// 	}

// 	product, err := services.GetProductById(id)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, response.Err(err))
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, response.Ok("suceed to get product", product))

// }

// func SearchProductsByPrice(c *gin.Context) {
// 	priceGt, err := strconv.Atoi(c.Query("priceGt"))

// 	if err != nil {
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Fail to parse price."})
// 		return
// 	}

// 	productsFiltered := services.SearchProductsByPrice(priceGt)

// 	if productsFiltered != nil {
// 		c.IndentedJSON(http.StatusOK, response.Ok("suceed to get products", productsFiltered))
// 	}

// 	c.IndentedJSON(http.StatusInternalServerError, response.Err(ErrNotFound))

// }

// const (
// 	layout = "02/01/2006"
// )

// func DateValidation(fl validator.FieldLevel) bool {
// 	_, err := time.Parse(layout, fl.Field().String())
// 	if err != nil {
// 		return false
// 	}
// 	return true
// }

// type Request struct {
// 	Name         string  `json:"name" validate:"required"`
// 	Quantity     int     `json:"quantity" validate:"required"`
// 	Code_value   string  `json:"code_value" validate:"required"`
// 	Is_published bool    `json:"is_published"`
// 	Expiration   string  `json:"expiration" validate:"required,date"`
// 	Price        float64 `json:"price" validate:"required"`
// }

// func CreateProduct(ctx *gin.Context) {
// 	// request
// 	var req Request

// 	if err := ctx.ShouldBind(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, response.Err(err))
// 		return
// 	}

// 	//validation
// 	validate := validator.New()
// 	validate.RegisterValidation("date", DateValidation)
// 	if err := validate.Struct(&req); err != nil {
// 		ctx.JSON(http.StatusUnprocessableEntity, response.Err(err))
// 		return
// 	}

// 	// process
// 	product, err := services.Create(req.Name, req.Quantity, req.Code_value, req.Is_published, req.Expiration, req.Price)

// 	if err != nil {
// 		if errors.Is(err, ErrAlreadyExist) {
// 			ctx.JSON(http.StatusConflict, response.Err(err))
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, response.Err(err))
// 		return
// 	}

// 	// response
// 	ctx.JSON(http.StatusCreated, response.Ok("suceed to create product", product))
// }
