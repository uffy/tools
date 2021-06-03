package cmd

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var json2csvCmd = &cobra.Command{
	Use:   "json2csv",
	Short: "convert json to csv",
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := cmd.Flag("file").Value.String()
		output := cmd.Flag("output").Value.String()
		output = strings.Replace(output, "{json_filename}", filename, 1)
		bs, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		lines := bytes.Split(bs, []byte{'\n'})
		if len(lines) == 0 {
			return errors.New("file is empty")
		}

		var headers []string
		headersSet := map[string]int{}

		record := make([]string, 0)    // for IDE lint
		records := make([][]string, 1) // first element for headers
		for _, line := range lines {
			lineMap := map[string]json.RawMessage{}
			if err := json.Unmarshal(line, &lineMap); err != nil {
				continue
			}
			for h, c := range lineMap {
				if _, ok := headersSet[h]; !ok {
					headers = append(headers, h)
					headersSet[h] = len(headers) - 1
				}
				i := headersSet[h]
				if i >= len(record) {
					newRecord := make([]string, len(headers))
					copy(newRecord, record)
					record = newRecord
				}

				if v, err := strconv.Unquote(string(c)); err == nil {
					record[i] = v
				} else {
					record[i] = string(c)
				}
			}

			records = append(records, record)
		}

		records[0] = headers

		outputFile, err := os.Create(output)
		if err != nil {
			return err
		}
		writer := csv.NewWriter(outputFile)
		if err := writer.WriteAll(records); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	json2csvCmd.PersistentFlags().StringP("file", "f", "", "json filename")
	json2csvCmd.PersistentFlags().StringP("output", "o", "{json_filename}.csv", "output filename")

	rootCmd.AddCommand(json2csvCmd)
}
