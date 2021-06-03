package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uffy/tools/pkg/json2csv"
)

var json2csvCmd = &cobra.Command{
	Use:   "json2csv",
	Short: "convert json to csv",
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := cmd.Flag("file").Value.String()
		output := cmd.Flag("output").Value.String()
		return json2csv.JSON2CSV(filename, output)
	},
}

func init() {
	json2csvCmd.PersistentFlags().StringP("file", "f", "", "json filename")
	json2csvCmd.PersistentFlags().StringP("output", "o", "{json_filename}.csv", "output filename")

	rootCmd.AddCommand(json2csvCmd)
}
