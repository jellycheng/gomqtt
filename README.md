# gomqtt
```

```

## Requirements
```
gomqtt library requires Go version >=1.14

```

## 拉取代码
```
go get -u github.com/jellycheng/gomqtt
    或者
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/jellycheng/gomqtt

直接拉取master分支代码：
    go get -u github.com/jellycheng/gomqtt@master

```

## 发布消息
```
package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jellycheng/gomqtt"
	"time"
)

func main() {
	// 连接的回调
	var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
		fmt.Println("连接成功")
	}
	// 连接丢失的回调
	var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
		fmt.Printf("连接丢失Connect lost: %v", err)
	}

	cfg1 := gomqtt.NewMqttConfig("tcp://broker.emqx.io:1883",
		gomqtt.WithMqttConfigClientID("go_mqtt_client_pub"))
	opts := cfg1.GetClientOptions()
	opts = cfg1.GetClientOptions()
	opts = cfg1.GetClientOptions()
	opts.OnConnect = connectHandler            //连接的回调
	opts.OnConnectionLost = connectLostHandler //连接丢失的回调
	client, err := gomqtt.NewMqttClient(opts)
	if err != nil {
		fmt.Println("new client error")
		return
	}

	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("发送消息信息： %d", i)
		err := gomqtt.Publish(client, "topic/test", gomqtt.QoS0, false, text)
		if err != nil {
			fmt.Println("消息发送失败：", err.Error())
		} else {
			fmt.Println("发送成功：", text)
		}
		time.Sleep(time.Second)
	}

}

```

## 订阅消息
```
package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jellycheng/gomqtt"
	"time"
)

func main() {
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

	cfg1 := gomqtt.NewMqttConfig("tcp://broker.emqx.io:1883",
		gomqtt.WithMqttConfigClientID("go_mqtt_client_sub01"))
	opts := cfg1.GetClientOptions()
	opts.OnConnect = connectHandler            //连接的回调
	opts.OnConnectionLost = connectLostHandler //连接丢失的回调
	opts.SetDefaultPublishHandler(messagePubHandler)
	client, err := gomqtt.NewMqttClient(opts)
	if err != nil {
		fmt.Println("new client error")
		return
	}
	// 订阅消息
	gomqtt.Subscribe(client, "topic/test", gomqtt.QoS1, nil)

	time.Sleep(5 * time.Minute)

}

```
