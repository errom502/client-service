package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/errom502/client-service/internal/client"
	"github.com/errom502/client-service/internal/config"
	"github.com/errom502/client-service/internal/handler"
	"github.com/errom502/client-service/internal/usecase"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	logger, err := initLogger(cfg)
	if err != nil {
		log.Fatalf("logger: %v", err)
	}

	verificationClient, err := client.NewVerificationClient(cfg.Gateway.Addr)
	if err != nil {
		logger.Fatal("grpc client init failed", zap.Error(err))
	}
	defer verificationClient.Close()

	uc := usecase.NewVerificationUsecase(verificationClient, logger)
	vh := handler.NewVerificationHandler(uc, logger)

	router := handler.NewRouter(vh, cfg.HTTP.FrontendDir)

	srv := &http.Server{
		Addr:    ":" + cfg.HTTP.Port,
		Handler: router,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		logger.Info("client-service started", zap.String("port", cfg.HTTP.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("server error", zap.Error(err))
		}
	}()

	<-ctx.Done()
	logger.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error("shutdown error", zap.Error(err))
	}
}

func initLogger(cfg *config.Config) (*zap.Logger, error) {
	// TODO: в бущем подключить к удаленной системе для просмотра логов
	level := parseLogLevel(cfg.LogLevel)

	zapConfig := zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: cfg.DevMode,
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:       "ts",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "msg",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
			EncodeCaller:  zapcore.ShortCallerEncoder,
		},

		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build zap config: %w", err)
	}

	logger = logger.With(
		zap.String("service", "client-service"),
		zap.String("env", cfg.Env),
	)

	return logger, err
}

func parseLogLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "fatal":
		return zapcore.FatalLevel
	case "panic":
		return zapcore.PanicLevel
	default:
		return zapcore.WarnLevel
	}
}
