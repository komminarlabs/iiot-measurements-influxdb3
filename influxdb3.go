package main

import (
	"context"
	"fmt"
	"time"

	"github.com/InfluxCommunity/influxdb3-go/influxdb3"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/influxdata/line-protocol/v2/lineprotocol"
)

// Measurement represents a single measurement
type Measurement struct {
	Measurement string    `lp:"measurement"`
	Timestamp   time.Time `lp:"timestamp"`

	// InfluxDB Line protocol's field set
	Speed       float64 `lp:"field,speed"`
	Temperature float32 `lp:"field,temperature"`

	// InfluxDB Line protocol's tag set
	Site string `lp:"tag,site"`
	Line string `lp:"tag,line"`
}

// influxdb represents an InfluxDB client
type influxdb struct {
	client influxdb3.Client
}

// NewClient creates a new InfluxDB client
func NewClient(ctx context.Context, cfg InfluxdbConfig) (*influxdb, error) {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10

	options := influxdb3.DefaultWriteOptions
	options.Precision = lineprotocol.Millisecond

	client, err := influxdb3.New(influxdb3.ClientConfig{
		Host:       cfg.Host,
		Token:      cfg.Token,
		Database:   cfg.Database,
		HTTPClient: retryClient.StandardClient(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create influx client: %w", err)
	}
	return &influxdb{client: *client}, nil
}

// BatchWrite writes a batch of measurements to InfluxDB3
func (i *influxdb) BatchWrite(ctx context.Context, measurements []*Measurement) error {
	fmt.Printf("ingesting %d measurement(s) to InfluxDb3\n", len(measurements))
	measurementsBatch := make([]any, len(measurements))
	for i, measurement := range measurements {
		measurementsBatch[i] = measurement
	}

	err := i.client.WriteData(context.Background(), measurementsBatch)
	if err != nil {
		return fmt.Errorf("failed to write measurement(s) to InfluxDB3: %w", err)
	}
	fmt.Printf("wrote %d measurement(s)\n", len(measurementsBatch))
	return nil
}
