package client

import "context"

type SSO interface {
	Connect(addr string) error
	UserToken(ctx context.Context, id string) error

}
