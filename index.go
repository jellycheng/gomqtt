package gomqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMqttClient(opts *mqtt.ClientOptions) (mqtt.Client, error) {
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return client, nil
}

// Publish 发消息
func Publish(client mqtt.Client, topic string, qos byte, retained bool, payload interface{}) error {
	token := client.Publish(topic, qos, retained, payload)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Subscribe 订阅消息/消费消息
func Subscribe(client mqtt.Client, topic string, qos byte, callback mqtt.MessageHandler) error {
	token := client.Subscribe(topic, qos, callback)
	//token.Wait()
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
