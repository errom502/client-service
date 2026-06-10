package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/errom502/client-service/internal/client"
	"github.com/errom502/client-service/internal/config"
	"github.com/errom502/client-service/internal/handler"
	"github.com/errom502/client-service/internal/usecase"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	verificationClient, err := client.NewVerificationClient(cfg.Gateway.Addr)
	if err != nil {
		log.Fatalf("grpc client: %v", err)
	}
	defer verificationClient.Close()

	uc := usecase.NewVerificationUsecase(verificationClient)
	vh := handler.NewVerificationHandler(uc)
	router := handler.NewRouter(vh, cfg.HTTP.FrontendDir)

	srv := &http.Server{
		Addr:    ":" + cfg.HTTP.Port,
		Handler: router,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("client-service listening on :%s", cfg.HTTP.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
}
