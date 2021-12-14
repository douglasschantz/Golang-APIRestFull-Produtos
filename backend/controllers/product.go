package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schantz/web/go-api-produtos/backend/database"
	models "github.com/schantz/web/go-api-produtos/backend/models/product"
	"github.com/schantz/web/go-api-produtos/backend/utils"
)

func AllProducts(c *gin.Context) {

	product := []models.Product{}

	database.ReturnDB().Select([]string{"id", "product_code", "COALESCE(description,' ') as description"}).Find(&product)

	for _, v := range product {
		fmt.Println(v.Id)
	}

	utils.RespondWithJSON(c, http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product models.Product

	json.NewDecoder(c.Request.Body).Decode(&product)

	database.ReturnDB().Select("product_code", "description").Create(&product)

	utils.RespondWithJSON(c, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func UpdateProduct(c *gin.Context) {
	var product models.Product

	id := c.Param("id")

	json.NewDecoder(c.Request.Body).Decode(&product)
	fmt.Println(id)
	fmt.Println(product)

	database.ReturnDB().Model(&product).Where("id = ?", id).Updates(models.Product{ProductCode: product.ProductCode, Description: product.Description})
	utils.RespondWithJSON(c, http.StatusOK, map[string]string{"message": "successfully update"})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	database.ReturnDB().Where("id = ?", id).Delete(models.Product{})
	utils.RespondWithJSON(c, http.StatusOK, map[string]string{"message": "successfully delete"})
}
