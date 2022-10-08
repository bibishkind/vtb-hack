package app

import (
	"coffee-layered-architecture/internal/config"
	handler2 "coffee-layered-architecture/internal/handler"
	postgres2 "coffee-layered-architecture/internal/postgres"
	"coffee-layered-architecture/internal/server"
	service2 "coffee-layered-architecture/internal/service"
	"coffee-layered-architecture/pkg/auth"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configPath string) {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, cfg.Postgres.URI)
	if err != nil {
		log.Fatal(err)
	}

	if err = pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("database has been connected")

	postgres := postgres2.NewPostgres(pool)
	tokenManager := auth.NewTokenManager(1*time.Minute, 2*time.Minute)

	service := service2.NewService(postgres, tokenManager)

	handler := handler2.NewHandler(cfg, service)

	srv := server.NewServer(cfg, handler.Init())
	go func() {
		if err = srv.Run(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	log.Printf("server listening on port %s\n", cfg.HTTP.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err = srv.Stop(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("server shutdown")

	pool.Close()
}
