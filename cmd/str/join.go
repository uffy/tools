package str

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func newJoin() *cobra.Command {
	var sepFlag string

	cmd := &cobra.Command{
		Use:  "join <value> [value...]",
		Long: "generate sql in clause, stdin is supported. separators: " + SupportedSeparatorsText,
		RunE: func(cmd *cobra.Command, args []string) error {
			args = NormalizeArgsOrParseFromStdIn(args)
			if len(args) == 0 {
				return cmd.Help()
			}

			rs := strings.Join(args, sepFlag)

			fmt.Println(rs)
			fmt.Println("total count:", len(args))

			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&sepFlag, "separator", "s", ",", "separator")

	return cmd
}
