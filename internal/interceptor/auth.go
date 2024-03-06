package interceptor

import (
	"context"
	"github.com/Murat993/chat-server/pkg/access_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	accessClient access_v1.AccessV1Client
}

func NewAuthInterceptor(accessClient access_v1.AccessV1Client) *AuthInterceptor {
	return &AuthInterceptor{accessClient: accessClient}
}

func (a *AuthInterceptor) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	val, ok := ctx.Value("accessToken").(string)
	if !ok {
		return nil, status.Errorf(codes.Internal, "could not retrieve token from context")
	}

	md := metadata.New(map[string]string{"Authorization": "Bearer " + val})
	ctx = metadata.NewOutgoingContext(ctx, md)

	_, err := a.accessClient.Check(ctx, &access_v1.CheckRequest{
		EndpointAddress: "",
	})

	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "you are not checked token")
	}

	return handler(ctx, req)
}
