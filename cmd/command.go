package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "uffy",
	Short: "Uffy tools",
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
