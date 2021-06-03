package json2csv

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func JSON2CSV(filename, output string) error {
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

	records := make([][]string, 1) // first element for headers
	for _, line := range lines {
		lineMap := map[string]json.RawMessage{}
		if err := json.Unmarshal(line, &lineMap); err != nil {
			continue
		}
		record := make([]string, len(headers)) // for IDE lint
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
}
