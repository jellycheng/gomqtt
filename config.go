package gomqtt

type MqttConfig struct {
	Broker   string //地址，如： tcp://broker.emqx.io:1883
	Username string //账号
	Pwd      string //密码
	ClientID string // 客户端id
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
