package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uffy/tools/cmd/sql"
	"github.com/uffy/tools/cmd/str"
	"github.com/uffy/tools/cmd/time"
	"github.com/uffy/tools/cmd/url"
)

// Execute executes the root command.
func Execute() error {
	return NewRoodCmd().Execute()
}

func NewRoodCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "uffy",
		Short: "Uffy tools",
	}

	cmd.AddCommand(httpServerCmd)
	cmd.AddCommand(json2csvCmd)
	cmd.AddCommand(randStrCmd)
	cmd.AddCommand(url.New())
	cmd.AddCommand(time.New())
	cmd.AddCommand(sql.New())
	cmd.AddCommand(str.New())

	return cmd
}
