package grpc

import (
	"context"

	"github.com/Somatic-KZ/sso-client/client"
	"github.com/Somatic-KZ/sso-client/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SSOClient struct {
	client protobuf.SSOClient
}

func NewClient() *SSOClient {
	return &SSOClient{}
}

func(s *SSOClient) Connect(addr string) error {
	cc, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	s.client = protobuf.NewSSOClient(cc)

	return nil
}

func (s *SSOClient) Disconnect() error {
	return nil
}
func (s *SSOClient) UserToken(ctx context.Context, id string) error {
	req := protobuf.UserTokenRequest{Id: id}

	_, err := s.client.UserToken(ctx, &req)
	switch status.Code(err) {
	case codes.NotFound:
		return client.ErrTokenNotFound
	case codes.Internal:
		return err
	default:
		return nil
	}
}
