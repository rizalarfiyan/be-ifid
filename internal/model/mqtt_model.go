package model

type MQTTData struct {
	DeviceId string     `json:"device_id"`
	Sensor   MQTTSensor `json:"sensor"`
}

type MQTTSensor struct {
	Humidity    float64 `json:"humidity"`
	Temperature float64 `json:"temperature"`
	Gas         float64 `json:"gas"`
}
