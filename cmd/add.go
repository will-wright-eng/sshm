package cmd

import (
	"github.com/spf13/cobra"
	"github.com/will-wright-eng/sshm/internal/config"
	"github.com/will-wright-eng/sshm/internal/manager"
	"github.com/will-wright-eng/sshm/internal/models"
)

func newAddCmd() *cobra.Command {
	var (
		hostname     string
		user         string
		identityFile string
		port         int
	)

	cmd := &cobra.Command{
		Use:   "add [name]",
		Short: "Add a new SSH host configuration",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := args[0]

			fileConfig := config.NewFileConfig(configPath)
			hostManager := manager.NewHostManager(fileConfig)

			entry := models.NewHostEntry(name, hostname, user, identityFile, port)
			return hostManager.AddHost(entry)
		},
	}

	cmd.Flags().StringVarP(&hostname, "hostname", "H", "", "Remote hostname or IP address")
	cmd.Flags().StringVarP(&user, "user", "u", "", "SSH username")
	cmd.Flags().StringVarP(&identityFile, "identity-file", "i", "", "SSH identity file path")
	cmd.Flags().IntVarP(&port, "port", "p", 22, "SSH port")

	cmd.MarkFlagRequired("hostname")
	cmd.MarkFlagRequired("user")

	return cmd
}
