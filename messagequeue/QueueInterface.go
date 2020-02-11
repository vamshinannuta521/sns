package messagequeue

type QueueInterface interface {
	GetQueueClient(*logrus.Entry) QueueClient
}

type QueueClient interface {
	Push(string, string) error
}