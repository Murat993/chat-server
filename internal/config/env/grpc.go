package env

import (
	"github.com/Murat993/chat-server/internal/config"
	"github.com/pkg/errors"
	"net"
	"os"
)

var _ config.GRPCConfig = (*grpcConfig)(nil)

const (
	grpcHostEnvName     = "GRPC_HOST"
	grpcPortEnvName     = "GRPC_PORT"
	grpcPortAuthEnvName = "GRPC_PORT_AUTH"
)

type grpcConfig struct {
	host     string
	port     string
	portAuth string
}

func NewGRPCConfig() (*grpcConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	portAuth := os.Getenv(grpcPortAuthEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port auth not found")
	}

	return &grpcConfig{
		host:     host,
		port:     port,
		portAuth: portAuth,
	}, nil
}

func (cfg *grpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func (cfg *grpcConfig) AddressAuth() string {
	return net.JoinHostPort(cfg.host, cfg.portAuth)
}
