package queue

type MessageQueue interface {
	Consume()
}
