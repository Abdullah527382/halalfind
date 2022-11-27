package controllers

import (
	"github.com/openfoodfacts/openfoodfacts-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFoods(c *gin.Context) {

	 // Retrieve the food id string from url
	 foodId := c.Param("foodid")
	 // Create a response message which includes the food id
	//  foodItem := fmt.Sprintf("Variable string %s content", foodId)
	 
	 api := openfoodfacts.NewClient("world", "", "")
	 product, err := api.Product(foodId)
	 // foodObject := {
	// Send the JSON object back to client
	c.JSON(http.StatusOK, gin.H{
		"message": product,
		"err": err,
	})
	return
}