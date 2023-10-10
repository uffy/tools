package sql

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

var separators = []string{"\r\n", "\n", ",", "\t", " "}
var supportedSeparatorsText = strings.Join(lo.Map(separators, func(s string, i int) string {
	return fmt.Sprintf("%q", s)
}), ", ")

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use: "sql",
	}

	cmd.AddCommand(&cobra.Command{
		Use:  "in <value> [value...]",
		Long: "generate sql in clause, stdin is supported. separators: " + supportedSeparatorsText,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				bs, err := io.ReadAll(os.Stdin)
				if err != nil {
					return cmd.Help()
				}

				for _, sep := range separators {
					if strings.Contains(string(bs), sep) {
						args = strings.Split(string(bs), sep)
						break
					}
				}
			}

			args = lo.Filter(args, func(s string, i int) bool {
				return s != ""
			})

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
