package kafka

import (
	"context"
	"time"

	kf "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

//to be changed as part of kafka client

type KafkaQueue struct {
	logger      *logrus.Entry
	endpointUrl string
}

func GetQueueClient(logg *logrus.Entry) *KafkaQueue {

	endpointUrl := "10.46.143.17:9092"

	return &KafkaQueue{logg, endpointUrl}

}

func (k *KafkaQueue) Push(message []byte, actionType string) error {

	conn, err := kf.DialLeader(context.Background(), "tcp", k.endpointUrl, actionType, 0)
	if err != nil {
		k.logger.Error(err)
		return err
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(kf.Message{Value: message})

	conn.Close()
	return nil
}
