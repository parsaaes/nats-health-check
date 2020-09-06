package config

// Default return default configuration
// nolint: gomnd
func Default() Config {
	return Config{
		NATS: NATS{
			URL: "nats://127.0.0.1:4222",
		},
		Streaming: Streaming{
			ClusterID: "4lie",
		},
	}
}
