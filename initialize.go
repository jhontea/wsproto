package main

import (
	"fmt"
	"log"
	"project/gowebsocket/model"
	"time"

	"github.com/nats-io/nats.go"
)

func InitNATS() {
	var err error

	userKey := model.KeyValue{
		Bucket:  "bucket-uk",
		TTL:     10,
		History: 1,
	}

	config := model.Config{
		Nats: model.NatsConfig{
			BrokerURL: natsURL,
			UsersKey:  userKey,
		},
	}
	natsJSClient, err = NewNatsJetstreamClient(config)
	if err != nil {
		fmt.Println("err nats js", err)
		panic(err)
	}

	natsKvUsersKey, err = NewNatsKeyValue(natsJSClient, config.Nats.UsersKey)
	if err != nil {
		fmt.Println("error NewRepositoryNats init Users Key", err)
		panic(err)
	}
}

func NewNatsJetstreamClient(config model.Config) (nats.JetStreamContext, error) {
	nc, err := nats.Connect(config.Nats.BrokerURL)
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	return js, nil
}

func NewNatsKeyValue(js nats.JetStreamContext, config model.KeyValue) (nats.KeyValue, error) {
	kv, err := keyValueBacket(js, config.Bucket, time.Duration(config.TTL)*time.Second, config.History)
	if err != nil {
		return nil, err
	}

	return kv, nil
}

func keyValueBacket(client nats.JetStreamContext, bucket string, ttl time.Duration, history int) (nats.KeyValue, error) {
	kv, err := client.KeyValue(bucket)
	if err != nil {
		kv, err = client.CreateKeyValue(
			&nats.KeyValueConfig{
				Bucket:  bucket,
				TTL:     ttl,
				History: uint8(history),
			},
		)
		if err != nil {
			log.Println("bucket", bucket)
			log.Println("Error create bucket")
			return nil, err
		}
	}
	return kv, nil
}
