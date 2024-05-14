package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	cfg, err := NewConfig(ctx)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	client, err := NewClient(ctx, *cfg)
	if err != nil {
		log.Fatalf("failed to create InfluxDB3 client: %v", err)
	}

	measurements := []*Measurement{
		{
			Measurement: "filler1",
			Timestamp:   time.Now(),
			Speed:       100.0,
			Temperature: 25.0,
			Site:        "site1",
			Line:        "line1",
		},
		{
			Measurement: "filler2",
			Timestamp:   time.Now(),
			Speed:       120.0,
			Temperature: 26.0,
			Site:        "site2",
			Line:        "line2",
		},
	}

	err = client.BatchWrite(ctx, measurements)
	if err != nil {
		log.Fatalf("failed to write batch: %v", err)
	}
}
