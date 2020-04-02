package app

import (
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const BROKER_URI = "192.168.1.100:1883"
const CLIENT_ID = "murgemachine-api"

const MQTT_TOPIC_PREPARATION = "murgemachine/preparation"
const MQTT_TOPIC_LIGHT = "murgemachine/light"

func connectMQTT() mqtt.Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(BROKER_URI)
	opts.SetClientID(CLIENT_ID)

	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {

	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}
