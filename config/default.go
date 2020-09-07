package config

import "time"

// Default return default configuration
// nolint: gomnd
func Default() Config {
	return Config{
		NATS: NATS{
			URL:           "nats://127.0.0.1:4222",
			MaxReconnect:  60,
			ReconnectWait: 1 * time.Second,
		},
		Streaming: Streaming{
			ClusterID: "charlie",
		},
	}
}
