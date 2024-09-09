package main

import (
	"log"
	"mas/monitor"
	"time"
)

func main() {
	config, err := monitor.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	for {
		metrics := monitor.CollectMetrics()

		// Log the metrics
		monitor.LogMetrics(metrics)

		// Check thresholds and alert if necessary
		monitor.CheckThresholds(metrics, config)

		// Sleep for the configured interval
		time.Sleep(time.Duration(config.Interval) * time.Second)
	}
}
