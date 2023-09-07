package gomqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"testing"
	"time"
)

func TestNewMqttConfigManage(t *testing.T) {
	cfg1 := NewMqttConfig("tcp://localhost:1883")

	cfgManage := NewMqttConfigManage()
	cfgManage.Set("default", cfg1)
	editUser := WithMqttConfigUsername("emqx")
	editUser(cfg1)

	if c, err := cfgManage.Get("default"); err == nil {
		fmt.Println(fmt.Sprintf("%+v", c))
	}

	cfg2 := NewMqttConfig("tcp://127.0.0.1:1883",
		WithMqttConfigClientID("hello-id1"),
		WithMqttConfigPwd("pubic"))
	cfgManage.Set("device01", cfg2)
	if c, err := cfgManage.Get("device01"); err == nil {
		fmt.Println(fmt.Sprintf("%+v", c))
	}

	all := cfgManage.GetAll()
	fmt.Println(fmt.Sprintf("%+v", all))
}

func TestPublish(t *testing.T) {
	// 连接的回调
	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("连接成功")
	}
	// 连接丢失的回调
	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("连接丢失Connect lost: %v", err)
	}

	cfg1 := NewMqttConfig("tcp://broker.emqx.io:1883",
		WithMqttConfigClientID("go_mqtt_client_pub"))
	opts := cfg1.GetClientOptions()
	opts = cfg1.GetClientOptions()
	opts = cfg1.GetClientOptions()
	opts.OnConnect = connectHandler            //连接的回调
	opts.OnConnectionLost = connectLostHandler //连接丢失的回调
	client, err := NewMqttClient(opts)
	if err != nil {
		fmt.Println("new client error")
		return
	}

	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("发送消息信息： %d", i)
		err := Publish(client, "topic/test", QoS0, false, text)
		if err != nil {
			fmt.Println("消息发送失败：", err.Error())
		} else {
			fmt.Println("发送成功：", text)
		}
		time.Sleep(time.Second)
	}
}

func TestSubscribe(t *testing.T) {
	// 连接的回调
	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("连接成功")
	}
	// 连接丢失的回调
	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("连接丢失Connect lost: %v", err)
	}
	// 全局 MQTT 消息处理
	var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("接收消息: %s from topic: %s\n", msg.Payload(), msg.Topic())
	}

	cfg1 := NewMqttConfig("tcp://broker.emqx.io:1883",
		WithMqttConfigClientID("go_mqtt_client_sub01"))
	opts := cfg1.GetClientOptions()
	opts.OnConnect = connectHandler            //连接的回调
	opts.OnConnectionLost = connectLostHandler //连接丢失的回调
	opts.SetDefaultPublishHandler(messagePubHandler)
	client, err := NewMqttClient(opts)
	if err != nil {
		fmt.Println("new client error")
		return
	}

	Subscribe(client, "topic/test", QoS1, nil)

	time.Sleep(5 * time.Minute)

}
