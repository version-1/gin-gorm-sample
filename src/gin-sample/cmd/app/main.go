package main

import (
	"gin-sample/internal/models"
	"gin-sample/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()
	db := pkg.Connect("development")
	defer db.Close()
	db.LogMode(true)

	r.GET("/products", func(c *gin.Context) {
		var products []models.Product
		db.Find(&products)
		c.JSON(200, products)
	})

	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var product models.Product
		db.First(&product, id)
		c.JSON(200, product)
	})

	r.POST("/products", func(c *gin.Context) {
		product := models.Product{}
		err := c.Bind(&product)
		if err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
			return
		}
		db.NewRecord(&product)
		res := db.Create(&product)
		if res.Error != nil {
			c.JSON(402, res.Error)
		} else {
			c.JSON(200, nil)
		}
	})

	r.PUT("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		product := models.Product{}
		db.First(&product, id)

		params := models.Product{}
		err := c.Bind(&params)
		if err != nil {
			c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
			return
		}
		db.Model(&product).Updates(params)
		c.JSON(200, product)
	})

	r.DELETE("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		db.Where("ID = ?", id).Delete(&models.Product{})
		c.JSON(200, nil)
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
