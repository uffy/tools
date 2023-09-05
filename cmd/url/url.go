package url

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/url"
)

func New() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "url",
		Short: "url-related",
	}

	cmd.AddCommand(&cobra.Command{
		Use: "encode [string]",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if len(args) == 0 {
				return cmd.Help()
			}
			v := url.QueryEscape(args[0])
			fmt.Println(v)
			return nil
		},
	}, &cobra.Command{
		Use: "decode [string]",
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			if len(args) == 0 {
				return cmd.Help()
			}
			v, err := url.QueryUnescape(args[0])
			if err != nil {
				return err
			}

			fmt.Println(v)
			return nil
		},
	})

	return cmd
}
