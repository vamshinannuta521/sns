package kafka

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

actionTopicMap := make(map[string]string)
actionTopicMap[""]

func Push(string message, string actionType) error{
	topic := getTopicForActionType(actionType)
	topic := "my-topic"
partition := 0

conn, _ := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)

conn.SetWriteDeadline(time.Now().Add(10*time.Second))
conn.WriteMessages(
    kafka.Message{Value: []byte("one!")},
    kafka.Message{Value: []byte("two!")},
    kafka.Message{Value: []byte("three!")},
)

conn.Close()
}
func Push(parent context.Context, key, value []byte) (err error) {
	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}

	return writer.WriteMessages(parent, message)
}

func getTopicForActionType(actionType) string,error{
	val, ok := actionTopicMap[actionType]
	if ok {
		return val, nil
	}
	retrun nil,nil
}