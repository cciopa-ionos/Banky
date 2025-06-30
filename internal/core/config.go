package core

import "os"

type Config struct {
	BankyPath string
}

func LoadConfig() *Config {
	cfg := &Config{
		BankyPath: "example",
	}

	if val := os.Getenv("BANKY_PATH"); val != "" {
		cfg.BankyPath = val
	}

	return cfg
}
