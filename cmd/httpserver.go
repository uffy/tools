package cmd

import (
	"log"
	"net"
	"net/http"

	"github.com/spf13/cobra"
)

var httpServerCmd = &cobra.Command{
	Use:   "server",
	Short: "run a http server",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := net.Listen("tcp", ":0")
		if err != nil {
			return err
		}
		log.Println("listing", l.Addr().String())
		return http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello, world!"))
		}))
	},
}
