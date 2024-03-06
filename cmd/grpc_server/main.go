package main

import (
	"context"
	"flag"
	"github.com/Murat993/chat-server/internal/app"
	"log"
)

var configPath string
var accessToken = flag.String("a", "", "access token")

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()

	ctxWithVal := context.WithValue(ctx, "accessToken", *accessToken)

	a, err := app.NewApp(ctxWithVal)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
