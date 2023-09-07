package gomqtt

const (
	QoS0 byte = 0 //最多一次，即：<=1，消息可能丢失
	QoS1 byte = 1 //至少一次，即：>=1，消息不会丢失，但可能重复
	QoS2 byte = 2 //一次，即：=1，消息不会丢失，也不会重复
)
