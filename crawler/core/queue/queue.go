package queue

type MessageQueue interface {
	Consume()
	Send(topic, key string, data interface{}) error
}
