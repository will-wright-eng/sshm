package cmd

import (
    "github.com/spf13/cobra"
    "path/filepath"
)

var (
    configPath string
    rootCmd    = &cobra.Command{
        Use:   "sshm",
        Short: "SSH config manager",
        Long:  `A tool to manage SSH config entries in a controlled block.`,
    }
)

func Execute() error {
    return rootCmd.Execute()
}

func init() {
    configPath = filepath.Join("~", ".ssh", "config")
    rootCmd.AddCommand(newAddCmd())
    rootCmd.AddCommand(newRemoveCmd())
    rootCmd.AddCommand(newListCmd())
}
