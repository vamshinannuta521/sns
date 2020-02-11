package messagequeue

type QueueFactory interface {
	GetQueueClient(*logrus.Entry) QueueClient
}

type QueueClient interface {
	Push(string, string) error
}