package auth

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ctxKey string

const clientKey ctxKey = "client"

type Client struct {
	Name string
}

type Authenticator struct {
	clients map[string]*Client
}

func NewAuthenticator(clients map[string]*Client) *Authenticator {
	return &Authenticator{clients: clients}
}

func (a *Authenticator) AuthMiddleware(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	if c, ok := a.clients[token]; ok {
		return context.WithValue(ctx, clientKey, c), nil
	}
	return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
}

func ClientFromContext(ctx context.Context) *Client {
	client, ok := ctx.Value(clientKey).(*Client)
	if !ok {
		return nil
	}
	return client
}
