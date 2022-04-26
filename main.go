package main

import (
	"flag"
	"log"
	"net/url"
	"runtime"

	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"

	datafeedV1 "go.stockbit.io/protos/gen/go/securities/transactional/datafeed/v1"
)

var client = flag.Int("client", 2, "client count")

// -------------------------------------

var addr = flag.String("addr", "localhost:8001", "http service address")
var scheme = "ws"
var natsURL = "nats://localhost:4222"
var stocks = []string{"BBCA", "BBNI", "TLKM", "ADMR"}

// -------------------------------------

var natsJSClient nats.JetStreamContext
var natsKvUsersKey nats.KeyValue

func main() {
	flag.Parse()
	log.SetFlags(0)

	InitNATS()

	for i := 0; i < *client; i++ {
		go func(i int) {
			ClientWSDial(i)
		}(i)
	}

	runtime.Goexit()
}

func ClientWSDial(i int) {
	u := url.URL{Scheme: scheme, Host: *addr, Path: "/ws"}
	log.Printf("client-%d connecting to %s\n", i, u.String())

	// generate wskey to nats kv
	natsKvUsersKey.Create("wskey", []byte("1"))

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	m := datafeedV1.WebsocketRequest{
		UserId: "1234",
		Key:    "wskey",
		Channel: &datafeedV1.WebsocketChannel{
			RunningTrade: make([]string, 0),
		},
	}
	m.Channel.RunningTrade = append(m.Channel.RunningTrade, stocks...)
	byteP, _ := proto.Marshal(&m)

	err = c.WriteMessage(websocket.BinaryMessage, byteP)

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("msg:", message)
			log.Println("err read:", err)
			return
		}

		var msg datafeedV1.WebsocketWrapMessageChannel
		_ = proto.Unmarshal(message, &msg)

		switch wrap := msg.MessageChannel.(type) {
		case *datafeedV1.WebsocketWrapMessageChannel_RunningTrade:
			log.Printf("recv msg: %+v\n", wrap.RunningTrade)
		}

	}
}
