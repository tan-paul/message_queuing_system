package producer

import (
	database "message_queuing_system/database"
	"strconv"

	"github.com/streadway/amqp"
)

type Producer struct{}

func (pr *Producer) InsertIntoQueue(product_id int64) error {
	ch, err := database.RabbitmqConn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	err = ch.Publish(
		"Product_Message_Exchange", // exchange name
		"Product_Message_Queue",    // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(strconv.FormatInt(product_id, 10)), //product id
		},
	)
	if err != nil {
		return err
	}
	return nil
}
