package cmd

import (
	"os"

	"github.com/4lie/nats-health-check/cmd/monitor"
	"github.com/4lie/nats-health-check/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ExitFailure status code.
const ExitFailure = 1

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cfg := config.New()

	root := &cobra.Command{
		Use:   "nats-health-check",
		Short: "A program for health check of each NATS server node",
	}

	monitor.Register(root, cfg)

	if err := root.Execute(); err != nil {
		logrus.Errorf("failed to execute root command: %s", err.Error())
		os.Exit(ExitFailure)
	}
}
