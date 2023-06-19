package url

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
)

var encodeCmd = &cobra.Command{
	Use:   "encode string",
	Short: "url.QueryEscape(args[0])",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) == 0 {
			return cmd.Help()
		}
		v := url.QueryEscape(args[0])
		fmt.Println(v)
		return nil
	},
}
