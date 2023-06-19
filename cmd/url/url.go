package url

import (
	"github.com/spf13/cobra"
)

func NewUrlCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "url",
		Short: "url-related",
	}

	cmd.AddCommand(encodeCmd)

	return cmd
}
