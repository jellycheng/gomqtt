package gomqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttConfig struct {
	Broker       string //mqtt地址，示例：tcp://broker.emqx.io:1883
	Username     string //账号
	Pwd          string //密码
	ClientID     string //客户端id
	clientOption *mqtt.ClientOptions
}

func (m *MqttConfig) GetClientOptions() *mqtt.ClientOptions {
	opts := m.clientOption
	if opts == nil {
		opts = mqtt.NewClientOptions()
		opts.AddBroker(m.Broker)
		opts.SetClientID(m.ClientID) // 客户端ID
		opts.SetUsername(m.Username) // 账号名
		opts.SetPassword(m.Pwd)      //密码
		m.clientOption = opts
	}

	return opts
}

type MqttConfigOption func(m *MqttConfig)

func NewMqttConfig(broker string, opts ...MqttConfigOption) *MqttConfig {
	ret := &MqttConfig{
		Broker: broker,
	}
	for _, opt := range opts {
		opt(ret)
	}
	return ret
}

func WithMqttConfigBroker(broker string) MqttConfigOption {
	return func(m *MqttConfig) {
		m.Broker = broker
	}
}

func WithMqttConfigUsername(username string) MqttConfigOption {
	return func(m *MqttConfig) {
		m.Username = username
	}
}

func WithMqttConfigPwd(pwd string) MqttConfigOption {
	return func(m *MqttConfig) {
		m.Pwd = pwd
	}
}

func WithMqttConfigClientID(id string) MqttConfigOption {
	return func(m *MqttConfig) {
		m.ClientID = id
	}
}
