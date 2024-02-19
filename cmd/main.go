package main

import (
	"context"
	"database/sql"
	"flag"
	sq "github.com/Masterminds/squirrel"
	"github.com/Murat993/chat-server/internal/config"
	"github.com/Murat993/chat-server/internal/config/env"
	"github.com/Murat993/chat-server/pkg/chat_v1"
	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	chat_v1.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	req.GetUsernames()
	return &chat_v1.CreateResponse{
		Id: 5,
	}, nil
}

func main() {
	flag.Parse()
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Конфиг grpc в config/grpc.go
	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to get grpc config: %v", err)
	}

	// Конфиг grpc в config/pg.go
	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address()) // tcp подключение по порту 50051
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	// Делаем запрос на вставку записи в таблицу chat
	builderInsert := sq.Insert("chat").
		PlaceholderFormat(sq.Dollar).
		Columns("form", "text").
		Values(gofakeit.City(), gofakeit.Address().Street).
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var chatID int
	err = pool.QueryRow(ctx, query, args...).Scan(&chatID)
	if err != nil {
		log.Fatalf("failed to insert chat: %v", err)
	}

	log.Printf("inserted chat with id: %d", chatID)

	// Делаем запрос на выборку записей из таблицы chat
	builderSelect := sq.Select("id", "form", "text", "created_at", "updated_at").
		From("chat").
		PlaceholderFormat(sq.Dollar).
		OrderBy("id ASC").
		Limit(10)

	query, args, err = builderSelect.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	rows, err := pool.Query(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to select chats: %v", err)
	}

	var id int
	var form, text string
	var createdAt time.Time
	var updatedAt sql.NullTime

	for rows.Next() {
		err = rows.Scan(&id, &form, &text, &createdAt, &updatedAt)
		if err != nil {
			log.Fatalf("failed to scan chat: %v", err)
		}

		log.Printf("id: %d, form: %s, text: %s, created_at: %v, updated_at: %v\n", id, form, text, createdAt, updatedAt)
	}

	s := grpc.NewServer()                      // Создаем новый grpc сервер
	reflection.Register(s)                     // Включается параметр. где со стороны клиента можно посмотреть какие есть ендпоинты
	chat_v1.RegisterChatV1Server(s, &server{}) // Регистрируем сервер grpc

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil { // Serve - блокирует и скрип вечно вист с портом который указан в lis
		log.Fatalf("failed to serve: %v", err)
	}
}
