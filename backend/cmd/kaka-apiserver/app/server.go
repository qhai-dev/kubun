package app

import (
	"github.com/qhai-dev/kaka/apiserver/rest"

	"github.com/spf13/cobra"
)

func NewAPIServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "apiserver",
		Short: "apiserver: easily bootstrap a secure API server",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
	return cmd
}

func run() {
	rest.Run()
}
