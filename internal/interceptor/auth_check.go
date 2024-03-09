package interceptor

import (
	"context"
	"github.com/Murat993/chat-server/pkg/access_v1"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

type AuthCheckInterceptor struct {
	accessClient access_v1.AccessV1Client
}

const (
	authPrefix = "Bearer "
)

func NewAuthCheckInterceptor(accessClient access_v1.AccessV1Client) *AuthCheckInterceptor {
	return &AuthCheckInterceptor{accessClient: accessClient}
}

func (a *AuthCheckInterceptor) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return nil, errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	md = metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := a.accessClient.Check(ctx, &access_v1.CheckRequest{
		EndpointAddress: "",
	})

	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "you are not checked token")
	}

	return handler(ctx, req)
}
