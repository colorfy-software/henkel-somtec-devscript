package mqttx

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"google.golang.org/protobuf/proto"
)

type Message interface {
	mqtt.Message
	UnmarshalProto() (proto.Message, error)
}

type message struct {
	mqtt.Message
}

func (msg *message) UnmarshalProto() (proto.Message, error) {
	return UnmarshalProto(msg.Payload())
}
