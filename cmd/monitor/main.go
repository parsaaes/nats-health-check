package monitor

import (
	"github.com/4lie/nats-health-check/config"
	"github.com/spf13/cobra"
)

func main(cfg config.Config) {
}

// Register server command.
func Register(root *cobra.Command, cfg config.Config) {
	root.AddCommand(
		&cobra.Command{
			Use:   "monitor",
			Short: "Monitors Given Topic in nats/nats-streaming",
			Run: func(cmd *cobra.Command, args []string) {
				main(cfg)
			},
		},
	)
}
