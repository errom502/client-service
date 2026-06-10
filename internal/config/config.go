package config

import "github.com/jessevdk/go-flags"

type Config struct {
	// LogLevel уровень логирования
	LogLevel string `long:"log-level" description:"Log level: panic, fatal, warn, info, debug" env:"LOG_LEVEL" default:"warn"`
	Env      string `long:"env" env:"ENV" default:"develop"`
	// DevMode режим отладки
	DevMode bool `long:"dev-mode" env:"DEV_MODE" description:"Developer mode"`
	HTTP    HTTP
	Gateway Gateway
}

type HTTP struct {
	Port        string `long:"http-port" env:"HTTP_PORT" default:"8080"`
	FrontendDir string `long:"http-frontend-dir" env:"HTTP_FRONTEND_DIR" default:"./frontend/dist"`
}

type Gateway struct {
	Addr string `long:"gateway-addr" env:"GATEWAY_ADDR" default:"localhost:9881"`
}

func New() (*Config, error) {
	cfg := &Config{}
	_, err := flags.Parse(cfg)
	return cfg, err
}
