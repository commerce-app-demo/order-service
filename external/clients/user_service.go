package clients

import (
	"context"

	userspb "github.com/commerce-app-demo/user-service/proto"
)

type UserClient struct {
	pbClient *userspb.UserServiceClient
	ctx      context.Context
}

func (cl *UserClient) ValidateUser(id string) (*userspb.User, error) {
	user, err := (*cl.pbClient).GetUser(cl.ctx, &userspb.GetUserRequest{
		Id: id,
	})

	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserClient(ctx context.Context, pb *userspb.UserServiceClient) (*UserClient, error) {
	return &UserClient{
		pbClient: pb,
		ctx:      ctx,
	}, nil
}
