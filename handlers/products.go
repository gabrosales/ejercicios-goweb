package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"ejercicios-goweb/pkg/response"
	"ejercicios-goweb/services"

	"github.com/gin-gonic/gin"
)

const SECRET_KEY = "ABC"

var (
	ErrUnauthorized = errors.New("error: invalid token")
	ErrParameter    = errors.New("error: wrong parameter")
	ErrNotFound     = errors.New("error: products not found")
)

func Ping(c *gin.Context) {
	result := services.Ping()
	c.JSON(http.StatusOK, response.Ok("suceed", result))
}

func GetAllProducts(c *gin.Context) {
	token := c.GetHeader("token")
	if token != SECRET_KEY {
		c.JSON(http.StatusUnauthorized, response.Err(ErrUnauthorized))
		return
	}

	products := services.GetAllProducts()

	c.JSON(http.StatusOK, response.Ok("succeed to get websites", products))
}

func GetProductById(c *gin.Context) {
	token := c.GetHeader("token")

	if token != SECRET_KEY {
		c.JSON(http.StatusUnauthorized, response.Err(ErrUnauthorized))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, response.Err(ErrParameter))
		return
	}

	product, err := services.GetProductById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, response.Err(err))
		return
	}

	c.IndentedJSON(http.StatusOK, response.Ok("suceed to get product", product))

}

func SearchProductsByPrice(c *gin.Context) {
	priceGt, err := strconv.Atoi(c.Query("priceGt"))

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Fail to parse price."})
		return
	}

	productsFiltered := services.SearchProductsByPrice(priceGt)

	if productsFiltered != nil {
		c.IndentedJSON(http.StatusOK, response.Ok("suceed to get products", productsFiltered))
	}

	c.IndentedJSON(http.StatusInternalServerError, response.Err(ErrNotFound))

}
