package app

import (
	"github.com/spf13/cobra"
)

func NewFoundationServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "foundation",
		Short: "foundation: bootstrap a secure API server",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
	return cmd
}

func run() {
	// todo
}
