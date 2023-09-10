package main

import (
	controller "message_queuing_system/controller"
	database "message_queuing_system/database"

	"github.com/gin-gonic/gin"
)

func main() {
	//Connecting with  mysql
	err := database.ConnectMysqlDB()
	if err != nil {
		panic(err)
	}
	defer database.DisConnectMysqlDB(database.Database_Conn)

	//Connecting with Rabbitmq
	err = database.ConnectRabbitMQServer()
	if err != nil {
		panic(err)
	}
	defer database.DisconnectRabbitMQServer(database.RabbitmqConn)

	// Creating a gin router with default middleware
	router := gin.Default()

	ctrl := new(controller.Controller)
	router.POST("/product_info", ctrl.ProcessProductDetails)

	router.Run(":8080") // listen and serve on port 8080
}
