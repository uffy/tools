package time

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"time"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use: "time [timestamp|date|datetime]",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}

			t := parseTime(args)
			if t.IsZero() {
				return fmt.Errorf("invalid time format")
			}

			fmt.Println(t.Unix())
			fmt.Println(t.Format(time.DateTime))
			return nil
		},
	}

	cmd.AddCommand(&cobra.Command{
		Use: "now",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(time.Now().Unix())
			fmt.Println(time.Now().Format(time.DateTime))
		},
	}, &cobra.Command{
		Use: "since",
		RunE: func(cmd *cobra.Command, args []string) error {
			t := parseTime(args)
			if t.IsZero() {
				return fmt.Errorf("invalid time format")
			}

			fmt.Printf("%s (%d) since: %s or %ds\n",
				t.Format(time.DateTime),
				t.Unix(),
				time.Since(t).String(),
				int64(time.Since(t).Seconds()),
			)

			return nil
		},
	})

	return cmd
}

func parseTime(args []string) time.Time {
	arg := args[0]
	if len(args) == 2 {
		arg = args[0] + " " + args[1]
	}

	ts, err := strconv.ParseUint(arg, 10, 64)
	if err == nil {
		return time.Unix(int64(ts), 0)
	}

	t, err := time.Parse(time.DateTime, arg)
	if err == nil {
		return t
	}

	t, err = time.Parse(time.DateOnly, arg)
	if err == nil {
		return t
	}

	return time.Time{}
}
