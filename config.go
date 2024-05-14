package main

import (
	"context"

	"github.com/caarlos0/env/v11"
)

// InfluxConfig holds the configuration for the InfluxDB client
type InfluxdbConfig struct {
	Host     string `env:"INFLUXDB_HOST"`
	Database string `env:"INFLUXDB_DATABASE"`
	Token    string `env:"INFLUXDB_TOKEN"`
}

// New instantiates a new InfluxConfig struct
func NewConfig(ctx context.Context) (*InfluxdbConfig, error) {
	cfg := InfluxdbConfig{}
	opts := env.Options{RequiredIfNoDef: true}

	err := env.ParseWithOptions(&cfg, opts)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
