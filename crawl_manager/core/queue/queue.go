package queue

type MessageQueue interface {
	Send(topic, key string, data interface{}) error
}
