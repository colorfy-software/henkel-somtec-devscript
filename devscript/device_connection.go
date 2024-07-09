package devscript

import (
	"crypto/tls"
	"fmt"
	"github.com/colorfy-software/henkel-somtec-devscript/libs/mqttx"
	iotdevpb "github.com/colorfy-software/henkel-somtec-devscript/proto/gen/colorfy/iot/device/v3"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/url"
	"time"
)

type DeviceConnection struct {
	ID       string
	Username string
	Password string
	Addr     string

	client      mqttx.Client
	updateTopic string
	eventTopic  string
}

func (conn *DeviceConnection) Connect() error {
	if conn.client != nil {
		return errors.New("already connected")
	}

	addr, err := url.Parse(conn.Addr)
	if err != nil {
		return err
	}

	conn.eventTopic = fmt.Sprintf("devices/%s/events", conn.ID)
	conn.updateTopic = fmt.Sprintf("devices/%s/updates", conn.ID)

	lastWill, err := mqttx.MarshalProto(&iotdevpb.PropertyUpdatedEvent{
		Property: &iotdevpb.Property{
			Name:       "connection_status",
			Value:      structpb.NewStringValue("disconnected"),
			UpdateTime: timestamppb.New(time.Now().Add(1 * time.Minute)),
		},
	})
	if err != nil {
		return err
	}

	clientOptions := &mqttx.ClientOptions{
		Servers:       []*url.URL{addr},
		ClientID:      conn.ID,
		Username:      conn.Username,
		Password:      conn.Password,
		AutoReconnect: true,
		WillEnabled:   true,
		WillQos:       1,
		WillPayload:   lastWill,
		WillRetained:  true,
		WillTopic:     conn.eventTopic,
		TLSConfig:     &tls.Config{InsecureSkipVerify: true},
	}

	client := mqttx.NewClient(clientOptions)

	if err := client.Connect(); err != nil {
		return err
	}

	conn.client = client

	return nil
}

func (conn *DeviceConnection) Disconnect() error {
	if conn.client == nil {
		return errors.New("not connected")
	}

	if err := conn.Publish(&iotdevpb.PropertyUpdatedEvent{
		Property: &iotdevpb.Property{
			Name:       "connection_status",
			Value:      structpb.NewStringValue("disconnected"),
			UpdateTime: timestamppb.New(time.Now().Add(1 * time.Minute)),
		},
	}); err != nil {
		return err
	}

	if err := conn.client.Disconnect(); err != nil {
		return err
	}
	conn.client = nil

	return nil
}

func (conn *DeviceConnection) Publish(msg proto.Message) error {
	b, err := mqttx.MarshalProto(msg)
	if err != nil {
		return err
	}

	if err := conn.client.Publish(conn.eventTopic, 1, false, b); err != nil {
		return err
	}

	return nil
}

func (conn *DeviceConnection) marshalProperty(property *Property) (*iotdevpb.Property, error) {
	value, err := structpb.NewValue(property.Value)
	if err != nil {
		return nil, err
	}

	if property.Writeable {
		return &iotdevpb.Property{
			Name:         property.Name,
			Status:       iotdevpb.Property_STATUS_CONFIRMED,
			Value:        value,
			DesiredValue: value,
			UpdateTime:   timestamppb.Now(),
		}, nil
	} else {
		return &iotdevpb.Property{
			Name:         property.Name,
			Status:       iotdevpb.Property_STATUS_CONFIRMED,
			Value:        value,
			DesiredValue: nil,
			UpdateTime:   timestamppb.Now(),
		}, nil
	}
}

func (conn *DeviceConnection) PublishPropertyUpdate(property *Property) error {
	pbProp, err := conn.marshalProperty(property)
	if err != nil {
		return err
	}
	return conn.Publish(&iotdevpb.PropertyUpdatedEvent{
		DeviceId: conn.ID,
		Property: pbProp,
	})
}

func (conn *DeviceConnection) PublishConnectedEvent(properties *DeviceProperties) error {
	pbProps := make([]*iotdevpb.Property, 0, len(properties.Properties))

	for _, prop := range properties.Properties {
		pbProp, err := conn.marshalProperty(prop)
		if err != nil {
			return err
		}
		pbProps = append(pbProps, pbProp)
	}

	return conn.Publish(&iotdevpb.DeviceConnectedEvent{
		DeviceId:   conn.ID,
		Properties: pbProps,
	})
}
