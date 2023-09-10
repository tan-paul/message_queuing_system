package main

import (
	"fmt"
	database "message_queuing_system/database"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := database.ConnectMysqlDB()
	if err != nil {
		panic(err)
	}
	defer database.DisConnectMysqlDB(database.Database_Conn)
	err = database.ConnectRabbitMQServer()
	if err != nil {
		panic(err)
	}
	defer database.DisconnectRabbitMQServer(database.RabbitmqConn)

	flag := false
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// handeling interrupts or running consumer
	for {
		select {
		case sig := <-sigs:
			fmt.Println("got signal :: ", sig, " hence exiting...")
			flag = true
			break
		default:
			consumer()

			// adding delay of 3 sec to avoid excess get call to queue ; it will effect interupt signal as well
			time.Sleep(3 * time.Second)
		}
		if flag {
			break
		}
	}
}
