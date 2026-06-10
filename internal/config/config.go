package config

import "github.com/jessevdk/go-flags"

type Config struct {
	Env     string `long:"env" env:"ENV" default:"develop"`
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
