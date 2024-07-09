package persil

import (
	"context"
	"fmt"
	"henkel-somtec-devscript/devscript"
)

type User struct {
	cli *devscript.ApiClient
}

func Login(username, password string) *User {
	user, err := login(username, password)
	if err != nil {
		panic(err)
	}
	return user
}

func login(username, password string) (*User, error) {
	ctx := context.Background()

	cli, err := devscript.NewApiClient(ctx, username, password)
	if err != nil {
		return nil, err
	}

	user, err := cli.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	if user.User.Product != "persil" {
		return nil, fmt.Errorf("user is not persil")
	}

	return &User{
		cli: cli,
	}, nil
}

func (u *User) ProvisionDevice() *Device {
	dev, err := u.provisionDevice()
	if err != nil {
		panic(err)
	}
	return dev
}

func (u *User) provisionDevice() (*Device, error) {
	res, err := u.cli.ProvisionNewDevice(context.Background())
	if err != nil {
		return nil, err
	}

	return newDevice(&devscript.DeviceConnection{
		ID:       res.ProvisionNewDevice.Device.ID,
		Username: res.ProvisionNewDevice.Connections.Mqtt.Username,
		Password: res.ProvisionNewDevice.Connections.Mqtt.Password,
		Addr:     res.ProvisionNewDevice.Connections.Mqtt.Address,
	}), nil
}
