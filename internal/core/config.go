package core

import "os"

type Config struct {
	OperationsPath string
	BankyPath      string
}

func LoadConfig() *Config {
	cfg := &Config{
		OperationsPath: "./banky/operations.json",
		BankyPath:      "./banky/banky.json",
	}

	// Override with env vars
	if val := os.Getenv("OPERATIONS_PATH"); val != "" {
		cfg.OperationsPath = val
	}

	if val := os.Getenv("BANKY_PATH"); val != "" {
		cfg.BankyPath = val
	}

	return cfg
}
