package database

import "github.com/streadway/amqp"

var RabbitmqConn *amqp.Connection

func ConnectRabbitMQServer() error {
	RabbitmqConn, err = amqp.Dial("amqp://guest:guest@127.0.0.1:5672")
	if err != nil {
		return err
	}
	return nil
}

func DisconnectRabbitMQServer(conn *amqp.Connection) error {
	err := conn.Close()
	if err != nil {
		return err
	}
	return nil
}
