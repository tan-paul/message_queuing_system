package controller

import (
	"fmt"
	model "message_queuing_system/model"
	producer "message_queuing_system/producer"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

var m = new(model.Models)

func (ctrl *Controller) ProcessProductDetails(c *gin.Context) {
	productdetails := model.RequestData{}
	err := c.ShouldBindJSON(&productdetails)
	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"Error": err})
		return
	}
	ok, err := m.IsAuthenticClient(productdetails.User_id)
	if err != nil {
		c.JSON(500, gin.H{"Error": "Internal Server Error"})
		return
	}
	if !ok {
		c.JSON(400, gin.H{"Error": "Invalid User"})
		return
	}
	if strings.Trim(productdetails.Product_name, " ") == "" {
		c.JSON(400, gin.H{"Error": "Product name required"})
		return
	}
	if strings.Trim(productdetails.Product_description, " ") == "" {
		c.JSON(400, gin.H{"Error": "Product description required"})
		return
	}
	if len(productdetails.Product_images) <= 0 {
		c.JSON(400, gin.H{"Error": "Product image urls required"})
		return
	}
	if productdetails.Product_price <= 0 {
		c.JSON(400, gin.H{"Error": "Product price invalid or not provided"})
		return
	}

	id, err := m.InsertIntoDatabase(productdetails)
	if err != nil {
		c.JSON(500, gin.H{"Error": err})
		return
	}
	fmt.Println("id for inserted product is :: ", id)
	pr := new(producer.Producer)
	err = pr.InsertIntoQueue(id)
	if err != nil {
		c.JSON(500, gin.H{"Error": err})
		return
	}
	c.JSON(201, gin.H{"Message": "Queued"})
	return
}
