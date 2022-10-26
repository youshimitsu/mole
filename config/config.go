package config

import (
	"flag"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Log -.
	Log struct {
		Level string `env-required:"true" json:"log_level" yaml:"logLevel"`
	}

	// KeepAliveConfig -.
	KeepAliveConfig struct {
		Interval uint
		CountMax uint
	}

	// Tunnel -.
	Tunnel struct {
		KeepAlive *KeepAliveConfig
		Tunnel    string
		Server    string
		Password  string
	}

	// TunnelConfig -.
	TunnelConfig struct {
		KeepAlive *KeepAliveConfig `env-required:"true" json:"keep_alive" yaml:"keepAlive"`

		Tunnels []Tunnel `env-required:"true" json:"tunnels" yaml:"tunnels"`
	}
)

func NewConfig() (*TunnelConfig, error) {
	cfg := &TunnelConfig{}

	cfgPath := flag.String("config", "./config/config.example.json", "TunnelConfig path (JSON, YAML)")

	flag.Parse()

	err := cleanenv.ReadConfig(*cfgPath, cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
