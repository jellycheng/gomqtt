package gomqtt

import (
	"fmt"
	"testing"
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
