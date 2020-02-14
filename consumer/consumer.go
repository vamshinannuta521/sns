package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"consumer/httpclient"
	"consumer/models"
	"consumer/smtp"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

var wg = sync.WaitGroup{}

var logger = logrus.NewEntry(logrus.New())

var client = httpclient.NewHTTPClient("admin", "Nutanix.123", &http.Transport{}, logger)

var kafkaAddr string

func init() {
	curDir, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}
	fileByte, err := ioutil.ReadFile(curDir + "/kafka.txt")
	if err != nil {
		logger.Fatal(err)
	}
	kafkaAddr = string(fileByte)

}

func main() {

	wg.Add(5)
	go consumeFromTopic("HTTP", &wg)
	go consumeFromTopic("SMS", &wg)
	go consumeFromTopic("EMAIL", &wg)
	go consumeFromTopic("CALM_SCALE_OUT", &wg)
	go consumeFromTopic("CALM_SCALE_DOWN", &wg)
	wg.Wait()
}

func consumeFromTopic(topic string, wg *sync.WaitGroup) {

	defer wg.Done()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaAddr},
		GroupID: "consumer-group-id",
		//	Partition: 0,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			logger.Error(err)
			break
		}

		logger.Infof("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		var action models.Action
		err = json.Unmarshal(m.Value, &action)
		if err != nil {
			logger.Error(err)
			continue
		}

		switch topic {
		case "HTTP":
			var actionSpec models.HTTPActionSpec
			err := json.Unmarshal([]byte(action.ActionSpec), &actionSpec)
			if err != nil {
				logger.Error(err)
				continue
			}
			_, _, err = client.Do(context.Background(), actionSpec.Method, actionSpec.URL, []byte(actionSpec.Body))
			if err != nil {
				logger.Error(err)
				continue
			}
		case "EMAIL":
			var actionSpec models.EmailActionSpec
			err := json.Unmarshal([]byte(action.ActionSpec), &actionSpec)
			if err != nil {
				logger.Error(err)
				continue
			}
			err = smtp.Send(&actionSpec)
			if err != nil {
				logger.Error(err)
				continue
			}

		case "CALM_SCALE_OUT":
			reqBody := `{"api_version":"3.0","metadata":{"project_reference":{"kind":"project","name":"default","uuid":"e53105fc-a498-4158-ac10-fa23d79ec66e"},"name":"SNS_TEST_APP","creation_time":"1581528249903526","spec_version":9,"kind":"app","last_update_time":"1581530108769015","uuid":"3dff19c9-9d4c-18b4-c65d-22dcd8f696ea"},"spec":{"target_uuid":"cb115218-2412-44f0-8f18-7432b8dfcb43","target_kind":"Application","args":[]}}`
			url := fmt.Sprintf("http://localhost:3000/api/nutanix/v3/apps/%s/actions/%s/run", "cb115218-2412-44f0-8f18-7432b8dfcb43", "ff61053e-15e2-4425-8526-b8c1c61e951e")
			_, _, err = client.Do(context.Background(), "POST", url, []byte(reqBody))
			if err != nil {
				logger.Error(err)
				continue
			}

		case "CALM_SCALE_DOWN":
			reqBody := `{"api_version":"3.0","metadata":{"project_reference":{"kind":"project","name":"default","uuid":"e53105fc-a498-4158-ac10-fa23d79ec66e"},"name":"SNS_TEST_APP","creation_time":"1581528249903526","spec_version":12,"kind":"app","last_update_time":"1581530985987996","uuid":"ef2bcc54-ed02-3d14-5fde-9e8254022399"},"spec":{"target_uuid":"cb115218-2412-44f0-8f18-7432b8dfcb43","target_kind":"Application","args":[]}}`
			url := fmt.Sprintf("http://localhost:3000/api/nutanix/v3/apps/%s/actions/%s/run", "cb115218-2412-44f0-8f18-7432b8dfcb43", "38ab5424-5e3d-4e81-88e9-aa71c7ce6c49")
			_, _, err = client.Do(context.Background(), "POST", url, []byte(reqBody))
			if err != nil {
				logger.Error(err)
				continue
			}

		default:
			logger.Infof("unexpected topic: %s", topic)
			continue

		}

	}

}
