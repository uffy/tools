package pkg

import (
	"fmt"
	"github.com/spf13/cobra"
	"golang.org/x/tools/go/packages"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use: "pkg",
	}

	cmd.AddCommand(&cobra.Command{
		Use: "load <package> [package...]",
		RunE: func(cmd *cobra.Command, args []string) error {
			pkgs, err := packages.Load(&packages.Config{
				Mode: packages.NeedTypesInfo,
			}, args...)

			if err != nil {
				return err
			}

			for _, pkg := range pkgs {
				fmt.Println(pkg.String())
			}

			return nil
		},
	})

	return cmd

}
