package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/will-wright-eng/sshm/internal/config"
	"github.com/will-wright-eng/sshm/internal/manager"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all managed SSH hosts",
		RunE: func(cmd *cobra.Command, args []string) error {
			fileConfig := config.NewFileConfig(configPath)
			hostManager := manager.NewHostManager(fileConfig)

			hosts, err := hostManager.ListHosts()
			if err != nil {
				return err
			}

			if len(hosts) == 0 {
				fmt.Println("No managed SSH hosts found")
				return nil
			}

			for _, host := range hosts {
				fmt.Printf("Name: %s\n  Hostname: %s\n  User: %s\n  Port: %d\n\n",
					host.Name, host.Hostname, host.User, host.Port)
			}
			return nil
		},
	}
	return cmd
}
