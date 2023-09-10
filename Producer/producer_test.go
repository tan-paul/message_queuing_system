package producer

import (
	database "message_queuing_system/database"

	"testing"
)

/*
Here 0-9 is product id, trying to push to Rabbitmq exchange :: Product_Message_Exchange with
routing key as :: Product_Message_Queue

note :- before pushing you have to make sure that you have one exchange,queue with the above name
to see the product id in UI of RMQ else gonna get error.
*/
func TestInsertIntoQueue(t *testing.T) {
	err := database.ConnectRabbitMQServer()
	if err != nil {
		panic(err)
	}
	p := new(Producer)
	defer database.DisconnectRabbitMQServer(database.RabbitmqConn)
	for product_id := 0; product_id < 10; product_id++ {
		if err := p.InsertIntoQueue(int64(product_id)); err != nil {
			t.Errorf("internal error ::  %q", product_id)
		}
	}
}
func BenchmarkInsertIntoQueue(b *testing.B) {
	err := database.ConnectRabbitMQServer()
	if err != nil {
		panic(err)
	}
	p := new(Producer)
	defer database.DisconnectRabbitMQServer(database.RabbitmqConn)
	for i := 0; i < b.N; i++ {
		p.InsertIntoQueue(int64(33))
	}
}
