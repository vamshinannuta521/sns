package kafka

import (
	"context"
	"time"

	kf "github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

//to be changed as part of kafka client

type KafkaQueue struct {
	actionTopicMap map[string]string
	logger         *logrus.Entry
	endpointUrl    string
}

func GetQueueClient(logg *logrus.Entry) *KafkaQueue {
	actionTopicMap := make(map[string]string)
	actionTopicMap["http"] = "HTTP"
	actionTopicMap["sms"] = "SMS"
	actionTopicMap["email"] = "EMAIL"

	endpointUrl := "10.46.143.17:9092"

	return &KafkaQueue{actionTopicMap, logg, endpointUrl}

}

func (k *KafkaQueue) Push(message []byte, actionType string) error {
	//log := logrus.NewEntry(logrus.New())
	topic, _ := k.getTopicForActionType(actionType)
	if topic == "" {
		topic = "Auro"
	}

	k.logger.Error(topic)

	conn, err := kf.DialLeader(context.Background(), "tcp", k.endpointUrl, topic, 0)
	if err != nil {
		k.logger.Error(err)
		return err
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(kf.Message{Value: message})

	k.logger.Error("helllo")
	conn.Close()
	return nil
}

func (k *KafkaQueue) getTopicForActionType(actionType string) (string, error) {
	val, ok := k.actionTopicMap[actionType]
	if ok {
		return val, nil
	}
	return "", nil
}
