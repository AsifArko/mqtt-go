package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"net/url"
	"os"
	"time"
)

func connect(clientId string, uri *url.URL) mqtt.Client {

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))

	opts.SetClientID(clientId)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func main() {

	os.Setenv("CLOUDMQTT_URL", "tcp://localhost:1883")
	uri, err := url.Parse(os.Getenv("CLOUDMQTT_URL"))
	if err != nil {
		log.Fatal(err)
	}
	topic := "test"

	c := connect("pub", uri)

	for {
		c.Publish(topic, 0, false, fmt.Sprintf("Message from Pub"))
	}

}
