package sql

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/uffy/tools/cmd/str"
	"strings"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use: "sql",
	}

	cmd.AddCommand(&cobra.Command{
		Use:  "in <value> [value...]",
		Long: "generate sql in clause, stdin is supported. separators: " + str.SupportedSeparatorsText,
		RunE: func(cmd *cobra.Command, args []string) error {
			args = str.NormalizeArgsOrParseFromStdIn(args)
			if len(args) == 0 {
				return cmd.Help()
			}

			inClause := strings.Join(lo.Map(args, func(s string, i int) string {
				return fmt.Sprintf("'%s'", s)
			}), ", ")

			fmt.Println(inClause)
			fmt.Println("total count:", len(args))

			return nil
		},
	})

	return cmd

}
