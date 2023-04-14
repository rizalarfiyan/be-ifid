package adapter

import (
	"be-ifid/config"
	"be-ifid/utils"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var mqttConn *mqtt.Client

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	utils.SafeError("MQTT connection lost: ", err)
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	utils.Success("MQTT connected")
}

func MQTTInit() {
	utils.Info("Connect MQTT server...")
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
		utils.SafeError("MQTT Connect problem: ", token.Error())
	}

	mqttConn = new(mqtt.Client)
	mqttConn = &client
}

func MQTTGet() *mqtt.Client {
	return mqttConn
}

func MQTTIsConnected() bool {
	if mqttConn == nil {
		return false
	}
	connected := (*mqttConn).IsConnected()
	if !connected {
		utils.SafeError("MQTT fails health check")
	}
	return connected
}
