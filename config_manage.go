package gomqtt

import (
	"errors"
	"sync"
)

type MqttConfigManage struct {
	lockObj         *sync.RWMutex
	mqttConfigGroup map[string]*MqttConfig
}

func (m *MqttConfigManage) Set(group string, mqttCfg *MqttConfig) *MqttConfigManage {
	m.lockObj.Lock()
	defer m.lockObj.Unlock()
	m.mqttConfigGroup[group] = mqttCfg
	return m
}

func (m MqttConfigManage) Get(group string) (*MqttConfig, error) {
	m.lockObj.RLock()
	defer m.lockObj.RUnlock()
	if ret, ok := m.mqttConfigGroup[group]; ok {
		return ret, nil
	} else {
		return nil, errors.New("配置不存在:" + group)
	}
}

func (m MqttConfigManage) GetAll() map[string]*MqttConfig {
	m.lockObj.RLock()
	defer m.lockObj.RUnlock()
	return m.mqttConfigGroup
}

func NewMqttConfigManage() *MqttConfigManage {
	ret := &MqttConfigManage{
		lockObj:         new(sync.RWMutex),
		mqttConfigGroup: make(map[string]*MqttConfig),
	}
	return ret
}
