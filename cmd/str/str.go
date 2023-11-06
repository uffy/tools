package str

import (
	"fmt"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

var separators = []string{"\r\n", "\n", ",", "\t", " "}
var SupportedSeparatorsText = strings.Join(lo.Map(separators, func(s string, i int) string {
	return fmt.Sprintf("%q", s)
}), ", ")

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use: "str",
	}

	cmd.AddCommand(newJoin())
	cmd.AddCommand(newMap())

	return cmd

}

func NormalizeArgsOrParseFromStdIn(args []string) []string {
	if len(args) == 0 {
		bs, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil
		}

		in := strings.TrimSpace(string(bs))

		for _, sep := range separators {
			if strings.Contains(in, sep) {
				args = strings.Split(in, sep)
				break
			}
		}
	}

	args = lo.Filter(args, func(s string, i int) bool {
		return s != ""
	})

	return args
}
