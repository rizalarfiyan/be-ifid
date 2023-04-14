package service

import (
	"be-ifid/config"
	"be-ifid/internal/model"
	"encoding/json"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTService interface {
	Subscibe()
}

type mqttService struct {
	conn mqtt.Client
	conf *config.Config
}

func NewMQTTService(conn mqtt.Client, conf *config.Config) MQTTService {
	return &mqttService{
		conn: conn,
		conf: conf,
	}
}

func (ctx *mqttService) Subscibe() {
	topic := ctx.conf.MQTT.Topic.Watch
	log.Println("Set subscribe topic: ", topic)

	token := ctx.conn.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		data := model.MQTTData{}
		err := json.Unmarshal([]byte(msg.Payload()), &data)
		if err != nil {
			return
		}
		fmt.Println(data)
	})

	token.Wait()
	log.Println("Subscribed to topic: ", topic)
}
