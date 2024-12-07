package cmd

import (
	"github.com/spf13/cobra"
	"github.com/will-wright-eng/sshm/internal/config"
	"github.com/will-wright-eng/sshm/internal/manager"
)

func newRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [name]",
		Short: "Remove an SSH host configuration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			fileConfig := config.NewFileConfig(configPath)
			hostManager := manager.NewHostManager(fileConfig)

			return hostManager.RemoveHost(name)
		},
	}
	return cmd
}
