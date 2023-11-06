package str

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"strings"
)

func newMap() *cobra.Command {
	var tmplFlag string
	var joinFlag string

	cmd := &cobra.Command{
		Use: "map <value> [value...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			args = NormalizeArgsOrParseFromStdIn(args)
			if len(args) == 0 {
				return cmd.Help()
			}

			args = lo.Map(args, func(s string, i int) string {
				return strings.ReplaceAll(tmplFlag, "{{.V}}", s)
			})

			rs := strings.Join(args, joinFlag)
			fmt.Println(rs)
			fmt.Println("total count:", len(args))

			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&tmplFlag, "template", "t", "value: {{.V}}", "template")
	cmd.PersistentFlags().StringVarP(&joinFlag, "join", "j", "\n", "join")

	return cmd
}
