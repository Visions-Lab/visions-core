package config

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	CronFile string `json:"cron_file"`
	LogLevel string `json:"log_level"`
	LogFile  string `json:"log_file"`
}

func Load(path string) (*AppConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg AppConfig
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
