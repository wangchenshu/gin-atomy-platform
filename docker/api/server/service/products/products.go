package products

import (
	"gin-atomy-platform/server/db"
	"gin-atomy-platform/server/model"

	"github.com/gin-gonic/gin"

	"net/http"
)

func getProducts() []model.Products {
	products := []model.Products{}
	db.Db.Find(&products)

	return products
}

func getProductsLike(name string, limit int) []model.Products {
	products := []model.Products{}
	db.Db.Where("name LIKE ?", "%"+name+"%").Limit(limit).Find(&products)

	return products
}

// GetProducts -
func GetProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, getProducts())
	}
}

// GetProductsLike -
func GetProductsLike(name string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, getProductsLike(name, 100))
	}
}
