package app

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func connectMQTT(brokerURI string, clientID string) mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURI)
	opts.SetClientID(clientID)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {

	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}
