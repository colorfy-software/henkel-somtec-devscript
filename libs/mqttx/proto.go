package mqttx

import (
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func MarshalProto(payload proto.Message) ([]byte, error) {
	anyMessage, err := anypb.New(payload)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(anyMessage)
}

func UnmarshalProto(message []byte) (proto.Message, error) {
	var anyMessage anypb.Any
	err := proto.Unmarshal(message, &anyMessage)
	if err != nil {
		return nil, err
	}
	dst, err := anyMessage.UnmarshalNew()
	if err != nil {
		return nil, errors.Wrapf(err, "could not unmarshal type url: '%s'", anyMessage.TypeUrl)
	}

	return dst, nil
}
