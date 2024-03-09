package app

import (
	"context"
	"github.com/Murat993/chat-server/internal/api/chat"
	"github.com/Murat993/chat-server/internal/client/db"
	"github.com/Murat993/chat-server/internal/client/db/pg"
	"github.com/Murat993/chat-server/internal/client/db/transaction"
	"github.com/Murat993/chat-server/internal/closer"
	"github.com/Murat993/chat-server/internal/config"
	"github.com/Murat993/chat-server/internal/config/env"
	"github.com/Murat993/chat-server/internal/repository"
	chatRepository "github.com/Murat993/chat-server/internal/repository/chat"
	"github.com/Murat993/chat-server/internal/service"
	chatService "github.com/Murat993/chat-server/internal/service/chat"
	"github.com/Murat993/chat-server/pkg/access_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig

	txManager      db.TxManager
	dbClient       db.Client
	chatRepository repository.ChatRepository

	chatService service.ChatService

	chatImpl *chat.Implementation

	accessClient access_v1.AccessV1Client
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()

		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		grpcConfig, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.grpcConfig = grpcConfig
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepository.NewRepository(s.DBClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(
			s.ChatRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(_ context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation()
	}

	return s.chatImpl
}

func (s *serviceProvider) connectGRPCClient(addressAuth string) (access_v1.AccessV1Client, error) {
	creds, err := credentials.NewClientTLSFromFile("certificates/service.pem", "")
	if err != nil {
		return nil, err
	}

	log.Printf("GRPC client is running on %s", addressAuth)

	conn, err := grpc.Dial(s.grpcConfig.AddressAuth(), grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	c := access_v1.NewAccessV1Client(conn)

	s.accessClient = c

	return s.accessClient, nil
}
