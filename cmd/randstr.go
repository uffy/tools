package cmd

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var base64UrlReplacer = strings.NewReplacer("-", "", "_", "")
var randStrCmd = &cobra.Command{
	Use:   "randstr",
	Short: "convert json to csv",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var l int64
		if len(args) > 0 {
			l, err = strconv.ParseInt(args[0], 10, 0)
			if err != nil {
				return err
			}
		}
		if l <= 0 {
			l = 40
		}

		buf := make([]byte, l)
		if _, err = rand.Read(buf); err != nil {
			return err
		}

		randStr := base64UrlReplacer.Replace(base64.RawURLEncoding.EncodeToString(buf))
		randStr = randStr[:l]

		fmt.Println(randStr)
		fmt.Println(strings.ToLower(randStr))
		fmt.Println(strings.ToUpper(randStr))
		return nil
	},
}
