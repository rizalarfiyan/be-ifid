package adapter

import (
	"be-ifid/config"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttConn *mqtt.Client

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	log.Println("Connection lost: ", err)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("MQTT connected")
}

func MQTTInit() {
	conf := config.Get()
	dsn := fmt.Sprintf("%s:%d", conf.MQTT.Server, conf.MQTT.Port)
	opts := mqtt.NewClientOptions()
	opts.AddBroker(dsn)
	opts.SetClientID(conf.MQTT.ClientId)
	opts.AutoReconnect = true
	opts.OnConnectionLost = connectLostHandler
	opts.OnConnect = connectHandler

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		log.Fatalln("MQTT Connect problem: ", token.Error())
	}

	mqttConn = new(mqtt.Client)
	mqttConn = &client
}

func MQTTGet() *mqtt.Client {
	return mqttConn
}
