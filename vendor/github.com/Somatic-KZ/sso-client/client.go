package client

import (
	"errors"

	"github.com/Somatic-KZ/sso-client/client"
	"github.com/Somatic-KZ/sso-client/client/grpc"
)

type Config struct {
	Protocol string
	Addr     string
}

func NewSSO(conf *Config) (client.SSO, error) {
	if conf.Protocol == "grpc" {
		cli := grpc.NewClient()
		if err := cli.Connect(conf.Addr); err != nil {
			return nil, err
		}
		return cli, nil
	}

	return nil, errors.New("unsupported protocol")
}
