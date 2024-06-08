package connection

import (
	"github.com/streadway/amqp"
)

func NewRabbitMQConnection(dsn string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CloseRabbitMQConnection(conn *amqp.Connection) {
	if conn != nil {
		_ = conn.Close()
	}
}
