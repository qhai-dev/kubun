package app

import (
	"fmt"

	"github.com/lithammer/dedent"
	"github.com/qhai-dev/ozma/foundation"
	"github.com/spf13/cobra"
)

func NewFoundationServerCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "foundation",
		Short: "foundation: easily bootstrap a secure Foundation server",
		Long: dedent.Dedent(`

				    ┌──────────────────────────────────────────────────────────┐
				    │ FOUNDATION                                               │
				    │ Easily bootstrap a secure Kubernetes cluster             │
				    │                                                          │
				    │ Please give us feedback at:                              │
				    │ https://github.com/kubernetes/kubeadm/issues             │
				    └──────────────────────────────────────────────────────────┘

				Example usage:

				    Create a two-machine cluster with one control-plane node
				    (which controls the cluster), and one worker node
				    (where your workloads, like Pods and Deployments run).

				    ┌──────────────────────────────────────────────────────────┐
				    │ On the first machine:                                    │
				    ├──────────────────────────────────────────────────────────┤
				    │ control-plane# foundation start                          │
				    └──────────────────────────────────────────────────────────┘

				    ┌──────────────────────────────────────────────────────────┐
				    │ On the second machine:                                   │
				    ├──────────────────────────────────────────────────────────┤
				    │ worker# foundation stop <arguments-returned-from-init>   │
				    └──────────────────────────────────────────────────────────┘

				    You can then repeat the second step on as many other machines as you like.

			`),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Starting Foundation server...")
			
			s, err := foundation.InitializeApp()
			if err != nil {
				panic(err)
			}

			s.Run()
		},
	}


	return cmd
}