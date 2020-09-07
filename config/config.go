package config

import (
	"strings"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/sirupsen/logrus"
)

const (
	// Prefix indicates environment variables prefix.
	Prefix = "nats_health_check_"
)

type (
	// Config holds all configurations.
	Config struct {
		NATS      NATS      `koanf:"nats"`
		Streaming Streaming `koanf:"streaming"`
	}

	// NATS configuration.
	NATS struct {
		URL           string        `koanf:"url"`
		MaxReconnect  int           `koanf:"max-reconnect"`
		ReconnectWait time.Duration `koanf:"reconnect-wait"`
	}

	// NATS Streaming Configuration.
	Streaming struct {
		ClusterID string `koanf:"cluster"`
	}
)

// New reads configuration with viper.
func New() Config {
	var instance Config

	k := koanf.New(".")

	// load default configuration from file
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		logrus.Fatalf("error loading default: %s", err)
	}

	// load configuration from file
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		logrus.Errorf("error loading config.yml: %s", err)
	}

	// load environment variables
	if err := k.Load(env.Provider(Prefix, ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, Prefix)), "_", ".", -1)
	}), nil); err != nil {
		logrus.Errorf("error loading environment variables: %s", err)
	}

	if err := k.Unmarshal("", &instance); err != nil {
		logrus.Fatalf("error unmarshalling config: %s", err)
	}

	logrus.Infof("following configuration is loaded:\n%+v", instance)

	return instance
}
